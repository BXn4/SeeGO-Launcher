package services

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"crypto/md5"

	"seegolauncher/internal/endpoints"
	"seegolauncher/internal/net"
	"seegolauncher/internal/paths"

	"github.com/charmbracelet/log"
)

type CacheService struct{}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Item struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	TotalPrice  float64 `json:"total_price"`
	Currency    string  `json:"currency"`
	CreatedAt   string  `json:"created_at"`
}

type ItemDetail struct {
	Data struct {
		Packages []Item `json:"packages"`
	} `json:"data"`
}

const (
	TermsVersionFileName = "terms.ver"
	TermsFileName        = "terms.chc"
	CategoriesFileName   = "categories.chc"
	ItemsFileName        = "items.chc"
)

func (s *CacheService) GetCachedTerms() (string, error) {
	dir, err := paths.GetCachePath()
	if err != nil {
		return "Error getting cache path", err
	}

	termsFile := filepath.Join(dir, TermsFileName)
	termsData, err := os.ReadFile(termsFile)

	return string(termsData), nil
}

func RequestCategories() (string, error) {
	if OA != "" {
		url := fmt.Sprintf("%s/%s/categories/", endpoints.Store, OA)
		response, err := net.Request(url)
		if err != nil {
			log.Errorf("Failed to request categories: %s", err)
			return "", err
		}
		return response, err
	}
	return "", nil
}

func RequestItems(categoryID int) (string, error) {
	if OA != "" {
		url := fmt.Sprintf("%s/%s/categories/%d?includePackages=1", endpoints.Store, OA, categoryID)
		response, err := net.Request(url)
		if err != nil {
			log.Errorf("Failed to request items: %s", err)
			return "", err
		}
		return response, err
	}
	return "", nil
}

func RequestSeeFileList() (string, error) {
	url := endpoints.FilesList
	response, err := net.Request(url)
	if err != nil {
		log.Errorf("Failed to request file list: %s", err)
		return "", err
	}
	return response, err
}

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

func CompareHashes(remote, local string) bool {
	remoteHash := GetMD5Hash(remote)
	localHash := GetMD5Hash(local)

	return remoteHash == localHash
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func CheckRemoteLocalHash() (bool, error) {
	termsDateRemote, err := net.Request(endpoints.TermsDate)
	if err != nil {
		log.Errorf("Failed to request remote hash: %s", err)
		return false, err
	}
	termsVersionPath, err := paths.GetCachedFilePath("", TermsVersionFileName)
	termsDateLocal, err := os.ReadFile(termsVersionPath)
	if err != nil {
		log.Errorf("Failed to read %s: %v", TermsVersionFileName, err)
		return false, nil
	}

	termsDataRemote, err := net.Request(endpoints.Terms)
	if err != nil {
		log.Errorf("Failed to request remote hash: %s", err)
		return false, nil
	}
	termsDataPath, err := paths.GetCachedFilePath("", TermsFileName)
	termsDataLocal, err := os.ReadFile(termsDataPath)
	if err != nil {
		log.Errorf("Failed to read %s: %v", TermsFileName, err)
		return false, nil
	}

	categoriesRemote, err := RequestCategories()
	if err != nil {
		log.Errorf("Failed to request remote hash: %s", err)
		return false, nil
	}
	categoriesPath, err := paths.GetCachedFilePath("store", CategoriesFileName)
	categoriesLocal, err := os.ReadFile(categoriesPath)
	if err != nil {
		log.Errorf("Failed to read %s: %v", CategoriesFileName, err)
		return false, nil
	}

	if !CompareHashes(termsDateRemote, string(termsDateLocal)) {
		return false, nil
	}
	if !CompareHashes(termsDataRemote, string(termsDataLocal)) {
		return false, nil
	}
	if !CompareHashes(categoriesRemote, string(categoriesLocal)) {
		return false, nil
	}

	log.Info("Hashes okay!")
	return true, nil
}

func LoadCache() bool {
	if err := paths.CheckDirs(); err != nil {
		log.Errorf("Failed to check cache directories: %v", err)
		return false
	}

	ok, err := CheckRemoteLocalHash()
	if err != nil {
		log.Errorf("Failed to check hashes: %v", err)
		return false
	}
	if ok {
		log.Info("Everything is cached!")
		return true
	}

	log.Info("Hash mismatch! Refreshing cache...")

	if err := refreshTerms(); err != nil {
		log.Errorf("Failed to refresh terms: %v", err)
		return false
	}

	if err := refreshCategories(); err != nil {
		log.Errorf("Failed to refresh categories: %v", err)
		return false
	}

	if err := refreshItems(); err != nil {
		log.Errorf("Failed to refresh items: %v", err)
		return false
	}

	log.Info("Everything is cached!")
	return true
}

func refreshTerms() error {
	remoteVersion, err := net.RequestTermsVersion()
	if err != nil {
		return fmt.Errorf("Failed to request terms version: %w", err)
	}

	versionPath, err := paths.GetCachedFilePath("", TermsVersionFileName)
	if err != nil {
		return err
	}

	termsPath, err := paths.GetCachedFilePath("", TermsFileName)
	if err != nil {
		return err
	}

	localVersion, err := os.ReadFile(versionPath)
	versionMissing := os.IsNotExist(err)
	if err != nil && !versionMissing {
		return fmt.Errorf("Failed to read %s: %w", TermsVersionFileName, err)
	}

	localTerms, err := os.ReadFile(termsPath)
	termsMissing := os.IsNotExist(err)
	if err != nil && !termsMissing {
		return fmt.Errorf("Failed to read %s: %w", TermsFileName, err)
	}

	versionStale := versionMissing || string(localVersion) != remoteVersion

	remoteContent, err := net.Request(endpoints.Terms)
	if err != nil {
		return fmt.Errorf("Failed to fetch terms content: %w", err)
	}

	contentStale := termsMissing || !CompareHashes(remoteContent, string(localTerms))

	if versionStale || contentStale {
		if err := writeCache("", TermsVersionFileName, remoteVersion); err != nil {
			return fmt.Errorf("Failed to write version: %w", err)
		}
		if err := writeCache("", TermsFileName, remoteContent); err != nil {
			return fmt.Errorf("Failed to write terms: %w", err)
		}
	}

	return nil
}

func refreshCategories() error {
	path, err := paths.GetCachedFilePath("store", CategoriesFileName)
	if err != nil {
		return err
	}

	remoteData, err := RequestCategories()
	if err != nil {
		return fmt.Errorf("Failed to fetch categories: %w", err)
	}

	localData, err := os.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("Failed to read %s: %w", CategoriesFileName, err)
	}

	if os.IsNotExist(err) || !CompareHashes(remoteData, string(localData)) {
		return writeCache("store", CategoriesFileName, remoteData)
	}

	return nil
}

