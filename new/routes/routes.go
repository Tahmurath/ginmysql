package routes

import (
	"resedist/controllers/user"
	"resedist/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	userService := services.NewUserService()
	userController := user.NewUserController(userService)

	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/:id", userController.GetUser)
		// مسیرهای دیگر کاربران
	}

	// می‌توانید مسیرهای دیگر مانند محصول، احراز هویت و غیره را به همین شکل اضافه کنید
}
