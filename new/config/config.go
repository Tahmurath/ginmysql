package config

import (
	"os"
)

var (
	ServerAddress string
	DatabaseURL   string
)

func LoadConfig() {
	ServerAddress = os.Getenv("SERVER_ADDRESS")
	DatabaseURL = os.Getenv("DATABASE_URL")

	if ServerAddress == "" {
		ServerAddress = ":8080" // آدرس پیش‌فرض سرور
	}
}
