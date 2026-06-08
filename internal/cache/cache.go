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

const (
	TermsFile string = "terms.ver"
)

type CacheData struct {
	TermsDate string
}

func getCachePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Could not get home directory: %w", err)
	}
	return filepath.Join(home, ".config", "seego-launcher", "cache"), nil
}

func getCacheFilePath(filename string) (string, error) {
	dir, err := getCachePath()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, filename), nil
}

func checkCacheDir() error {
	dir, err := getCachePath()
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

func LoadCache() {
	if err := checkCacheDir(); err != nil {
		fmt.Println("Error creating cache directory:", err)
		return
	}
	path, err := getCachePath()
	if err != nil {
		fmt.Println("Error getting cache path:", err)
		return
	}

	if _, err := os.Stat(path + "/" + TermsFile); os.IsNotExist(err) {
		log.Infof("%s is not found, requesting it.", TermsFile)
		response, err := request(endpoints.TermsDate)
		if err != nil {
			log.Errorf("Failed to request terms: %s", err)
			return
		}
		if err := os.WriteFile(path+"/"+TermsFile, []byte(response), 0644); err != nil {
			log.Errorf("Failed to write %s: %s", TermsFile, err)
			return
		}
	}
}
