package vcmp

/*
#cgo CFLAGS: -I${SRCDIR}/../include
#include "plugin.h"
#include <stdlib.h>
#include <string.h>

// cgo //export symbols — must use these addresses, not raw Go func pointers.
extern uint8_t OnServerInitialise(void);
extern void OnServerShutdown(void);
extern void OnServerFrame(float elapsedTime);
extern uint8_t OnPluginCommand(uint32_t commandIdentifier, const char* message);
extern uint8_t OnIncomingConnection(char* playerName, size_t nameBufferSize, const char* userPassword, const char* ipAddress);
extern void OnClientScriptData(int32_t playerId, const uint8_t* data, size_t size);

extern void OnPlayerConnect(int32_t playerId);
extern void OnPlayerDisconnect(int32_t playerId, vcmpDisconnectReason reason);
extern uint8_t OnPlayerRequestClass(int32_t playerId, int32_t offset);
extern uint8_t OnPlayerRequestSpawn(int32_t playerId);
extern void OnPlayerSpawn(int32_t playerId);
extern void OnPlayerDeath(int32_t playerId, int32_t killerId, int32_t reason, vcmpBodyPart bodyPart);
extern void OnPlayerUpdate(int32_t playerId, vcmpPlayerUpdate updateType);

extern uint8_t OnPlayerRequestEnterVehicle(int32_t playerId, int32_t vehicleId, int32_t slotIndex);
extern void OnPlayerEnterVehicle(int32_t playerId, int32_t vehicleId, int32_t slotIndex);
extern void OnPlayerExitVehicle(int32_t playerId, int32_t vehicleId);

extern void OnPlayerNameChange(int32_t playerId, const char* oldName, const char* newName);
extern void OnPlayerStateChange(int32_t playerId, vcmpPlayerState oldState, vcmpPlayerState newState);
extern void OnPlayerActionChange(int32_t playerId, int32_t oldAction, int32_t newAction);
extern void OnPlayerOnFireChange(int32_t playerId, uint8_t isOnFire);
extern void OnPlayerCrouchChange(int32_t playerId, uint8_t isCrouching);
extern void OnPlayerGameKeysChange(int32_t playerId, uint32_t oldKeys, uint32_t newKeys);
extern void OnPlayerBeginTyping(int32_t playerId);
extern void OnPlayerEndTyping(int32_t playerId);
extern void OnPlayerAwayChange(int32_t playerId, uint8_t isAway);

extern uint8_t OnPlayerMessage(int32_t playerId, const char* message);
extern uint8_t OnPlayerCommand(int32_t playerId, const char* message);
extern uint8_t OnPlayerPrivateMessage(int32_t playerId, int32_t targetPlayerId, const char* message);

extern void OnPlayerKeyBindDown(int32_t playerId, int32_t bindId);
extern void OnPlayerKeyBindUp(int32_t playerId, int32_t bindId);
extern void OnPlayerSpectate(int32_t playerId, int32_t targetPlayerId);
extern void OnPlayerCrashReport(int32_t playerId, const char* report);
extern void OnPlayerModuleList(int32_t playerId, const char* list);

extern void OnVehicleUpdate(int32_t vehicleId, vcmpVehicleUpdate updateType);
extern void OnVehicleExplode(int32_t vehicleId);
extern void OnVehicleRespawn(int32_t vehicleId);

extern void OnObjectShot(int32_t objectId, int32_t playerId, int32_t weaponId);
extern void OnObjectTouched(int32_t objectId, int32_t playerId);

extern uint8_t OnPickupPickAttempt(int32_t pickupId, int32_t playerId);
extern void OnPickupPicked(int32_t pickupId, int32_t playerId);
extern void OnPickupRespawn(int32_t pickupId);

extern void OnCheckpointEntered(int32_t checkPointId, int32_t playerId);
extern void OnCheckpointExited(int32_t checkPointId, int32_t playerId);

extern void OnEntityPoolChange(vcmpEntityPool entityType, int32_t entityId, uint8_t isDeleted);
extern void OnServerPerformanceReport(size_t entryCount, const char** descriptions, uint64_t* times);

static void vcmp_register_callbacks(PluginCallbacks *calls) {
	calls->structSize = sizeof(PluginCallbacks);

	calls->OnServerInitialise = OnServerInitialise;
	calls->OnServerShutdown = OnServerShutdown;
	calls->OnServerFrame = OnServerFrame;
	calls->OnPluginCommand = OnPluginCommand;
	calls->OnIncomingConnection = OnIncomingConnection;
	calls->OnClientScriptData = OnClientScriptData;

	calls->OnPlayerConnect = OnPlayerConnect;
	calls->OnPlayerDisconnect = OnPlayerDisconnect;
	calls->OnPlayerRequestClass = OnPlayerRequestClass;
	calls->OnPlayerRequestSpawn = OnPlayerRequestSpawn;
	calls->OnPlayerSpawn = OnPlayerSpawn;
	calls->OnPlayerDeath = OnPlayerDeath;
	calls->OnPlayerUpdate = OnPlayerUpdate;

	calls->OnPlayerRequestEnterVehicle = OnPlayerRequestEnterVehicle;
	calls->OnPlayerEnterVehicle = OnPlayerEnterVehicle;
	calls->OnPlayerExitVehicle = OnPlayerExitVehicle;

	calls->OnPlayerNameChange = OnPlayerNameChange;
	calls->OnPlayerStateChange = OnPlayerStateChange;
	calls->OnPlayerActionChange = OnPlayerActionChange;
	calls->OnPlayerOnFireChange = OnPlayerOnFireChange;
	calls->OnPlayerCrouchChange = OnPlayerCrouchChange;
	calls->OnPlayerGameKeysChange = OnPlayerGameKeysChange;
	calls->OnPlayerBeginTyping = OnPlayerBeginTyping;
	calls->OnPlayerEndTyping = OnPlayerEndTyping;
	calls->OnPlayerAwayChange = OnPlayerAwayChange;

	calls->OnPlayerMessage = OnPlayerMessage;
	calls->OnPlayerCommand = OnPlayerCommand;
	calls->OnPlayerPrivateMessage = OnPlayerPrivateMessage;

	calls->OnPlayerKeyBindDown = OnPlayerKeyBindDown;
	calls->OnPlayerKeyBindUp = OnPlayerKeyBindUp;
	calls->OnPlayerSpectate = OnPlayerSpectate;
	calls->OnPlayerCrashReport = OnPlayerCrashReport;
	calls->OnPlayerModuleList = OnPlayerModuleList;

	calls->OnVehicleUpdate = OnVehicleUpdate;
	calls->OnVehicleExplode = OnVehicleExplode;
	calls->OnVehicleRespawn = OnVehicleRespawn;

	calls->OnObjectShot = OnObjectShot;
	calls->OnObjectTouched = OnObjectTouched;

	calls->OnPickupPickAttempt = OnPickupPickAttempt;
	calls->OnPickupPicked = OnPickupPicked;
	calls->OnPickupRespawn = OnPickupRespawn;

	calls->OnCheckpointEntered = OnCheckpointEntered;
	calls->OnCheckpointExited = OnCheckpointExited;

	calls->OnEntityPoolChange = OnEntityPoolChange;
	calls->OnServerPerformanceReport = OnServerPerformanceReport;
}
*/
import "C"

