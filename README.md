# vcmp-go-plugin

Public Go SDK for [VC:MP](https://vc-mp.com/) 0.4 native server plugins.

## Install

```bash
go get github.com/masteroz/vcmp-go-plugin/vcmp
```

Requires CGO and `include/plugin.h` (bundled in this repo, or run `make deps` to fetch the latest from upstream).

## Build a plugin

Your gamemode repo imports this module and sets hooks in `init()`:

```go
func init() {
    vcmp.MetaProvider = func() vcmp.PluginMeta {
        return vcmp.PluginMeta{Name: "MyMode", Version: 0x00010000}
    }
    vcmp.OnLoad = func() {
        // wire vcmp.Events, start gamemode logic
    }
}
```

```go
package main

import "github.com/masteroz/vcmp-go-plugin/vcmp"

func main() {}
```

```bash
CGO_ENABLED=1 go build -buildmode=c-shared -o myplugin04rel64.so .
```

Copy the `.so` next to your VC:MP server and add `plugins myplugin04rel64` to `server.cfg`.

## API surface

- **`vcmp.API`** — typed wrappers (`API.Player.GiveWeapon`, `API.Vehicle.Create`, …)
- **`vcmp.Events`** — register VC:MP callback handlers before players connect
- **`vcmp.Init`** — bind natives, set plugin info, register all callbacks

See [`examples/blank`](examples/blank) for a minimal plugin template.

## Blank example

```bash
make build
# or: cd examples/blank && CGO_ENABLED=1 go build -buildmode=c-shared -o ../../goplugin04rel64.so .
```
