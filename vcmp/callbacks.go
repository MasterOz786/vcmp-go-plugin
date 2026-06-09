package vcmp

/*
#cgo CFLAGS: -I${SRCDIR}/../include
#include "plugin.h"
#include <stdlib.h>
#include <string.h>

static void vcmp_copy_player_name(char *dst, size_t buflen, const char *src) {
	if (dst != NULL && buflen > 0) {
		strncpy(dst, src, buflen - 1);
		dst[buflen - 1] = '\0';
	}
}
*/
import "C"

import "unsafe"

//export OnServerInitialise
func OnServerInitialise() C.uint8_t {
	if Events.OnServerStart != nil {
		return C.uint8_t(Events.OnServerStart())
	}
	return C.uint8_t(FilterAllow)
}

//export OnServerShutdown
func OnServerShutdown() {
	if Events.OnServerStop != nil {
		Events.OnServerStop()
	}
}

//export OnServerFrame
func OnServerFrame(elapsedTime C.float) {
	if Events.OnServerFrame != nil {
		Events.OnServerFrame(float32(elapsedTime))
	}
}

//export OnPluginCommand
func OnPluginCommand(commandIdentifier C.uint32_t, message *C.char) C.uint8_t {
	if Events.OnPluginCommand != nil {
		return C.uint8_t(Events.OnPluginCommand(uint32(commandIdentifier), C.GoString(message)))
	}
	return C.uint8_t(FilterAllow)
}

//export OnIncomingConnection
func OnIncomingConnection(playerName *C.char, nameBufferSize C.size_t, userPassword *C.char, ipAddress *C.char) C.uint8_t {
	if Events.OnIncomingConnection != nil {
		updated := Events.OnIncomingConnection(
			C.GoString(playerName),
			C.GoString(userPassword),
			C.GoString(ipAddress),
		)
		cName := C.CString(updated)
		defer C.free(unsafe.Pointer(cName))
		C.vcmp_copy_player_name(playerName, nameBufferSize, cName)
	}
	return C.uint8_t(FilterAllow)
}

//export OnClientScriptData
func OnClientScriptData(playerID C.int32_t, data *C.uint8_t, size C.size_t) {
	if Events.OnClientScriptData == nil {
		return
	}
	var payload []byte
	if data != nil && size > 0 {
		payload = C.GoBytes(unsafe.Pointer(data), C.int(size))
	}
	Events.OnClientScriptData(int(playerID), payload)
}

//export OnPlayerConnect
func OnPlayerConnect(playerID C.int32_t) {
	if Events.OnPlayerConnect != nil {
		Events.OnPlayerConnect(int(playerID))
	}
}

//export OnPlayerDisconnect
func OnPlayerDisconnect(playerID C.int32_t, reason C.vcmpDisconnectReason) {
	if Events.OnPlayerDisconnect != nil {
		Events.OnPlayerDisconnect(int(playerID), DisconnectReason(reason))
	}
}

//export OnPlayerRequestClass
func OnPlayerRequestClass(playerID C.int32_t, offset C.int32_t) C.uint8_t {
	if Events.OnPlayerRequestClass != nil {
		return C.uint8_t(Events.OnPlayerRequestClass(int(playerID), int(offset)))
	}
	return C.uint8_t(FilterAllow)
}

//export OnPlayerRequestSpawn
func OnPlayerRequestSpawn(playerID C.int32_t) C.uint8_t {
	if Events.OnPlayerRequestSpawn != nil {
		return C.uint8_t(Events.OnPlayerRequestSpawn(int(playerID)))
	}
	return C.uint8_t(FilterAllow)
}

//export OnPlayerSpawn
func OnPlayerSpawn(playerID C.int32_t) {
	if Events.OnPlayerSpawn != nil {
		Events.OnPlayerSpawn(int(playerID))
	}
}

//export OnPlayerDeath
func OnPlayerDeath(playerID, killerID, reason C.int32_t, bodyPart C.vcmpBodyPart) {
	if Events.OnPlayerDeath != nil {
		Events.OnPlayerDeath(int(playerID), int(killerID), int(reason), BodyPart(bodyPart))
	}
}

//export OnPlayerUpdate
func OnPlayerUpdate(playerID C.int32_t, updateType C.vcmpPlayerUpdate) {
	if Events.OnPlayerUpdate != nil {
		Events.OnPlayerUpdate(int(playerID), PlayerUpdate(updateType))
	}
}

//export OnPlayerRequestEnterVehicle
func OnPlayerRequestEnterVehicle(playerID, vehicleID, slotIndex C.int32_t) C.uint8_t {
	if Events.OnPlayerRequestEnterVehicle != nil {
		return C.uint8_t(Events.OnPlayerRequestEnterVehicle(int(playerID), int(vehicleID), int(slotIndex)))
	}
	return C.uint8_t(FilterAllow)
}

