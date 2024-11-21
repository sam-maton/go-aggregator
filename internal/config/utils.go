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
