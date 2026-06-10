package vcmp

/*
#cgo CFLAGS: -I${SRCDIR}/../include
#include "plugin.h"

extern PluginFuncs *g_pf;

static vcmpError vcmp_set_pickup_world(int32_t pickupId, int32_t world) {
	if (g_pf && g_pf->SetPickupWorld) return g_pf->SetPickupWorld(pickupId, world);
	return vcmpErrorNoSuchEntity;
}
static int32_t vcmp_get_pickup_world(int32_t pickupId) {
	if (g_pf && g_pf->GetPickupWorld) return g_pf->GetPickupWorld(pickupId);
	return 0;
}
static vcmpError vcmp_set_pickup_alpha(int32_t pickupId, int32_t alpha) {
	if (g_pf && g_pf->SetPickupAlpha) return g_pf->SetPickupAlpha(pickupId, alpha);
	return vcmpErrorNoSuchEntity;
}
static int32_t vcmp_get_pickup_alpha(int32_t pickupId) {
	if (g_pf && g_pf->GetPickupAlpha) return g_pf->GetPickupAlpha(pickupId);
	return 0;
}
static vcmpError vcmp_set_pickup_automatic(int32_t pickupId, uint8_t toggle) {
	if (g_pf && g_pf->SetPickupIsAutomatic) return g_pf->SetPickupIsAutomatic(pickupId, toggle);
	return vcmpErrorNoSuchEntity;
}
static uint8_t vcmp_is_pickup_automatic(int32_t pickupId) {
	if (g_pf && g_pf->IsPickupAutomatic) return g_pf->IsPickupAutomatic(pickupId);
	return 0;
}
static vcmpError vcmp_refresh_pickup(int32_t pickupId) {
	if (g_pf && g_pf->RefreshPickup) return g_pf->RefreshPickup(pickupId);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_set_pickup_position(int32_t pickupId, float x, float y, float z) {
	if (g_pf && g_pf->SetPickupPosition) return g_pf->SetPickupPosition(pickupId, x, y, z);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_get_pickup_position(int32_t pickupId, float *x, float *y, float *z) {
	if (g_pf && g_pf->GetPickupPosition) return g_pf->GetPickupPosition(pickupId, x, y, z);
	return vcmpErrorNoSuchEntity;
}
static int32_t vcmp_get_pickup_model(int32_t pickupId) {
	if (g_pf && g_pf->GetPickupModel) return g_pf->GetPickupModel(pickupId);
	return 0;
}
static int32_t vcmp_get_pickup_quantity(int32_t pickupId) {
	if (g_pf && g_pf->GetPickupQuantity) return g_pf->GetPickupQuantity(pickupId);
	return 0;
}
static vcmpError vcmp_set_pickup_option(int32_t pickupId, vcmpPickupOption option, uint8_t toggle) {
	if (g_pf && g_pf->SetPickupOption) return g_pf->SetPickupOption(pickupId, option, toggle);
	return vcmpErrorNoSuchEntity;
}
static uint8_t vcmp_get_pickup_option(int32_t pickupId, vcmpPickupOption option) {
	if (g_pf && g_pf->GetPickupOption) return g_pf->GetPickupOption(pickupId, option);
	return 0;
}

static uint8_t vcmp_is_checkpoint_sphere(int32_t cpId) {
	if (g_pf && g_pf->IsCheckPointSphere) return g_pf->IsCheckPointSphere(cpId);
	return 0;
}
static vcmpError vcmp_set_checkpoint_world(int32_t cpId, int32_t world) {
	if (g_pf && g_pf->SetCheckPointWorld) return g_pf->SetCheckPointWorld(cpId, world);
	return vcmpErrorNoSuchEntity;
}
static int32_t vcmp_get_checkpoint_world(int32_t cpId) {
	if (g_pf && g_pf->GetCheckPointWorld) return g_pf->GetCheckPointWorld(cpId);
	return 0;
}
static vcmpError vcmp_set_checkpoint_colour(int32_t cpId, int32_t r, int32_t g, int32_t b, int32_t a) {
	if (g_pf && g_pf->SetCheckPointColour) return g_pf->SetCheckPointColour(cpId, r, g, b, a);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_get_checkpoint_colour(int32_t cpId, int32_t *r, int32_t *g, int32_t *b, int32_t *a) {
	if (g_pf && g_pf->GetCheckPointColour) return g_pf->GetCheckPointColour(cpId, r, g, b, a);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_set_checkpoint_position(int32_t cpId, float x, float y, float z) {
	if (g_pf && g_pf->SetCheckPointPosition) return g_pf->SetCheckPointPosition(cpId, x, y, z);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_get_checkpoint_position(int32_t cpId, float *x, float *y, float *z) {
	if (g_pf && g_pf->GetCheckPointPosition) return g_pf->GetCheckPointPosition(cpId, x, y, z);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_set_checkpoint_radius(int32_t cpId, float radius) {
	if (g_pf && g_pf->SetCheckPointRadius) return g_pf->SetCheckPointRadius(cpId, radius);
	return vcmpErrorNoSuchEntity;
}
static float vcmp_get_checkpoint_radius(int32_t cpId) {
	if (g_pf && g_pf->GetCheckPointRadius) return g_pf->GetCheckPointRadius(cpId);
	return 0;
}
static int32_t vcmp_get_checkpoint_owner(int32_t cpId) {
	if (g_pf && g_pf->GetCheckPointOwner) return g_pf->GetCheckPointOwner(cpId);
	return -1;
}

static int32_t vcmp_get_object_model(int32_t objectId) {
	if (g_pf && g_pf->GetObjectModel) return g_pf->GetObjectModel(objectId);
	return 0;
}
static vcmpError vcmp_set_object_world(int32_t objectId, int32_t world) {
	if (g_pf && g_pf->SetObjectWorld) return g_pf->SetObjectWorld(objectId, world);
	return vcmpErrorNoSuchEntity;
}
static int32_t vcmp_get_object_world(int32_t objectId) {
	if (g_pf && g_pf->GetObjectWorld) return g_pf->GetObjectWorld(objectId);
	return 0;
}
static vcmpError vcmp_set_object_alpha(int32_t objectId, int32_t alpha, uint32_t duration) {
	if (g_pf && g_pf->SetObjectAlpha) return g_pf->SetObjectAlpha(objectId, alpha, duration);
	return vcmpErrorNoSuchEntity;
}
static int32_t vcmp_get_object_alpha(int32_t objectId) {
	if (g_pf && g_pf->GetObjectAlpha) return g_pf->GetObjectAlpha(objectId);
	return 0;
}
static vcmpError vcmp_move_object_to(int32_t objectId, float x, float y, float z, uint32_t duration) {
	if (g_pf && g_pf->MoveObjectTo) return g_pf->MoveObjectTo(objectId, x, y, z, duration);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_set_object_position(int32_t objectId, float x, float y, float z) {
	if (g_pf && g_pf->SetObjectPosition) return g_pf->SetObjectPosition(objectId, x, y, z);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_get_object_position(int32_t objectId, float *x, float *y, float *z) {
	if (g_pf && g_pf->GetObjectPosition) return g_pf->GetObjectPosition(objectId, x, y, z);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_rotate_object_to_euler(int32_t objectId, float x, float y, float z, uint32_t duration) {
	if (g_pf && g_pf->RotateObjectToEuler) return g_pf->RotateObjectToEuler(objectId, x, y, z, duration);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_get_object_rotation_euler(int32_t objectId, float *x, float *y, float *z) {
	if (g_pf && g_pf->GetObjectRotationEuler) return g_pf->GetObjectRotationEuler(objectId, x, y, z);
	return vcmpErrorNoSuchEntity;
}
static vcmpError vcmp_set_object_shot_report(int32_t objectId, uint8_t toggle) {
	if (g_pf && g_pf->SetObjectShotReportEnabled) return g_pf->SetObjectShotReportEnabled(objectId, toggle);
	return vcmpErrorNoSuchEntity;
}
static uint8_t vcmp_is_object_shot_report(int32_t objectId) {
	if (g_pf && g_pf->IsObjectShotReportEnabled) return g_pf->IsObjectShotReportEnabled(objectId);
	return 0;
}
static vcmpError vcmp_set_object_touched_report(int32_t objectId, uint8_t toggle) {
	if (g_pf && g_pf->SetObjectTouchedReportEnabled) return g_pf->SetObjectTouchedReportEnabled(objectId, toggle);
	return vcmpErrorNoSuchEntity;
}
static uint8_t vcmp_is_object_touched_report(int32_t objectId) {
	if (g_pf && g_pf->IsObjectTouchedReportEnabled) return g_pf->IsObjectTouchedReportEnabled(objectId);
	return 0;
}
*/
import "C"

