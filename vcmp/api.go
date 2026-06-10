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
	Radio      RadioAPI
	Entity     EntityAPI
	Weapon     WeaponAPI
	Handling   HandlingAPI
	Admin      AdminAPI
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
	Radio:      RadioAPI{},
	Entity:     EntityAPI{},
	Weapon:     WeaponAPI{},
	Handling:   HandlingAPI{},
	Admin:      AdminAPI{},
}

type ServerAPI struct{}

func (ServerAPI) Log(msg string)                         { bridgeLog(msg) }
func (ServerAPI) Time() uint64                           { return bridgeServerTimeMs() }
func (ServerAPI) Name() string                           { return bridgeGetServerName() }
func (ServerAPI) SetName(name string)                    { bridgeSetServerName(name) }
func (ServerAPI) GameModeText() string                   { return bridgeGetGameModeText() }
func (ServerAPI) SetGameModeText(text string)            { bridgeSetGameModeText(text) }
func (ServerAPI) Option(option ServerOption) bool        { return bridgeGetServerOption(option) }
func (ServerAPI) SetOption(option ServerOption, on bool) { bridgeSetServerOption(option, on) }
func (ServerAPI) SetSpawnPosition(pos Vec3)              { bridgeSetSpawnPos(pos) }
func (ServerAPI) SetSpawnCameraPosition(pos Vec3)        { bridgeSetSpawnCameraPosition(pos) }
func (ServerAPI) SetSpawnCameraLookAt(pos Vec3)          { bridgeSetSpawnCameraLookAt(pos) }
func (ServerAPI) SetKillCommandDelay(delay int)          { bridgeSetKillCommandDelay(delay) }
func (ServerAPI) KillCommandDelay() int                  { return bridgeGetKillCommandDelay() }
func (ServerAPI) ForceAllSelect()                        { bridgeForceAllSelect() }
func (ServerAPI) Version() uint32                        { return bridgeServerVersion() }
func (ServerAPI) Settings() (ServerSettings, error)      { return bridgeServerSettings() }
func (ServerAPI) PluginCount() uint32                    { return bridgePluginCount() }
func (ServerAPI) FindPlugin(name string) int             { return bridgeFindPlugin(name) }
func (ServerAPI) LastError() error                       { return bridgeLastError() }
func (ServerAPI) Shutdown()                              { bridgeShutdownServer() }
func (ServerAPI) SetMaxPlayers(max uint32) error         { return bridgeSetMaxPlayers(max) }
func (ServerAPI) MaxPlayers() uint32                     { return bridgeGetMaxPlayers() }
func (ServerAPI) SetPassword(password string) error      { return bridgeSetServerPassword(password) }
func (ServerAPI) Password() string                       { return bridgeGetServerPassword() }
func (ServerAPI) SetWastedSettings(s WastedSettings)     { bridgeSetWastedSettings(s) }
func (ServerAPI) WastedSettings() WastedSettings         { return bridgeGetWastedSettings() }
func (ServerAPI) SetTimeRate(rate int)                   { bridgeSetTimeRate(rate) }
func (ServerAPI) TimeRate() int                          { return bridgeGetTimeRate() }
func (ServerAPI) SetMinute(minute int)                   { bridgeSetMinute(minute) }
func (ServerAPI) Minute() int                            { return bridgeGetMinute() }
func (ServerAPI) SetGameSpeed(s float32)                 { bridgeSetGameSpeed(s) }
func (ServerAPI) GameSpeed() float32                     { return bridgeGetGameSpeed() }
func (ServerAPI) SetWaterLevel(l float32)                { bridgeSetWaterLevel(l) }
func (ServerAPI) WaterLevel() float32                    { return bridgeGetWaterLevel() }
func (ServerAPI) SetMaxFlightAltitude(h float32)         { bridgeSetMaxFlightAltitude(h) }
func (ServerAPI) MaxFlightAltitude() float32             { return bridgeGetMaxFlightAltitude() }
func (ServerAPI) SetVehiclesForcedRespawnHeight(h float32) { bridgeSetVehiclesForcedRespawnHeight(h) }
func (ServerAPI) VehiclesForcedRespawnHeight() float32   { return bridgeGetVehiclesForcedRespawnHeight() }
func (ServerAPI) SetFallTimer(t uint16)                  { bridgeSetFallTimer(t) }
func (ServerAPI) FallTimer() uint16                      { return bridgeGetFallTimer() }

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
func (WorldAPI) Hour() int              { return bridgeGetHour() }
func (WorldAPI) SetWeather(weather int) { bridgeSetWeather(weather) }
func (WorldAPI) Weather() int           { return bridgeGetWeather() }
func (WorldAPI) SetGravity(g float32)   { bridgeSetGravity(g) }
func (WorldAPI) Gravity() float32       { return bridgeGetGravity() }

