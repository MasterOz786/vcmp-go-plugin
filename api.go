package main

import "fmt"

// API is the root namespace for typed VC:MP SDK calls.
var API = struct {
	Server     ServerAPI
	Player     PlayerAPI
	Vehicle    VehicleAPI
	Object     ObjectAPI
	Pickup     PickupAPI
	Checkpoint CheckpointAPI
	Blip       BlipAPI
	KeyBind    KeyBindAPI
	World      WorldAPI
}{
	Server:     ServerAPI{},
	Player:     PlayerAPI{},
	Vehicle:    VehicleAPI{},
	Object:     ObjectAPI{},
	Pickup:     PickupAPI{},
	Checkpoint: CheckpointAPI{},
	Blip:       BlipAPI{},
	KeyBind:    KeyBindAPI{},
	World:      WorldAPI{},
}

type ServerAPI struct{}

func (ServerAPI) Log(msg string)                         { bridgeLog(msg) }
func (ServerAPI) Time() uint64                           { return bridgeTime() }
func (ServerAPI) SetName(name string)                    { bridgeSetServerName(name) }
func (ServerAPI) SetGameModeText(text string)            { bridgeSetGameModeText(text) }
func (ServerAPI) SetOption(option ServerOption, on bool) { bridgeSetServerOption(option, on) }
func (ServerAPI) SetSpawnPosition(pos Vec3)              { bridgeSetSpawnPos(pos) }

func (ServerAPI) AddPlayerClass(teamID int, colour uint32, model int, pos Vec3, angle float32, weapons ...int) int {
	w := make([]int, 6)
	for i := range w {
		if i < len(weapons) {
			w[i] = weapons[i]
		}
	}
	return bridgeAddPlayerClass(teamID, colour, model, pos, angle, w)
}

func (ServerAPI) Broadcast(colour uint32, msg string) {
	for id := 0; id < MaxPlayers; id++ {
		if API.Player.IsConnected(id) {
			API.Player.SendMessage(id, colour, msg)
		}
	}
}

type WorldAPI struct{}

func (WorldAPI) SetBounds(maxX, minX, maxY, minY float32) { bridgeSetWorldBounds(maxX, minX, maxY, minY) }
func (WorldAPI) SetHour(hour int)                           { bridgeSetHour(hour) }
func (WorldAPI) SetWeather(weather int)                     { bridgeSetWeather(weather) }
func (WorldAPI) SetGravity(gravity float32)                 { bridgeSetGravity(gravity) }
func (WorldAPI) CreateExplosion(world, explosionType int, pos Vec3, responsiblePlayer int, atGround bool) {
	bridgeCreateExplosion(world, explosionType, pos, responsiblePlayer, atGround)
}

type PlayerAPI struct{}

func (PlayerAPI) Name(playerID int) string        { return bridgePlayerName(playerID) }
func (PlayerAPI) IDFromName(name string) int      { return bridgePlayerIDFromName(name) }
func (PlayerAPI) IsConnected(playerID int) bool   { return bridgeIsPlayerConnected(playerID) }
func (PlayerAPI) IsAdmin(playerID int) bool     { return bridgeIsPlayerAdmin(playerID) }
func (PlayerAPI) SetAdmin(playerID int, admin bool) { bridgeSetPlayerAdmin(playerID, admin) }
func (PlayerAPI) Kick(playerID int)               { bridgeKickPlayer(playerID) }
func (PlayerAPI) ForceSpawn(playerID int)         { bridgeForceSpawn(playerID) }
func (PlayerAPI) SendMessage(playerID int, colour uint32, msg string) {
	bridgeSendClientMessage(playerID, colour, msg)
}
func (PlayerAPI) SendGameMessage(playerID int, msgType int, msg string) {
	bridgeSendGameMessage(playerID, msgType, msg)
}
func (PlayerAPI) GiveMoney(playerID, amount int)  { bridgeGiveMoney(playerID, amount) }
func (PlayerAPI) SetScore(playerID, score int)    { bridgeSetScore(playerID, score) }
func (PlayerAPI) Score(playerID int) int          { return bridgeGetScore(playerID) }
func (PlayerAPI) Money(playerID int) int          { return bridgeGetMoney(playerID) }
func (PlayerAPI) SetHealth(playerID int, health float32) { bridgeSetHealth(playerID, health) }
func (PlayerAPI) Health(playerID int) float32     { return bridgeGetHealth(playerID) }
func (PlayerAPI) SetArmour(playerID int, armour float32) { bridgeSetArmour(playerID, armour) }
func (PlayerAPI) SetPosition(playerID int, pos Vec3) { bridgeSetPlayerPos(playerID, pos) }
func (PlayerAPI) Position(playerID int) Vec3      { return bridgeGetPlayerPos(playerID) }
func (PlayerAPI) SetSkin(playerID, skin int)      { bridgeSetSkin(playerID, skin) }
func (PlayerAPI) SetTeam(playerID, team int)      { bridgeSetTeam(playerID, team) }
func (PlayerAPI) GiveWeapon(playerID, weapon, ammo int) { bridgeGiveWeapon(playerID, weapon, ammo) }
func (PlayerAPI) PutInVehicle(playerID, vehicleID, slot int, makeRoom, warp bool) {
	bridgePutInVehicle(playerID, vehicleID, slot, makeRoom, warp)
}
func (PlayerAPI) RemoveFromVehicle(playerID int) { bridgeRemoveFromVehicle(playerID) }
func (PlayerAPI) FormatPosition(pos Vec3) string {
	return fmt.Sprintf("%.2f, %.2f, %.2f", pos.X, pos.Y, pos.Z)
}