//export OnPlayerEnterVehicle
func OnPlayerEnterVehicle(playerID, vehicleID, slot C.int32_t) {
	if Events.OnPlayerEnterVehicle != nil {
		Events.OnPlayerEnterVehicle(int(playerID), int(vehicleID), int(slot))
	}
}

//export OnPlayerExitVehicle
func OnPlayerExitVehicle(playerID, vehicleID C.int32_t) {
	if Events.OnPlayerExitVehicle != nil {
		Events.OnPlayerExitVehicle(int(playerID), int(vehicleID))
	}
}

//export OnPlayerNameChange
func OnPlayerNameChange(playerID C.int32_t, oldName, newName *C.char) {
	if Events.OnPlayerNameChange != nil {
		Events.OnPlayerNameChange(int(playerID), C.GoString(oldName), C.GoString(newName))
	}
}

//export OnPlayerStateChange
func OnPlayerStateChange(playerID C.int32_t, oldState, newState C.vcmpPlayerState) {
	if Events.OnPlayerStateChange != nil {
		Events.OnPlayerStateChange(int(playerID), PlayerState(oldState), PlayerState(newState))
	}
}

//export OnPlayerActionChange
func OnPlayerActionChange(playerID C.int32_t, oldAction, newAction C.int32_t) {
	if Events.OnPlayerActionChange != nil {
		Events.OnPlayerActionChange(int(playerID), int(oldAction), int(newAction))
	}
}

//export OnPlayerOnFireChange
func OnPlayerOnFireChange(playerID C.int32_t, isOnFire C.uint8_t) {
	if Events.OnPlayerOnFireChange != nil {
		Events.OnPlayerOnFireChange(int(playerID), isOnFire != 0)
	}
}

//export OnPlayerCrouchChange
func OnPlayerCrouchChange(playerID C.int32_t, isCrouching C.uint8_t) {
	if Events.OnPlayerCrouchChange != nil {
		Events.OnPlayerCrouchChange(int(playerID), isCrouching != 0)
	}
}

//export OnPlayerGameKeysChange
func OnPlayerGameKeysChange(playerID C.int32_t, oldKeys, newKeys C.uint32_t) {
	if Events.OnPlayerGameKeysChange != nil {
		Events.OnPlayerGameKeysChange(int(playerID), uint32(oldKeys), uint32(newKeys))
	}
}

//export OnPlayerBeginTyping
func OnPlayerBeginTyping(playerID C.int32_t) {
	if Events.OnPlayerBeginTyping != nil {
		Events.OnPlayerBeginTyping(int(playerID))
	}
}

//export OnPlayerEndTyping
func OnPlayerEndTyping(playerID C.int32_t) {
	if Events.OnPlayerEndTyping != nil {
		Events.OnPlayerEndTyping(int(playerID))
	}
}

//export OnPlayerAwayChange
func OnPlayerAwayChange(playerID C.int32_t, isAway C.uint8_t) {
	if Events.OnPlayerAwayChange != nil {
		Events.OnPlayerAwayChange(int(playerID), isAway != 0)
	}
}

//export OnPlayerMessage
func OnPlayerMessage(playerID C.int32_t, message *C.char) C.uint8_t {
	if Events.OnPlayerMessage != nil {
		return C.uint8_t(Events.OnPlayerMessage(int(playerID), C.GoString(message)))
	}
	return C.uint8_t(FilterAllow)
}

//export OnPlayerCommand
func OnPlayerCommand(playerID C.int32_t, message *C.char) C.uint8_t {
	if Events.OnPlayerCommand != nil {
		return C.uint8_t(Events.OnPlayerCommand(int(playerID), C.GoString(message)))
	}
	return C.uint8_t(FilterAllow)
}

//export OnPlayerPrivateMessage
func OnPlayerPrivateMessage(playerID, targetPlayerID C.int32_t, message *C.char) C.uint8_t {
	if Events.OnPlayerPrivateMessage != nil {
		return C.uint8_t(Events.OnPlayerPrivateMessage(int(playerID), int(targetPlayerID), C.GoString(message)))
	}
	return C.uint8_t(FilterAllow)
}

//export OnPlayerKeyBindDown
func OnPlayerKeyBindDown(playerID, bindID C.int32_t) {
	if Events.OnPlayerKeyBindDown != nil {
		Events.OnPlayerKeyBindDown(int(playerID), int(bindID))
	}
}

//export OnPlayerKeyBindUp
func OnPlayerKeyBindUp(playerID, bindID C.int32_t) {
	if Events.OnPlayerKeyBindUp != nil {
		Events.OnPlayerKeyBindUp(int(playerID), int(bindID))
	}
}

