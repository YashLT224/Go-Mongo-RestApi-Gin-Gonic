package routes

import (
	controller "Go_mongo/controllers"
	"Go_mongo/middleware"

	"github.com/gin-gonic/gin"
)

//way1
//UserRoutes function
// func UserRoutes(incomingRoutes *gin.Engine) {
// 	// incomingRoutes.Use(middleware.Authentication())
// 	incomingRoutes.GET("/users", middleware.Authentication(), controller.GetUsers())
// 	incomingRoutes.GET("/users/:user_id", middleware.Authentication(), controller.GetUser())
// }

func UserRoutes(incomingRoutes *gin.RouterGroup) {

	incomingRoutes.GET("/", middleware.Authentication(), controller.GetUsers())
	incomingRoutes.GET("/:user_id", middleware.Authentication(), controller.GetUser())
}
