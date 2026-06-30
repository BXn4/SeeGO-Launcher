package net

import (
	"io"
	"net/http"
	"seegolauncher/internal/endpoints"
)

func Request(url string) (string, error) {
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

func RequestTermsVersion() (string, error) {
	response, err := Request(endpoints.TermsDate)
	if err != nil {
		return "Failed to request terms date", err
	}
	return response, err
}
