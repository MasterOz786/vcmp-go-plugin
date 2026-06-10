package vcmp

/*
#cgo CFLAGS: -I${SRCDIR}/../include
#include "plugin.h"
#include <stdlib.h>
#include <string.h>

extern PluginFuncs *g_pf;

static uint32_t vcmp_get_server_version(void) {
	if (g_pf && g_pf->GetServerVersion) return g_pf->GetServerVersion();
	return 0;
}
static vcmpError vcmp_get_server_settings(ServerSettings *out) {
	if (g_pf && g_pf->GetServerSettings) return g_pf->GetServerSettings(out);
	return vcmpErrorNullArgument;
}
static uint32_t vcmp_get_number_of_plugins(void) {
	if (g_pf && g_pf->GetNumberOfPlugins) return g_pf->GetNumberOfPlugins();
	return 0;
}
static int32_t vcmp_find_plugin(const char *name) {
	if (g_pf && g_pf->FindPlugin) return g_pf->FindPlugin(name);
	return -1;
}
static vcmpError vcmp_get_last_error(void) {
	if (g_pf && g_pf->GetLastError) return g_pf->GetLastError();
	return vcmpErrorNone;
}
static void vcmp_shutdown_server(void) {
	if (g_pf && g_pf->ShutdownServer) g_pf->ShutdownServer();
}
static vcmpError vcmp_set_max_players(uint32_t maxPlayers) {
	if (g_pf && g_pf->SetMaxPlayers) return g_pf->SetMaxPlayers(maxPlayers);
	return vcmpErrorArgumentOutOfBounds;
}
static uint32_t vcmp_get_max_players(void) {
	if (g_pf && g_pf->GetMaxPlayers) return g_pf->GetMaxPlayers();
	return 0;
}
static vcmpError vcmp_set_server_password(const char *password) {
	if (g_pf && g_pf->SetServerPassword) return g_pf->SetServerPassword(password);
	return vcmpErrorNullArgument;
}
static void vcmp_get_server_password(char *buf, size_t buflen) {
	if (buf && buflen > 0) buf[0] = '\0';
	if (g_pf && g_pf->GetServerPassword && buf && buflen > 0) g_pf->GetServerPassword(buf, buflen);
}
static void vcmp_set_wasted_settings(uint32_t deathTimer, uint32_t fadeTimer, float fadeIn, float fadeOut, uint32_t fadeColour, uint32_t corpseFadeStart, uint32_t corpseFadeTime) {
	if (g_pf && g_pf->SetWastedSettings) g_pf->SetWastedSettings(deathTimer, fadeTimer, fadeIn, fadeOut, fadeColour, corpseFadeStart, corpseFadeTime);
}
static void vcmp_get_wasted_settings(uint32_t *deathTimer, uint32_t *fadeTimer, float *fadeIn, float *fadeOut, uint32_t *fadeColour, uint32_t *corpseFadeStart, uint32_t *corpseFadeTime) {
	if (g_pf && g_pf->GetWastedSettings) g_pf->GetWastedSettings(deathTimer, fadeTimer, fadeIn, fadeOut, fadeColour, corpseFadeStart, corpseFadeTime);
}
static void vcmp_set_time_rate(int32_t rate) { if (g_pf && g_pf->SetTimeRate) g_pf->SetTimeRate(rate); }
static int32_t vcmp_get_time_rate(void) { if (g_pf && g_pf->GetTimeRate) return g_pf->GetTimeRate(); return 0; }
static void vcmp_set_minute(int32_t minute) { if (g_pf && g_pf->SetMinute) g_pf->SetMinute(minute); }
static int32_t vcmp_get_minute(void) { if (g_pf && g_pf->GetMinute) return g_pf->GetMinute(); return 0; }
static void vcmp_set_game_speed(float speed) { if (g_pf && g_pf->SetGameSpeed) g_pf->SetGameSpeed(speed); }
static float vcmp_get_game_speed(void) { if (g_pf && g_pf->GetGameSpeed) return g_pf->GetGameSpeed(); return 0; }
static void vcmp_set_water_level(float level) { if (g_pf && g_pf->SetWaterLevel) g_pf->SetWaterLevel(level); }
static float vcmp_get_water_level(void) { if (g_pf && g_pf->GetWaterLevel) return g_pf->GetWaterLevel(); return 0; }
static void vcmp_set_max_flight_altitude(float height) { if (g_pf && g_pf->SetMaximumFlightAltitude) g_pf->SetMaximumFlightAltitude(height); }
static float vcmp_get_max_flight_altitude(void) { if (g_pf && g_pf->GetMaximumFlightAltitude) return g_pf->GetMaximumFlightAltitude(); return 0; }
static void vcmp_set_vehicles_forced_respawn_height(float height) { if (g_pf && g_pf->SetVehiclesForcedRespawnHeight) g_pf->SetVehiclesForcedRespawnHeight(height); }
static float vcmp_get_vehicles_forced_respawn_height(void) { if (g_pf && g_pf->GetVehiclesForcedRespawnHeight) return g_pf->GetVehiclesForcedRespawnHeight(); return 0; }
static void vcmp_set_fall_timer(uint16_t timeRate) { if (g_pf && g_pf->SetFallTimer) g_pf->SetFallTimer(timeRate); }
static uint16_t vcmp_get_fall_timer(void) { if (g_pf && g_pf->GetFallTimer) return g_pf->GetFallTimer(); return 0; }
static void vcmp_hide_map_object(int32_t modelId, int16_t tenthX, int16_t tenthY, int16_t tenthZ) {
	if (g_pf && g_pf->HideMapObject) g_pf->HideMapObject(modelId, tenthX, tenthY, tenthZ);
}
static void vcmp_show_map_object(int32_t modelId, int16_t tenthX, int16_t tenthY, int16_t tenthZ) {
	if (g_pf && g_pf->ShowMapObject) g_pf->ShowMapObject(modelId, tenthX, tenthY, tenthZ);
}
static void vcmp_show_all_map_objects(void) { if (g_pf && g_pf->ShowAllMapObjects) g_pf->ShowAllMapObjects(); }
static vcmpError vcmp_reset_weapon_data_value(int32_t weaponId, int32_t fieldId) {
	if (g_pf && g_pf->ResetWeaponDataValue) return g_pf->ResetWeaponDataValue(weaponId, fieldId);
	return vcmpErrorArgumentOutOfBounds;
}
static uint8_t vcmp_is_weapon_data_modified(int32_t weaponId, int32_t fieldId) {
	if (g_pf && g_pf->IsWeaponDataValueModified) return g_pf->IsWeaponDataValueModified(weaponId, fieldId);
	return 0;
}
static vcmpError vcmp_reset_weapon_data(int32_t weaponId) {
	if (g_pf && g_pf->ResetWeaponData) return g_pf->ResetWeaponData(weaponId);
	return vcmpErrorArgumentOutOfBounds;
}
static void vcmp_reset_all_weapon_data(void) { if (g_pf && g_pf->ResetAllWeaponData) g_pf->ResetAllWeaponData(); }
static vcmpError vcmp_get_key_bind_data(int32_t bindId, uint8_t *onRelease, int32_t *k1, int32_t *k2, int32_t *k3) {
	if (g_pf && g_pf->GetKeyBindData) return g_pf->GetKeyBindData(bindId, onRelease, k1, k2, k3);
	return vcmpErrorNoSuchEntity;
}
static void vcmp_get_player_uid2(int32_t playerId, char *buf, size_t buflen) {
	if (buf && buflen > 0) buf[0] = '\0';
	if (g_pf && g_pf->GetPlayerUID2 && buf && buflen > 0) g_pf->GetPlayerUID2(playerId, buf, buflen);
}
static void vcmp_ban_ip(char *ip) { if (g_pf && g_pf->BanIP) g_pf->BanIP(ip); }
static uint8_t vcmp_unban_ip(char *ip) { if (g_pf && g_pf->UnbanIP) return g_pf->UnbanIP(ip); return 0; }
static uint8_t vcmp_is_ip_banned(char *ip) { if (g_pf && g_pf->IsIPBanned) return g_pf->IsIPBanned(ip); return 0; }
static uint32_t vcmp_get_player_key(int32_t playerId) {
	if (g_pf && g_pf->GetPlayerKey) return g_pf->GetPlayerKey(playerId);
	return 0;
}
static double vcmp_get_player_fps(int32_t playerId) {
	if (g_pf && g_pf->GetPlayerFPS) return g_pf->GetPlayerFPS(playerId);
	return 0;
}
static vcmpError vcmp_set_player_secondary_world(int32_t playerId, int32_t world) {
	if (g_pf && g_pf->SetPlayerSecondaryWorld) return g_pf->SetPlayerSecondaryWorld(playerId, world);
	return vcmpErrorNoSuchEntity;
}
static int32_t vcmp_get_player_secondary_world(int32_t playerId) {
	if (g_pf && g_pf->GetPlayerSecondaryWorld) return g_pf->GetPlayerSecondaryWorld(playerId);
	return 0;
}
static int32_t vcmp_get_player_unique_world(int32_t playerId) {
	if (g_pf && g_pf->GetPlayerUniqueWorld) return g_pf->GetPlayerUniqueWorld(playerId);
	return 0;
}
static uint8_t vcmp_is_player_world_compatible(int32_t playerId, int32_t world) {
	if (g_pf && g_pf->IsPlayerWorldCompatible) return g_pf->IsPlayerWorldCompatible(playerId, world);
	return 0;
}
static uint8_t vcmp_is_player_typing(int32_t playerId) {
	if (g_pf && g_pf->IsPlayerTyping) return g_pf->IsPlayerTyping(playerId);
	return 0;
}
static vcmpError vcmp_set_player_immunity_flags(int32_t playerId, uint32_t flags) {
	if (g_pf && g_pf->SetPlayerImmunityFlags) return g_pf->SetPlayerImmunityFlags(playerId, flags);
	return vcmpErrorNoSuchEntity;
}
static uint32_t vcmp_get_player_immunity_flags(int32_t playerId) {
	if (g_pf && g_pf->GetPlayerImmunityFlags) return g_pf->GetPlayerImmunityFlags(playerId);
	return 0;
}
static float vcmp_get_player_armour(int32_t playerId) {
	if (g_pf && g_pf->GetPlayerArmour) return g_pf->GetPlayerArmour(playerId);
	return 0;
}
static vcmpError vcmp_set_player_speed(int32_t playerId, float x, float y, float z) {
	if (g_pf && g_pf->SetPlayerSpeed) return g_pf->SetPlayerSpeed(playerId, x, y, z);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_get_player_speed(int32_t playerId, float *x, float *y, float *z) {
	if (g_pf && g_pf->GetPlayerSpeed) return g_pf->GetPlayerSpeed(playerId, x, y, z);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_add_player_speed(int32_t playerId, float x, float y, float z) {
	if (g_pf && g_pf->AddPlayerSpeed) return g_pf->AddPlayerSpeed(playerId, x, y, z);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_set_player_alpha(int32_t playerId, int32_t alpha, uint32_t fadeTime) {
	if (g_pf && g_pf->SetPlayerAlpha) return g_pf->SetPlayerAlpha(playerId, alpha, fadeTime);
	return vcmpErrorNoSuchEntity;
}
static int32_t vcmp_get_player_alpha(int32_t playerId) {
	if (g_pf && g_pf->GetPlayerAlpha) return g_pf->GetPlayerAlpha(playerId);
	return 0;
}
static vcmpError vcmp_get_player_aim_position(int32_t playerId, float *x, float *y, float *z) {
	if (g_pf && g_pf->GetPlayerAimPosition) return g_pf->GetPlayerAimPosition(playerId, x, y, z);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_get_player_aim_direction(int32_t playerId, float *x, float *y, float *z) {
	if (g_pf && g_pf->GetPlayerAimDirection) return g_pf->GetPlayerAimDirection(playerId, x, y, z);
	return vcmpErrorNoSuchEntity;
}
static uint8_t vcmp_is_player_on_fire(int32_t playerId) {
	if (g_pf && g_pf->IsPlayerOnFire) return g_pf->IsPlayerOnFire(playerId);
	return 0;
}
static uint8_t vcmp_is_player_crouching(int32_t playerId) {
	if (g_pf && g_pf->IsPlayerCrouching) return g_pf->IsPlayerCrouching(playerId);
	return 0;
}
static int32_t vcmp_get_player_action(int32_t playerId) {
	if (g_pf && g_pf->GetPlayerAction) return g_pf->GetPlayerAction(playerId);
	return 0;
}
static uint32_t vcmp_get_player_game_keys(int32_t playerId) {
	if (g_pf && g_pf->GetPlayerGameKeys) return g_pf->GetPlayerGameKeys(playerId);
	return 0;
}
static vcmpError vcmp_set_player_animation(int32_t playerId, int32_t groupId, int32_t animationId) {
	if (g_pf && g_pf->SetPlayerAnimation) return g_pf->SetPlayerAnimation(playerId, groupId, animationId);
	return vcmpErrorNoSuchEntity;
}
static int32_t vcmp_get_player_standing_on_vehicle(int32_t playerId) {
	if (g_pf && g_pf->GetPlayerStandingOnVehicle) return g_pf->GetPlayerStandingOnVehicle(playerId);
	return -1;
}
static int32_t vcmp_get_player_standing_on_object(int32_t playerId) {
	if (g_pf && g_pf->GetPlayerStandingOnObject) return g_pf->GetPlayerStandingOnObject(playerId);
	return -1;
}
static uint8_t vcmp_is_player_away(int32_t playerId) {
	if (g_pf && g_pf->IsPlayerAway) return g_pf->IsPlayerAway(playerId);
	return 0;
}
static vcmpError vcmp_redirect_player_to_server(int32_t playerId, const char *ip, uint32_t port, const char *nick, const char *serverPassword, const char *userPassword) {
	if (g_pf && g_pf->RedirectPlayerToServer) return g_pf->RedirectPlayerToServer(playerId, ip, port, nick, serverPassword, userPassword);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_set_vehicle_rotation(int32_t vehicleId, float x, float y, float z, float w) {
	if (g_pf && g_pf->SetVehicleRotation) return g_pf->SetVehicleRotation(vehicleId, x, y, z, w);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_get_vehicle_rotation(int32_t vehicleId, float *x, float *y, float *z, float *w) {
	if (g_pf && g_pf->GetVehicleRotation) return g_pf->GetVehicleRotation(vehicleId, x, y, z, w);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_set_vehicle_turn_speed(int32_t vehicleId, float x, float y, float z, uint8_t add, uint8_t relative) {
	if (g_pf && g_pf->SetVehicleTurnSpeed) return g_pf->SetVehicleTurnSpeed(vehicleId, x, y, z, add, relative);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_get_vehicle_turn_speed(int32_t vehicleId, float *x, float *y, float *z, uint8_t relative) {
	if (g_pf && g_pf->GetVehicleTurnSpeed) return g_pf->GetVehicleTurnSpeed(vehicleId, x, y, z, relative);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_get_vehicle_spawn_position(int32_t vehicleId, float *x, float *y, float *z) {
	if (g_pf && g_pf->GetVehicleSpawnPosition) return g_pf->GetVehicleSpawnPosition(vehicleId, x, y, z);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_set_vehicle_spawn_rotation_euler(int32_t vehicleId, float x, float y, float z) {
	if (g_pf && g_pf->SetVehicleSpawnRotationEuler) return g_pf->SetVehicleSpawnRotationEuler(vehicleId, x, y, z);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_get_vehicle_spawn_rotation_euler(int32_t vehicleId, float *x, float *y, float *z) {
	if (g_pf && g_pf->GetVehicleSpawnRotationEuler) return g_pf->GetVehicleSpawnRotationEuler(vehicleId, x, y, z);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_set_vehicle_spawn_rotation(int32_t vehicleId, float x, float y, float z, float w) {
	if (g_pf && g_pf->SetVehicleSpawnRotation) return g_pf->SetVehicleSpawnRotation(vehicleId, x, y, z, w);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_get_vehicle_spawn_rotation(int32_t vehicleId, float *x, float *y, float *z, float *w) {
	if (g_pf && g_pf->GetVehicleSpawnRotation) return g_pf->GetVehicleSpawnRotation(vehicleId, x, y, z, w);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_set_vehicle_idle_respawn_timer(int32_t vehicleId, uint32_t millis) {
	if (g_pf && g_pf->SetVehicleIdleRespawnTimer) return g_pf->SetVehicleIdleRespawnTimer(vehicleId, millis);
	return vcmpErrorNoSuchEntity;
}
static uint32_t vcmp_get_vehicle_idle_respawn_timer(int32_t vehicleId) {
	if (g_pf && g_pf->GetVehicleIdleRespawnTimer) return g_pf->GetVehicleIdleRespawnTimer(vehicleId);
	return 0;
}
static vcmpError vcmp_set_vehicle_damage_data(int32_t vehicleId, uint32_t damageData) {
	if (g_pf && g_pf->SetVehicleDamageData) return g_pf->SetVehicleDamageData(vehicleId, damageData);
	return vcmpErrorNoSuchEntity;
}
static uint32_t vcmp_get_vehicle_damage_data(int32_t vehicleId) {
	if (g_pf && g_pf->GetVehicleDamageData) return g_pf->GetVehicleDamageData(vehicleId);
	return 0;
}
static vcmpError vcmp_get_vehicle_turret_rotation(int32_t vehicleId, float *h, float *v) {
	if (g_pf && g_pf->GetVehicleTurretRotation) return g_pf->GetVehicleTurretRotation(vehicleId, h, v);
	return vcmpErrorNoSuchEntity;
}
static void vcmp_reset_all_vehicle_handlings(void) { if (g_pf && g_pf->ResetAllVehicleHandlings) g_pf->ResetAllVehicleHandlings(); }
static uint8_t vcmp_exists_handling_rule(int32_t modelIndex, int32_t ruleIndex) {
	if (g_pf && g_pf->ExistsHandlingRule) return g_pf->ExistsHandlingRule(modelIndex, ruleIndex);
	return 0;
}
static vcmpError vcmp_set_handling_rule(int32_t modelIndex, int32_t ruleIndex, double value) {
	if (g_pf && g_pf->SetHandlingRule) return g_pf->SetHandlingRule(modelIndex, ruleIndex, value);
	return vcmpErrorArgumentOutOfBounds;
}
static double vcmp_get_handling_rule(int32_t modelIndex, int32_t ruleIndex) {
	if (g_pf && g_pf->GetHandlingRule) return g_pf->GetHandlingRule(modelIndex, ruleIndex);
	return 0;
}
static vcmpError vcmp_reset_handling_rule(int32_t modelIndex, int32_t ruleIndex) {
	if (g_pf && g_pf->ResetHandlingRule) return g_pf->ResetHandlingRule(modelIndex, ruleIndex);
	return vcmpErrorArgumentOutOfBounds;
}
static vcmpError vcmp_reset_handling(int32_t modelIndex) {
	if (g_pf && g_pf->ResetHandling) return g_pf->ResetHandling(modelIndex);
	return vcmpErrorArgumentOutOfBounds;
}
static uint8_t vcmp_exists_inst_handling_rule(int32_t vehicleId, int32_t ruleIndex) {
	if (g_pf && g_pf->ExistsInstHandlingRule) return g_pf->ExistsInstHandlingRule(vehicleId, ruleIndex);
	return 0;
}
static vcmpError vcmp_set_inst_handling_rule(int32_t vehicleId, int32_t ruleIndex, double value) {
	if (g_pf && g_pf->SetInstHandlingRule) return g_pf->SetInstHandlingRule(vehicleId, ruleIndex, value);
	return vcmpErrorNoSuchEntity;
}
static double vcmp_get_inst_handling_rule(int32_t vehicleId, int32_t ruleIndex) {
	if (g_pf && g_pf->GetInstHandlingRule) return g_pf->GetInstHandlingRule(vehicleId, ruleIndex);
	return 0;
}
static vcmpError vcmp_reset_inst_handling_rule(int32_t vehicleId, int32_t ruleIndex) {
	if (g_pf && g_pf->ResetInstHandlingRule) return g_pf->ResetInstHandlingRule(vehicleId, ruleIndex);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_reset_inst_handling(int32_t vehicleId) {
	if (g_pf && g_pf->ResetInstHandling) return g_pf->ResetInstHandling(vehicleId);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_set_pickup_auto_timer(int32_t pickupId, uint32_t durationMillis) {
	if (g_pf && g_pf->SetPickupAutoTimer) return g_pf->SetPickupAutoTimer(pickupId, durationMillis);
	return vcmpErrorNoSuchEntity;
}
static uint32_t vcmp_get_pickup_auto_timer(int32_t pickupId) {
	if (g_pf && g_pf->GetPickupAutoTimer) return g_pf->GetPickupAutoTimer(pickupId);
	return 0;
}
static vcmpError vcmp_move_object_by(int32_t objectId, float x, float y, float z, uint32_t duration) {
	if (g_pf && g_pf->MoveObjectBy) return g_pf->MoveObjectBy(objectId, x, y, z, duration);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_rotate_object_by_euler(int32_t objectId, float x, float y, float z, uint32_t duration) {
	if (g_pf && g_pf->RotateObjectByEuler) return g_pf->RotateObjectByEuler(objectId, x, y, z, duration);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_rotate_object_to(int32_t objectId, float x, float y, float z, float w, uint32_t duration) {
	if (g_pf && g_pf->RotateObjectTo) return g_pf->RotateObjectTo(objectId, x, y, z, w, duration);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_rotate_object_by(int32_t objectId, float x, float y, float z, float w, uint32_t duration) {
	if (g_pf && g_pf->RotateObjectBy) return g_pf->RotateObjectBy(objectId, x, y, z, w, duration);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_get_object_rotation(int32_t objectId, float *x, float *y, float *z, float *w) {
	if (g_pf && g_pf->GetObjectRotation) return g_pf->GetObjectRotation(objectId, x, y, z, w);
	return vcmpErrorNoSuchEntity;
}
*/
import "C"

