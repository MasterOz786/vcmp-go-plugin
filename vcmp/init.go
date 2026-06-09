package vcmp

/*
#cgo CFLAGS: -I${SRCDIR}/../include
#include "plugin.h"
#include <stdlib.h>
#include <string.h>
*/
import "C"

import (
	"reflect"
	"unsafe"
)

// PluginMeta is reported to the VC:MP server in PluginInfo.
type PluginMeta struct {
	Name    string
	Version uint32
}

// Init binds native APIs, fills plugin metadata, and registers all SDK callbacks.
// Returns false when required pointers are nil.
func Init(funcs *C.PluginFuncs, calls *C.PluginCallbacks, info *C.PluginInfo, meta PluginMeta) bool {
	if funcs == nil || calls == nil || info == nil {
		return false
	}
	Bind(funcs)
	setPluginInfo(info, meta)
	RegisterCallbacks(calls)
	return true
}

func setPluginInfo(info *C.PluginInfo, meta PluginMeta) {
	info.structSize = C.uint32_t(C.sizeof_PluginInfo)
	name := cString(meta.Name)
	defer freeCString(name)
	C.strncpy(&info.name[0], name, 31)
	info.pluginVersion = C.uint32_t(meta.Version)
	info.apiMajorVersion = C.PLUGIN_API_MAJOR
	info.apiMinorVersion = C.PLUGIN_API_MINOR
}

// RegisterCallbacks wires all //export handlers into the VC:MP callback table.
func RegisterCallbacks(calls *C.PluginCallbacks) {
	calls.structSize = C.uint32_t(C.sizeof_PluginCallbacks)

	setCallback(&calls.OnServerInitialise, OnServerInitialise)
	setCallback(&calls.OnServerShutdown, OnServerShutdown)
	setCallback(&calls.OnServerFrame, OnServerFrame)
	setCallback(&calls.OnPluginCommand, OnPluginCommand)
	setCallback(&calls.OnIncomingConnection, OnIncomingConnection)
	setCallback(&calls.OnClientScriptData, OnClientScriptData)

	setCallback(&calls.OnPlayerConnect, OnPlayerConnect)
	setCallback(&calls.OnPlayerDisconnect, OnPlayerDisconnect)
	setCallback(&calls.OnPlayerRequestClass, OnPlayerRequestClass)
	setCallback(&calls.OnPlayerRequestSpawn, OnPlayerRequestSpawn)
	setCallback(&calls.OnPlayerSpawn, OnPlayerSpawn)
	setCallback(&calls.OnPlayerDeath, OnPlayerDeath)
	setCallback(&calls.OnPlayerUpdate, OnPlayerUpdate)

	setCallback(&calls.OnPlayerRequestEnterVehicle, OnPlayerRequestEnterVehicle)
	setCallback(&calls.OnPlayerEnterVehicle, OnPlayerEnterVehicle)
	setCallback(&calls.OnPlayerExitVehicle, OnPlayerExitVehicle)

	setCallback(&calls.OnPlayerNameChange, OnPlayerNameChange)
	setCallback(&calls.OnPlayerStateChange, OnPlayerStateChange)
	setCallback(&calls.OnPlayerActionChange, OnPlayerActionChange)
	setCallback(&calls.OnPlayerOnFireChange, OnPlayerOnFireChange)
	setCallback(&calls.OnPlayerCrouchChange, OnPlayerCrouchChange)
	setCallback(&calls.OnPlayerGameKeysChange, OnPlayerGameKeysChange)
	setCallback(&calls.OnPlayerBeginTyping, OnPlayerBeginTyping)
	setCallback(&calls.OnPlayerEndTyping, OnPlayerEndTyping)
	setCallback(&calls.OnPlayerAwayChange, OnPlayerAwayChange)

	setCallback(&calls.OnPlayerMessage, OnPlayerMessage)
	setCallback(&calls.OnPlayerCommand, OnPlayerCommand)
	setCallback(&calls.OnPlayerPrivateMessage, OnPlayerPrivateMessage)

	setCallback(&calls.OnPlayerKeyBindDown, OnPlayerKeyBindDown)
	setCallback(&calls.OnPlayerKeyBindUp, OnPlayerKeyBindUp)
	setCallback(&calls.OnPlayerSpectate, OnPlayerSpectate)
	setCallback(&calls.OnPlayerCrashReport, OnPlayerCrashReport)
	setCallback(&calls.OnPlayerModuleList, OnPlayerModuleList)

	setCallback(&calls.OnVehicleUpdate, OnVehicleUpdate)
	setCallback(&calls.OnVehicleExplode, OnVehicleExplode)
	setCallback(&calls.OnVehicleRespawn, OnVehicleRespawn)

	setCallback(&calls.OnObjectShot, OnObjectShot)
	setCallback(&calls.OnObjectTouched, OnObjectTouched)

	setCallback(&calls.OnPickupPickAttempt, OnPickupPickAttempt)
	setCallback(&calls.OnPickupPicked, OnPickupPicked)
	setCallback(&calls.OnPickupRespawn, OnPickupRespawn)

	setCallback(&calls.OnCheckpointEntered, OnCheckpointEntered)
	setCallback(&calls.OnCheckpointExited, OnCheckpointExited)

	setCallback(&calls.OnEntityPoolChange, OnEntityPoolChange)
	setCallback(&calls.OnServerPerformanceReport, OnServerPerformanceReport)
}

func setCallback[T any](dst *T, fn any) {
	*dst = *(*T)(unsafe.Pointer(reflect.ValueOf(fn).Pointer()))
}
