package main

import (
	"fmt"
	"strings"

	"github.com/masteroz/vcmp-go-plugin/vcmp"
)

// register wires vcmp.Events directly to safari.Engine (no event queue).
func (p *Plugin) register() {
	if p.engine == nil {
		return
	}

	vcmp.Events.OnServerStart = func() vcmp.FilterResult {
		p.engine.OnServerStart()
		_ = vcmp.API.KeyBind.Register(1, false, 0x31, 0, 0) // "1" -> pack 1
		_ = vcmp.API.KeyBind.Register(2, false, 0x32, 0, 0) // "2" -> pack 2
		_ = vcmp.API.KeyBind.Register(3, false, 0x33, 0, 0) // "3" -> pack 3
		_ = vcmp.API.KeyBind.Register(4, false, 0x48, 0, 0) // "H" -> hydra camera cycle
		return vcmp.FilterAllow
	}

	vcmp.Events.OnServerStop = func() {
		p.shutdown()
	}

	vcmp.Events.OnServerFrame = func(elapsed float32) {
		p.engine.OnServerFrame(elapsed)
	}

	vcmp.Events.OnPlayerConnect = func(playerID int) {
		p.engine.OnConnect(playerID)
	}

	vcmp.Events.OnPlayerDisconnect = func(playerID int, _ vcmp.DisconnectReason) {
		p.engine.OnDisconnect(playerID)
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
		p.engine.OnSpawn(playerID)
	}

	vcmp.Events.OnPlayerStateChange = func(playerID int, oldState, newState vcmp.PlayerState) {
		if oldState == vcmp.PlayerStateUnspawned && newState == vcmp.PlayerStateNormal {
			p.engine.OnSpawn(playerID)
		}
	}

	vcmp.Events.OnPlayerDeath = func(playerID, killerID int, _ int, _ vcmp.BodyPart) {
		p.engine.OnDeath(playerID, killerID)
	}

	vcmp.Events.OnPlayerCommand = func(playerID int, command string) vcmp.FilterResult {
		if res := p.engine.HandleCommand(playerID, command); res.Deny {
			return vcmp.FilterDeny
		}
		return vcmp.FilterAllow
	}

	vcmp.Events.OnPlayerMessage = func(playerID int, message string) vcmp.FilterResult {
		if strings.HasPrefix(strings.TrimSpace(message), "/") {
			if res := p.engine.HandleCommand(playerID, message); res.Deny {
				return vcmp.FilterDeny
			}
		}
		return vcmp.FilterAllow
	}

	vcmp.Events.OnVehicleExplode = func(vehicleID int) {
		p.engine.OnVehicleExplode(vehicleID)
	}

	vcmp.Events.OnPickupPickAttempt = func(pickupID, playerID int) vcmp.FilterResult {
		if p.engine.HandlePickupPickAttempt(pickupID, playerID) {
			return vcmp.FilterAllow
		}
		return vcmp.FilterDeny
	}

	vcmp.Events.OnPickupPicked = func(pickupID, playerID int) {
		p.engine.OnPickupPicked(pickupID, playerID)
	}

	vcmp.Events.OnPlayerKeyBindDown = func(playerID, bindID int) {
		p.engine.OnPlayerKeyBind(playerID, bindID, false)
	}

	vcmp.Events.OnPlayerRequestEnterVehicle = func(playerID, vehicleID, slot int) vcmp.FilterResult {
		if p.engine.HandleEnterVehicleRequest(playerID, vehicleID, slot) {
			return vcmp.FilterAllow
		}
		return vcmp.FilterDeny
	}

	vcmp.Events.OnPlayerEnterVehicle = func(playerID, vehicleID, slot int) {
		p.engine.OnPlayerEnterVehicle(playerID, vehicleID, slot)
	}

	vcmp.Events.OnPlayerExitVehicle = func(playerID, vehicleID int) {
		p.engine.OnPlayerExitVehicle(playerID, vehicleID)
	}

	vcmp.Events.OnIncomingConnection = func(name, password, ip string) string {
		_, _, _ = password, ip, name
		return name
	}

	vcmp.Events.OnPlayerCrashReport = func(playerID int, report string) {
		vcmp.API.Server.Log(fmt.Sprintf("[safari] crash report from %d: %s", playerID, report))
	}

	vcmp.Events.OnClientScriptData = func(playerID int, data []byte) {
		p.engine.HandleClientScriptData(playerID, data)
	}
}
