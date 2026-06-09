package main

import "fmt"

type Plugin struct {
	cfg Config
}

func newPlugin(cfg Config) *Plugin {
	return &Plugin{cfg: cfg}
}

func (p *Plugin) initialise() {
	if p.cfg.ServerName != "" {
		API.Server.SetName(p.cfg.ServerName)
	}
	if p.cfg.GameModeText != "" {
		API.Server.SetGameModeText(p.cfg.GameModeText)
	}
	API.Server.Log(fmt.Sprintf("[%s] SDK ready — register handlers on `events` and use `API`", PluginName))
}

func (p *Plugin) shutdown() {
	API.Server.Log(fmt.Sprintf("[%s] SDK shutdown", PluginName))
}