// PluginMeta is reported to the VC:MP server in PluginInfo.
type PluginMeta struct {
	Name    string
	Version uint32
}

// Init binds native APIs, fills plugin metadata, and registers all SDK callbacks.
// Returns false when required pointers are nil.
func Init(funcs *C.PluginFuncs, calls *C.PluginCallbacks, info *C.PluginInfo, meta PluginMeta) bool {
	if funcs == nil || calls == nil || info == nil {
		return false
	}
	Bind(funcs)
	setPluginInfo(info, meta)
	RegisterCallbacks(calls)
	return true
}

func setPluginInfo(info *C.PluginInfo, meta PluginMeta) {
	info.structSize = C.uint32_t(C.sizeof_PluginInfo)
	name := cString(meta.Name)
	defer freeCString(name)
	C.strncpy(&info.name[0], name, 31)
	info.pluginVersion = C.uint32_t(meta.Version)
	info.apiMajorVersion = C.PLUGIN_API_MAJOR
	info.apiMinorVersion = C.PLUGIN_API_MINOR
}

// RegisterCallbacks wires all //export handlers into the VC:MP callback table.
// Uses cgo export symbols so VC:MP calls the runtime-aware trampolines.
func RegisterCallbacks(calls *C.PluginCallbacks) {
	C.vcmp_register_callbacks(calls)
}
