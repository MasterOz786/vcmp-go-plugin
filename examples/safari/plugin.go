package main

import (
	"fmt"

	"github.com/masteroz/vcmp-go-plugin/vcmp"
	"github.com/masteroz/vcmp-go-server/safari"
)

type Plugin struct {
	engine *safari.Engine
	store  *safari.Store
	db     *safari.DBWorker
}

var plug *Plugin

func newPlugin(cfg Config) *Plugin {
	safariCfg := safari.LoadConfig()
	mapCfg, err := safari.LoadMap(safariCfg.MapFile)
	if err != nil {
		vcmp.API.Server.Log(fmt.Sprintf("[safari] map load failed (%s): %v — using defaults", safariCfg.MapFile, err))
		mapCfg = defaultSafariMap()
	}

	store, err := safari.OpenStore(safariCfg.DBPath)
	if err != nil {
		vcmp.API.Server.Log(fmt.Sprintf("[safari] database open failed: %v", err))
		return &Plugin{}
	}

	db := safari.NewDBWorker(store, 128)
	db.Start()

	gameMode := cfg.GameModeText
	if gameMode == "" {
		gameMode = "Project Safari: Hydra Warfare"
	}

	engine := safari.NewEngine(safari.VCMPAPI{}, db, safariCfg, mapCfg, cfg.ServerName, gameMode)
	engine.Start()

	vcmp.API.Server.Log(fmt.Sprintf("[safari] gamemode initialized (map=%s db=%s)", safariCfg.MapFile, safariCfg.DBPath))

	return &Plugin{engine: engine, store: store, db: db}
}

func (p *Plugin) shutdown() {
	if p.engine != nil {
		p.engine.Stop()
	}
	if p.db != nil {
		p.db.Stop()
	}
	if p.store != nil {
		_ = p.store.Close()
	}
}

func defaultSafariMap() safari.MapConfig {
	return safari.MapConfig{
		HydraStart:   safari.Vec3{X: -974, Y: -106, Z: 11.2},
		HydraAngle:   90,
		World:        0,
		Waypoints: []safari.Vec3{
			{X: -1020, Y: 85, Z: 14.5},
			{X: -880, Y: 115, Z: 14.5},
			{X: -720, Y: 95, Z: 14.5},
		},
		EscortSpawns: []safari.Vec3{
			{X: -1145, Y: 15, Z: 14.5},
			{X: -1095, Y: -55, Z: 11.2},
		},
		DefendSpawns: []safari.Vec3{
			{X: -920, Y: 175, Z: 22.0},
			{X: -780, Y: -75, Z: 11.2},
		},
	}
}
