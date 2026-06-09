package main

/*
#cgo CFLAGS: -I${SRCDIR}/include
#include "plugin.h"
#include <stdlib.h>
#include <string.h>
*/
import "C"

import (
	"reflect"
	"unsafe"
)

var plug *Plugin

//export VcmpPluginInit
func VcmpPluginInit(
	pluginFuncsIn *C.PluginFuncs,
	pluginCalls *C.PluginCallbacks,
	pluginInfo *C.PluginInfo,
) C.uint {
	if pluginFuncsIn == nil || pluginCalls == nil || pluginInfo == nil {
		return 0
	}

	bindPluginAPI(pluginFuncsIn)
	plug = newPlugin(loadConfig())

	pluginInfo.structSize = C.uint32_t(C.sizeof_PluginInfo)
	name := C.CString(PluginName)
	defer C.free(unsafe.Pointer(name))
	C.strncpy(&pluginInfo.name[0], name, 31)
	pluginInfo.pluginVersion = 0x00020000
	pluginInfo.apiMajorVersion = C.PLUGIN_API_MAJOR
	pluginInfo.apiMinorVersion = C.PLUGIN_API_MINOR

	pluginCalls.structSize = C.uint32_t(C.sizeof_PluginCallbacks)
	registerCallbacks(pluginCalls)

	return 1
}

func registerCallbacks(calls *C.PluginCallbacks) {
	setCallback(&calls.OnServerInitialise, OnServerInitialise)
	setCallback(&calls.OnServerShutdown, OnServerShutdown)
	setCallback(&calls.OnServerFrame, OnServerFrame)
	setCallback(&calls.OnPlayerConnect, OnPlayerConnect)
	setCallback(&calls.OnPlayerDisconnect, OnPlayerDisconnect)
	setCallback(&calls.OnPlayerRequestClass, OnPlayerRequestClass)
	setCallback(&calls.OnPlayerRequestSpawn, OnPlayerRequestSpawn)
	setCallback(&calls.OnPlayerSpawn, OnPlayerSpawn)
	setCallback(&calls.OnPlayerDeath, OnPlayerDeath)
	setCallback(&calls.OnPlayerUpdate, OnPlayerUpdate)
	setCallback(&calls.OnPlayerMessage, OnPlayerMessage)
	setCallback(&calls.OnPlayerCommand, OnPlayerCommand)
	setCallback(&calls.OnPlayerPrivateMessage, OnPlayerPrivateMessage)
	setCallback(&calls.OnPlayerRequestEnterVehicle, OnPlayerRequestEnterVehicle)
	setCallback(&calls.OnPlayerEnterVehicle, OnPlayerEnterVehicle)
	setCallback(&calls.OnPlayerExitVehicle, OnPlayerExitVehicle)
	setCallback(&calls.OnPlayerKeyBindDown, OnPlayerKeyBindDown)
	setCallback(&calls.OnPlayerKeyBindUp, OnPlayerKeyBindUp)
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
}

func setCallback[T any](dst *T, fn any) {
	*dst = *(*T)(unsafe.Pointer(reflect.ValueOf(fn).Pointer()))
}

func main() {}
