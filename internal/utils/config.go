package utils

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
)

type Config struct {
	Language string `json:"language"`
}

var DefaultConfig = Config{
	Language: "hu",
}

func getPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".config", "seego-launhcer", "config.json"), nil
}

func LoadConfig() (*Config, error) {
	path, err := getPath()
	log.Debugf("Loading config from: %s", path)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Info("The config.json is not exist, creating one with default config")
		return createDefaultConfig()
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	println(config.Language)

	if config.Language != "en" && config.Language != "hu" {
		log.Warnf("The language is invalid, using default one: %s", DefaultConfig.Language)
		config.Language = DefaultConfig.Language

		SaveConfig(&config)
	}

	log.Info("App config:", "language", config.Language)

	return &config, nil
}

func createDefaultConfig() (*Config, error) {
	path, err := getPath()
	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return nil, err
	}

	config := &Config{
		Language: DefaultConfig.Language,
	}

	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return nil, err
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return nil, err
	}

	log.Info("Created a new config.json")
	SaveConfig(config)

	return config, nil
}

func SaveConfig(config *Config) error {
	path, err := getPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	log.Debug(config.Language)

	log.Info("Saved config")

	return os.WriteFile(path, data, 0644)
}
