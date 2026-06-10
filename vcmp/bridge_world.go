package vcmp

/*
#cgo CFLAGS: -I${SRCDIR}/../include
#include "plugin.h"

extern PluginFuncs *g_pf;

static vcmpError vcmp_play_sound(int32_t worldId, int32_t soundId, float x, float y, float z) {
	if (g_pf && g_pf->PlaySound) return g_pf->PlaySound(worldId, soundId, x, y, z);
	return vcmpErrorArgumentOutOfBounds;
}
static void vcmp_set_world_bounds(float maxX, float minX, float maxY, float minY) {
	if (g_pf && g_pf->SetWorldBounds) g_pf->SetWorldBounds(maxX, minX, maxY, minY);
}
static void vcmp_get_world_bounds(float *maxX, float *minX, float *maxY, float *minY) {
	if (g_pf && g_pf->GetWorldBounds) g_pf->GetWorldBounds(maxX, minX, maxY, minY);
}
static vcmpError vcmp_get_coord_blip_info(int32_t index, int32_t *world, float *x, float *y, float *z, int32_t *scale, uint32_t *colour, int32_t *sprite) {
	if (g_pf && g_pf->GetCoordBlipInfo) return g_pf->GetCoordBlipInfo(index, world, x, y, z, scale, colour, sprite);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_set_vehicle_speed(int32_t vehicleId, float x, float y, float z, uint8_t add, uint8_t relative) {
	if (g_pf && g_pf->SetVehicleSpeed) return g_pf->SetVehicleSpeed(vehicleId, x, y, z, add, relative);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_get_vehicle_speed(int32_t vehicleId, float *x, float *y, float *z, uint8_t relative) {
	if (g_pf && g_pf->GetVehicleSpeed) return g_pf->GetVehicleSpeed(vehicleId, x, y, z, relative);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_get_vehicle_rotation_euler(int32_t vehicleId, float *x, float *y, float *z) {
	if (g_pf && g_pf->GetVehicleRotationEuler) return g_pf->GetVehicleRotationEuler(vehicleId, x, y, z);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_set_vehicle_spawn_position(int32_t vehicleId, float x, float y, float z) {
	if (g_pf && g_pf->SetVehicleSpawnPosition) return g_pf->SetVehicleSpawnPosition(vehicleId, x, y, z);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_set_vehicle_lights_data(int32_t vehicleId, uint32_t lightsData) {
	if (g_pf && g_pf->SetVehicleLightsData) return g_pf->SetVehicleLightsData(vehicleId, lightsData);
	return vcmpErrorNoSuchEntity;
}
static uint32_t vcmp_get_vehicle_lights_data(int32_t vehicleId) {
	if (g_pf && g_pf->GetVehicleLightsData) return g_pf->GetVehicleLightsData(vehicleId);
	return 0;
}
static vcmpError vcmp_set_player_spectate_target(int32_t playerId, int32_t targetId) {
	if (g_pf && g_pf->SetPlayerSpectateTarget) return g_pf->SetPlayerSpectateTarget(playerId, targetId);
	return vcmpErrorNoSuchEntity;
}
static int32_t vcmp_get_player_spectate_target(int32_t playerId) {
	if (g_pf && g_pf->GetPlayerSpectateTarget) return g_pf->GetPlayerSpectateTarget(playerId);
	return -1;
}
*/
import "C"

import "errors"

type WorldBounds struct {
	MaxX, MinX, MaxY, MinY float32
}

type BlipInfo struct {
	World    int
	Position Vec3
	Scale    int
	Colour   uint32
	Sprite   int
}

func bridgePlaySound(world, soundID int, pos Vec3) error {
	return bridgeError(C.vcmp_play_sound(C.int32_t(world), C.int32_t(soundID), C.float(pos.X), C.float(pos.Y), C.float(pos.Z)))
}

func bridgeSetWorldBounds(bounds WorldBounds) {
	C.vcmp_set_world_bounds(C.float(bounds.MaxX), C.float(bounds.MinX), C.float(bounds.MaxY), C.float(bounds.MinY))
}

