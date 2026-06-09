package main

/*
#cgo CFLAGS: -I${SRCDIR}/../../include
#include "plugin.h"
*/
import "C"

import (
	"encoding/json"
	"os"

	"github.com/masteroz/vcmp-go-plugin/vcmp"
)

const configFile = "goserver.json"

type config struct {
	ServerName   string `json:"server_name"`
	GameModeText string `json:"gamemode_text"`
}

func loadConfig() config {
	cfg := config{ServerName: "VC:MP Go SDK", GameModeText: "Go SDK Host"}
	data, err := os.ReadFile(configFile)
	if err != nil {
		return cfg
	}
	_ = json.Unmarshal(data, &cfg)
	if cfg.ServerName == "" {
		cfg.ServerName = "VC:MP Go SDK"
	}
	return cfg
}

func init() {
	vcmp.MetaProvider = func() vcmp.PluginMeta {
		return vcmp.PluginMeta{Name: "GoSDK", Version: 0x00020000}
	}
	vcmp.OnLoad = func() {
		cfg := loadConfig()
		if cfg.ServerName != "" {
			vcmp.API.Server.SetName(cfg.ServerName)
		}
		if cfg.GameModeText != "" {
			vcmp.API.Server.SetGameModeText(cfg.GameModeText)
		}
		vcmp.API.Server.Log("[GoSDK] ready — register handlers on vcmp.Events in hooks.go")
	}
}

func main() {}