func refreshItems() error {
	itemsPath, err := paths.GetCachedFilePath("store", ItemsFileName)
	if err != nil {
		return err
	}

	if _, err := os.ReadFile(itemsPath); err == nil {
		return nil // already cached
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("Failed to read %s: %w", ItemsFileName, err)
	}

	log.Infof("%s not found, requesting it", ItemsFileName)

	cats, err := loadCachedCategories()
	if err != nil {
		return err
	}

	items, err := fetchItemsConcurrently(cats)
	if err != nil {
		return err
	}

	itemsJSON, err := json.Marshal(items)
	if err != nil {
		return fmt.Errorf("Failed to marshal items: %w", err)
	}

	return writeCache("store", ItemsFileName, string(itemsJSON))
}

func loadCachedCategories() ([]Category, error) {
	path, err := paths.GetCachedFilePath("store", CategoriesFileName)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read %s: %w", CategoriesFileName, err)
	}

	var payload struct {
		Data []Category `json:"data"`
	}
	if err := json.Unmarshal(data, &payload); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal categories: %w", err)
	}

	return payload.Data, nil
}

func fetchItemsConcurrently(cats []Category) ([]Item, error) {
	type result struct {
		items []Item
		err   error
	}

	results := make([]result, len(cats))

	var wg sync.WaitGroup
	for i, c := range cats {
		wg.Add(1)
		go func(i int, c Category) {
			defer wg.Done()

			data, err := RequestItems(c.ID)
			if err != nil {
				results[i] = result{err: err}
				return
			}

			var payload ItemDetail
			if err := json.Unmarshal([]byte(data), &payload); err != nil {
				results[i] = result{err: fmt.Errorf("Failed to unmarshal items for category %d: %w", c.ID, err)}
				return
			}

			results[i] = result{items: payload.Data.Packages}
			time.Sleep(500 * time.Millisecond) // rate limit
		}(i, c)
	}
	wg.Wait()

	var items []Item
	for _, r := range results {
		if r.err != nil {
			return nil, r.err
		}
		items = append(items, r.items...)
	}

	return items, nil
}