import (
	"unsafe"
)

type ServerSettings struct {
	ServerName string
	MaxPlayers uint32
	Port       uint32
	Flags      uint32
}

type WastedSettings struct {
	DeathTimer      uint32
	FadeTimer       uint32
	FadeInSpeed     float32
	FadeOutSpeed    float32
	FadeColour      uint32
	CorpseFadeStart uint32
	CorpseFadeTime  uint32
}

type KeyBindData struct {
	OnRelease bool
	KeyOne    int
	KeyTwo    int
	KeyThree  int
}

type MapObjectCoord struct {
	Model int
	TenthX, TenthY, TenthZ int16
}

func bridgeServerVersion() uint32 { return uint32(C.vcmp_get_server_version()) }

func bridgeServerSettings() (ServerSettings, error) {
	var s C.ServerSettings
	s.structSize = C.uint32_t(C.sizeof_ServerSettings)
	if err := bridgeError(C.vcmp_get_server_settings(&s)); err != nil {
		return ServerSettings{}, err
	}
	return ServerSettings{
		ServerName: C.GoString(&s.serverName[0]),
		MaxPlayers: uint32(s.maxPlayers),
		Port:       uint32(s.port),
		Flags:      uint32(s.flags),
	}, nil
}

func bridgePluginCount() uint32 { return uint32(C.vcmp_get_number_of_plugins()) }

