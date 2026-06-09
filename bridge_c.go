package main

/*
#cgo CFLAGS: -I${SRCDIR}/include
#include "plugin.h"
#include <stdlib.h>
#include <string.h>

static PluginFuncs *g_pf;

static void vcmp_set_funcs(PluginFuncs *pf) { g_pf = pf; }

static void vcmp_log_msg(const char *msg) {
	if (g_pf && g_pf->LogMessage) g_pf->LogMessage("%s", msg);
}

static uint64_t vcmp_get_time(void) {
	return g_pf && g_pf->GetTime ? g_pf->GetTime() : 0;
}

static void vcmp_set_server_name(const char *name) {
	if (g_pf && g_pf->SetServerName) g_pf->SetServerName(name);
}

static void vcmp_set_gamemode_text(const char *text) {
	if (g_pf && g_pf->SetGameModeText) g_pf->SetGameModeText(text);
}

static void vcmp_set_server_option(vcmpServerOption option, uint8_t toggle) {
	if (g_pf && g_pf->SetServerOption) g_pf->SetServerOption(option, toggle);
}

static void vcmp_set_world_bounds(float maxX, float minX, float maxY, float minY) {
	if (g_pf && g_pf->SetWorldBounds) g_pf->SetWorldBounds(maxX, minX, maxY, minY);
}

static void vcmp_set_hour(int32_t hour) {
	if (g_pf && g_pf->SetHour) g_pf->SetHour(hour);
}

static void vcmp_set_weather(int32_t weather) {
	if (g_pf && g_pf->SetWeather) g_pf->SetWeather(weather);
}

static void vcmp_set_gravity(float gravity) {
	if (g_pf && g_pf->SetGravity) g_pf->SetGravity(gravity);
}

static void vcmp_create_explosion(int32_t world, int32_t type, float x, float y, float z, int32_t playerId, uint8_t ground) {
	if (g_pf && g_pf->CreateExplosion) g_pf->CreateExplosion(world, type, x, y, z, playerId, ground);
}

static void vcmp_set_spawn_pos(float x, float y, float z) {
	if (g_pf && g_pf->SetSpawnPlayerPosition) g_pf->SetSpawnPlayerPosition(x, y, z);
}

static int32_t vcmp_add_player_class(
	int32_t teamId, uint32_t colour, int32_t modelIndex,
	float x, float y, float z, float angle,
	int32_t w1, int32_t w1a, int32_t w2, int32_t w2a, int32_t w3, int32_t w3a
) {
	if (g_pf && g_pf->AddPlayerClass) {
		return g_pf->AddPlayerClass(teamId, colour, modelIndex, x, y, z, angle, w1, w1a, w2, w2a, w3, w3a);
	}
	return -1;
}

static void vcmp_send_client_message(int32_t playerId, uint32_t colour, const char *msg) {
	if (g_pf && g_pf->SendClientMessage) g_pf->SendClientMessage(playerId, colour, "%s", msg);
}

static void vcmp_send_game_message(int32_t playerId, int32_t type, const char *msg) {
	if (g_pf && g_pf->SendGameMessage) g_pf->SendGameMessage(playerId, type, "%s", msg);
}

static void vcmp_get_player_name(int32_t playerId, char *buf, size_t buflen) {
	if (buf && buflen > 0) buf[0] = '\0';
	if (g_pf && g_pf->GetPlayerName && buf && buflen > 0) {
		g_pf->GetPlayerName(playerId, buf, (int32_t)buflen);
	}
}

static int32_t vcmp_get_player_id_from_name(const char *name) {
	if (g_pf && g_pf->GetPlayerIdFromName) return g_pf->GetPlayerIdFromName(name);
	return -1;
}

static uint8_t vcmp_is_player_connected(int32_t playerId) {
	if (g_pf && g_pf->IsPlayerConnected) return g_pf->IsPlayerConnected(playerId);
	return 0;
}

static uint8_t vcmp_is_player_admin(int32_t playerId) {
	if (g_pf && g_pf->IsPlayerAdmin) return g_pf->IsPlayerAdmin(playerId);
	return 0;
}

static void vcmp_set_player_admin(int32_t playerId, uint8_t toggle) {
	if (g_pf && g_pf->SetPlayerAdmin) g_pf->SetPlayerAdmin(playerId, toggle);
}

static void vcmp_kick_player(int32_t playerId) {
	if (g_pf && g_pf->KickPlayer) g_pf->KickPlayer(playerId);
}

static void vcmp_force_spawn(int32_t playerId) {
	if (g_pf && g_pf->ForcePlayerSpawn) g_pf->ForcePlayerSpawn(playerId);
}

static void vcmp_give_player_money(int32_t playerId, int32_t amount) {
	if (g_pf && g_pf->GivePlayerMoney) g_pf->GivePlayerMoney(playerId, amount);
}

static void vcmp_set_player_score(int32_t playerId, int32_t score) {
	if (g_pf && g_pf->SetPlayerScore) g_pf->SetPlayerScore(playerId, score);
}

static int32_t vcmp_get_player_score(int32_t playerId) {
	if (g_pf && g_pf->GetPlayerScore) return g_pf->GetPlayerScore(playerId);
	return 0;
}

static int32_t vcmp_get_player_money(int32_t playerId) {
	if (g_pf && g_pf->GetPlayerMoney) return g_pf->GetPlayerMoney(playerId);
	return 0;
}

static void vcmp_set_player_health(int32_t playerId, float health) {
	if (g_pf && g_pf->SetPlayerHealth) g_pf->SetPlayerHealth(playerId, health);
}

static float vcmp_get_player_health(int32_t playerId) {
	if (g_pf && g_pf->GetPlayerHealth) return g_pf->GetPlayerHealth(playerId);
	return 0;
}

static void vcmp_set_player_armour(int32_t playerId, float armour) {
	if (g_pf && g_pf->SetPlayerArmour) g_pf->SetPlayerArmour(playerId, armour);
}

static void vcmp_set_player_position(int32_t playerId, float x, float y, float z) {
	if (g_pf && g_pf->SetPlayerPosition) g_pf->SetPlayerPosition(playerId, x, y, z);
}

static void vcmp_get_player_position(int32_t playerId, float *x, float *y, float *z) {
	if (g_pf && g_pf->GetPlayerPosition && x && y && z) g_pf->GetPlayerPosition(playerId, x, y, z);
}

static void vcmp_set_player_skin(int32_t playerId, int32_t skin) {
	if (g_pf && g_pf->SetPlayerSkin) g_pf->SetPlayerSkin(playerId, skin);
}

static void vcmp_set_player_team(int32_t playerId, int32_t team) {
	if (g_pf && g_pf->SetPlayerTeam) g_pf->SetPlayerTeam(playerId, team);
}

static void vcmp_give_player_weapon(int32_t playerId, int32_t weapon, int32_t ammo) {
	if (g_pf && g_pf->GivePlayerWeapon) g_pf->GivePlayerWeapon(playerId, weapon, ammo);
}

static void vcmp_put_player_in_vehicle(int32_t playerId, int32_t vehicleId, int32_t slot, uint8_t makeRoom, uint8_t warp) {
	if (g_pf && g_pf->PutPlayerInVehicle) g_pf->PutPlayerInVehicle(playerId, vehicleId, slot, makeRoom, warp);
}

static void vcmp_remove_player_from_vehicle(int32_t playerId) {
	if (g_pf && g_pf->RemovePlayerFromVehicle) g_pf->RemovePlayerFromVehicle(playerId);
}

static int32_t vcmp_create_vehicle(int32_t model, int32_t world, float x, float y, float z, float angle, int32_t c1, int32_t c2) {
	if (g_pf && g_pf->CreateVehicle) return g_pf->CreateVehicle(model, world, x, y, z, angle, c1, c2);
	return -1;
}

static void vcmp_delete_vehicle(int32_t vehicleId) {
	if (g_pf && g_pf->DeleteVehicle) g_pf->DeleteVehicle(vehicleId);
}

static void vcmp_respawn_vehicle(int32_t vehicleId) {
	if (g_pf && g_pf->RespawnVehicle) g_pf->RespawnVehicle(vehicleId);
}

static void vcmp_set_vehicle_position(int32_t vehicleId, float x, float y, float z, uint8_t removeOccupants) {
	if (g_pf && g_pf->SetVehiclePosition) g_pf->SetVehiclePosition(vehicleId, x, y, z, removeOccupants);
}

static void vcmp_get_vehicle_position(int32_t vehicleId, float *x, float *y, float *z) {
	if (g_pf && g_pf->GetVehiclePosition && x && y && z) g_pf->GetVehiclePosition(vehicleId, x, y, z);
}

static int32_t vcmp_create_object(int32_t model, int32_t world, float x, float y, float z, int32_t alpha) {
	if (g_pf && g_pf->CreateObject) return g_pf->CreateObject(model, world, x, y, z, alpha);
	return -1;
}

static void vcmp_delete_object(int32_t objectId) {
	if (g_pf && g_pf->DeleteObject) g_pf->DeleteObject(objectId);
}

static void vcmp_set_object_position(int32_t objectId, float x, float y, float z) {
	if (g_pf && g_pf->SetObjectPosition) g_pf->SetObjectPosition(objectId, x, y, z);
}

static int32_t vcmp_create_pickup(int32_t model, int32_t world, int32_t qty, float x, float y, float z, int32_t alpha, uint8_t automatic) {
	if (g_pf && g_pf->CreatePickup) return g_pf->CreatePickup(model, world, qty, x, y, z, alpha, automatic);
	return -1;
}

static void vcmp_delete_pickup(int32_t pickupId) {
	if (g_pf && g_pf->DeletePickup) g_pf->DeletePickup(pickupId);
}

static int32_t vcmp_create_checkpoint(int32_t playerId, int32_t world, uint8_t sphere, float x, float y, float z, int32_t r, int32_t g, int32_t b, int32_t a, float radius) {
	if (g_pf && g_pf->CreateCheckPoint) return g_pf->CreateCheckPoint(playerId, world, sphere, x, y, z, r, g, b, a, radius);
	return -1;
}

static void vcmp_delete_checkpoint(int32_t checkpointId) {
	if (g_pf && g_pf->DeleteCheckPoint) g_pf->DeleteCheckPoint(checkpointId);
}

static int32_t vcmp_create_coord_blip(int32_t index, int32_t world, float x, float y, float z, int32_t scale, uint32_t colour, int32_t sprite) {
	if (g_pf && g_pf->CreateCoordBlip) return g_pf->CreateCoordBlip(index, world, x, y, z, scale, colour, sprite);
	return -1;
}

static void vcmp_destroy_coord_blip(int32_t index) {
	if (g_pf && g_pf->DestroyCoordBlip) g_pf->DestroyCoordBlip(index);
}

static int32_t vcmp_get_keybind_unused_slot(void) {
	if (g_pf && g_pf->GetKeyBindUnusedSlot) return g_pf->GetKeyBindUnusedSlot();
	return -1;
}

static void vcmp_register_keybind(int32_t bindId, uint8_t onRelease, int32_t k1, int32_t k2, int32_t k3) {
	if (g_pf && g_pf->RegisterKeyBind) g_pf->RegisterKeyBind(bindId, onRelease, k1, k2, k3);
}

static void vcmp_remove_keybind(int32_t bindId) {
	if (g_pf && g_pf->RemoveKeyBind) g_pf->RemoveKeyBind(bindId);
}
*/
import "C"

