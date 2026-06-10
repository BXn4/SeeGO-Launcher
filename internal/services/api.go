package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type API struct{}

var OA string // setting it during the build

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

type CategoryReponse struct {
	Data []Category `json:"data"`
}

func (a *API) GetCategories() ([]Category, error) {
	// {"data":[{"id":3241741,
	if OA != "" {
		url := fmt.Sprintf("https://headless.tebex.io/api/accounts/%s/categories/", OA)
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
		url := fmt.Sprintf("https://headless.tebex.io/api/accounts/%s/categories/%d?includePackages=1", OA, categoryID)
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
