package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"seegolauncher/internal/endpoints"
	"seegolauncher/internal/net"
	"seegolauncher/internal/paths"
	"seegolauncher/internal/utils"
	"sync"
	"time"

	"github.com/charmbracelet/log"
)

type CacheService struct{}

type NewsItem struct {
	Title     string
	Date      string
	Content   string
	Image     string
	ImageName string
}

const (
	TermsFileName = "terms.chc"
	NewsFileName  = "news.chc"
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

func (s *CacheService) GetAllNews() ([]*NewsItem, error) {
	return getAllNews()
}

func (s *CacheService) GetNewsImage(name string) []byte {
	image, err := getNewsImage(name)
	if err != nil {
		return nil
	}

	return image
}

func RefreshNews() (bool, error) {
	localLatestNew, err := getLatestNew()
	if err != nil {
		return false, err
	}

	body, err := net.RequestNewsFeed("simplefeed", 1, 0)
	if err != nil {
		return false, fmt.Errorf("Failed to fetch news content: %w", err)
	}

	var items []json.RawMessage
	if err := json.Unmarshal([]byte(body), &items); err != nil {
		return false, fmt.Errorf("Failed to parse remote news: %w", err)
	}

	var remoteLatestNew *NewsItem
	for _, raw := range items {
		var entry []string
		if err := json.Unmarshal(raw, &entry); err != nil {
			continue
		}
		if len(entry) == 1 {
			continue
		}
		if len(entry) < 2 {
			return false, fmt.Errorf("Corrupted remote news entry")
		}
		remoteLatestNew = &NewsItem{
			Title: entry[0],
			Date:  entry[1],
		}
		break
	}

	if remoteLatestNew == nil {
		return false, fmt.Errorf("No valid remote news entry found")
	}

	if localLatestNew.Date != remoteLatestNew.Date {
		log.Info("News was updated! Refreshing cache...")
		if err := refreshNews(); err != nil {
			return false, fmt.Errorf("Failed to refresh news: %w", err)
		}
	} else {
		log.Debug("Latest new")
		return false, nil
	}

	return true, nil
}

func getNewsImage(name string) ([]byte, error) {
	dir, err := paths.GetCachePath()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(dir, "news", name)

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	log.Infof("Loaded image %s", path)

	return data, nil
}

func getAllNews() ([]*NewsItem, error) {
	newsPath, err := paths.GetCachedFilePath("news", NewsFileName)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(newsPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read %s: %w", NewsFileName, err)
	}

	var items []json.RawMessage
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("Failed to parse %s: %w", NewsFileName, err)
	}
	if len(items) == 0 {
		return nil, fmt.Errorf("No news found")
	}

	allNews := make([]*NewsItem, 0, len(items))

	for _, item := range items {
		var entry []string
		if err := json.Unmarshal(item, &entry); err != nil {
			return nil, fmt.Errorf("Failed to parse news entry: %w", err)
		}
		if len(entry) < 5 {
			return nil, fmt.Errorf("Corrupted news")
		}
		news := &NewsItem{
			Title:     entry[0],
			Date:      entry[1],
			Content:   entry[2],
			ImageName: entry[3],
		}
		if news.ImageName != "" {
			_, err := paths.GetCachedFilePath("news", news.ImageName)
			if err != nil {
				return nil, err
			}
		}
		allNews = append(allNews, news)
	}
	return allNews, nil
}

func getLatestNew() (*NewsItem, error) {
	newsPath, err := paths.GetCachedFilePath("news", NewsFileName)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(newsPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read %s: %w", NewsFileName, err)
	}

	var items []json.RawMessage
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("Failed to parse %s: %w", NewsFileName, err)
	}
	if len(items) == 0 {
		return nil, fmt.Errorf("No news found")
	}

	var entry []string
	if err := json.Unmarshal(items[0], &entry); err != nil {
		return nil, fmt.Errorf("Failed to parse latest news entry: %w", err)
	}
	if len(entry) < 5 {
		return nil, fmt.Errorf("Corrupted news")
	}

	news := &NewsItem{
		Title:     entry[0],
		Date:      entry[1],
		Content:   entry[2],
		ImageName: entry[3],
	}

	if news.ImageName != "" {
		_, err := paths.GetCachedFilePath("news", news.ImageName)
		if err != nil {
			return nil, err
		}
	}

	return news, nil
}

