package cache

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"seegolauncher/internal/endpoints"
	"seegolauncher/internal/net"

	"github.com/charmbracelet/log"
)

// cache all news (cache/news), store it in json with timestamp
// due to i want to make the app work offline too (with connection disabled),
// i need to store the images
// always show the latest on top
// and also, create categories. changelogs etc, but i dont think that theyre using categories
const (
	TermsVersionFileName = "terms.ver"
	TermsFileName        = "terms.chc"
)

func GetCachePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Could not get home directory: %w", err)
	}

	if runtime.GOOS == "windows" {
		return filepath.Join(home, "ProgramData", "seego-launcher", "cache"), nil
	}
	return filepath.Join(home, ".config", "seego-launcher", "cache"), nil
}

func getCachedFilePath(dir, name string) (string, error) {
	path, err := GetCachePath()
	if err != nil {
		return "", fmt.Errorf("Could not get home directory: %w", err)
	}
	return filepath.Join(path, dir, name), nil
}

func checkCacheDir() error {
	dir, err := GetCachePath()
	if err != nil {
		return err
	}

	return os.MkdirAll(dir, 0755)
}

func checkNewsDir() error {
	dir, err := GetCachePath()
	if err != nil {
		return err
	}
	dir = filepath.Join(dir, "news")
	return os.MkdirAll(dir, 0755)
}

func checkShopDir() error {
	dir, err := GetCachePath()
	if err != nil {
		return err
	}
	dir = filepath.Join(dir, "shop")
	return os.MkdirAll(dir, 0755)
}

func checkDirs() error {
	err := checkCacheDir()
	if err != nil {
		return fmt.Errorf("Failed to check the cache dir: %s", err)
	}

	err = checkNewsDir()
	if err != nil {
		return fmt.Errorf("Failed to check the news dir: %s", err)
	}

	err = checkShopDir()
	if err != nil {
		return fmt.Errorf("Failed to check the shop dir: %s", err)
	}

	return nil
}

func writeCache(dir, filename, v string) error {
	cacheDir, err := GetCachePath()
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

func RequestTermsDate() (string, error) {
	response, err := net.Request(endpoints.TermsDate)
	if err != nil {
		log.Errorf("Failed to request terms date: %s", err)
		return "", err
	}

	return response, err
}

func LoadCache() bool {
	if err := checkCacheDir(); err != nil {
		log.Errorf("Error creating cache directory: %v", err)
		return false
	}

	if err := checkDirs(); err != nil {
		log.Errorf("Error during checking cache directory: %v", err)
		return false
	}

	termsDateresponse, _ := RequestTermsDate()
	termsDatePath, err := getCachedFilePath("", TermsVersionFileName)
	if err != nil {
		return false
	}

	termsDateData, err := os.ReadFile(termsDatePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Infof("%s not found, requesting it", TermsVersionFileName)
			writeCache("", TermsVersionFileName, string(termsDateresponse))
		} else {
			log.Errorf("Failed to read %s: %v", TermsVersionFileName, err)
			return false
		}
	}

	termsFilePath, err := getCachedFilePath("", TermsFileName)
	if err != nil {
		return false
	}

	_, err = os.ReadFile(termsFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Infof("%s not found, requesting it", TermsFileName)
			termsDataResponse, _ := net.Request(endpoints.Terms)
			writeCache("", TermsFileName, string(termsDataResponse))
		} else {
			log.Errorf("Failed to read %s: %v", TermsFileName, err)
			return false
		}
	}

	if string(termsDateData) != termsDateresponse {
		writeCache("", TermsVersionFileName, string(termsDateresponse))

		response, _ := net.Request(endpoints.Terms)
		writeCache("", TermsFileName, string(response))
	}

	log.Info("Everything is cached!")

	return true
}