func bridgeFindPlugin(name string) int {
	c := cString(name)
	defer freeCString(c)
	return int(C.vcmp_find_plugin(c))
}

func bridgeLastError() error { return bridgeError(C.vcmp_get_last_error()) }

func bridgeShutdownServer() { C.vcmp_shutdown_server() }

func bridgeSetMaxPlayers(max uint32) error {
	return bridgeError(C.vcmp_set_max_players(C.uint32_t(max)))
}

func bridgeGetMaxPlayers() uint32 { return uint32(C.vcmp_get_max_players()) }

func bridgeSetServerPassword(password string) error {
	c := cString(password)
	defer freeCString(c)
	return bridgeError(C.vcmp_set_server_password(c))
}

func bridgeGetServerPassword() string {
	buf := (*[128]C.char)(C.malloc(128))
	defer C.free(unsafe.Pointer(buf))
	C.vcmp_get_server_password(&buf[0], 128)
	return C.GoString(&buf[0])
}

func bridgeSetWastedSettings(s WastedSettings) {
	C.vcmp_set_wasted_settings(
		C.uint32_t(s.DeathTimer), C.uint32_t(s.FadeTimer),
		C.float(s.FadeInSpeed), C.float(s.FadeOutSpeed),
		C.uint32_t(s.FadeColour), C.uint32_t(s.CorpseFadeStart), C.uint32_t(s.CorpseFadeTime),
	)
}

