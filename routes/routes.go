package routes

import (
	"vida/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// routeName := os.Getenv("HOME_ROUTE")

	// routes := map[string]gin.HandlerFunc{
	//     "HomeHandler":  Index,
	// }
	// handler, exists := routes[routeName]
	// if !exists {
	// 	fmt.Println("Route handler not found!")
	// 	return
	// }

	r.GET("/", controllers.Index)
	//fmt.Println(os.Getenv("HOME_ROUTE"))

	PostsRoutes(r)

}
