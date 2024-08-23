package routes

import (
	"fmt"
	"resedist/controllers/post"
	"resedist/services"

	"github.com/gin-gonic/gin"
)

func PostsRoutes(r *gin.Engine) {

	posts := r.Group("/posts")

	posts.Use(func(ctx *gin.Context) {
		fmt.Println("posts middleware")
		ctx.Next()
	})

	postService := services.NewPostService()
	postController := post.NewPostController(postService)

	posts.POST("/", postController.GetPost)
	// posts.GET("/", PostController.PostIndex)
	// posts.GET("/:id", PostController.PostShow)
	// posts.PUT("/:id", PostController.PostUpdate)
	// posts.DELETE("/:id", PostController.PostDelete)
}