func bridgeSetPickupWorld(pickupID, world int) error {
	return bridgeError(C.vcmp_set_pickup_world(C.int32_t(pickupID), C.int32_t(world)))
}

func bridgeGetPickupWorld(pickupID int) int {
	return int(C.vcmp_get_pickup_world(C.int32_t(pickupID)))
}

func bridgeSetPickupAlpha(pickupID, alpha int) error {
	return bridgeError(C.vcmp_set_pickup_alpha(C.int32_t(pickupID), C.int32_t(alpha)))
}

func bridgeGetPickupAlpha(pickupID int) int {
	return int(C.vcmp_get_pickup_alpha(C.int32_t(pickupID)))
}

func bridgeSetPickupAutomatic(pickupID int, automatic bool) error {
	return bridgeError(C.vcmp_set_pickup_automatic(C.int32_t(pickupID), boolToU8(automatic)))
}

func bridgeIsPickupAutomatic(pickupID int) bool {
	return C.vcmp_is_pickup_automatic(C.int32_t(pickupID)) != 0
}

func bridgeRefreshPickup(pickupID int) error {
	return bridgeError(C.vcmp_refresh_pickup(C.int32_t(pickupID)))
}

func bridgeSetPickupPosition(pickupID int, pos Vec3) error {
	return bridgeError(C.vcmp_set_pickup_position(C.int32_t(pickupID), C.float(pos.X), C.float(pos.Y), C.float(pos.Z)))
}

