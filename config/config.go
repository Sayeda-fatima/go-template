package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type DBConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

type Config struct {
	Port     string   `json:"port"`
	DB       DBConfig `json:"db"`
	LogLevel string   `json:"log_level"`
}

var AppConfig Config

func Load() error {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	configFile := fmt.Sprintf("config/%s.json", env)
	file, err := os.Open(configFile)
	if err != nil {
		return fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&AppConfig); err != nil {
		return fmt.Errorf("failed to decode config JSON: %w", err)
	}

	fmt.Printf("Loaded %s environment config\n", env)
	return nil
}
