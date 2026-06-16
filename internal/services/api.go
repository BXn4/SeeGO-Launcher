package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"seegolauncher/internal/endpoints"
)

type API struct{}

var OA string

type ServerDetail struct {
	Players int `json:"players"`
	Admins  int `json:"admins"`
	Queue   int `json:"queue"`
	Slots   int `json:"slots"`
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
