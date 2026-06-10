param(
    [ValidateSet("blank", "safari")]
    [string]$Example = "blank",
    [ValidateSet("native", "linux", "windows")]
    [string]$Target = "windows",
    [switch]$All,
    [switch]$Deps,
    [switch]$Clean
)

$ErrorActionPreference = "Stop"
$Root = $PSScriptRoot
$PluginDir = Join-Path $Root "plugins"
$Header = Join-Path $Root "include\plugin.h"

$PluginNames = @{
    blank  = "goplugin04rel64"
    safari = "goserver04rel64"
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

if ($All) {
    Invoke-Build -Ex "blank" -OsTarget $Target
    Invoke-Build -Ex "safari" -OsTarget $Target
    exit 0
}

Invoke-Build -Ex $Example -OsTarget $Target