func (WorldAPI) CreateExplosion(world, explosionType int, pos Vec3, responsiblePlayer int, atGround bool) error {
	return bridgeCreateExplosion(world, explosionType, pos, responsiblePlayer, atGround)
}
func (WorldAPI) PlaySound(world, soundID int, pos Vec3) error {
	return bridgePlaySound(world, soundID, pos)
}
func (WorldAPI) SetBounds(bounds WorldBounds) { bridgeSetWorldBounds(bounds) }
func (WorldAPI) Bounds() WorldBounds          { return bridgeGetWorldBounds() }
func (WorldAPI) HideMapObject(coord MapObjectCoord) { bridgeHideMapObject(coord) }
func (WorldAPI) ShowMapObject(coord MapObjectCoord) { bridgeShowMapObject(coord) }
func (WorldAPI) ShowAllMapObjects()                 { bridgeShowAllMapObjects() }

type PlayerAPI struct{}

func (PlayerAPI) Name(playerID int) string             { return bridgePlayerName(playerID) }
func (PlayerAPI) IDFromName(name string) int           { return bridgePlayerIDFromName(name) }
func (PlayerAPI) IsConnected(playerID int) bool          { return bridgeIsConnected(playerID) }
func (PlayerAPI) IsSpawned(playerID int) bool            { return bridgeIsPlayerSpawned(playerID) }
func (PlayerAPI) State(playerID int) PlayerState         { return bridgeGetPlayerState(playerID) }
func (PlayerAPI) IsAdmin(playerID int) bool              { return bridgeIsAdmin(playerID) }
func (PlayerAPI) SetAdmin(playerID int, admin bool)      { bridgeSetAdmin(playerID, admin) }
func (PlayerAPI) UID(playerID int) string                { return bridgePlayerUID(playerID) }
func (PlayerAPI) UID2(playerID int) string               { return bridgePlayerUID2(playerID) }
func (PlayerAPI) IP(playerID int) string                 { return bridgeGetPlayerIP(playerID) }
func (PlayerAPI) Ping(playerID int) int                  { return bridgeGetPlayerPing(playerID) }
func (PlayerAPI) FPS(playerID int) float64               { return bridgeGetPlayerFPS(playerID) }
func (PlayerAPI) Key(playerID int) uint32                { return bridgeGetPlayerKey(playerID) }
func (PlayerAPI) IsTyping(playerID int) bool             { return bridgeIsPlayerTyping(playerID) }
func (PlayerAPI) SetName(playerID int, name string) error { return bridgeSetPlayerName(playerID, name) }
func (PlayerAPI) Class(playerID int) int                 { return bridgeGetPlayerClass(playerID) }
func (PlayerAPI) Team(playerID int) int                  { return bridgePlayerTeam(playerID) }
func (PlayerAPI) SetTeam(playerID, team int)             { bridgeSetPlayerTeam(playerID, team) }
func (PlayerAPI) Skin(playerID int) int                  { return bridgeGetPlayerSkin(playerID) }
func (PlayerAPI) SetSkin(playerID, skinID int) error     { return bridgeSetPlayerSkin(playerID, skinID) }
func (PlayerAPI) Colour(playerID int) uint32            { return bridgeGetPlayerColour(playerID) }
func (PlayerAPI) SetColour(playerID int, colour uint32) error {
	return bridgeSetPlayerColour(playerID, colour)
}
func (PlayerAPI) World(playerID int) int                 { return bridgePlayerWorld(playerID) }
func (PlayerAPI) SetWorld(playerID, world int)           { bridgeSetPlayerWorld(playerID, world) }
func (PlayerAPI) SetSecondaryWorld(playerID, world int) error {
	return bridgeSetPlayerSecondaryWorld(playerID, world)
}
func (PlayerAPI) SecondaryWorld(playerID int) int  { return bridgeGetPlayerSecondaryWorld(playerID) }
func (PlayerAPI) UniqueWorld(playerID int) int     { return bridgeGetPlayerUniqueWorld(playerID) }
func (PlayerAPI) WorldCompatible(playerID, world int) bool {
	return bridgeIsPlayerWorldCompatible(playerID, world)
}
func (PlayerAPI) VehicleID(playerID int) int             { return bridgePlayerVehicleID(playerID) }
func (PlayerAPI) VehicleSlot(playerID int) int           { return bridgeGetPlayerInVehicleSlot(playerID) }
func (PlayerAPI) InVehicleStatus(playerID int) PlayerVehicle {
	return bridgeGetPlayerInVehicleStatus(playerID)
}
func (PlayerAPI) IsStreamedFor(checkedPlayerID, forPlayerID int) bool {
	return bridgeIsPlayerStreamedForPlayer(checkedPlayerID, forPlayerID)
}
func (PlayerAPI) Option(playerID int, option PlayerOption) bool {
	return bridgeGetPlayerOption(playerID, option)
}
func (PlayerAPI) SetOption(playerID int, option PlayerOption, on bool) error {
	return bridgeSetPlayerOption(playerID, option, on)
}
func (PlayerAPI) Kick(playerID int) error                { return bridgeKickPlayer(playerID) }
func (PlayerAPI) Ban(playerID int) error                 { return bridgeBanPlayer(playerID) }
func (PlayerAPI) Kill(playerID int) error                { return bridgeKillPlayer(playerID) }
func (PlayerAPI) ForceSpawn(playerID int) error          { return bridgeForcePlayerSpawn(playerID) }
func (PlayerAPI) ForceSelect(playerID int) error         { return bridgeForcePlayerSelect(playerID) }
func (PlayerAPI) Redirect(playerID int, ip string, port uint32, nick, serverPassword, userPassword string) error {
	return bridgeRedirectPlayerToServer(playerID, ip, port, nick, serverPassword, userPassword)
}
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
func (PlayerAPI) SetMoney(playerID, amount int) error  { return bridgeSetPlayerMoney(playerID, amount) }
func (PlayerAPI) Money(playerID int) int               { return bridgeGetPlayerMoney(playerID) }
func (PlayerAPI) SetScore(playerID, score int)           { bridgeSetPlayerScore(playerID, score) }
func (PlayerAPI) Score(playerID int) int                   { return bridgeGetPlayerScore(playerID) }
func (PlayerAPI) SetWantedLevel(playerID, level int) error {
	return bridgeSetPlayerWantedLevel(playerID, level)
}
func (PlayerAPI) WantedLevel(playerID int) int { return bridgeGetPlayerWantedLevel(playerID) }
func (PlayerAPI) SetHealth(playerID int, health float32) error {
	return bridgeSetPlayerHealth(playerID, health)
}
func (PlayerAPI) Health(playerID int) float32 { return bridgeGetPlayerHealth(playerID) }
func (PlayerAPI) SetArmour(playerID int, armour float32) error {
	return bridgeSetPlayerArmour(playerID, armour)
}
func (PlayerAPI) Armour(playerID int) float32 { return bridgeGetPlayerArmour(playerID) }
func (PlayerAPI) SetImmunityFlags(playerID int, flags uint32) error {
	return bridgeSetPlayerImmunityFlags(playerID, flags)
}
func (PlayerAPI) ImmunityFlags(playerID int) uint32 { return bridgeGetPlayerImmunityFlags(playerID) }
func (PlayerAPI) SetPosition(playerID int, pos Vec3) error { return bridgeSetPlayerPosition(playerID, pos) }
func (PlayerAPI) SetSpeed(playerID int, speed Vec3) error  { return bridgeSetPlayerSpeed(playerID, speed) }
func (PlayerAPI) Speed(playerID int) (Vec3, error)       { return bridgeGetPlayerSpeed(playerID) }
func (PlayerAPI) AddSpeed(playerID int, delta Vec3) error  { return bridgeAddPlayerSpeed(playerID, delta) }
func (PlayerAPI) SetAlpha(playerID, alpha int, fadeMs uint32) error {
	return bridgeSetPlayerAlpha(playerID, alpha, fadeMs)
}
func (PlayerAPI) Alpha(playerID int) int { return bridgeGetPlayerAlpha(playerID) }
func (PlayerAPI) AimPosition(playerID int) (Vec3, error) { return bridgeGetPlayerAimPosition(playerID) }
func (PlayerAPI) AimDirection(playerID int) (Vec3, error)  { return bridgeGetPlayerAimDirection(playerID) }
func (PlayerAPI) IsOnFire(playerID int) bool     { return bridgeIsPlayerOnFire(playerID) }
func (PlayerAPI) IsCrouching(playerID int) bool  { return bridgeIsPlayerCrouching(playerID) }
func (PlayerAPI) Action(playerID int) int        { return bridgeGetPlayerAction(playerID) }
func (PlayerAPI) GameKeys(playerID int) uint32   { return bridgeGetPlayerGameKeys(playerID) }
func (PlayerAPI) SetAnimation(playerID, groupID, animationID int) error {
	return bridgeSetPlayerAnimation(playerID, groupID, animationID)
}
func (PlayerAPI) StandingOnVehicle(playerID int) int { return bridgeGetPlayerStandingOnVehicle(playerID) }
func (PlayerAPI) StandingOnObject(playerID int) int  { return bridgeGetPlayerStandingOnObject(playerID) }
func (PlayerAPI) IsAway(playerID int) bool           { return bridgeIsPlayerAway(playerID) }
func (PlayerAPI) Set3DArrowFor(playerID, targetPlayerID int, enabled bool) error {
	return bridgeSetPlayer3DArrowForPlayer(playerID, targetPlayerID, enabled)
}
func (PlayerAPI) Has3DArrowFor(playerID, targetPlayerID int) bool {
	return bridgeGetPlayer3DArrowForPlayer(playerID, targetPlayerID)
}
func (PlayerAPI) Position(playerID int) Vec3               { return bridgePlayerPos(playerID) }
func (PlayerAPI) SetHeading(playerID int, angle float32) error {
	return bridgeSetPlayerHeading(playerID, angle)
}
func (PlayerAPI) Heading(playerID int) float32 { return bridgeGetPlayerHeading(playerID) }
func (PlayerAPI) GiveWeapon(playerID, weapon, ammo int)  { bridgeGiveWeapon(playerID, weapon, ammo) }
func (PlayerAPI) SetWeapon(playerID, weapon, ammo int) error {
	return bridgeSetPlayerWeapon(playerID, weapon, ammo)
}
func (PlayerAPI) Weapon(playerID int) int      { return bridgeGetPlayerWeapon(playerID) }
func (PlayerAPI) WeaponAmmo(playerID int) int  { return bridgeGetPlayerWeaponAmmo(playerID) }
func (PlayerAPI) WeaponSlot(playerID int) int  { return bridgeGetPlayerWeaponSlot(playerID) }
func (PlayerAPI) SetWeaponSlot(playerID, slot int) error {
	return bridgeSetPlayerWeaponSlot(playerID, slot)
}
func (PlayerAPI) RemoveAllWeapons(playerID int)            { bridgeRemoveAllWeapons(playerID) }
func (PlayerAPI) WeaponAtSlot(playerID, slot int) int       { return bridgeGetPlayerWeaponAtSlot(playerID, slot) }
func (PlayerAPI) AmmoAtSlot(playerID, slot int) int         { return bridgeGetPlayerAmmoAtSlot(playerID, slot) }
func (PlayerAPI) RemoveWeapon(playerID, weaponID int) error { return bridgeRemovePlayerWeapon(playerID, weaponID) }
func (PlayerAPI) SetCamera(playerID int, pos, lookAt Vec3) error {
	return bridgeSetCameraPosition(playerID, pos, lookAt)
}
func (PlayerAPI) RestoreCamera(playerID int) error { return bridgeRestoreCamera(playerID) }
func (PlayerAPI) CameraLocked(playerID int) bool   { return bridgeIsCameraLocked(playerID) }
func (PlayerAPI) InterpolateCameraLookAt(playerID int, lookAt Vec3, interpMs uint32) error {
	return bridgeInterpolateCameraLookAt(playerID, lookAt, interpMs)
}
func (PlayerAPI) SetSpectateTarget(playerID, targetID int) error {
	return bridgeSetPlayerSpectateTarget(playerID, targetID)
}
func (PlayerAPI) SpectateTarget(playerID int) int { return bridgeGetPlayerSpectateTarget(playerID) }
func (PlayerAPI) SetDrunkHandling(playerID int, level uint32) error {
	return bridgeSetPlayerDrunkHandling(playerID, level)
}
func (PlayerAPI) DrunkHandling(playerID int) uint32 { return bridgeGetPlayerDrunkHandling(playerID) }
func (PlayerAPI) SetDrunkVisuals(playerID int, level uint8) error {
	return bridgeSetPlayerDrunkVisuals(playerID, level)
}
func (PlayerAPI) DrunkVisuals(playerID int) uint8 { return bridgeGetPlayerDrunkVisuals(playerID) }
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
func (VehicleAPI) Exists(vehicleID int) bool {
	return bridgeCheckEntityExists(EntityPoolVehicle, vehicleID)
}
func (VehicleAPI) IsStreamedFor(vehicleID, playerID int) bool {
	return bridgeIsVehicleStreamedForPlayer(vehicleID, playerID)
}
func (VehicleAPI) SyncSource(vehicleID int) int { return bridgeGetVehicleSyncSource(vehicleID) }
func (VehicleAPI) SyncType(vehicleID int) VehicleSync { return bridgeGetVehicleSyncType(vehicleID) }
func (VehicleAPI) Option(vehicleID int, option VehicleOption) bool {
	return bridgeGetVehicleOption(vehicleID, option)
}
func (VehicleAPI) SetOption(vehicleID int, option VehicleOption, on bool) error {
	return bridgeSetVehicleOption(vehicleID, option, on)
}
func (VehicleAPI) World(vehicleID int) int { return bridgeGetVehicleWorld(vehicleID) }
func (VehicleAPI) SetWorld(vehicleID, world int) error {
	return bridgeSetVehicleWorld(vehicleID, world)
}
func (VehicleAPI) Model(vehicleID int) int { return bridgeGetVehicleModel(vehicleID) }
func (VehicleAPI) Occupant(vehicleID, slot int) int {
	return bridgeVehicleOccupant(vehicleID, slot)
}
func (VehicleAPI) SetPosition(vehicleID int, pos Vec3, removeOccupants bool) {
	_ = removeOccupants
	bridgeSetVehiclePosition(vehicleID, pos)
}
func (VehicleAPI) Position(vehicleID int) Vec3 { return bridgeVehiclePos(vehicleID) }
func (VehicleAPI) SetRotationEuler(vehicleID int, rot Vec3) error {
	return bridgeSetVehicleRotationEuler(vehicleID, rot)
}
func (VehicleAPI) Health(vehicleID int) float32 { return bridgeVehicleHealth(vehicleID) }
func (VehicleAPI) SetHealth(vehicleID int, health float32) { bridgeSetVehicleHealth(vehicleID, health) }
func (VehicleAPI) SetColour(vehicleID, primary, secondary int) error {
	return bridgeSetVehicleColour(vehicleID, primary, secondary)
}
func (VehicleAPI) Colour(vehicleID int) (primary, secondary int, err error) {
	return bridgeGetVehicleColour(vehicleID)
}
func (VehicleAPI) SetRadio(vehicleID, radioID int) error {
	return bridgeSetVehicleRadio(vehicleID, radioID)
}
func (VehicleAPI) Radio(vehicleID int) int { return bridgeGetVehicleRadio(vehicleID) }
func (VehicleAPI) SetImmunityFlags(vehicleID int, flags uint32) error {
	return bridgeSetVehicleImmunityFlags(vehicleID, flags)
}
func (VehicleAPI) ImmunityFlags(vehicleID int) uint32 {
	return bridgeGetVehicleImmunityFlags(vehicleID)
}
func (VehicleAPI) IsWrecked(vehicleID int) bool { return bridgeIsVehicleWrecked(vehicleID) }
func (VehicleAPI) SetSpeed(vehicleID int, speed Vec3, add, relative bool) error {
	return bridgeSetVehicleSpeed(vehicleID, speed, add, relative)
}
func (VehicleAPI) Speed(vehicleID int, relative bool) (Vec3, error) {
	return bridgeGetVehicleSpeed(vehicleID, relative)
}
func (VehicleAPI) RotationEuler(vehicleID int) (Vec3, error) {
	return bridgeGetVehicleRotationEuler(vehicleID)
}
func (VehicleAPI) SetSpawnPosition(vehicleID int, pos Vec3) error {
	return bridgeSetVehicleSpawnPosition(vehicleID, pos)
}
func (VehicleAPI) SetLightsData(vehicleID int, lights uint32) error {
	return bridgeSetVehicleLightsData(vehicleID, lights)
}
func (VehicleAPI) LightsData(vehicleID int) uint32 { return bridgeGetVehicleLightsData(vehicleID) }
func (VehicleAPI) SetRotation(vehicleID int, rot Quat) error {
	return bridgeSetVehicleRotation(vehicleID, rot)
}
func (VehicleAPI) Rotation(vehicleID int) (Quat, error) { return bridgeGetVehicleRotation(vehicleID) }
func (VehicleAPI) SetTurnSpeed(vehicleID int, speed Vec3, add, relative bool) error {
	return bridgeSetVehicleTurnSpeed(vehicleID, speed, add, relative)
}
func (VehicleAPI) TurnSpeed(vehicleID int, relative bool) (Vec3, error) {
	return bridgeGetVehicleTurnSpeed(vehicleID, relative)
}
func (VehicleAPI) SpawnPosition(vehicleID int) (Vec3, error) {
	return bridgeGetVehicleSpawnPosition(vehicleID)
}
func (VehicleAPI) SetSpawnRotationEuler(vehicleID int, rot Vec3) error {
	return bridgeSetVehicleSpawnRotationEuler(vehicleID, rot)
}
func (VehicleAPI) SetSpawnRotation(vehicleID int, rot Quat) error {
	return bridgeSetVehicleSpawnRotation(vehicleID, rot)
}
func (VehicleAPI) SpawnRotationEuler(vehicleID int) (Vec3, error) {
	return bridgeGetVehicleSpawnRotationEuler(vehicleID)
}
func (VehicleAPI) SpawnRotation(vehicleID int) (Quat, error) {
	return bridgeGetVehicleSpawnRotation(vehicleID)
}
func (VehicleAPI) SetIdleRespawnTimer(vehicleID int, millis uint32) error {
	return bridgeSetVehicleIdleRespawnTimer(vehicleID, millis)
}
func (VehicleAPI) IdleRespawnTimer(vehicleID int) uint32 {
	return bridgeGetVehicleIdleRespawnTimer(vehicleID)
}
func (VehicleAPI) SetPartStatus(vehicleID, partID, status int) error {
	return bridgeSetVehiclePartStatus(vehicleID, partID, status)
}
func (VehicleAPI) PartStatus(vehicleID, partID int) int {
	return bridgeGetVehiclePartStatus(vehicleID, partID)
}
func (VehicleAPI) SetTyreStatus(vehicleID, tyreID, status int) error {
	return bridgeSetVehicleTyreStatus(vehicleID, tyreID, status)
}
func (VehicleAPI) TyreStatus(vehicleID, tyreID int) int {
	return bridgeGetVehicleTyreStatus(vehicleID, tyreID)
}
func (VehicleAPI) SetDamageData(vehicleID int, data uint32) error {
	return bridgeSetVehicleDamageData(vehicleID, data)
}
func (VehicleAPI) DamageData(vehicleID int) uint32 { return bridgeGetVehicleDamageData(vehicleID) }
func (VehicleAPI) TurretRotation(vehicleID int) (horizontal, vertical float32, err error) {
	return bridgeGetVehicleTurretRotation(vehicleID)
}
func (VehicleAPI) Set3DArrowFor(vehicleID, targetPlayerID int, enabled bool) error {
	return bridgeSetVehicle3DArrowForPlayer(vehicleID, targetPlayerID, enabled)
}
func (VehicleAPI) Has3DArrowFor(vehicleID, targetPlayerID int) bool {
	return bridgeGetVehicle3DArrowForPlayer(vehicleID, targetPlayerID)
}
func (VehicleAPI) Break(vehicleID int) { bridgeBreakVehicle(vehicleID) }

