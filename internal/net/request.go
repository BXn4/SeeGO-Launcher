package net

import (
	"fmt"
	"io"
	"net/http"
	"seegolauncher/internal/endpoints"

	"github.com/charmbracelet/log"
)

func Request(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Error(err)
		return err.Error(), err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
		return err.Error(), err
	}
	return string(body), nil
}

func RequestTermsVersion() (string, error) {
	response, err := Request(endpoints.TermsDate)
	if err != nil {
		return err.Error(), err
	}
	return response, err
}

func RequestCategories(account string) (string, error) {
	url := fmt.Sprintf("%s/%s/categories/", endpoints.Store, account)
	response, err := Request(url)
	if err != nil {
		return err.Error(), err
	}
	return response, err
}

func RequestItems(account string, categoryID int) (string, error) {
	url := fmt.Sprintf("%s/%s/categories/%d?includePackages=1", endpoints.Store, account, categoryID)
	response, err := Request(url)
	if err != nil {
		return err.Error(), err
	}
	return response, err
}

func RequestSeeFileList() (string, error) {
	url := endpoints.FilesList
	response, err := Request(url)
	if err != nil {
		return err.Error(), err
	}
	return response, err
}

func RequestNewsFeed(feed string, count, page int) (string, error) {
	// "simplefeed" to the titles comment, and "feed" to the content
	// im not going to use the simplefeed
	url := fmt.Sprintf("%s/%s/launcher_rpg/%d/%d", endpoints.News, feed, count, page)

	response, err := Request(url)
	if err != nil {
		return err.Error(), err
	}

	return response, err
}

func RequestNewsImage(name string) (string, error) {
	url := fmt.Sprintf("%s/%s", endpoints.NewsImage, name)

	response, err := Request(url)
	if err != nil {
		return err.Error(), err
	}

	return response, err
}
