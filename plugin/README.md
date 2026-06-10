# Safari plugin entry (native binary)

This package compiles to `goserver04rel64.dll` / `.so`.

| File | Role |
|------|------|
| `main.go` | `VcmpPluginInit` via `vcmp.OnLoad` |
| `wiring.go` | `vcmp.Events` → `safari.Engine` |
| `plugin.go` | Plugin lifecycle, DB, engine bootstrap |
| `config.go` | `goserver.json` host settings |

Gamemode rules live in **`vcmp-go-server/safari/`** — imported as a Go library.

## Build

```powershell
cd D:\vcmp-go-plugin
.\build.ps1 -Test -StopServer

# Or from vcmp-go-server:
cd D:\vcmp-go-server
.\build.ps1
```

`server.cfg` must load `goserver04rel64`.
