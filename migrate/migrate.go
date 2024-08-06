package main

import (
	"vida/initializers"
	"vida/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDb()
}

func main() {
	initializers.DB.AutoMigrate(models.Post{})
}