import "unsafe"

func bindPluginAPI(pf *C.PluginFuncs) {
	C.vcmp_set_funcs(pf)
}

func cString(s string) *C.char {
	return C.CString(s)
}

func freeCString(p *C.char) {
	C.free(unsafe.Pointer(p))
}

func bridgeLog(msg string) {
	cmsg := cString(msg)
	defer freeCString(cmsg)
	C.vcmp_log_msg(cmsg)
}

func bridgeTime() uint64 { return uint64(C.vcmp_get_time()) }

func bridgeSetServerName(name string) {
	c := cString(name)
	defer freeCString(c)
	C.vcmp_set_server_name(c)
}

func bridgeSetGameModeText(text string) {
	c := cString(text)
	defer freeCString(c)
	C.vcmp_set_gamemode_text(c)
}

func bridgeSetServerOption(option ServerOption, enabled bool) {
	t := C.uint8_t(0)
	if enabled {
		t = 1
	}
	C.vcmp_set_server_option(C.vcmpServerOption(option), t)
}

func bridgeSetWorldBounds(maxX, minX, maxY, minY float32) {
	C.vcmp_set_world_bounds(C.float(maxX), C.float(minX), C.float(maxY), C.float(minY))
}

func bridgeSetHour(hour int) { C.vcmp_set_hour(C.int32_t(hour)) }
func bridgeSetWeather(w int) { C.vcmp_set_weather(C.int32_t(w)) }
func bridgeSetGravity(g float32) { C.vcmp_set_gravity(C.float(g)) }

