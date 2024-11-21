package config

import (
	"encoding/json"
	"fmt"
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

func (conf *Config) SetUser(userName string) {
	conf.UserName = userName

	err := write(*conf)

	if err != nil {
		fmt.Println(err)
	}
}
