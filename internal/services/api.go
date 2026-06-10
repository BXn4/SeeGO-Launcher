package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"seegolauncher/internal/endpoints"
)

type API struct{}

var OA string

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

type ServerDetail struct {
	Players int `json:"players"`
	Admins  int `json:"admins"`
	Queue   int `json:"queue"`
	Slots   int `json:"slots"`
}

type CategoryReponse struct {
	Data []Category `json:"data"`
}

func (a *API) GetCategories() ([]Category, error) {
	// {"data":[{"id":3241741,
	// https://headless.tebex.io/api/accounts/
	if OA != "" {
		url := fmt.Sprintf("%s/%s/categories/", endpoints.Store, OA)
		response, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("Failed to fetch categories: %w", err)
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("Unexpected status code: %d", response.StatusCode)
		}

		var result CategoryReponse
		if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
			return nil, fmt.Errorf("Failed to decode response: %w", err)
		}

		/* b, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(b)) */

		return result.Data, nil
	}
	return nil, nil
}

func (a *API) GetItems(categoryID int) ([]Item, error) {
	if OA != "" {
		url := fmt.Sprintf("%s/%s/categories/%d?includePackages=1", endpoints.Store, OA, categoryID)
		response, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("Failed to fetch items: %w", err)
		}
		defer response.Body.Close()
		if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("Unexpected status code: %d", response.StatusCode)
		}
		var result ItemDetail
		if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
			return nil, fmt.Errorf("Failed to decode response: %w", err)
		}

		/* b, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(b)) */

		return result.Data.Packages, nil
	}
	return nil, nil
}

func (a *API) GetServerPlayers() (ServerDetail, error) {
	url := fmt.Sprintf(endpoints.Players)
	response, err := http.Get(url)
	if err != nil {
		return ServerDetail{}, fmt.Errorf("Failed to fetch players: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return ServerDetail{}, fmt.Errorf("Unexpected status code: %d", response.StatusCode)
	}
	var result ServerDetail
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return ServerDetail{}, fmt.Errorf("Failed to decode response: %w", err)
	}

	return result, nil
}