func bridgeGetWastedSettings() WastedSettings {
	var death, fade, colour, corpseStart, corpseTime C.uint32_t
	var fadeIn, fadeOut C.float
	C.vcmp_get_wasted_settings(&death, &fade, &fadeIn, &fadeOut, &colour, &corpseStart, &corpseTime)
	return WastedSettings{
		DeathTimer: uint32(death), FadeTimer: uint32(fade),
		FadeInSpeed: float32(fadeIn), FadeOutSpeed: float32(fadeOut),
		FadeColour: uint32(colour), CorpseFadeStart: uint32(corpseStart), CorpseFadeTime: uint32(corpseTime),
	}
}

func bridgeSetTimeRate(rate int) { C.vcmp_set_time_rate(C.int32_t(rate)) }
func bridgeGetTimeRate() int     { return int(C.vcmp_get_time_rate()) }
func bridgeSetMinute(minute int) { C.vcmp_set_minute(C.int32_t(minute)) }
func bridgeGetMinute() int       { return int(C.vcmp_get_minute()) }
func bridgeSetGameSpeed(s float32) { C.vcmp_set_game_speed(C.float(s)) }
func bridgeGetGameSpeed() float32  { return float32(C.vcmp_get_game_speed()) }
func bridgeSetWaterLevel(l float32) { C.vcmp_set_water_level(C.float(l)) }
func bridgeGetWaterLevel() float32  { return float32(C.vcmp_get_water_level()) }
func bridgeSetMaxFlightAltitude(h float32) { C.vcmp_set_max_flight_altitude(C.float(h)) }
func bridgeGetMaxFlightAltitude() float32    { return float32(C.vcmp_get_max_flight_altitude()) }
func bridgeSetVehiclesForcedRespawnHeight(h float32) { C.vcmp_set_vehicles_forced_respawn_height(C.float(h)) }
func bridgeGetVehiclesForcedRespawnHeight() float32  { return float32(C.vcmp_get_vehicles_forced_respawn_height()) }
func bridgeSetFallTimer(t uint16) { C.vcmp_set_fall_timer(C.uint16_t(t)) }
func bridgeGetFallTimer() uint16    { return uint16(C.vcmp_get_fall_timer()) }

