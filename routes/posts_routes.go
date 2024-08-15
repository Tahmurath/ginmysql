package routes

import (
	"vida/controllers"

	"github.com/gin-gonic/gin"
)

func PostsRoutes(r *gin.Engine) {

	posts := r.Group("/posts")

	posts.POST("/", controllers.PostsCreate)
	posts.GET("/", controllers.PostIndex)
	posts.GET("/:id", controllers.PostShow)
	posts.PUT("/:id", controllers.PostUpdate)
	posts.DELETE("/:id", controllers.PostDelete)
}
