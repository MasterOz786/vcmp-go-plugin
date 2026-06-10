# vcmp-go-plugin

Public Go SDK for [VC:MP](https://vc-mp.com/) 0.4 native server plugins, plus the Safari gamemode native binary.

## Install

```bash
go get github.com/masteroz/vcmp-go-plugin/vcmp
```

Requires CGO and `include/plugin.h` (bundled in this repo, or run `make deps` to fetch the latest from upstream).

## Layout

| Path | Role |
|------|------|
| `vcmp/` | SDK — API wrappers, events, CGO bridge |
| `plugin/` | Safari plugin `main` — builds `goserver04rel64.dll` / `.so` |
| `plugins/` | Build output (gitignored) |

Gamemode rules live in the sibling [`vcmp-go-server`](https://github.com/masteroz/vcmp-go-server) repo (`safari/` library).

## Build Safari plugin

```powershell
cd D:\vcmp-go-plugin
.\build.ps1 -Test -StopServer          # test library, stop server, build, deploy
.\build.ps1 -Test -StopServer -StartServer
```

```bash
make build                 # → plugins/goserver04rel64.so
make build-windows         # → plugins/goserver04rel64.dll
```

Deploy target (default): sibling `../vcmp-go-server/plugins/`. Override with `-ServerRoot`.

## Write your own plugin

Set hooks in `init()` and build with `-buildmode=c-shared`:

```go
func init() {
    vcmp.MetaProvider = func() vcmp.PluginMeta {
        return vcmp.PluginMeta{Name: "MyMode", Version: 0x00010000}
    }
    vcmp.OnLoad = func() {
        // wire vcmp.Events
    }
}
```

Copy the built `.dll` / `.so` into the VC:MP server's `plugins/` folder and add it to `server.cfg`.
