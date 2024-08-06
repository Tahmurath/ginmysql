package main

import (
	"vida/controllers"
	"vida/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDb()
}
func main() {

	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostIndex)

	r.Run() // listen and serve on 0.0.0.0:8080
}