type ObjectAPI struct{}

func (ObjectAPI) Create(model, world int, pos Vec3, alpha int) int {
	return bridgeCreateObject(model, world, pos, alpha)
}
func (ObjectAPI) Delete(objectID int) error { return bridgeDeleteObject(objectID) }
func (ObjectAPI) Exists(objectID int) bool {
	return bridgeCheckEntityExists(EntityPoolObject, objectID)
}
func (ObjectAPI) IsStreamedFor(objectID, playerID int) bool {
	return bridgeIsObjectStreamedForPlayer(objectID, playerID)
}
func (ObjectAPI) Model(objectID int) int { return bridgeGetObjectModel(objectID) }
func (ObjectAPI) World(objectID int) int { return bridgeGetObjectWorld(objectID) }
func (ObjectAPI) SetWorld(objectID, world int) error {
	return bridgeSetObjectWorld(objectID, world)
}
func (ObjectAPI) Alpha(objectID int) int { return bridgeGetObjectAlpha(objectID) }
func (ObjectAPI) SetAlpha(objectID, alpha int, durationMs uint32) error {
	return bridgeSetObjectAlpha(objectID, alpha, durationMs)
}
func (ObjectAPI) Position(objectID int) (Vec3, error) { return bridgeGetObjectPosition(objectID) }
func (ObjectAPI) SetPosition(objectID int, pos Vec3) error {
	return bridgeSetObjectPosition(objectID, pos)
}
func (ObjectAPI) MoveTo(objectID int, pos Vec3, durationMs uint32) error {
	return bridgeMoveObjectTo(objectID, pos, durationMs)
}
func (ObjectAPI) RotationEuler(objectID int) (Vec3, error) {
	return bridgeGetObjectRotationEuler(objectID)
}
func (ObjectAPI) RotateToEuler(objectID int, rot Vec3, durationMs uint32) error {
	return bridgeRotateObjectToEuler(objectID, rot, durationMs)
}
func (ObjectAPI) RotateTo(objectID int, rot Quat, durationMs uint32) error {
	return bridgeRotateObjectTo(objectID, rot, durationMs)
}
func (ObjectAPI) SetShotReport(objectID int, on bool) error {
	return bridgeSetObjectShotReport(objectID, on)
}
func (ObjectAPI) ShotReportEnabled(objectID int) bool { return bridgeIsObjectShotReport(objectID) }
func (ObjectAPI) SetTouchedReport(objectID int, on bool) error {
	return bridgeSetObjectTouchedReport(objectID, on)
}
func (ObjectAPI) TouchedReportEnabled(objectID int) bool {
	return bridgeIsObjectTouchedReport(objectID)
}
func (ObjectAPI) MoveBy(objectID int, offset Vec3, durationMs uint32) error {
	return bridgeMoveObjectBy(objectID, offset, durationMs)
}
func (ObjectAPI) RotateByEuler(objectID int, rot Vec3, durationMs uint32) error {
	return bridgeRotateObjectByEuler(objectID, rot, durationMs)
}
func (ObjectAPI) RotateBy(objectID int, rot Quat, durationMs uint32) error {
	return bridgeRotateObjectBy(objectID, rot, durationMs)
}
func (ObjectAPI) Rotation(objectID int) (Quat, error) { return bridgeGetObjectRotation(objectID) }

