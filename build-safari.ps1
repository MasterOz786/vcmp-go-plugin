# Full Safari pipeline: test library -> stop server -> build DLL -> deploy to vcmp-go-server.
param(
    [switch]$StartServer,
    [switch]$NoTest,
    [switch]$NoStop
)

$params = @{
    Example        = "safari"
    Target         = "windows"
    DeployToServer = $true
    StopServer     = -not $NoStop
    Test           = -not $NoTest
    StartServer    = $StartServer
}

& "$PSScriptRoot\build.ps1" @params
