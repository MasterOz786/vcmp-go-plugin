package vcmp

/*
#cgo CFLAGS: -I${SRCDIR}/../include
#include "plugin.h"
*/
import "C"

import "fmt"

// OnLoad runs after Init succeeds. Gamemodes register event handlers here.
var OnLoad func()

// MetaProvider supplies plugin name/version shown in VC:MP. Required for custom plugins.
var MetaProvider func() PluginMeta

func defaultMeta() PluginMeta {
	return PluginMeta{Name: "GoSDK", Version: 0x00020000}
}

func formatPluginVersion(version uint32) string {
	return fmt.Sprintf("%d.%d.%d", (version>>16)&0xFF, (version>>8)&0xFF, version&0xFF)
}

func logPluginLoaded(meta PluginMeta) {
	bridgeLog(fmt.Sprintf(
		"[plugin] loaded %s v%s (API %d.%d)",
		meta.Name,
		formatPluginVersion(meta.Version),
		int(C.PLUGIN_API_MAJOR),
		int(C.PLUGIN_API_MINOR),
	))
}

//export VcmpPluginInit
func VcmpPluginInit(
	funcs *C.PluginFuncs,
	calls *C.PluginCallbacks,
	info *C.PluginInfo,
) C.uint {
	meta := defaultMeta()
	if MetaProvider != nil {
		meta = MetaProvider()
	}
	if !Init(funcs, calls, info, meta) {
		return 0
	}
	if OnLoad != nil {
		OnLoad()
	}
	logPluginLoaded(meta)
	return 1
}
