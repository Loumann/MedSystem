package configs

import (
	"encoding/json"
	"os"
)

const configPath = "configs/config.json"

type Config struct {
	Host     string `json:"host" binding:"required"`
	Port     string `json:"port" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	DBName   string `json:"db_name" binding:"required"`
	SSLMode  string `json:"ssl_mode" binding:"required"`
}

func GetConfig() (*Config, error) {
	var config Config

	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