func bridgeHideMapObject(coord MapObjectCoord) {
	C.vcmp_hide_map_object(C.int32_t(coord.Model), C.int16_t(coord.TenthX), C.int16_t(coord.TenthY), C.int16_t(coord.TenthZ))
}
func bridgeShowMapObject(coord MapObjectCoord) {
	C.vcmp_show_map_object(C.int32_t(coord.Model), C.int16_t(coord.TenthX), C.int16_t(coord.TenthY), C.int16_t(coord.TenthZ))
}
func bridgeShowAllMapObjects() { C.vcmp_show_all_map_objects() }

func bridgeResetWeaponDataValue(weaponID, fieldID int) error {
	return bridgeError(C.vcmp_reset_weapon_data_value(C.int32_t(weaponID), C.int32_t(fieldID)))
}
func bridgeIsWeaponDataModified(weaponID, fieldID int) bool {
	return C.vcmp_is_weapon_data_modified(C.int32_t(weaponID), C.int32_t(fieldID)) != 0
}
func bridgeResetWeaponData(weaponID int) error {
	return bridgeError(C.vcmp_reset_weapon_data(C.int32_t(weaponID)))
}
func bridgeResetAllWeaponData() { C.vcmp_reset_all_weapon_data() }

func bridgeGetKeyBindData(bindID int) (KeyBindData, error) {
	var onRelease C.uint8_t
	var k1, k2, k3 C.int32_t
	if err := bridgeError(C.vcmp_get_key_bind_data(C.int32_t(bindID), &onRelease, &k1, &k2, &k3)); err != nil {
		return KeyBindData{}, err
	}
	return KeyBindData{OnRelease: onRelease != 0, KeyOne: int(k1), KeyTwo: int(k2), KeyThree: int(k3)}, nil
}

