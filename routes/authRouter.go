package routes

import (
	controller "Go_mongo/controllers"

	"github.com/gin-gonic/gin"
)

//UserRoutes function
// func AuthRoutes(incomingRoutes *gin.Engine) {
// 	incomingRoutes.POST("/users/signup", controller.SignUp())
// 	incomingRoutes.POST("/users/login", controller.Login())
// }

func AuthRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.POST("/signup", controller.SignUp())
	incomingRoutes.POST("/login", controller.Login())
}
