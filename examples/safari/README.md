# Safari plugin entry (native binary)

This is the **only** package that compiles to `goserver04rel64.dll` / `.so`.

| File | Role |
|------|------|
| `main.go` | `VcmpPluginInit` via `vcmp.OnLoad` |
| `wiring.go` | `vcmp.Events` → `safari.Engine` |
| `plugin.go` | Plugin lifecycle, DB, engine bootstrap |
| `config.go` | `goserver.json` host settings |

Gamemode rules (teams, Hydra, commands) live in **`../../../vcmp-go-server/safari/`** — imported as a Go library, not compiled separately.

## Build

```powershell
cd D:\vcmp-go-plugin
.\build-safari.ps1           # test + stop server + build + deploy
.\build-safari.ps1 -StartServer
```

Or from `vcmp-go-server`: `.\build.ps1`

Do **not** copy `goplugin04rel64.dll` (blank example) into the Safari server — `server.cfg` must load `goserver04rel64`.
