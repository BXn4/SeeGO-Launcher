package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"seegolauncher/internal/paths"
	"slices"
	"sync"

	"github.com/charmbracelet/log"
)

var (
	configInstance *Config
	configOnce     sync.Once
)

type ConfigData struct {
	Language      string `json:"language"`
	TermsAccepted bool   `json:"terms_accepted"`
	Theme         string `json:"theme"`
}

var DefaultConfig = ConfigData{
	Language:      "hu",
	TermsAccepted: false,
	Theme:         "dark",
}

var validLanguages = []string{"en", "hu"}
var validThemes = []string{"dark", "light"}

type Config struct {
	data *ConfigData
}

// getters
func (s *Config) GetLanguage() string {
	return s.data.Language
}

func (s *Config) GetLanguages() []string {
	return validLanguages
}

func (s *Config) GetTermsAccepted() bool {
	return s.data.TermsAccepted
}

func (s *Config) GetConfig() ConfigData {
	return *s.data
}

func (s *Config) GetTheme() string {
	return s.data.Theme
}

// setters
func (s *Config) SetLanguage(lang string) error {
	if !isValidLanguage(lang) {
		s.data.Language = "hu"
		saveConfig(s.data)
		return fmt.Errorf("Invalid language: %s, saved default", lang)
	}
	s.data.Language = lang
	return saveConfig(s.data)
}

func (s *Config) SetTheme(theme string) error {
	if !isValidTheme(theme) {
		s.data.Theme = "dark"
		saveConfig(s.data)
		return fmt.Errorf("Invalid theme: %s, saved default", theme)
	}
	s.data.Theme = theme
	return saveConfig(s.data)
}

func (s *Config) SetTermsAccepted() error {
	s.data.TermsAccepted = true
	return saveConfig(s.data)
}

func ConfigService() *Config {
	configOnce.Do(func() {
		data, err := LoadConfig()
		if err != nil {
			log.Warn("Failed to load config, using defaults", "err", err)
			defaultCopy := DefaultConfig
			data = &defaultCopy
		}
		configInstance = &Config{data: data}
	})
	return configInstance
}

func isValidLanguage(lang string) bool {
	if slices.Contains(validLanguages, lang) {
		return true
	}
	return false
}

func isValidTheme(theme string) bool {
	if slices.Contains(validThemes, theme) {
		return true
	}
	return false
}

func getConfigPath() (string, error) {
	home, err := paths.GetRootPath()
	if err != nil {
		return err.Error(), err
	}
	return filepath.Join(home, "config.json"), nil
}

func LoadConfig() (*ConfigData, error) {
	path, err := getConfigPath()
	if err != nil {
		return nil, err
	}
	log.Debugf("Loading config from: %s", path)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Info("config.json not found, creating and using default")
		return createDefaultConfig()
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read config: %w", err)
	}

	var config ConfigData
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("Failed to parse config: %w", err)
	}

	if !isValidLanguage(config.Language) {
		log.Warnf("Invalid language '%s', using default '%s'", config.Language, DefaultConfig.Language)
		config.Language = DefaultConfig.Language
		if err := saveConfig(&config); err != nil {
			log.Warn("Failed to save config", "err", err)
		}
	}

	if !isValidTheme(config.Theme) {
		log.Warnf("Invalid theme '%s', using default '%s'", config.Theme, DefaultConfig.Theme)
		config.Theme = DefaultConfig.Theme
		if err := saveConfig(&config); err != nil {
			log.Warn("Failed to save config", "err", err)
		}
	}

	log.Info("Config loaded", "language", config.Language)
	log.Info("Config loaded", "terms", config.TermsAccepted)
	log.Info("Config loaded", "theme", config.Theme)
	return &config, nil
}

func createDefaultConfig() (*ConfigData, error) {
	config := &ConfigData{Language: DefaultConfig.Language, TermsAccepted: DefaultConfig.TermsAccepted, Theme: DefaultConfig.Theme}
	if err := saveConfig(config); err != nil {
		return nil, fmt.Errorf("Failed to create default config: %w", err)
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
		return fmt.Errorf("Failed to create config directory: %w", err)
	}
	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return fmt.Errorf("Failed to serialize config: %w", err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("Failed to write config file: %w", err)
	}
	log.Info("Config saved", "language", config.Language)
	log.Info("Config saved", "terms", config.TermsAccepted)
	log.Info("Config saved", "theme", config.Theme)
	return nil
}