func bridgeGetPickupPosition(pickupID int) (Vec3, error) {
	var x, y, z C.float
	if err := bridgeError(C.vcmp_get_pickup_position(C.int32_t(pickupID), &x, &y, &z)); err != nil {
		return Vec3{}, err
	}
	return Vec3{X: float32(x), Y: float32(y), Z: float32(z)}, nil
}

func bridgeGetPickupModel(pickupID int) int {
	return int(C.vcmp_get_pickup_model(C.int32_t(pickupID)))
}

func bridgeGetPickupQuantity(pickupID int) int {
	return int(C.vcmp_get_pickup_quantity(C.int32_t(pickupID)))
}

func bridgeSetPickupOption(pickupID int, option PickupOption, on bool) error {
	return bridgeError(C.vcmp_set_pickup_option(C.int32_t(pickupID), C.vcmpPickupOption(option), boolToU8(on)))
}

func bridgeGetPickupOption(pickupID int, option PickupOption) bool {
	return C.vcmp_get_pickup_option(C.int32_t(pickupID), C.vcmpPickupOption(option)) != 0
}

func bridgeIsCheckPointSphere(checkpointID int) bool {
	return C.vcmp_is_checkpoint_sphere(C.int32_t(checkpointID)) != 0
}

func bridgeSetCheckPointWorld(checkpointID, world int) error {
	return bridgeError(C.vcmp_set_checkpoint_world(C.int32_t(checkpointID), C.int32_t(world)))
}

func bridgeGetCheckPointWorld(checkpointID int) int {
	return int(C.vcmp_get_checkpoint_world(C.int32_t(checkpointID)))
}

func bridgeSetCheckPointColour(checkpointID, r, g, b, alpha int) error {
	return bridgeError(C.vcmp_set_checkpoint_colour(C.int32_t(checkpointID), C.int32_t(r), C.int32_t(g), C.int32_t(b), C.int32_t(alpha)))
}

func bridgeGetCheckPointColour(checkpointID int) (r, g, b, alpha int, err error) {
	var cr, cg, cb, ca C.int32_t
	if e := bridgeError(C.vcmp_get_checkpoint_colour(C.int32_t(checkpointID), &cr, &cg, &cb, &ca)); e != nil {
		return 0, 0, 0, 0, e
	}
	return int(cr), int(cg), int(cb), int(ca), nil
}

