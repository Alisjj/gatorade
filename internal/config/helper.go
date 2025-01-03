package config

import (
	"encoding/json"
	"os"
)

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	path := homeDir + "/.gatorconfig.json"
	if err != nil {
		return "", err
	}

	return path, nil
}
func write(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	if err = os.WriteFile(path, data, 0644); err != nil {
		return err
	}

	return nil
}
