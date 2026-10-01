//go:build decoy

package config

import (
	"errors"
	"os"
	"strings"
)

type Config struct {
	ChatID string
}

func Load() (*Config, error) {
	chatID := strings.TrimSpace(os.Getenv("XTR_CHAT_ID"))
	if chatID == "" {
		return nil, errors.New("XTR_CHAT_ID is required")
	}
	return &Config{ChatID: chatID}, nil
}
