package services

import (
	"encoding/json"
	"seegolauncher/data"
	"sync"
)

var localization = &Localization{
	cache: make(map[string]map[string]string),
}

type Localization struct {
	cache map[string]map[string]string
	mutex sync.RWMutex
}

func LocalizationService() *Localization {
	return localization
}

func (s *Localization) load(lang string) (map[string]string, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if v, ok := s.cache[lang]; ok {
		return v, nil
	}

	data, err := data.Locales.ReadFile("locales/" + lang + ".json")
	if err != nil {
		return nil, err
	}

	var m map[string]string
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}

	s.cache[lang] = m
	return m, nil
}

func (s *Localization) Get(key string, lang string) string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	data, err := s.load(lang)
	if err != nil {
		return key
	}

	if val, ok := data[key]; ok {
		return val
	}

	return key
}