func bridgeSetCheckPointPosition(checkpointID int, pos Vec3) error {
	return bridgeError(C.vcmp_set_checkpoint_position(C.int32_t(checkpointID), C.float(pos.X), C.float(pos.Y), C.float(pos.Z)))
}

func bridgeGetCheckPointPosition(checkpointID int) (Vec3, error) {
	var x, y, z C.float
	if err := bridgeError(C.vcmp_get_checkpoint_position(C.int32_t(checkpointID), &x, &y, &z)); err != nil {
		return Vec3{}, err
	}
	return Vec3{X: float32(x), Y: float32(y), Z: float32(z)}, nil
}

func bridgeSetCheckPointRadius(checkpointID int, radius float32) error {
	return bridgeError(C.vcmp_set_checkpoint_radius(C.int32_t(checkpointID), C.float(radius)))
}

func bridgeGetCheckPointRadius(checkpointID int) float32 {
	return float32(C.vcmp_get_checkpoint_radius(C.int32_t(checkpointID)))
}

func bridgeGetCheckPointOwner(checkpointID int) int {
	return int(C.vcmp_get_checkpoint_owner(C.int32_t(checkpointID)))
}

func bridgeGetObjectModel(objectID int) int {
	return int(C.vcmp_get_object_model(C.int32_t(objectID)))
}

func bridgeSetObjectWorld(objectID, world int) error {
	return bridgeError(C.vcmp_set_object_world(C.int32_t(objectID), C.int32_t(world)))
}

func bridgeGetObjectWorld(objectID int) int {
	return int(C.vcmp_get_object_world(C.int32_t(objectID)))
}

func bridgeSetObjectAlpha(objectID, alpha int, durationMs uint32) error {
	return bridgeError(C.vcmp_set_object_alpha(C.int32_t(objectID), C.int32_t(alpha), C.uint32_t(durationMs)))
}

func bridgeGetObjectAlpha(objectID int) int {
	return int(C.vcmp_get_object_alpha(C.int32_t(objectID)))
}

func bridgeMoveObjectTo(objectID int, pos Vec3, durationMs uint32) error {
	return bridgeError(C.vcmp_move_object_to(C.int32_t(objectID), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.uint32_t(durationMs)))
}

func bridgeSetObjectPosition(objectID int, pos Vec3) error {
	return bridgeError(C.vcmp_set_object_position(C.int32_t(objectID), C.float(pos.X), C.float(pos.Y), C.float(pos.Z)))
}

func bridgeGetObjectPosition(objectID int) (Vec3, error) {
	var x, y, z C.float
	if err := bridgeError(C.vcmp_get_object_position(C.int32_t(objectID), &x, &y, &z)); err != nil {
		return Vec3{}, err
	}
	return Vec3{X: float32(x), Y: float32(y), Z: float32(z)}, nil
}

func bridgeRotateObjectToEuler(objectID int, rot Vec3, durationMs uint32) error {
	return bridgeError(C.vcmp_rotate_object_to_euler(C.int32_t(objectID), C.float(rot.X), C.float(rot.Y), C.float(rot.Z), C.uint32_t(durationMs)))
}

func bridgeGetObjectRotationEuler(objectID int) (Vec3, error) {
	var x, y, z C.float
	if err := bridgeError(C.vcmp_get_object_rotation_euler(C.int32_t(objectID), &x, &y, &z)); err != nil {
		return Vec3{}, err
	}
	return Vec3{X: float32(x), Y: float32(y), Z: float32(z)}, nil
}

func bridgeSetObjectShotReport(objectID int, on bool) error {
	return bridgeError(C.vcmp_set_object_shot_report(C.int32_t(objectID), boolToU8(on)))
}

func bridgeIsObjectShotReport(objectID int) bool {
	return C.vcmp_is_object_shot_report(C.int32_t(objectID)) != 0
}

func bridgeSetObjectTouchedReport(objectID int, on bool) error {
	return bridgeError(C.vcmp_set_object_touched_report(C.int32_t(objectID), boolToU8(on)))
}

func bridgeIsObjectTouchedReport(objectID int) bool {
	return C.vcmp_is_object_touched_report(C.int32_t(objectID)) != 0
}
