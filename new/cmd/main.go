package main

import (
	"log"
	"resedist/config"
	"resedist/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// بارگذاری فایل .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// بارگذاری تنظیمات
	config.LoadConfig()

	// ایجاد یک نمونه جدید از Gin
	router := gin.Default()

	// تنظیم مسیرها
	routes.RegisterRoutes(router)

	// اجرای سرور
	err = router.Run()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