func bridgePlayerUID2(playerID int) string {
	buf := (*[128]C.char)(C.malloc(128))
	defer C.free(unsafe.Pointer(buf))
	C.vcmp_get_player_uid2(C.int32_t(playerID), &buf[0], 128)
	return C.GoString(&buf[0])
}

func bridgeBanIP(ip string) {
	c := cString(ip)
	defer freeCString(c)
	C.vcmp_ban_ip(c)
}
func bridgeUnbanIP(ip string) bool {
	c := cString(ip)
	defer freeCString(c)
	return C.vcmp_unban_ip(c) != 0
}
func bridgeIsIPBanned(ip string) bool {
	c := cString(ip)
	defer freeCString(c)
	return C.vcmp_is_ip_banned(c) != 0
}

func bridgeGetPlayerKey(playerID int) uint32 { return uint32(C.vcmp_get_player_key(C.int32_t(playerID))) }
func bridgeGetPlayerFPS(playerID int) float64 { return float64(C.vcmp_get_player_fps(C.int32_t(playerID))) }

func bridgeSetPlayerSecondaryWorld(playerID, world int) error {
	return bridgeError(C.vcmp_set_player_secondary_world(C.int32_t(playerID), C.int32_t(world)))
}
func bridgeGetPlayerSecondaryWorld(playerID int) int {
	return int(C.vcmp_get_player_secondary_world(C.int32_t(playerID)))
}
func bridgeGetPlayerUniqueWorld(playerID int) int {
	return int(C.vcmp_get_player_unique_world(C.int32_t(playerID)))
}
func bridgeIsPlayerWorldCompatible(playerID, world int) bool {
	return C.vcmp_is_player_world_compatible(C.int32_t(playerID), C.int32_t(world)) != 0
}
func bridgeIsPlayerTyping(playerID int) bool { return C.vcmp_is_player_typing(C.int32_t(playerID)) != 0 }

func bridgeSetPlayerImmunityFlags(playerID int, flags uint32) error {
	return bridgeError(C.vcmp_set_player_immunity_flags(C.int32_t(playerID), C.uint32_t(flags)))
}
func bridgeGetPlayerImmunityFlags(playerID int) uint32 {
	return uint32(C.vcmp_get_player_immunity_flags(C.int32_t(playerID)))
}

func bridgeSetPlayerSpeed(playerID int, speed Vec3) error {
	return bridgeError(C.vcmp_set_player_speed(C.int32_t(playerID), C.float(speed.X), C.float(speed.Y), C.float(speed.Z)))
}
func bridgeGetPlayerSpeed(playerID int) (Vec3, error) {
	var x, y, z C.float
	if err := bridgeError(C.vcmp_get_player_speed(C.int32_t(playerID), &x, &y, &z)); err != nil {
		return Vec3{}, err
	}
	return Vec3{X: float32(x), Y: float32(y), Z: float32(z)}, nil
}
func bridgeAddPlayerSpeed(playerID int, delta Vec3) error {
	return bridgeError(C.vcmp_add_player_speed(C.int32_t(playerID), C.float(delta.X), C.float(delta.Y), C.float(delta.Z)))
}
func bridgeSetPlayerAlpha(playerID, alpha int, fadeMs uint32) error {
	return bridgeError(C.vcmp_set_player_alpha(C.int32_t(playerID), C.int32_t(alpha), C.uint32_t(fadeMs)))
}
func bridgeGetPlayerAlpha(playerID int) int { return int(C.vcmp_get_player_alpha(C.int32_t(playerID))) }

func bridgeGetPlayerAimPosition(playerID int) (Vec3, error) {
	var x, y, z C.float
	if err := bridgeError(C.vcmp_get_player_aim_position(C.int32_t(playerID), &x, &y, &z)); err != nil {
		return Vec3{}, err
	}
	return Vec3{X: float32(x), Y: float32(y), Z: float32(z)}, nil
}
func bridgeGetPlayerAimDirection(playerID int) (Vec3, error) {
	var x, y, z C.float
	if err := bridgeError(C.vcmp_get_player_aim_direction(C.int32_t(playerID), &x, &y, &z)); err != nil {
		return Vec3{}, err
	}
	return Vec3{X: float32(x), Y: float32(y), Z: float32(z)}, nil
}

