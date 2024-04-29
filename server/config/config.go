package config

import (
	"os"
)

var (
	Port        string
	DatabaseURL string
)

func LoadConfig() {
	Port = os.Getenv("PORT")
	DatabaseURL = os.Getenv("DATABASE_URL")
}