type VehicleAPI struct{}

func (VehicleAPI) Create(model, world int, pos Vec3, angle float32, primaryColour, secondaryColour int) int {
	return bridgeCreateVehicle(model, world, pos, angle, primaryColour, secondaryColour)
}
func (VehicleAPI) Delete(vehicleID int)              { bridgeDeleteVehicle(vehicleID) }
func (VehicleAPI) Respawn(vehicleID int)           { bridgeRespawnVehicle(vehicleID) }
func (VehicleAPI) SetPosition(vehicleID int, pos Vec3, removeOccupants bool) {
	bridgeSetVehiclePos(vehicleID, pos, removeOccupants)
}
func (VehicleAPI) Position(vehicleID int) Vec3 { return bridgeGetVehiclePos(vehicleID) }

type ObjectAPI struct{}

func (ObjectAPI) Create(model, world int, pos Vec3, alpha int) int {
	return bridgeCreateObject(model, world, pos, alpha)
}
func (ObjectAPI) Delete(objectID int)             { bridgeDeleteObject(objectID) }
func (ObjectAPI) SetPosition(objectID int, pos Vec3) { bridgeSetObjectPos(objectID, pos) }

type PickupAPI struct{}

func (PickupAPI) Create(model, world, quantity int, pos Vec3, alpha int, automatic bool) int {
	return bridgeCreatePickup(model, world, quantity, pos, alpha, automatic)
}
func (PickupAPI) Delete(pickupID int) { bridgeDeletePickup(pickupID) }

type CheckpointAPI struct{}

func (CheckpointAPI) Create(playerID, world int, sphere bool, pos Vec3, rgba [4]int32, radius float32) int {
	return bridgeCreateCheckpoint(playerID, world, sphere, pos, rgba, radius)
}
func (CheckpointAPI) Delete(checkpointID int) { bridgeDeleteCheckpoint(checkpointID) }

type BlipAPI struct{}

func (BlipAPI) Create(index, world int, pos Vec3, scale int, colour uint32, sprite int) int {
	return bridgeCreateBlip(index, world, pos, scale, colour, sprite)
}
func (BlipAPI) Destroy(index int) { bridgeDestroyBlip(index) }

type KeyBindAPI struct{}

func (KeyBindAPI) UnusedSlot() int { return bridgeKeyBindUnusedSlot() }
func (KeyBindAPI) Register(bindID int, onRelease bool, keys ...int) {
	k := make([]int, 3)
	for i := range k {
		if i < len(keys) {
			k[i] = keys[i]
		}
	}
	bridgeRegisterKeyBind(bindID, onRelease, k[0], k[1], k[2])
}
func (KeyBindAPI) Remove(bindID int) { bridgeRemoveKeyBind(bindID) }
