package main

import (
	"encoding/json"
	"os"
)

const configFile = "goserver.json"

type Config struct {
	ServerName   string `json:"server_name"`
	GameModeText string `json:"gamemode_text"`
}

func defaultConfig() Config {
	return Config{
		ServerName:   "VC:MP Go SDK",
		GameModeText: "Go SDK Host",
	}
}

func loadConfig() Config {
	cfg := defaultConfig()
	data, err := os.ReadFile(configFile)
	if err != nil {
		return cfg
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return defaultConfig()
	}
	if cfg.ServerName == "" {
		cfg.ServerName = defaultConfig().ServerName
	}
	return cfg
}
