package main

/*
#cgo CFLAGS: -I${SRCDIR}/include
#include "plugin.h"
*/
import "C"

//export OnServerInitialise
func OnServerInitialise() C.uint8_t {
	plug.initialise()
	if events.OnServerStart != nil {
		return C.uint8_t(events.OnServerStart())
	}
	return C.uint8_t(FilterAllow)
}

//export OnServerShutdown
func OnServerShutdown() {
	if events.OnServerStop != nil {
		events.OnServerStop()
	}
	plug.shutdown()
}

//export OnServerFrame
func OnServerFrame(elapsed C.float) {
	if events.OnServerFrame != nil {
		events.OnServerFrame(float32(elapsed))
	}
}

//export OnPlayerConnect
func OnPlayerConnect(playerID C.int32_t) {
	if events.OnPlayerConnect != nil {
		events.OnPlayerConnect(int(playerID))
	}
}

//export OnPlayerDisconnect
func OnPlayerDisconnect(playerID C.int32_t, reason C.vcmpDisconnectReason) {
	if events.OnPlayerDisconnect != nil {
		events.OnPlayerDisconnect(int(playerID), DisconnectReason(reason))
	}
}

//export OnPlayerRequestClass
func OnPlayerRequestClass(playerID C.int32_t, offset C.int32_t) C.uint8_t {
	if events.OnPlayerRequestClass != nil {
		return C.uint8_t(events.OnPlayerRequestClass(int(playerID), int(offset)))
	}
	return C.uint8_t(FilterAllow)
}

//export OnPlayerRequestSpawn
func OnPlayerRequestSpawn(playerID C.int32_t) C.uint8_t {
	if events.OnPlayerRequestSpawn != nil {
		return C.uint8_t(events.OnPlayerRequestSpawn(int(playerID)))
	}
	return C.uint8_t(FilterAllow)
}

//export OnPlayerSpawn
func OnPlayerSpawn(playerID C.int32_t) {
	if events.OnPlayerSpawn != nil {
		events.OnPlayerSpawn(int(playerID))
	}
}

//export OnPlayerDeath
func OnPlayerDeath(playerID, killerID C.int32_t, reason C.int32_t, bodyPart C.vcmpBodyPart) {
	if events.OnPlayerDeath != nil {
		events.OnPlayerDeath(int(playerID), int(killerID), int(reason), int(bodyPart))
	}
}

//export OnPlayerUpdate
func OnPlayerUpdate(playerID C.int32_t, updateType C.vcmpPlayerUpdate) {
	if events.OnPlayerUpdate != nil {
		events.OnPlayerUpdate(int(playerID), int(updateType))
	}
}

//export OnPlayerMessage
func OnPlayerMessage(playerID C.int32_t, message *C.char) C.uint8_t {
	if events.OnPlayerMessage != nil {
		return C.uint8_t(events.OnPlayerMessage(int(playerID), C.GoString(message)))
	}
	return C.uint8_t(FilterAllow)
}

//export OnPlayerCommand
func OnPlayerCommand(playerID C.int32_t, message *C.char) C.uint8_t {
	if events.OnPlayerCommand != nil {
		return C.uint8_t(events.OnPlayerCommand(int(playerID), C.GoString(message)))
	}
	return C.uint8_t(FilterAllow)
}

//export OnPlayerPrivateMessage
func OnPlayerPrivateMessage(playerID, targetID C.int32_t, message *C.char) C.uint8_t {
	if events.OnPlayerPrivateMessage != nil {
		return C.uint8_t(events.OnPlayerPrivateMessage(int(playerID), int(targetID), C.GoString(message)))
	}
	return C.uint8_t(FilterAllow)
}

//export OnPlayerRequestEnterVehicle
func OnPlayerRequestEnterVehicle(playerID, vehicleID, slot C.int32_t) C.uint8_t {
	if events.OnPlayerRequestEnterVehicle != nil {
		return C.uint8_t(events.OnPlayerRequestEnterVehicle(int(playerID), int(vehicleID), int(slot)))
	}
	return C.uint8_t(FilterAllow)
}

