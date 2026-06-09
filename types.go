package main

/*
#cgo CFLAGS: -I${SRCDIR}/include
#include "plugin.h"
*/
import "C"

const (
	MaxPlayers = 100
	PluginName = "GoSDK"

	ColourWhite  uint32 = 0xFFFFFFFF
	ColourGreen  uint32 = 0xFF00FF00
	ColourYellow uint32 = 0xFFFFFF00
	ColourRed    uint32 = 0xFFFF4040
	ColourCyan   uint32 = 0xFF00FFFF
)

type Vec3 struct {
	X float32
	Y float32
	Z float32
}

type ServerOption int

const (
	ServerOptionSyncFrameLimiter ServerOption = ServerOption(C.vcmpServerOptionSyncFrameLimiter)
	ServerOptionFrameLimiter     ServerOption = ServerOption(C.vcmpServerOptionFrameLimiter)
	ServerOptionFriendlyFire     ServerOption = ServerOption(C.vcmpServerOptionFriendlyFire)
	ServerOptionJoinMessages     ServerOption = ServerOption(C.vcmpServerOptionJoinMessages)
	ServerOptionDeathMessages    ServerOption = ServerOption(C.vcmpServerOptionDeathMessages)
	ServerOptionUseClasses       ServerOption = ServerOption(C.vcmpServerOptionUseClasses)
	ServerOptionShowNameTags     ServerOption = ServerOption(C.vcmpServerOptionShowNameTags)
)

type PlayerOption int

const (
	PlayerOptionControllable PlayerOption = PlayerOption(C.vcmpPlayerOptionControllable)
	PlayerOptionDriveBy      PlayerOption = PlayerOption(C.vcmpPlayerOptionDriveBy)
	PlayerOptionCanAttack    PlayerOption = PlayerOption(C.vcmpPlayerOptionCanAttack)
)

type EntityPool int

const (
	EntityPoolVehicle    EntityPool = EntityPool(C.vcmpEntityPoolVehicle)
	EntityPoolObject     EntityPool = EntityPool(C.vcmpEntityPoolObject)
	EntityPoolPickup     EntityPool = EntityPool(C.vcmpEntityPoolPickup)
	EntityPoolBlip       EntityPool = EntityPool(C.vcmpEntityPoolBlip)
	EntityPoolCheckPoint EntityPool = EntityPool(C.vcmpEntityPoolCheckPoint)
)

type DisconnectReason int

const (
	DisconnectTimeout  DisconnectReason = DisconnectReason(C.vcmpDisconnectReasonTimeout)
	DisconnectQuit     DisconnectReason = DisconnectReason(C.vcmpDisconnectReasonQuit)
	DisconnectKick     DisconnectReason = DisconnectReason(C.vcmpDisconnectReasonKick)
	DisconnectCrash    DisconnectReason = DisconnectReason(C.vcmpDisconnectReasonCrash)
	DisconnectAntiCheat DisconnectReason = DisconnectReason(C.vcmpDisconnectReasonAntiCheat)
)

type FilterResult uint8

const (
	FilterAllow FilterResult = 1
	FilterDeny  FilterResult = 0
)
