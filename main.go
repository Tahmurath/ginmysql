package main

import (
	"fmt"
	"vida/initializers"
	"vida/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	//gin.SetMode(gin.ReleaseMode)
	initializers.LoadEnvVars()
	initializers.ConnectDb()
}
func main() {

	r := gin.Default()
	r.Use(func(ctx *gin.Context) {
		fmt.Println("Global middleware")
		ctx.Next()
	})

	routes.SetupRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