type PickupAPI struct{}

func (PickupAPI) Create(model, world, quantity int, pos Vec3, alpha int, automatic bool) int {
	return bridgeCreatePickup(model, world, quantity, pos, alpha, automatic)
}
func (PickupAPI) Delete(pickupID int) error { return bridgeDeletePickup(pickupID) }
func (PickupAPI) Exists(pickupID int) bool {
	return bridgeCheckEntityExists(EntityPoolPickup, pickupID)
}
func (PickupAPI) IsStreamedFor(pickupID, playerID int) bool {
	return bridgeIsPickupStreamedForPlayer(pickupID, playerID)
}
func (PickupAPI) World(pickupID int) int { return bridgeGetPickupWorld(pickupID) }
func (PickupAPI) SetWorld(pickupID, world int) error { return bridgeSetPickupWorld(pickupID, world) }
func (PickupAPI) Alpha(pickupID int) int             { return bridgeGetPickupAlpha(pickupID) }
func (PickupAPI) SetAlpha(pickupID, alpha int) error { return bridgeSetPickupAlpha(pickupID, alpha) }
func (PickupAPI) Automatic(pickupID int) bool        { return bridgeIsPickupAutomatic(pickupID) }
func (PickupAPI) SetAutomatic(pickupID int, on bool) error {
	return bridgeSetPickupAutomatic(pickupID, on)
}
func (PickupAPI) Refresh(pickupID int) error { return bridgeRefreshPickup(pickupID) }
func (PickupAPI) Position(pickupID int) (Vec3, error) { return bridgeGetPickupPosition(pickupID) }
func (PickupAPI) SetPosition(pickupID int, pos Vec3) error {
	return bridgeSetPickupPosition(pickupID, pos)
}
func (PickupAPI) Model(pickupID int) int    { return bridgeGetPickupModel(pickupID) }
func (PickupAPI) Quantity(pickupID int) int { return bridgeGetPickupQuantity(pickupID) }
func (PickupAPI) Option(pickupID int, option PickupOption) bool {
	return bridgeGetPickupOption(pickupID, option)
}
func (PickupAPI) SetOption(pickupID int, option PickupOption, on bool) error {
	return bridgeSetPickupOption(pickupID, option, on)
}
func (PickupAPI) SetAutoTimer(pickupID int, durationMs uint32) error {
	return bridgeSetPickupAutoTimer(pickupID, durationMs)
}
func (PickupAPI) AutoTimer(pickupID int) uint32 { return bridgeGetPickupAutoTimer(pickupID) }

