param(
    [ValidateSet("blank", "safari")]
    [string]$Example = "blank",
    [ValidateSet("native", "linux", "windows")]
    [string]$Target = "windows",
    [switch]$All,
    [switch]$Deps,
    [switch]$Clean,
    [switch]$DeployToServer,
    [switch]$StopServer,
    [switch]$StartServer,
    [switch]$Test,
    [switch]$Full
)

$ErrorActionPreference = "Stop"
$Root = $PSScriptRoot
$PluginDir = Join-Path $Root "plugins"
$Header = Join-Path $Root "include\plugin.h"
$ServerRoot = Join-Path (Split-Path $Root -Parent) "vcmp-go-server"

$PluginNames = @{
    blank  = "goplugin04rel64"
    safari = "goserver04rel64"
}

# Safari dev workflow: test library, stop server, build, deploy.
if ($Full) {
    $Example = "safari"
    $Target = "windows"
    $DeployToServer = $true
    $StopServer = $true
    $Test = $true
}

if ($Example -eq "safari" -and (Test-Path $ServerRoot) -and -not $DeployToServer -and -not $Clean -and -not $Deps -and -not $All) {
    $DeployToServer = $true
}

function Ensure-Deps {
    if (-not (Test-Path $Header)) {
        Write-Host "Fetching include/plugin.h..."
        Push-Location (Join-Path $Root "scripts")
        try {
            go run fetch_plugin.go
        } finally {
            Pop-Location
        }
        if (-not (Test-Path $Header)) {
            throw "deps failed: $Header not found"
        }
    }
}

function Stop-VcmpServer {
    $procs = Get-Process -Name server64 -ErrorAction SilentlyContinue
    if (-not $procs) {
        return $false
    }
    Write-Host "Stopping server64..."
    $procs | Stop-Process -Force
    Start-Sleep -Seconds 1
    return $true
}

function Start-VcmpServer {
    if (-not (Test-Path $ServerRoot)) {
        Write-Warning "Start skipped: $ServerRoot not found"
        return
    }
    $exe = Join-Path $ServerRoot "server64.exe"
    if (-not (Test-Path $exe)) {
        Write-Warning "Start skipped: server64.exe not found in $ServerRoot"
        return
    }
    Write-Host "Starting server64.exe..."
    Start-Process -FilePath $exe -WorkingDirectory $ServerRoot
}

function Invoke-ServerTests {
    if (-not (Test-Path $ServerRoot)) {
        Write-Warning "Tests skipped: $ServerRoot not found"
        return
    }
    Write-Host "Running tests in vcmp-go-server..."
    Push-Location $ServerRoot
    try {
        go test ./...
        if ($LASTEXITCODE -ne 0) { throw "go test failed in vcmp-go-server" }
        Write-Host "Tests OK"
    } finally {
        Pop-Location
    }
}

function Invoke-Build {
    param(
        [string]$Ex,
        [string]$OsTarget
    )

    Ensure-Deps
    New-Item -ItemType Directory -Force -Path $PluginDir | Out-Null

    $name = $PluginNames[$Ex]
    $exampleDir = Join-Path $Root "examples\$Ex"

    if ($Ex -eq "safari") {
        Push-Location $exampleDir
        try { go mod tidy } finally { Pop-Location }
    }

    $env:CGO_ENABLED = "1"
    $env:GOOS = $null
    $env:GOARCH = $null
    $env:CC = $null

    switch ($OsTarget) {
        "linux" {
            $env:GOOS = "linux"
            $env:GOARCH = "amd64"
            $out = Join-Path $PluginDir "$name.so"
        }
        "windows" {
            $env:GOOS = "windows"
            $env:GOARCH = "amd64"
            $env:CC = "x86_64-w64-mingw32-gcc"
            $out = Join-Path $PluginDir "$name.dll"
        }
        default {
            if ($IsWindows -or $env:OS -match "Windows") {
                $out = Join-Path $PluginDir "$name.dll"
            } else {
                $out = Join-Path $PluginDir "$name.so"
            }
        }
    }

    Write-Host "Building $Ex -> $out"
    Push-Location $exampleDir
    try {
        go build -buildmode=c-shared -o $out .
        if ($LASTEXITCODE -ne 0) { throw "go build failed" }
    } finally {
        Pop-Location
        Remove-Item Env:GOOS -ErrorAction SilentlyContinue
        Remove-Item Env:GOARCH -ErrorAction SilentlyContinue
        Remove-Item Env:CC -ErrorAction SilentlyContinue
    }

    Write-Host "OK: $out"
    return $out
}

function Deploy-PluginToServer {
    param([string]$PluginPath)

    if (-not (Test-Path $ServerRoot)) {
        Write-Warning "Deploy skipped: $ServerRoot not found"
        return
    }
    $destDir = Join-Path $ServerRoot "plugins"
    New-Item -ItemType Directory -Force -Path $destDir | Out-Null
    $dest = Join-Path $destDir (Split-Path $PluginPath -Leaf)

    try {
        Copy-Item -Force $PluginPath $dest
    } catch {
        Write-Host "Deploy locked — stopping server64 and retrying..."
        Stop-VcmpServer
        Copy-Item -Force $PluginPath $dest
    }

    $info = Get-Item $dest
    Write-Host "Deployed -> $dest ($([math]::Round($info.Length / 1MB, 2)) MB, $($info.LastWriteTime))"
}

if ($Clean) {
    Remove-Item -Force -ErrorAction SilentlyContinue `
        (Join-Path $PluginDir "goplugin04rel64.so"),
        (Join-Path $PluginDir "goplugin04rel64.dll"),
        (Join-Path $PluginDir "goserver04rel64.so"),
        (Join-Path $PluginDir "goserver04rel64.dll"),
        (Join-Path $Root "goplugin04rel64.h"),
        (Join-Path $Root "goserver04rel64.h")
    Write-Host "Cleaned plugin outputs"
    exit 0
}

if ($Deps) {
    Ensure-Deps
    Write-Host "Deps OK: $Header"
    exit 0
}

if ($StopServer) {
    Stop-VcmpServer | Out-Null
}

if ($Test) {
    Invoke-ServerTests
}

if ($All) {
    $blankOut = Invoke-Build -Ex "blank" -OsTarget $Target
    $safariOut = Invoke-Build -Ex "safari" -OsTarget $Target
    if ($DeployToServer -and $safariOut) { Deploy-PluginToServer $safariOut }
    if ($StartServer) { Start-VcmpServer }
    exit 0
}

$out = Invoke-Build -Ex $Example -OsTarget $Target
if ($DeployToServer -and $out) {
    Deploy-PluginToServer $out
}
if ($StartServer) {
    Start-VcmpServer
}
