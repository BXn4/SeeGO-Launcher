package paths

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func GetRootPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	if runtime.GOOS == "windows" {
		return filepath.Join(home, "../", "ProgramData", "seego-launcher"), nil
	}
	return filepath.Join(home, ".config", "seego-launcher"), nil
}

func GetCachePath() (string, error) {
	path, err := GetRootPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(path, "cache"), nil
}

func GetCachedFilePath(dir, name string) (string, error) {
	path, err := GetCachePath()
	if err != nil {
		return "", fmt.Errorf("Could not get home directory: %w", err)
	}
	return filepath.Join(path, dir, name), nil
}

func GetSeePath() (string, error) {
	path, err := GetRootPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(path, "see"), nil
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

func checkStoreDir() error {
	dir, err := GetCachePath()
	if err != nil {
		return err
	}
	dir = filepath.Join(dir, "store")
	return os.MkdirAll(dir, 0755)
}

func checkSeeDir() error {
	dir, err := GetRootPath()
	if err != nil {
		return err
	}
	dir = filepath.Join(dir, "see")
	return os.MkdirAll(dir, 0755)
} // see files will be stored here (SeeLenium, SeeRPGClient, srpg_logo)
// the launcher will starts the seerpg client, and injects the payload, and creates a sub process between this app (later, this is important)

func CheckDirs() error {
	if err := checkCacheDir(); err != nil {
		return fmt.Errorf("Failed to check the cache dir: %s", err)
	}
	if err := checkNewsDir(); err != nil {
		return fmt.Errorf("Failed to check the news dir: %s", err)
	}
	if err := checkStoreDir(); err != nil {
		return fmt.Errorf("Failed to check the store dir: %s", err)
	}
	if err := checkSeeDir(); err != nil {
		return fmt.Errorf("Failed to check the see dir: %s", err)
	}
	return nil
}
