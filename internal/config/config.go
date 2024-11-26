package config

import (
	"encoding/json"
	"os"
)

const configName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configPath := homePath + "/" + configName

	return configPath, nil
}

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

func write(conf Config) error {
	data, err := json.Marshal(conf)

	if err != nil {
		return err
	}

	filePath, err := getConfigFilePath()

	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0666)

	if err != nil {
		return err
	}

	return nil
}

func (conf *Config) SetUser(userName string) error {
	conf.UserName = userName

	err := write(*conf)

	if err != nil {
		return err
	}

	return nil
}
