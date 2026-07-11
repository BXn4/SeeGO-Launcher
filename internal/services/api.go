package services

import (
	"encoding/json"
	"fmt"
	"seegolauncher/internal/endpoints"
	"seegolauncher/internal/net"
)

type API struct{}

var OA string = ""

type ServerDetail struct {
	Players int `json:"players"`
	Admins  int `json:"admins"`
	Queue   int `json:"queue"`
	Slots   int `json:"slots"`
}

func (a *API) GetServerPlayers() (ServerDetail, error) {
	response, err := net.Request(endpoints.Players)
	if err != nil {
		return ServerDetail{}, fmt.Errorf("Failed to fetch players: %w", err)
	}

	var result ServerDetail
	if err := json.Unmarshal([]byte(response), &result); err != nil {
		return ServerDetail{}, fmt.Errorf("Failed to decode response: %w", err)
	}

	return result, nil
}
