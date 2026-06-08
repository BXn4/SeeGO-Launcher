package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/charmbracelet/log"
)

type ConfigData struct {
	Language string `json:"language"`
}

var DefaultConfig = ConfigData{
	Language: "hu",
}

var validLanguages = []string{"en", "hu"}

type Config struct {
	data *ConfigData
}

func ConfigService() *Config {
	data, err := LoadConfig()
	if err != nil {
		log.Warn("Failed to load config, using defaults", "err", err)
		defaultCopy := DefaultConfig
		data = &defaultCopy
	}
	return &Config{data: data}
}

func (s *Config) GetLanguage() string {
	return s.data.Language
}

func (s *Config) SetLanguage(lang string) error {
	if !isValidLanguage(lang) {
		return fmt.Errorf("invalid language: %s", lang)
	}
	s.data.Language = lang
	return saveConfig(s.data)
}

func (s *Config) GetConfig() ConfigData {
	return *s.data
}

func isValidLanguage(lang string) bool {
	if slices.Contains(validLanguages, lang) {
		return true
	}
	return false
}

func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get home directory: %w", err)
	}
	return filepath.Join(home, ".config", "seego-launcher", "config.json"), nil
}

func LoadConfig() (*ConfigData, error) {
	path, err := getConfigPath()
	if err != nil {
		return nil, err
	}
	log.Debugf("Loading config from: %s", path)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Info("config.json not found, creating default")
		return createDefaultConfig()
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var config ConfigData
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	if !isValidLanguage(config.Language) {
		log.Warnf("Invalid language '%s', reverting to default '%s'", config.Language, DefaultConfig.Language)
		config.Language = DefaultConfig.Language
		if err := saveConfig(&config); err != nil {
			log.Warn("Failed to save corrected config", "err", err)
		}
	}

	log.Info("Config loaded", "language", config.Language)
	return &config, nil
}

func createDefaultConfig() (*ConfigData, error) {
	config := &ConfigData{Language: DefaultConfig.Language}
	if err := saveConfig(config); err != nil {
		return nil, fmt.Errorf("failed to create default config: %w", err)
	}
	log.Info("Created default config.json")
	return config, nil
}

func saveConfig(config *ConfigData) error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}
	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to serialize config: %w", err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	log.Info("Config saved", "language", config.Language)
	return nil
}
