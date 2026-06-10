param(
    [ValidateSet("native", "linux", "windows")]
    [string]$Target = "windows",
    [string]$ServerRoot = "",
    [switch]$NoDeploy,
    [switch]$StopServer,
    [switch]$StartServer,
    [switch]$Test,
    [switch]$Deps,
    [switch]$Clean
)

$ErrorActionPreference = "Stop"
$Root = $PSScriptRoot
$PluginDir = Join-Path $Root "plugins"
$PluginSrc = Join-Path $Root "plugin"
$Header = Join-Path $Root "include\plugin.h"
$PluginName = "goserver04rel64"

if (-not $ServerRoot) {
    $ServerRoot = Join-Path (Split-Path $Root -Parent) "server"
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
    Write-Host "Running tests in server..."
    Push-Location $ServerRoot
    try {
        go test ./...
        if ($LASTEXITCODE -ne 0) { throw "go test failed in server" }
        Write-Host "Tests OK"
    } finally {
        Pop-Location
    }
}

function Invoke-Build {
    param([string]$OsTarget)

    Ensure-Deps
    New-Item -ItemType Directory -Force -Path $PluginDir | Out-Null

    Push-Location $PluginSrc
    try { go mod tidy } finally { Pop-Location }

    $env:CGO_ENABLED = "1"
    $env:GOOS = $null
    $env:GOARCH = $null
    $env:CC = $null

    switch ($OsTarget) {
        "linux" {
            $env:GOOS = "linux"
            $env:GOARCH = "amd64"
            $out = Join-Path $PluginDir "$PluginName.so"
        }
        "windows" {
            $env:GOOS = "windows"
            $env:GOARCH = "amd64"
            $env:CC = "x86_64-w64-mingw32-gcc"
            $out = Join-Path $PluginDir "$PluginName.dll"
        }
        default {
            if ($IsWindows -or $env:OS -match "Windows") {
                $out = Join-Path $PluginDir "$PluginName.dll"
            } else {
                $out = Join-Path $PluginDir "$PluginName.so"
            }
        }
    }

    Write-Host "Building Safari plugin -> $out"
    Push-Location $PluginSrc
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

    $clientScript = Join-Path $ServerRoot "store\script\main.nut"
    if (Test-Path $clientScript) {
        Write-Host "Client script OK: $clientScript"
    } else {
        Write-Warning "Missing $clientScript — Hydra camera will not work until store/script/main.nut is next to server64.exe"
    }
}

if ($Clean) {
    Remove-Item -Force -ErrorAction SilentlyContinue `
        (Join-Path $PluginDir "goserver04rel64.so"),
        (Join-Path $PluginDir "goserver04rel64.dll"),
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

$out = Invoke-Build -OsTarget $Target

$deploy = -not $NoDeploy
if ($deploy -and $out) {
    Deploy-PluginToServer $out
}

if ($StartServer) {
    Start-VcmpServer
}
