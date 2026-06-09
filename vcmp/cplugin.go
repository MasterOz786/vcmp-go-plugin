package vcmp

/*
#cgo CFLAGS: -I${SRCDIR}/../include
#include "plugin.h"
*/
import "C"

// OnLoad runs after Init succeeds. Gamemodes register event handlers here.
var OnLoad func()

// MetaProvider supplies plugin name/version shown in VC:MP. Required for custom plugins.
var MetaProvider func() PluginMeta

func defaultMeta() PluginMeta {
	return PluginMeta{Name: "GoSDK", Version: 0x00020000}
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
	return 1
}