//export OnPlayerEnterVehicle
func OnPlayerEnterVehicle(playerID, vehicleID, slot C.int32_t) {
	if events.OnPlayerEnterVehicle != nil {
		events.OnPlayerEnterVehicle(int(playerID), int(vehicleID), int(slot))
	}
}

//export OnPlayerExitVehicle
func OnPlayerExitVehicle(playerID, vehicleID C.int32_t) {
	if events.OnPlayerExitVehicle != nil {
		events.OnPlayerExitVehicle(int(playerID), int(vehicleID))
	}
}

//export OnPlayerKeyBindDown
func OnPlayerKeyBindDown(playerID, bindID C.int32_t) {
	if events.OnPlayerKeyBindDown != nil {
		events.OnPlayerKeyBindDown(int(playerID), int(bindID))
	}
}

//export OnPlayerKeyBindUp
func OnPlayerKeyBindUp(playerID, bindID C.int32_t) {
	if events.OnPlayerKeyBindUp != nil {
		events.OnPlayerKeyBindUp(int(playerID), int(bindID))
	}
}

//export OnVehicleUpdate
func OnVehicleUpdate(vehicleID C.int32_t, updateType C.vcmpVehicleUpdate) {
	if events.OnVehicleUpdate != nil {
		events.OnVehicleUpdate(int(vehicleID), int(updateType))
	}
}

//export OnVehicleExplode
func OnVehicleExplode(vehicleID C.int32_t) {
	if events.OnVehicleExplode != nil {
		events.OnVehicleExplode(int(vehicleID))
	}
}

//export OnVehicleRespawn
func OnVehicleRespawn(vehicleID C.int32_t) {
	if events.OnVehicleRespawn != nil {
		events.OnVehicleRespawn(int(vehicleID))
	}
}

//export OnObjectShot
func OnObjectShot(objectID, playerID, weaponID C.int32_t) {
	if events.OnObjectShot != nil {
		events.OnObjectShot(int(objectID), int(playerID), int(weaponID))
	}
}

//export OnObjectTouched
func OnObjectTouched(objectID, playerID C.int32_t) {
	if events.OnObjectTouched != nil {
		events.OnObjectTouched(int(objectID), int(playerID))
	}
}

//export OnPickupPickAttempt
func OnPickupPickAttempt(pickupID, playerID C.int32_t) C.uint8_t {
	if events.OnPickupPickAttempt != nil {
		return C.uint8_t(events.OnPickupPickAttempt(int(pickupID), int(playerID)))
	}
	return C.uint8_t(FilterAllow)
}

//export OnPickupPicked
func OnPickupPicked(pickupID, playerID C.int32_t) {
	if events.OnPickupPicked != nil {
		events.OnPickupPicked(int(pickupID), int(playerID))
	}
}

//export OnPickupRespawn
func OnPickupRespawn(pickupID C.int32_t) {
	if events.OnPickupRespawn != nil {
		events.OnPickupRespawn(int(pickupID))
	}
}

//export OnCheckpointEntered
func OnCheckpointEntered(checkpointID, playerID C.int32_t) {
	if events.OnCheckpointEntered != nil {
		events.OnCheckpointEntered(int(checkpointID), int(playerID))
	}
}

//export OnCheckpointExited
func OnCheckpointExited(checkpointID, playerID C.int32_t) {
	if events.OnCheckpointExited != nil {
		events.OnCheckpointExited(int(checkpointID), int(playerID))
	}
}

//export OnEntityPoolChange
func OnEntityPoolChange(pool C.vcmpEntityPool, entityID C.int32_t, deleted C.uint8_t) {
	if events.OnEntityPoolChange != nil {
		events.OnEntityPoolChange(EntityPool(pool), int(entityID), deleted != 0)
	}
}
