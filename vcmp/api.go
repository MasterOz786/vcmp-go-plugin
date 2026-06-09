package vcmp

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
func (ServerAPI) Time() uint64                           { return bridgeServerTimeMs() }
func (ServerAPI) Name() string                           { return bridgeGetServerName() }
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

func (WorldAPI) SetHour(hour int)       { bridgeSetHour(hour) }
func (WorldAPI) SetWeather(weather int) { bridgeSetWeather(weather) }
func (WorldAPI) SetGravity(g float32)   { bridgeSetGravity(g) }

func (WorldAPI) CreateExplosion(world, explosionType int, pos Vec3, responsiblePlayer int, atGround bool) error {
	return bridgeCreateExplosion(world, explosionType, pos, responsiblePlayer, atGround)
}

type PlayerAPI struct{}

func (PlayerAPI) Name(playerID int) string             { return bridgePlayerName(playerID) }
func (PlayerAPI) IDFromName(name string) int           { return bridgePlayerIDFromName(name) }
func (PlayerAPI) IsConnected(playerID int) bool          { return bridgeIsConnected(playerID) }
func (PlayerAPI) IsAdmin(playerID int) bool              { return bridgeIsAdmin(playerID) }
func (PlayerAPI) SetAdmin(playerID int, admin bool)      { bridgeSetAdmin(playerID, admin) }
func (PlayerAPI) UID(playerID int) string                { return bridgePlayerUID(playerID) }
func (PlayerAPI) Team(playerID int) int                  { return bridgePlayerTeam(playerID) }
func (PlayerAPI) SetTeam(playerID, team int)             { bridgeSetPlayerTeam(playerID, team) }
func (PlayerAPI) World(playerID int) int                 { return bridgePlayerWorld(playerID) }
func (PlayerAPI) SetWorld(playerID, world int)           { bridgeSetPlayerWorld(playerID, world) }
func (PlayerAPI) VehicleID(playerID int) int             { return bridgePlayerVehicleID(playerID) }
func (PlayerAPI) Kick(playerID int) error                { return bridgeKickPlayer(playerID) }
func (PlayerAPI) ForceSpawn(playerID int) error          { return bridgeForcePlayerSpawn(playerID) }
func (PlayerAPI) SendMessage(playerID int, colour uint32, msg string) {
	bridgeSendClientMessage(playerID, colour, msg)
}
func (PlayerAPI) SendGameMessage(playerID int, msgType int, msg string) error {
	return bridgeSendGameMessage(playerID, msgType, msg)
}
func (PlayerAPI) SendScriptData(playerID int, data []byte) error {
	return bridgeSendClientScriptData(playerID, data)
}
func (PlayerAPI) GiveMoney(playerID, amount int) error { return bridgeGivePlayerMoney(playerID, amount) }
func (PlayerAPI) SetScore(playerID, score int)           { bridgeSetPlayerScore(playerID, score) }
func (PlayerAPI) Score(playerID int) int                   { return bridgeGetPlayerScore(playerID) }
func (PlayerAPI) SetHealth(playerID int, health float32) error {
	return bridgeSetPlayerHealth(playerID, health)
}
func (PlayerAPI) Health(playerID int) float32 { return bridgeGetPlayerHealth(playerID) }
func (PlayerAPI) SetArmour(playerID int, armour float32) error {
	return bridgeSetPlayerArmour(playerID, armour)
}
func (PlayerAPI) SetPosition(playerID int, pos Vec3) error { return bridgeSetPlayerPosition(playerID, pos) }
func (PlayerAPI) Position(playerID int) Vec3               { return bridgePlayerPos(playerID) }
func (PlayerAPI) SetHeading(playerID int, angle float32) error {
	return bridgeSetPlayerHeading(playerID, angle)
}
func (PlayerAPI) GiveWeapon(playerID, weapon, ammo int)  { bridgeGiveWeapon(playerID, weapon, ammo) }
func (PlayerAPI) RemoveAllWeapons(playerID int)            { bridgeRemoveAllWeapons(playerID) }
func (PlayerAPI) PutInVehicle(playerID, vehicleID, slot int, makeRoom, warp bool) {
	bridgePutInVehicle(playerID, vehicleID, slot, makeRoom, warp)
}
func (PlayerAPI) RemoveFromVehicle(playerID int) error { return bridgeRemovePlayerFromVehicle(playerID) }
func (PlayerAPI) FormatPosition(pos Vec3) string {
	return fmt.Sprintf("%.2f, %.2f, %.2f", pos.X, pos.Y, pos.Z)
}

type VehicleAPI struct{}

func (VehicleAPI) Create(model, world int, pos Vec3, angle float32, primaryColour, secondaryColour int) int {
	return bridgeCreateVehicle(model, world, pos, angle, primaryColour, secondaryColour)
}
func (VehicleAPI) Delete(vehicleID int)  { bridgeDeleteVehicle(vehicleID) }
func (VehicleAPI) Respawn(vehicleID int) error { return bridgeRespawnVehicle(vehicleID) }
func (VehicleAPI) Explode(vehicleID int) error { return bridgeExplodeVehicle(vehicleID) }
func (VehicleAPI) SetPosition(vehicleID int, pos Vec3, removeOccupants bool) {
	_ = removeOccupants
	bridgeSetVehiclePosition(vehicleID, pos)
}
func (VehicleAPI) Position(vehicleID int) Vec3 { return bridgeVehiclePos(vehicleID) }
func (VehicleAPI) Health(vehicleID int) float32 { return bridgeVehicleHealth(vehicleID) }
func (VehicleAPI) SetHealth(vehicleID int, health float32) { bridgeSetVehicleHealth(vehicleID, health) }
func (VehicleAPI) Break(vehicleID int) { bridgeBreakVehicle(vehicleID) }

type ObjectAPI struct{}

func (ObjectAPI) Create(model, world int, pos Vec3, alpha int) int {
	return bridgeCreateObject(model, world, pos, alpha)
}
func (ObjectAPI) Delete(objectID int) error { return bridgeDeleteObject(objectID) }

type PickupAPI struct{}

func (PickupAPI) Create(model, world, quantity int, pos Vec3, alpha int, automatic bool) int {
	return bridgeCreatePickup(model, world, quantity, pos, alpha, automatic)
}
func (PickupAPI) Delete(pickupID int) error { return bridgeDeletePickup(pickupID) }

type CheckpointAPI struct{}

func (CheckpointAPI) Create(playerID, world int, sphere bool, pos Vec3, r, g, b, alpha int, radius float32) int {
	return bridgeCreateCheckPoint(playerID, world, sphere, pos, r, g, b, alpha, radius)
}
func (CheckpointAPI) Delete(checkpointID int) error { return bridgeDeleteCheckPoint(checkpointID) }

type BlipAPI struct{}

func (BlipAPI) Create(index, world int, pos Vec3, scale int, colour uint32, sprite int) int {
	return bridgeCreateCoordBlip(index, world, pos, scale, colour, sprite)
}
func (BlipAPI) Destroy(index int) error { return bridgeDestroyCoordBlip(index) }

type KeyBindAPI struct{}

func (KeyBindAPI) Register(bindID int, onRelease bool, keyOne, keyTwo, keyThree int) error {
	return bridgeRegisterKeyBind(bindID, onRelease, keyOne, keyTwo, keyThree)
}
func (KeyBindAPI) Remove(bindID int) error { return bridgeRemoveKeyBind(bindID) }
