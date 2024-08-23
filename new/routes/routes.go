package routes

import (
	"resedist/controllers/post"
	"resedist/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	postService := services.NewPostService()
	postController := post.NewPostController(postService)

	postRoutes := router.Group("/posts")
	{
		postRoutes.GET("/:id", postController.GetPost)
		// مسیرهای دیگر کاربران
	}

	// می‌توانید مسیرهای دیگر مانند محصول، احراز هویت و غیره را به همین شکل اضافه کنید
}
