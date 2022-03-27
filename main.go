package main

import (
	"Go_mongo/middleware"
	"Go_mongo/routes"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.CORSMiddleware())

	// routes.AuthRoutes(router)
	// routes.UserRoutes(router)

	v1 := router.Group("/v1")
	{
		auth := v1.Group("/auth")
		routes.AuthRoutes(auth)

		user := v1.Group("/user")
		routes.UserRoutes(user)
	}

	// API-2
	router.GET("/api-1", func(c *gin.Context) {

		c.JSON(200, gin.H{"success": "Access granted for api-1"})

	})

	// API-1
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}
