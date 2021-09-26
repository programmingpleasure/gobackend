package main

import (
	"errors"
	"os"
)

const (
	apiKeyEnv = "API_KEY"
)

type config struct {
	APIKey string
}

func loadConfig() (config, error) {
	apiKey := os.Getenv(apiKeyEnv)
	if apiKey == "" {
		return config{}, errors.New("empty API_KEY environment variable")
	}

	return config{
		APIKey: os.Getenv(apiKeyEnv),
	}, nil
}