type CheckpointAPI struct{}

func (CheckpointAPI) Create(playerID, world int, sphere bool, pos Vec3, r, g, b, alpha int, radius float32) int {
	return bridgeCreateCheckPoint(playerID, world, sphere, pos, r, g, b, alpha, radius)
}
func (CheckpointAPI) Delete(checkpointID int) error { return bridgeDeleteCheckPoint(checkpointID) }
func (CheckpointAPI) Exists(checkpointID int) bool {
	return bridgeCheckEntityExists(EntityPoolCheckPoint, checkpointID)
}
func (CheckpointAPI) IsStreamedFor(checkpointID, playerID int) bool {
	return bridgeIsCheckPointStreamedForPlayer(checkpointID, playerID)
}
func (CheckpointAPI) Sphere(checkpointID int) bool { return bridgeIsCheckPointSphere(checkpointID) }
func (CheckpointAPI) World(checkpointID int) int   { return bridgeGetCheckPointWorld(checkpointID) }
func (CheckpointAPI) SetWorld(checkpointID, world int) error {
	return bridgeSetCheckPointWorld(checkpointID, world)
}
func (CheckpointAPI) SetColour(checkpointID, r, g, b, alpha int) error {
	return bridgeSetCheckPointColour(checkpointID, r, g, b, alpha)
}
func (CheckpointAPI) Colour(checkpointID int) (r, g, b, alpha int, err error) {
	return bridgeGetCheckPointColour(checkpointID)
}
func (CheckpointAPI) Position(checkpointID int) (Vec3, error) {
	return bridgeGetCheckPointPosition(checkpointID)
}
func (CheckpointAPI) SetPosition(checkpointID int, pos Vec3) error {
	return bridgeSetCheckPointPosition(checkpointID, pos)
}
func (CheckpointAPI) SetRadius(checkpointID int, radius float32) error {
	return bridgeSetCheckPointRadius(checkpointID, radius)
}
func (CheckpointAPI) Radius(checkpointID int) float32 { return bridgeGetCheckPointRadius(checkpointID) }
func (CheckpointAPI) Owner(checkpointID int) int    { return bridgeGetCheckPointOwner(checkpointID) }