//export OnPlayerSpectate
func OnPlayerSpectate(playerID, targetPlayerID C.int32_t) {
	if Events.OnPlayerSpectate != nil {
		Events.OnPlayerSpectate(int(playerID), int(targetPlayerID))
	}
}

//export OnPlayerCrashReport
func OnPlayerCrashReport(playerID C.int32_t, report *C.char) {
	if Events.OnPlayerCrashReport != nil {
		Events.OnPlayerCrashReport(int(playerID), C.GoString(report))
	}
}

//export OnPlayerModuleList
func OnPlayerModuleList(playerID C.int32_t, list *C.char) {
	if Events.OnPlayerModuleList != nil {
		Events.OnPlayerModuleList(int(playerID), C.GoString(list))
	}
}

//export OnVehicleUpdate
func OnVehicleUpdate(vehicleID C.int32_t, updateType C.vcmpVehicleUpdate) {
	if Events.OnVehicleUpdate != nil {
		Events.OnVehicleUpdate(int(vehicleID), VehicleUpdate(updateType))
	}
}

//export OnVehicleExplode
func OnVehicleExplode(vehicleID C.int32_t) {
	if Events.OnVehicleExplode != nil {
		Events.OnVehicleExplode(int(vehicleID))
	}
}

//export OnVehicleRespawn
func OnVehicleRespawn(vehicleID C.int32_t) {
	if Events.OnVehicleRespawn != nil {
		Events.OnVehicleRespawn(int(vehicleID))
	}
}

//export OnObjectShot
func OnObjectShot(objectID, playerID, weaponID C.int32_t) {
	if Events.OnObjectShot != nil {
		Events.OnObjectShot(int(objectID), int(playerID), int(weaponID))
	}
}

//export OnObjectTouched
func OnObjectTouched(objectID, playerID C.int32_t) {
	if Events.OnObjectTouched != nil {
		Events.OnObjectTouched(int(objectID), int(playerID))
	}
}

//export OnPickupPickAttempt
func OnPickupPickAttempt(pickupID, playerID C.int32_t) C.uint8_t {
	if Events.OnPickupPickAttempt != nil {
		return C.uint8_t(Events.OnPickupPickAttempt(int(pickupID), int(playerID)))
	}
	return C.uint8_t(FilterAllow)
}

//export OnPickupPicked
func OnPickupPicked(pickupID, playerID C.int32_t) {
	if Events.OnPickupPicked != nil {
		Events.OnPickupPicked(int(pickupID), int(playerID))
	}
}

//export OnPickupRespawn
func OnPickupRespawn(pickupID C.int32_t) {
	if Events.OnPickupRespawn != nil {
		Events.OnPickupRespawn(int(pickupID))
	}
}

//export OnCheckpointEntered
func OnCheckpointEntered(checkPointID, playerID C.int32_t) {
	if Events.OnCheckpointEntered != nil {
		Events.OnCheckpointEntered(int(checkPointID), int(playerID))
	}
}

//export OnCheckpointExited
func OnCheckpointExited(checkPointID, playerID C.int32_t) {
	if Events.OnCheckpointExited != nil {
		Events.OnCheckpointExited(int(checkPointID), int(playerID))
	}
}

//export OnEntityPoolChange
func OnEntityPoolChange(entityType C.vcmpEntityPool, entityID C.int32_t, isDeleted C.uint8_t) {
	if Events.OnEntityPoolChange != nil {
		Events.OnEntityPoolChange(EntityPool(entityType), int(entityID), isDeleted != 0)
	}
}

//export OnServerPerformanceReport
func OnServerPerformanceReport(entryCount C.size_t, descriptions **C.char, times *C.uint64_t) {
	if Events.OnServerPerformanceReport == nil {
		return
	}
	n := int(entryCount)
	if n == 0 {
		Events.OnServerPerformanceReport(nil, nil)
		return
	}
	descs := cStringSlice(descriptions, n)
	ts := cUint64Slice(times, n)
	Events.OnServerPerformanceReport(descs, ts)
}

func cStringSlice(pp **C.char, n int) []string {
	if pp == nil || n == 0 {
		return nil
	}
	hdr := (*[1 << 30]*C.char)(unsafe.Pointer(pp))[:n:n]
	out := make([]string, n)
	for i, p := range hdr {
		if p != nil {
			out[i] = C.GoString(p)
		}
	}
	return out
}

func cUint64Slice(p *C.uint64_t, n int) []uint64 {
	if p == nil || n == 0 {
		return nil
	}
	hdr := (*[1 << 30]C.uint64_t)(unsafe.Pointer(p))[:n:n]
	out := make([]uint64, n)
	for i, v := range hdr {
		out[i] = uint64(v)
	}
	return out
}
