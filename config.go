// config.go

package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config represents the application's configuration.
type Config struct {
	GitHubToken string `json:"github_token"`
	TableName   string `json:"table_name"`
	Owner       string `json:"owner"`
	Repo        string `json:"repo"`
}

// ConfigFileName is the name of the config file.
const ConfigFileName = "config.json"

// GetConfigFilePath returns the path to the config file in the user's home directory.
func GetConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configPath := filepath.Join(homeDir, ConfigFileName)
	return configPath, nil
}

// CreateOrReadConfigFile creates a new configuration file if does not exist, otherwise reads the existing file.
func CreateOrReadConfigFile() (*Config, error) {
	configPath, err := GetConfigFilePath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		config := Config{}
		err = WriteConfig(&config)
		if err != nil {
			return nil, err
		}
		return &config, nil
	}

	return ReadConfig()
}

// ReadConfig reads the configuration file and returns the Config struct.
func ReadConfig() (*Config, error) {
	configPath, err := GetConfigFilePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// WriteConfig writes the Config struct to the configuration file.
func WriteConfig(config *Config) error {
	configPath, err := GetConfigFilePath()

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, data, 0600)
	if err != nil {
		return err
	}

	return nil
}
