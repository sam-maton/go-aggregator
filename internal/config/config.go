package config

import (
	"encoding/json"
	"os"
)

func Read() (Config, error) {

	configPath, err := getConfigFilePath()

	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(configPath)

	if err != nil {
		return Config{}, err
	}

	newConfig := Config{}

	err = json.Unmarshal(data, &newConfig)

	if err != nil {
		return Config{}, err
	}

	return newConfig, nil
}

func (conf *Config) SetUser() {

}
