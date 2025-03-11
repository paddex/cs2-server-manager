package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Logging   LoggingConfig
	WebServer WebServerConfig
}

type LoggingConfig struct {
	ErrorLog        string `json:"error_log"`
	AccessLog       string `json:"access_log"`
	LogLevel        string `json:"log_level"`
	EnableAccessLog bool   `json:"enable_access_log"`
}

type WebServerConfig struct {
	Host     string
	Port     int
	BasePath string
}

func ReadConfig() (*Config, error) {
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		return nil, fmt.Errorf("Error beim Lesen der Config-Datei: %w", err)
	}

	var cfg Config
	err = json.Unmarshal(configFile, &cfg)
	if err != nil {
		return nil, fmt.Errorf("Error beim Parsen der Config-Datei: %w", err)
	}
	if cfg.WebServer.BasePath == "/" {
		cfg.WebServer.BasePath = ""
	}

	return &cfg, nil
}
