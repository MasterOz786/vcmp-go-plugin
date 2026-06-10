package main

import (
	"fmt"
	"strings"

	"github.com/masteroz/vcmp-go-plugin/vcmp"
	"github.com/masteroz/vcmp-go-server/safari"
)

func (p *Plugin) register() {
	if p.engine == nil {
		return
	}

	vcmp.Events.OnServerStart = func() vcmp.FilterResult {
		p.engine.OnServerStart()
		_ = vcmp.API.KeyBind.Register(1, false, 0x31, 0, 0) // "1" -> pack 1
		_ = vcmp.API.KeyBind.Register(2, false, 0x32, 0, 0) // "2" -> pack 2
		return vcmp.FilterAllow
	}

	vcmp.Events.OnServerStop = func() {
		p.shutdown()
	}

	vcmp.Events.OnServerFrame = func(_ float32) {
		p.engine.OnServerFrame()
	}

	vcmp.Events.OnPlayerConnect = func(playerID int) {
		p.engine.Enqueue(safari.NewConnectEvent(playerID))
	}

	vcmp.Events.OnPlayerDisconnect = func(playerID int, _ vcmp.DisconnectReason) {
		p.engine.Enqueue(safari.NewDisconnectEvent(playerID))
	}

	vcmp.Events.OnPlayerRequestSpawn = func(playerID int) vcmp.FilterResult {
		if p.engine.HandleRequestSpawn(playerID) {
			return vcmp.FilterAllow
		}
		return vcmp.FilterDeny
	}

	vcmp.Events.OnPlayerRequestClass = func(playerID int, classIndex int) vcmp.FilterResult {
		if p.engine.HandleRequestClass(playerID, classIndex) {
			return vcmp.FilterAllow
		}
		return vcmp.FilterDeny
	}

	vcmp.Events.OnPlayerSpawn = func(playerID int) {
		p.engine.SchedulePlayerLoadout(playerID)
		p.engine.Enqueue(safari.NewSpawnEvent(playerID))
	}

	vcmp.Events.OnPlayerStateChange = func(playerID int, oldState, newState vcmp.PlayerState) {
		if oldState == vcmp.PlayerStateUnspawned && newState == vcmp.PlayerStateNormal {
			p.engine.SchedulePlayerLoadout(playerID)
			p.engine.Enqueue(safari.NewSpawnEvent(playerID))
		}
	}

	vcmp.Events.OnPlayerDeath = func(playerID, killerID int, _ int, _ vcmp.BodyPart) {
		p.engine.Enqueue(safari.NewDeathEvent(playerID, killerID))
	}

	vcmp.Events.OnPlayerCommand = func(playerID int, command string) vcmp.FilterResult {
		if p.engine.HandleCommandSync(playerID, command) {
			return vcmp.FilterDeny
		}
		return vcmp.FilterAllow
	}

	vcmp.Events.OnPlayerMessage = func(playerID int, message string) vcmp.FilterResult {
		if strings.HasPrefix(strings.TrimSpace(message), "/") {
			if p.engine.HandleCommandSync(playerID, message) {
				return vcmp.FilterDeny
			}
		}
		return vcmp.FilterAllow
	}

	vcmp.Events.OnVehicleExplode = func(vehicleID int) {
		p.engine.Enqueue(safari.NewVehicleExplodeEvent(vehicleID))
	}

	vcmp.Events.OnVehicleUpdate = func(vehicleID int, updateType vcmp.VehicleUpdate) {
		p.engine.Enqueue(safari.NewVehicleUpdateEvent(vehicleID, int(updateType)))
	}

	vcmp.Events.OnVehicleRespawn = func(vehicleID int) {
		p.engine.Enqueue(safari.NewVehicleRespawnEvent(vehicleID))
	}

	vcmp.Events.OnPickupPickAttempt = func(pickupID, playerID int) vcmp.FilterResult {
		if p.engine.HandlePickupPickAttemptSync(pickupID, playerID) {
			return vcmp.FilterAllow
		}
		return vcmp.FilterDeny
	}

	vcmp.Events.OnPickupPicked = func(pickupID, playerID int) {
		p.engine.Enqueue(safari.NewPickupPickedEvent(pickupID, playerID))
	}

	vcmp.Events.OnCheckpointEntered = func(checkpointID, playerID int) {
		p.engine.Enqueue(safari.NewCheckpointEnteredEvent(checkpointID, playerID))
	}

	vcmp.Events.OnCheckpointExited = func(checkpointID, playerID int) {
		p.engine.Enqueue(safari.NewCheckpointExitedEvent(checkpointID, playerID))
	}

	vcmp.Events.OnPlayerKeyBindDown = func(playerID, bindID int) {
		p.engine.Enqueue(safari.NewKeyBindEvent(playerID, bindID, false))
	}

	vcmp.Events.OnPlayerKeyBindUp = func(playerID, bindID int) {
		p.engine.Enqueue(safari.NewKeyBindEvent(playerID, bindID, true))
	}

	vcmp.Events.OnObjectShot = func(objectID, playerID, weaponID int) {
		p.engine.Enqueue(safari.NewObjectShotEvent(objectID, playerID, weaponID))
	}

	vcmp.Events.OnObjectTouched = func(objectID, playerID int) {
		p.engine.Enqueue(safari.NewObjectTouchedEvent(objectID, playerID))
	}

	vcmp.Events.OnPickupRespawn = func(pickupID int) {
		p.engine.Enqueue(safari.NewPickupRespawnEvent(pickupID))
	}

	vcmp.Events.OnEntityPoolChange = func(entityType vcmp.EntityPool, entityID int, isDeleted bool) {
		p.engine.Enqueue(safari.NewEntityPoolChangeEvent(int(entityType), entityID, isDeleted))
	}

	vcmp.Events.OnPlayerUpdate = func(playerID int, updateType vcmp.PlayerUpdate) {
		p.engine.Enqueue(safari.NewPlayerUpdateEvent(playerID, int(updateType)))
	}

	vcmp.Events.OnPlayerRequestEnterVehicle = func(playerID, vehicleID, slot int) vcmp.FilterResult {
		if p.engine.HandleEnterVehicleRequestSync(playerID, vehicleID, slot) {
			return vcmp.FilterAllow
		}
		return vcmp.FilterDeny
	}

	vcmp.Events.OnPlayerEnterVehicle = func(playerID, vehicleID, slot int) {
		p.engine.Enqueue(safari.NewPlayerEnterVehicleEvent(playerID, vehicleID, slot))
	}

	vcmp.Events.OnPlayerExitVehicle = func(playerID, vehicleID int) {
		p.engine.Enqueue(safari.NewPlayerExitVehicleEvent(playerID, vehicleID))
	}

	vcmp.Events.OnIncomingConnection = func(name, password, ip string) string {
		_, _, _ = password, ip, name
		return name
	}

	vcmp.Events.OnPlayerNameChange = func(playerID int, oldName, newName string) {
		_, _, _ = playerID, oldName, newName
	}

	vcmp.Events.OnPlayerActionChange = func(playerID, oldAction, newAction int) {
		_, _, _ = playerID, oldAction, newAction
	}

	vcmp.Events.OnPlayerOnFireChange = func(playerID int, isOnFire bool) {
		_, _ = playerID, isOnFire
	}

	vcmp.Events.OnPlayerCrouchChange = func(playerID int, isCrouching bool) {
		_, _ = playerID, isCrouching
	}

	vcmp.Events.OnPlayerGameKeysChange = func(playerID int, oldKeys, newKeys uint32) {
		_, _, _ = playerID, oldKeys, newKeys
	}

	vcmp.Events.OnPlayerBeginTyping = func(playerID int) { _ = playerID }
	vcmp.Events.OnPlayerEndTyping = func(playerID int)   { _ = playerID }
	vcmp.Events.OnPlayerAwayChange = func(playerID int, isAway bool) {
		_, _ = playerID, isAway
	}

	vcmp.Events.OnPlayerPrivateMessage = func(playerID, targetPlayerID int, message string) vcmp.FilterResult {
		_, _, _ = playerID, targetPlayerID, message
		return vcmp.FilterAllow
	}

	vcmp.Events.OnPlayerSpectate = func(playerID, targetPlayerID int) {
		_, _ = playerID, targetPlayerID
	}

	vcmp.Events.OnPlayerCrashReport = func(playerID int, report string) {
		vcmp.API.Server.Log(fmt.Sprintf("[safari] crash report from %d: %s", playerID, report))
	}

	vcmp.Events.OnPluginCommand = func(commandID uint32, message string) vcmp.FilterResult {
		_, _ = commandID, message
		return vcmp.FilterAllow
	}

	vcmp.Events.OnServerPerformanceReport = func(descriptions []string, times []uint64) {
		_, _ = descriptions, times
	}
}