func bridgeIsPlayerOnFire(playerID int) bool    { return C.vcmp_is_player_on_fire(C.int32_t(playerID)) != 0 }
func bridgeIsPlayerCrouching(playerID int) bool  { return C.vcmp_is_player_crouching(C.int32_t(playerID)) != 0 }
func bridgeGetPlayerAction(playerID int) int     { return int(C.vcmp_get_player_action(C.int32_t(playerID))) }
func bridgeGetPlayerGameKeys(playerID int) uint32 { return uint32(C.vcmp_get_player_game_keys(C.int32_t(playerID))) }

func bridgeSetPlayerAnimation(playerID, groupID, animationID int) error {
	return bridgeError(C.vcmp_set_player_animation(C.int32_t(playerID), C.int32_t(groupID), C.int32_t(animationID)))
}
func bridgeGetPlayerStandingOnVehicle(playerID int) int {
	return int(C.vcmp_get_player_standing_on_vehicle(C.int32_t(playerID)))
}
func bridgeGetPlayerStandingOnObject(playerID int) int {
	return int(C.vcmp_get_player_standing_on_object(C.int32_t(playerID)))
}
func bridgeIsPlayerAway(playerID int) bool { return C.vcmp_is_player_away(C.int32_t(playerID)) != 0 }

func bridgeRedirectPlayerToServer(playerID int, ip string, port uint32, nick, serverPassword, userPassword string) error {
	cIP, cNick, cSP, cUP := cString(ip), cString(nick), cString(serverPassword), cString(userPassword)
	defer freeCString(cIP)
	defer freeCString(cNick)
	defer freeCString(cSP)
	defer freeCString(cUP)
	return bridgeError(C.vcmp_redirect_player_to_server(C.int32_t(playerID), cIP, C.uint32_t(port), cNick, cSP, cUP))
}

func bridgeSetVehicleRotation(vehicleID int, rot Quat) error {
	return bridgeError(C.vcmp_set_vehicle_rotation(C.int32_t(vehicleID), C.float(rot.X), C.float(rot.Y), C.float(rot.Z), C.float(rot.W)))
}
func bridgeGetVehicleRotation(vehicleID int) (Quat, error) {
	var x, y, z, w C.float
	if err := bridgeError(C.vcmp_get_vehicle_rotation(C.int32_t(vehicleID), &x, &y, &z, &w)); err != nil {
		return Quat{}, err
	}
	return Quat{X: float32(x), Y: float32(y), Z: float32(z), W: float32(w)}, nil
}
func bridgeSetVehicleTurnSpeed(vehicleID int, speed Vec3, add, relative bool) error {
	return bridgeError(C.vcmp_set_vehicle_turn_speed(C.int32_t(vehicleID), C.float(speed.X), C.float(speed.Y), C.float(speed.Z), boolToU8(add), boolToU8(relative)))
}
func bridgeGetVehicleTurnSpeed(vehicleID int, relative bool) (Vec3, error) {
	var x, y, z C.float
	if err := bridgeError(C.vcmp_get_vehicle_turn_speed(C.int32_t(vehicleID), &x, &y, &z, boolToU8(relative))); err != nil {
		return Vec3{}, err
	}
	return Vec3{X: float32(x), Y: float32(y), Z: float32(z)}, nil
}
func bridgeGetVehicleSpawnPosition(vehicleID int) (Vec3, error) {
	var x, y, z C.float
	if err := bridgeError(C.vcmp_get_vehicle_spawn_position(C.int32_t(vehicleID), &x, &y, &z)); err != nil {
		return Vec3{}, err
	}
	return Vec3{X: float32(x), Y: float32(y), Z: float32(z)}, nil
}
func bridgeSetVehicleSpawnRotationEuler(vehicleID int, rot Vec3) error {
	return bridgeError(C.vcmp_set_vehicle_spawn_rotation_euler(C.int32_t(vehicleID), C.float(rot.X), C.float(rot.Y), C.float(rot.Z)))
}
func bridgeSetVehicleSpawnRotation(vehicleID int, rot Quat) error {
	return bridgeError(C.vcmp_set_vehicle_spawn_rotation(C.int32_t(vehicleID), C.float(rot.X), C.float(rot.Y), C.float(rot.Z), C.float(rot.W)))
}
func bridgeGetVehicleSpawnRotation(vehicleID int) (Quat, error) {
	var x, y, z, w C.float
	if err := bridgeError(C.vcmp_get_vehicle_spawn_rotation(C.int32_t(vehicleID), &x, &y, &z, &w)); err != nil {
		return Quat{}, err
	}
	return Quat{X: float32(x), Y: float32(y), Z: float32(z), W: float32(w)}, nil
}