func bridgeCreateExplosion(world, typ int, pos Vec3, playerID int, ground bool) {
	g := C.uint8_t(0)
	if ground {
		g = 1
	}
	C.vcmp_create_explosion(C.int32_t(world), C.int32_t(typ), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.int32_t(playerID), g)
}

func bridgeSetSpawnPos(pos Vec3) {
	C.vcmp_set_spawn_pos(C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func bridgeAddPlayerClass(teamID int, colour uint32, model int, pos Vec3, angle float32, w []int) int {
	return int(C.vcmp_add_player_class(
		C.int32_t(teamID), C.uint32_t(colour), C.int32_t(model),
		C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(angle),
		C.int32_t(w[0]), C.int32_t(w[1]), C.int32_t(w[2]), C.int32_t(w[3]), C.int32_t(w[4]), C.int32_t(w[5]),
	))
}

func bridgeSendClientMessage(playerID int, colour uint32, msg string) {
	c := cString(msg)
	defer freeCString(c)
	C.vcmp_send_client_message(C.int32_t(playerID), C.uint32_t(colour), c)
}

func bridgeSendGameMessage(playerID int, msgType int, msg string) {
	c := cString(msg)
	defer freeCString(c)
	C.vcmp_send_game_message(C.int32_t(playerID), C.int32_t(msgType), c)
}

func bridgePlayerName(playerID int) string {
	buf := (*[128]C.char)(C.malloc(128))
	defer C.free(unsafe.Pointer(buf))
	C.vcmp_get_player_name(C.int32_t(playerID), &buf[0], 128)
	return C.GoString(&buf[0])
}

func bridgePlayerIDFromName(name string) int {
	c := cString(name)
	defer freeCString(c)
	return int(C.vcmp_get_player_id_from_name(c))
}

func bridgeIsPlayerConnected(playerID int) bool {
	return C.vcmp_is_player_connected(C.int32_t(playerID)) != 0
}

func bridgeIsPlayerAdmin(playerID int) bool {
	return C.vcmp_is_player_admin(C.int32_t(playerID)) != 0
}

func bridgeSetPlayerAdmin(playerID int, admin bool) {
	t := C.uint8_t(0)
	if admin {
		t = 1
	}
	C.vcmp_set_player_admin(C.int32_t(playerID), t)
}

func bridgeKickPlayer(playerID int) { C.vcmp_kick_player(C.int32_t(playerID)) }
func bridgeForceSpawn(playerID int) { C.vcmp_force_spawn(C.int32_t(playerID)) }
func bridgeGiveMoney(playerID, amount int) { C.vcmp_give_player_money(C.int32_t(playerID), C.int32_t(amount)) }
func bridgeSetScore(playerID, score int) { C.vcmp_set_player_score(C.int32_t(playerID), C.int32_t(score)) }
func bridgeGetScore(playerID int) int { return int(C.vcmp_get_player_score(C.int32_t(playerID))) }
func bridgeGetMoney(playerID int) int { return int(C.vcmp_get_player_money(C.int32_t(playerID))) }
func bridgeSetHealth(playerID int, h float32) { C.vcmp_set_player_health(C.int32_t(playerID), C.float(h)) }
func bridgeGetHealth(playerID int) float32 { return float32(C.vcmp_get_player_health(C.int32_t(playerID))) }
func bridgeSetArmour(playerID int, a float32) { C.vcmp_set_player_armour(C.int32_t(playerID), C.float(a)) }
func bridgeSetPlayerPos(playerID int, pos Vec3) {
	C.vcmp_set_player_position(C.int32_t(playerID), C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func bridgeGetPlayerPos(playerID int) Vec3 {
	var x, y, z C.float
	C.vcmp_get_player_position(C.int32_t(playerID), &x, &y, &z)
	return Vec3{X: float32(x), Y: float32(y), Z: float32(z)}
}

func bridgeSetSkin(playerID, skin int) { C.vcmp_set_player_skin(C.int32_t(playerID), C.int32_t(skin)) }
func bridgeSetTeam(playerID, team int) { C.vcmp_set_player_team(C.int32_t(playerID), C.int32_t(team)) }
func bridgeGiveWeapon(playerID, weapon, ammo int) {
	C.vcmp_give_player_weapon(C.int32_t(playerID), C.int32_t(weapon), C.int32_t(ammo))
}

func bridgePutInVehicle(playerID, vehicleID, slot int, makeRoom, warp bool) {
	mr, w := C.uint8_t(0), C.uint8_t(0)
	if makeRoom {
		mr = 1
	}
	if warp {
		w = 1
	}
	C.vcmp_put_player_in_vehicle(C.int32_t(playerID), C.int32_t(vehicleID), C.int32_t(slot), mr, w)
}

func bridgeRemoveFromVehicle(playerID int) { C.vcmp_remove_player_from_vehicle(C.int32_t(playerID)) }

func bridgeCreateVehicle(model, world int, pos Vec3, angle float32, c1, c2 int) int {
	return int(C.vcmp_create_vehicle(C.int32_t(model), C.int32_t(world), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(angle), C.int32_t(c1), C.int32_t(c2)))
}

func bridgeDeleteVehicle(id int) { C.vcmp_delete_vehicle(C.int32_t(id)) }
func bridgeRespawnVehicle(id int) { C.vcmp_respawn_vehicle(C.int32_t(id)) }

func bridgeSetVehiclePos(id int, pos Vec3, removeOccupants bool) {
	ro := C.uint8_t(0)
	if removeOccupants {
		ro = 1
	}
	C.vcmp_set_vehicle_position(C.int32_t(id), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), ro)
}

func bridgeGetVehiclePos(id int) Vec3 {
	var x, y, z C.float
	C.vcmp_get_vehicle_position(C.int32_t(id), &x, &y, &z)
	return Vec3{X: float32(x), Y: float32(y), Z: float32(z)}
}

func bridgeCreateObject(model, world int, pos Vec3, alpha int) int {
	return int(C.vcmp_create_object(C.int32_t(model), C.int32_t(world), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.int32_t(alpha)))
}

func bridgeDeleteObject(id int) { C.vcmp_delete_object(C.int32_t(id)) }
func bridgeSetObjectPos(id int, pos Vec3) {
	C.vcmp_set_object_position(C.int32_t(id), C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func bridgeCreatePickup(model, world, qty int, pos Vec3, alpha int, auto bool) int {
	a := C.uint8_t(0)
	if auto {
		a = 1
	}
	return int(C.vcmp_create_pickup(C.int32_t(model), C.int32_t(world), C.int32_t(qty), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.int32_t(alpha), a))
}

func bridgeDeletePickup(id int) { C.vcmp_delete_pickup(C.int32_t(id)) }

func bridgeCreateCheckpoint(playerID, world int, sphere bool, pos Vec3, rgba [4]int32, radius float32) int {
	s := C.uint8_t(0)
	if sphere {
		s = 1
	}
	return int(C.vcmp_create_checkpoint(C.int32_t(playerID), C.int32_t(world), s, C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.int32_t(rgba[0]), C.int32_t(rgba[1]), C.int32_t(rgba[2]), C.int32_t(rgba[3]), C.float(radius)))
}

func bridgeDeleteCheckpoint(id int) { C.vcmp_delete_checkpoint(C.int32_t(id)) }

func bridgeCreateBlip(index, world int, pos Vec3, scale int, colour uint32, sprite int) int {
	return int(C.vcmp_create_coord_blip(C.int32_t(index), C.int32_t(world), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.int32_t(scale), C.uint32_t(colour), C.int32_t(sprite)))
}

func bridgeDestroyBlip(index int) { C.vcmp_destroy_coord_blip(C.int32_t(index)) }
func bridgeKeyBindUnusedSlot() int { return int(C.vcmp_get_keybind_unused_slot()) }

func bridgeRegisterKeyBind(bindID int, onRelease bool, k1, k2, k3 int) {
	or := C.uint8_t(0)
	if onRelease {
		or = 1
	}
	C.vcmp_register_keybind(C.int32_t(bindID), or, C.int32_t(k1), C.int32_t(k2), C.int32_t(k3))
}

func bridgeRemoveKeyBind(bindID int) { C.vcmp_remove_keybind(C.int32_t(bindID)) }