func bridgeGetWorldBounds() WorldBounds {
	var maxX, minX, maxY, minY C.float
	C.vcmp_get_world_bounds(&maxX, &minX, &maxY, &minY)
	return WorldBounds{
		MaxX: float32(maxX), MinX: float32(minX),
		MaxY: float32(maxY), MinY: float32(minY),
	}
}

func bridgeGetCoordBlipInfo(index int) (BlipInfo, error) {
	var world, scale, sprite C.int32_t
	var x, y, z C.float
	var colour C.uint32_t
	if err := bridgeError(C.vcmp_get_coord_blip_info(C.int32_t(index), &world, &x, &y, &z, &scale, &colour, &sprite)); err != nil {
		return BlipInfo{}, err
	}
	return BlipInfo{
		World:    int(world),
		Position: Vec3{X: float32(x), Y: float32(y), Z: float32(z)},
		Scale:    int(scale),
		Colour:   uint32(colour),
		Sprite:   int(sprite),
	}, nil
}

func bridgeSetVehicleSpeed(vehicleID int, speed Vec3, add, relative bool) error {
	return bridgeError(C.vcmp_set_vehicle_speed(
		C.int32_t(vehicleID), C.float(speed.X), C.float(speed.Y), C.float(speed.Z),
		boolToU8(add), boolToU8(relative),
	))
}

func bridgeGetVehicleSpeed(vehicleID int, relative bool) (Vec3, error) {
	var x, y, z C.float
	if err := bridgeError(C.vcmp_get_vehicle_speed(C.int32_t(vehicleID), &x, &y, &z, boolToU8(relative))); err != nil {
		return Vec3{}, err
	}
	return Vec3{X: float32(x), Y: float32(y), Z: float32(z)}, nil
}

func bridgeGetVehicleRotationEuler(vehicleID int) (Vec3, error) {
	var x, y, z C.float
	if err := bridgeError(C.vcmp_get_vehicle_rotation_euler(C.int32_t(vehicleID), &x, &y, &z)); err != nil {
		return Vec3{}, err
	}
	return Vec3{X: float32(x), Y: float32(y), Z: float32(z)}, nil
}

func bridgeSetVehicleSpawnPosition(vehicleID int, pos Vec3) error {
	return bridgeError(C.vcmp_set_vehicle_spawn_position(C.int32_t(vehicleID), C.float(pos.X), C.float(pos.Y), C.float(pos.Z)))
}

func bridgeSetVehicleLightsData(vehicleID int, lights uint32) error {
	return bridgeError(C.vcmp_set_vehicle_lights_data(C.int32_t(vehicleID), C.uint32_t(lights)))
}

func bridgeGetVehicleLightsData(vehicleID int) uint32 {
	return uint32(C.vcmp_get_vehicle_lights_data(C.int32_t(vehicleID)))
}

func bridgeSetPlayerSpectateTarget(playerID, targetID int) error {
	return bridgeError(C.vcmp_set_player_spectate_target(C.int32_t(playerID), C.int32_t(targetID)))
}

func bridgeGetPlayerSpectateTarget(playerID int) int {
	return int(C.vcmp_get_player_spectate_target(C.int32_t(playerID)))
}

var errPluginAPINotAvailable = errors.New("vcmp: not available in VC:MP 0.4 plugin.h")

func bridgeInterpolateCameraLookAt(playerID int, lookAt Vec3, interpMs uint32) error {
	return errPluginAPINotAvailable
}

func bridgeSetPlayerDrunkHandling(playerID int, level uint32) error {
	return errPluginAPINotAvailable
}

func bridgeGetPlayerDrunkHandling(playerID int) uint32 {
	return 0
}

func bridgeSetPlayerDrunkVisuals(playerID int, level uint8) error {
	return errPluginAPINotAvailable
}

func bridgeGetPlayerDrunkVisuals(playerID int) uint8 {
	return 0
}
