package main

// Events holds optional Go handlers for VC:MP callbacks.
// Register handlers before the server starts (during package init or from init()).
// Unset handlers use permissive defaults (allow actions/messages).
type Events struct {
	OnServerStart func() FilterResult
	OnServerStop  func()

	OnServerFrame func(elapsed float32)

	OnPlayerConnect           func(playerID int)
	OnPlayerDisconnect        func(playerID int, reason DisconnectReason)
	OnPlayerRequestClass      func(playerID int, offset int) FilterResult
	OnPlayerRequestSpawn      func(playerID int) FilterResult
	OnPlayerSpawn             func(playerID int)
	OnPlayerDeath             func(playerID, killerID int, reason int, bodyPart int)
	OnPlayerUpdate            func(playerID int, updateType int)
	OnPlayerMessage           func(playerID int, message string) FilterResult
	OnPlayerCommand           func(playerID int, command string) FilterResult
	OnPlayerPrivateMessage    func(playerID, targetID int, message string) FilterResult
	OnPlayerRequestEnterVehicle func(playerID, vehicleID, slot int) FilterResult
	OnPlayerEnterVehicle      func(playerID, vehicleID, slot int)
	OnPlayerExitVehicle       func(playerID, vehicleID int)
	OnPlayerKeyBindDown       func(playerID, bindID int)
	OnPlayerKeyBindUp         func(playerID, bindID int)

	OnVehicleUpdate  func(vehicleID int, updateType int)
	OnVehicleExplode func(vehicleID int)
	OnVehicleRespawn func(vehicleID int)

	OnObjectShot    func(objectID, playerID, weaponID int)
	OnObjectTouched func(objectID, playerID int)

	OnPickupPickAttempt func(pickupID, playerID int) FilterResult
	OnPickupPicked      func(pickupID, playerID int)
	OnPickupRespawn     func(pickupID int)

	OnCheckpointEntered func(checkpointID, playerID int)
	OnCheckpointExited  func(checkpointID, playerID int)

	OnEntityPoolChange func(pool EntityPool, entityID int, deleted bool)
}

var events Events
