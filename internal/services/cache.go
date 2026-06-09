package services

import (
	"os"
	"path/filepath"
	"seegolauncher/internal/cache"

	"github.com/charmbracelet/log"
)

type CacheService struct {
}

func (s *CacheService) GetCachedTerms() (error, string) {
	cacheDir, err := cache.GetCachePath()
	if err != nil {
		log.Errorf("Error getting cache path: %v", err)
		return err, ""
	}

	termsFile := filepath.Join(cacheDir, cache.TermsFile)
	termsData, err := os.ReadFile(termsFile)

	return nil, string(termsData)
}
