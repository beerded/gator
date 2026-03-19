package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName string = ".gatorconfig.json"

type Config struct {
	DBUrl			 string		`json:"db_url"`
	CurrentUserName  string		`json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("%w",err)
	}
	configFilePath := path + "/" + configFileName
	return configFilePath, nil
}

func Read() (*Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return nil, fmt.Errorf("%w",err)
	}
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("%w",err)
	}

	cfg := Config{}
	// unmarshal the json in the homedir
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("%w",err)
	}
	return &cfg, nil
}

func (cfg *Config) SetUser(user string) error {
	if cfg == nil {
		return fmt.Errorf("Nil Config")
	}
	cfg.CurrentUserName = user
	err := write(cfg)
	if err != nil {
		return err
	}
	return nil
}

func (cfg *Config) Print() error {
	if cfg == nil {
		return fmt.Errorf("Nil Config")
	}
	fmt.Println("**** PRINTING CONFIG ****")
	fmt.Printf("    db_url: %s\n", cfg.DBUrl)
	fmt.Printf("    current_user_name: %s\n", cfg.CurrentUserName)
	return nil
}

func write(cfg *Config) error {
	//Marshal the struct into json bytes and then write it to the
	//cfg filepath
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("%w",err)
	}
	//Marshal the data
	jsonData, err := json.Marshal(*cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(configFilePath, jsonData, 0666)
	if err != nil {
		return err
	}
	fmt.Println("Wrote config")

	return nil
}
