package config

import "os"

const configName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configPath := homePath + "/" + configName

	return configPath, nil
}