func bridgeGetVehicleSpawnRotationEuler(vehicleID int) (Vec3, error) {
	var x, y, z C.float
	if err := bridgeError(C.vcmp_get_vehicle_spawn_rotation_euler(C.int32_t(vehicleID), &x, &y, &z)); err != nil {
		return Vec3{}, err
	}
	return Vec3{X: float32(x), Y: float32(y), Z: float32(z)}, nil
}
func bridgeSetVehicleIdleRespawnTimer(vehicleID int, millis uint32) error {
	return bridgeError(C.vcmp_set_vehicle_idle_respawn_timer(C.int32_t(vehicleID), C.uint32_t(millis)))
}
func bridgeGetVehicleIdleRespawnTimer(vehicleID int) uint32 {
	return uint32(C.vcmp_get_vehicle_idle_respawn_timer(C.int32_t(vehicleID)))
}
func bridgeSetVehicleDamageData(vehicleID int, data uint32) error {
	return bridgeError(C.vcmp_set_vehicle_damage_data(C.int32_t(vehicleID), C.uint32_t(data)))
}
func bridgeGetVehicleDamageData(vehicleID int) uint32 {
	return uint32(C.vcmp_get_vehicle_damage_data(C.int32_t(vehicleID)))
}
func bridgeGetVehicleTurretRotation(vehicleID int) (horizontal, vertical float32, err error) {
	var h, v C.float
	if e := bridgeError(C.vcmp_get_vehicle_turret_rotation(C.int32_t(vehicleID), &h, &v)); e != nil {
		return 0, 0, e
	}
	return float32(h), float32(v), nil
}

func bridgeResetAllVehicleHandlings() { C.vcmp_reset_all_vehicle_handlings() }
func bridgeExistsHandlingRule(modelIndex, ruleIndex int) bool {
	return C.vcmp_exists_handling_rule(C.int32_t(modelIndex), C.int32_t(ruleIndex)) != 0
}
func bridgeSetHandlingRule(modelIndex, ruleIndex int, value float64) error {
	return bridgeError(C.vcmp_set_handling_rule(C.int32_t(modelIndex), C.int32_t(ruleIndex), C.double(value)))
}
func bridgeGetHandlingRule(modelIndex, ruleIndex int) float64 {
	return float64(C.vcmp_get_handling_rule(C.int32_t(modelIndex), C.int32_t(ruleIndex)))
}
func bridgeResetHandlingRule(modelIndex, ruleIndex int) error {
	return bridgeError(C.vcmp_reset_handling_rule(C.int32_t(modelIndex), C.int32_t(ruleIndex)))
}
func bridgeResetHandling(modelIndex int) error {
	return bridgeError(C.vcmp_reset_handling(C.int32_t(modelIndex)))
}
func bridgeExistsInstHandlingRule(vehicleID, ruleIndex int) bool {
	return C.vcmp_exists_inst_handling_rule(C.int32_t(vehicleID), C.int32_t(ruleIndex)) != 0
}
func bridgeSetInstHandlingRule(vehicleID, ruleIndex int, value float64) error {
	return bridgeError(C.vcmp_set_inst_handling_rule(C.int32_t(vehicleID), C.int32_t(ruleIndex), C.double(value)))
}
func bridgeGetInstHandlingRule(vehicleID, ruleIndex int) float64 {
	return float64(C.vcmp_get_inst_handling_rule(C.int32_t(vehicleID), C.int32_t(ruleIndex)))
}
func bridgeResetInstHandlingRule(vehicleID, ruleIndex int) error {
	return bridgeError(C.vcmp_reset_inst_handling_rule(C.int32_t(vehicleID), C.int32_t(ruleIndex)))
}
func bridgeResetInstHandling(vehicleID int) error {
	return bridgeError(C.vcmp_reset_inst_handling(C.int32_t(vehicleID)))
}

func bridgeSetPickupAutoTimer(pickupID int, durationMs uint32) error {
	return bridgeError(C.vcmp_set_pickup_auto_timer(C.int32_t(pickupID), C.uint32_t(durationMs)))
}
func bridgeGetPickupAutoTimer(pickupID int) uint32 {
	return uint32(C.vcmp_get_pickup_auto_timer(C.int32_t(pickupID)))
}

func bridgeMoveObjectBy(objectID int, offset Vec3, durationMs uint32) error {
	return bridgeError(C.vcmp_move_object_by(C.int32_t(objectID), C.float(offset.X), C.float(offset.Y), C.float(offset.Z), C.uint32_t(durationMs)))
}
func bridgeRotateObjectTo(objectID int, rot Quat, durationMs uint32) error {
	return bridgeError(C.vcmp_rotate_object_to(C.int32_t(objectID), C.float(rot.X), C.float(rot.Y), C.float(rot.Z), C.float(rot.W), C.uint32_t(durationMs)))
}
func bridgeRotateObjectBy(objectID int, rot Quat, durationMs uint32) error {
	return bridgeError(C.vcmp_rotate_object_by(C.int32_t(objectID), C.float(rot.X), C.float(rot.Y), C.float(rot.Z), C.float(rot.W), C.uint32_t(durationMs)))
}
func bridgeRotateObjectByEuler(objectID int, rot Vec3, durationMs uint32) error {
	return bridgeError(C.vcmp_rotate_object_by_euler(C.int32_t(objectID), C.float(rot.X), C.float(rot.Y), C.float(rot.Z), C.uint32_t(durationMs)))
}
func bridgeGetObjectRotation(objectID int) (Quat, error) {
	var x, y, z, w C.float
	if err := bridgeError(C.vcmp_get_object_rotation(C.int32_t(objectID), &x, &y, &z, &w)); err != nil {
		return Quat{}, err
	}
	return Quat{X: float32(x), Y: float32(y), Z: float32(z), W: float32(w)}, nil
}

func bridgeSetVehicle3DArrowForPlayer(vehicleID, targetPlayerID int, enabled bool) error {
	return errPluginAPINotAvailable
}
func bridgeGetVehicle3DArrowForPlayer(vehicleID, targetPlayerID int) bool {
	return false
}
func bridgeSetPlayer3DArrowForPlayer(playerID, targetPlayerID int, enabled bool) error {
	return errPluginAPINotAvailable
}
func bridgeGetPlayer3DArrowForPlayer(playerID, targetPlayerID int) bool {
	return false
}