type BlipAPI struct{}

func (BlipAPI) Create(index, world int, pos Vec3, scale int, colour uint32, sprite int) int {
	return bridgeCreateCoordBlip(index, world, pos, scale, colour, sprite)
}
func (BlipAPI) Destroy(index int) error { return bridgeDestroyCoordBlip(index) }
func (BlipAPI) Info(index int) (BlipInfo, error) { return bridgeGetCoordBlipInfo(index) }

type KeyBindAPI struct{}

func (KeyBindAPI) UnusedSlot() int { return bridgeGetKeyBindUnusedSlot() }

func (KeyBindAPI) Register(bindID int, onRelease bool, keyOne, keyTwo, keyThree int) error {
	return bridgeRegisterKeyBind(bindID, onRelease, keyOne, keyTwo, keyThree)
}
func (KeyBindAPI) Remove(bindID int) error { return bridgeRemoveKeyBind(bindID) }
func (KeyBindAPI) RemoveAll()              { bridgeRemoveAllKeyBinds() }
func (KeyBindAPI) Data(bindID int) (KeyBindData, error) { return bridgeGetKeyBindData(bindID) }

type WeaponAPI struct{}

func (WeaponAPI) SetDataValue(weaponID, fieldID int, value float64) error {
	return bridgeSetWeaponDataValue(weaponID, fieldID, value)
}
func (WeaponAPI) DataValue(weaponID, fieldID int) float64 {
	return bridgeGetWeaponDataValue(weaponID, fieldID)
}
func (WeaponAPI) ResetDataValue(weaponID, fieldID int) error {
	return bridgeResetWeaponDataValue(weaponID, fieldID)
}
func (WeaponAPI) IsDataModified(weaponID, fieldID int) bool {
	return bridgeIsWeaponDataModified(weaponID, fieldID)
}
func (WeaponAPI) ResetData(weaponID int) error { return bridgeResetWeaponData(weaponID) }
func (WeaponAPI) ResetAllData()                { bridgeResetAllWeaponData() }

