package cache

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"seegolauncher/internal/endpoints"

	"github.com/charmbracelet/log"
)

// cache all news (cache/news), store it in json with timestamp
// due to i want to make the app work offline too (with connection disabled),
// i need to store the images
// always show the latest on top
// and also, create categories. changelogs etc, but i dont think that theyre using categories
const (
	TermsDate = "terms.ver"
	TermsFile = "terms.chc"
)

func GetCachePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Could not get home directory: %w", err)
	}
	return filepath.Join(home, ".config", "seego-launcher", "cache"), nil
}

func getNewsCachePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Could not get home directory: %w", err)
	}
	return filepath.Join(home, ".config", "seego-launcher", "cache", "news"), nil
}

func getCacheFilePath(filename string) (string, error) {
	dir, err := GetCachePath()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, filename), nil
}

func checkCacheDir() error {
	dir, err := GetCachePath()
	if err != nil {
		return err
	}
	return os.MkdirAll(dir, 0755)
}

func checkNewsCacheDir() error {
	dir, err := getNewsCachePath()
	if err != nil {
		return err
	}
	return os.MkdirAll(dir, 0755)
}

func request(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func RequestTermsDate() (string, error) {
	response, err := request(endpoints.TermsDate)
	if err != nil {
		log.Errorf("Failed to request terms: %s", err)
		return "", err
	}

	return response, err
}

func WriteTermsDate(v string) error {
	cacheDir, err := GetCachePath()
	if err != nil {
		log.Errorf("Error getting cache path: %v", err)
		return err
	}

	termsFile := filepath.Join(cacheDir, TermsDate)

	if err := os.WriteFile(termsFile, []byte(v), 0644); err != nil {
		log.Errorf("Failed to write %s: %v", TermsDate, err)
		return err
	}
	return nil
}

func RequestTerms() (string, error) {
	response, err := request(endpoints.Terms)
	if err != nil {
		log.Errorf("Failed to request terms: %s", err)
		return "", err
	}

	return response, err
}

func WriteTerms(v string) error {
	cacheDir, err := GetCachePath()
	if err != nil {
		log.Errorf("Error getting cache path: %v", err)
		return err
	}

	termsFile := filepath.Join(cacheDir, TermsFile)

	if err := os.WriteFile(termsFile, []byte(v), 0644); err != nil {
		log.Errorf("Failed to write %s: %v", TermsFile, err)
		return err
	}
	return nil
}

func LoadCache() {
	if err := checkCacheDir(); err != nil {
		log.Errorf("Error creating cache directory: %v", err)
		return
	}

	if err := checkNewsCacheDir(); err != nil {
		log.Errorf("Error creating news cache directory: %v", err)
		return
	}

	cacheDir, err := GetCachePath()
	if err != nil {
		log.Errorf("Error getting cache path: %v", err)
		return
	}

	termsDatePath := filepath.Join(cacheDir, TermsDate)
	termsFilePath := filepath.Join(cacheDir, TermsFile)
	termsDateresponse, _ := RequestTermsDate()

	termsDateData, err := os.ReadFile(termsDatePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Infof("%s not found, requesting it", TermsDate)
			WriteTermsDate(termsDateresponse)
		} else {
			log.Errorf("Failed to read %s: %v", TermsDate, err)
			return
		}
	}

	_, err = os.ReadFile(termsFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Infof("%s not found, requesting it", TermsFile)
			termsDataResponse, _ := RequestTerms()
			WriteTerms(termsDataResponse)
		} else {
			log.Errorf("Failed to read %s: %v", TermsFile, err)
			return
		}
	}

	if string(termsDateData) != termsDateresponse {
		WriteTermsDate(string(termsDateresponse))

		response, _ := RequestTerms()
		WriteTerms(response)
	}
}
