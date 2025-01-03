package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	CurrentUserName string `json:"current_user_name"`
	DbUrl           string `json:"db_url"`
}

func Read() (Config, error) {
	var conf Config
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return conf, err
	}

	if err = json.Unmarshal(data, &conf); err != nil {
		return Config{}, err
	}

	return conf, nil

}

func (c Config) SetUser(name string) error {
	c.CurrentUserName = name
	if err := write(c); err != nil {
		return err
	}
	return nil
}
