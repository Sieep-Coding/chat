package config

import (
	"fmt"
	"os"
)

var (
	Port        string
	DatabaseURL string
)

func LoadConfig() error {
	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8080" // Default port if not set
	}

	DatabaseURL = os.Getenv("DATABASE_URL")
	if DatabaseURL == "" {
		return fmt.Errorf("DATABASE_URL environment variable not set")
	}

	return nil
}