type HandlingAPI struct{}

func (HandlingAPI) ResetAll() { bridgeResetAllVehicleHandlings() }
func (HandlingAPI) ExistsRule(modelIndex, ruleIndex int) bool {
	return bridgeExistsHandlingRule(modelIndex, ruleIndex)
}
func (HandlingAPI) SetRule(modelIndex, ruleIndex int, value float64) error {
	return bridgeSetHandlingRule(modelIndex, ruleIndex, value)
}
func (HandlingAPI) Rule(modelIndex, ruleIndex int) float64 {
	return bridgeGetHandlingRule(modelIndex, ruleIndex)
}
func (HandlingAPI) ResetRule(modelIndex, ruleIndex int) error {
	return bridgeResetHandlingRule(modelIndex, ruleIndex)
}
func (HandlingAPI) Reset(modelIndex int) error { return bridgeResetHandling(modelIndex) }
func (HandlingAPI) ExistsInstRule(vehicleID, ruleIndex int) bool {
	return bridgeExistsInstHandlingRule(vehicleID, ruleIndex)
}
func (HandlingAPI) SetInstRule(vehicleID, ruleIndex int, value float64) error {
	return bridgeSetInstHandlingRule(vehicleID, ruleIndex, value)
}
func (HandlingAPI) InstRule(vehicleID, ruleIndex int) float64 {
	return bridgeGetInstHandlingRule(vehicleID, ruleIndex)
}
func (HandlingAPI) ResetInstRule(vehicleID, ruleIndex int) error {
	return bridgeResetInstHandlingRule(vehicleID, ruleIndex)
}
func (HandlingAPI) ResetInst(vehicleID int) error { return bridgeResetInstHandling(vehicleID) }

type AdminAPI struct{}

func (AdminAPI) BanIP(ip string)             { bridgeBanIP(ip) }
func (AdminAPI) UnbanIP(ip string) bool       { return bridgeUnbanIP(ip) }
func (AdminAPI) IsIPBanned(ip string) bool    { return bridgeIsIPBanned(ip) }

type RadioAPI struct{}

func (RadioAPI) Add(id int, name, url string, listed bool) error {
	return bridgeAddRadioStream(id, name, url, listed)
}
func (RadioAPI) Remove(id int) error { return bridgeRemoveRadioStream(id) }

type EntityAPI struct{}

func (EntityAPI) Exists(pool EntityPool, id int) bool {
	return bridgeCheckEntityExists(pool, id)
}
