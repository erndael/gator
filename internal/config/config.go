package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, configFileName), nil
}

func write(config Config) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	if err := os.WriteFile(configFilePath, data, 0666); err != nil {
		return err
	}
	return nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return write(*c)
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
