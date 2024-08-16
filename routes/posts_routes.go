package routes

import (
	"fmt"
	"vida/controllers"

	"github.com/gin-gonic/gin"
)

func PostsRoutes(r *gin.Engine) {

	posts := r.Group("/posts")

	posts.Use(func(ctx *gin.Context) {
		fmt.Println("posts middleware")
		ctx.Next()
	})

	posts.POST("/", controllers.PostsCreate)
	posts.GET("/", controllers.PostIndex)
	posts.GET("/:id", controllers.PostShow)
	posts.PUT("/:id", controllers.PostUpdate)
	posts.DELETE("/:id", controllers.PostDelete)
}