func CheckHashes() error {
	termsRemote, err := net.Request(endpoints.Terms)
	if err != nil {
		return err
	}
	termsLocalPath, err := paths.GetCachedFilePath("", TermsFileName)
	if err != nil {
		return err
	}
	termsLocal, err := os.ReadFile(termsLocalPath)
	if err != nil {
		return err
	}
	if !utils.CompareHashes(termsRemote, string(termsLocal)) {
		return fmt.Errorf("The hashes different!")
	}

	latestNewsRemote, err := net.RequestNewsFeed("feed", 1, 0)
	if err != nil {
		return err
	}

	var remoteItems []json.RawMessage
	if err := json.Unmarshal([]byte(latestNewsRemote), &remoteItems); err != nil {
		return fmt.Errorf("Failed to parse remote news: %w", err)
	}
	if len(remoteItems) == 0 {
		return fmt.Errorf("No remote news found")
	}

	var remoteEntry []string
	if err := json.Unmarshal(remoteItems[0], &remoteEntry); err != nil {
		return fmt.Errorf("Failed to parse remote news: %w", err)
	}
	if len(remoteEntry) < 2 {
		return fmt.Errorf("Corrupted remote news")
	}
	remoteDate := remoteEntry[1]

	latestNew, err := getLatestNew()
	if err != nil {
		return err
	}

	if remoteDate != latestNew.Date {
		return fmt.Errorf("The news date is different!")
	}

	log.Info("Hashes okay!")

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

		log.Debug("Terms done!")

		// only need to refresh, if the cached latest new different from the remote
		if err := refreshNews(); err != nil {
			log.Errorf("Failed to refresh news: %v", err)
			return err
		}

		log.Debug("News done!")
	}

	return nil
}

func refreshTerms() error {
	log.Debug("Refreshing terms...")
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

func refreshNews() error {
	// 1/0 latest <- 0 always latest
	// each news have ["next"]
	// fetch until theres no ["next"] in it.
	// 1/i until no ["next"]
	log.Debug("Refreshing news...")
	page := 0
	count := 1

	var allNews []json.RawMessage
	var imageNames []string

	bar := utils.CreateProgressBar(-1, "Downloading news")

	for {
		body, err := net.RequestNewsFeed("feed", count, page)
		if err != nil {
			bar.Finish()
			return fmt.Errorf("Failed to fetch news content: %w", err)
		}
		var items []json.RawMessage
		if err := json.Unmarshal([]byte(body), &items); err != nil {
			bar.Finish()
			return fmt.Errorf("Failed to parse news content: %w", err)
		}
		continueFetch := false
		for _, raw := range items {
			var marker []string
			if err := json.Unmarshal(raw, &marker); err == nil && len(marker) == 1 && marker[0] == "next" {
				continueFetch = true
				continue
			}

			var entry []string
			if err := json.Unmarshal(raw, &entry); err == nil && len(entry) >= 4 && entry[3] != "" {
				imageNames = append(imageNames, entry[3])
			}
			allNews = append(allNews, raw)
		}
		if !continueFetch {
			break
		}
		page++
		bar.Set(page)
		// log.Debugf("Page: %d Count: %d", page, count)
		time.Sleep(250 * time.Millisecond)
	}

	bar.Finish()

	if len(imageNames) > 0 {
		log.Debugf("Downloading news images")
		if err := downloadNewsImages(imageNames); err != nil {
			return fmt.Errorf("Failed to download news images: %w", err)
		}
	}

	joined, err := json.Marshal(allNews)
	if err != nil {
		return fmt.Errorf("Failed to marshal news content: %w", err)
	}
	if err := writeCache("news", NewsFileName, string(joined)); err != nil {
		return fmt.Errorf("Failed to write news: %w", err)
	}
	return nil
}

func downloadNewsImages(names []string) error {
	if len(names) == 0 {
		return nil
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(names))

	sem := make(chan struct{}, 4)

	bar := utils.CreateProgressBar(len(names), "Downloading news images")
	for _, name := range names {
		wg.Add(1)
		go func(imgName string) {
			defer wg.Done()
			defer bar.Add(1)

			sem <- struct{}{}
			defer func() { <-sem }()

			imagesPath, err := paths.GetCachedFilePath("news", imgName)
			if err != nil {
				bar.Finish()
				errChan <- err
				return
			}

			if _, err := os.Stat(imagesPath); err == nil {
				bar.Finish()
				return
			}

			image, err := net.Request(endpoints.NewsImage + imgName)
			if err != nil {
				bar.Finish()
				errChan <- fmt.Errorf("failed to download %s: %w", imgName, err)
				return
			}

			if err := os.WriteFile(imagesPath, []byte(image), 0644); err != nil {
				bar.Finish()
				errChan <- fmt.Errorf("failed to save %s: %w", imgName, err)
				return
			}
		}(name)
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		bar.Finish()
		return <-errChan
	}

	bar.Finish()

	return nil
}
