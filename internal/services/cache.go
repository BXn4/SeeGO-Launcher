package services

import (
	"fmt"
	"os"
	"path/filepath"
	"seegolauncher/internal/endpoints"
	"seegolauncher/internal/net"
	"seegolauncher/internal/paths"
	"seegolauncher/internal/utils"

	"github.com/charmbracelet/log"
)

type CacheService struct{}

const (
	TermsFileName = "terms.chc"
)

func writeCache(dir, filename, v string) error {
	cacheDir, err := paths.GetCachePath()
	if err != nil {
		log.Errorf("Error getting cache path: %v", err)
		return err
	}

	file := filepath.Join(cacheDir, dir, filename)
	if err := os.WriteFile(file, []byte(v), 0644); err != nil {
		log.Errorf("Failed to write %s: %v", file, err)
		return err
	}
	return nil
}

func (s *CacheService) GetCachedTerms() (string, error) {
	dir, err := paths.GetCachePath()
	termsFile := filepath.Join(dir, TermsFileName)
	termsData, err := os.ReadFile(termsFile)

	if err != nil {
		log.Error(err)
		return err.Error(), err
	}

	return string(termsData), nil
}

func CheckHashes() error {
	termsRemote, err := net.Request(endpoints.Terms)
	if err != nil {
		return err
	}
	termsLocalPath, err := paths.GetCachedFilePath("", TermsFileName)
	termsLocal, err := os.ReadFile(termsLocalPath)
	if err != nil {
		return err
	}

	if !utils.CompareHashes(termsRemote, string(termsLocal)) {
		return fmt.Errorf("The hashes different!")
	}

	return nil
}

func LoadCache() error {
	if err := paths.CheckDirs(); err != nil {
		return err
	}

	if err := CheckHashes(); err != nil {
		log.Info("Hash mismatch! Refreshing cache...")

		if err := refreshTerms(); err != nil {
			log.Errorf("Failed to refresh terms: %v", err)
			return err
		}
		return err
	}

	return nil
}

func refreshTerms() error {
	termsPath, err := paths.GetCachedFilePath("", TermsFileName)
	if err != nil {
		return err
	}

	_, err = os.ReadFile(termsPath)
	termsMissing := os.IsNotExist(err)
	if err != nil && !termsMissing {
		return fmt.Errorf("Failed to read %s: %w", TermsFileName, err)
	}

	if termsMissing {
		terms, err := net.Request(endpoints.Terms)
		if err != nil {
			return fmt.Errorf("Failed to fetch terms content: %w", err)
		}

		if err := writeCache("", TermsFileName, terms); err != nil {
			return fmt.Errorf("Failed to write terms: %w", err)
		}
	}

	return nil
}
