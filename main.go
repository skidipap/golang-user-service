package main

import (
	"example/users-service/controllers"
	"example/users-service/initializers"
	"example/users-service/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"data": "ok"})
	c.Abort()
	return
}

func main() {
	initializers.ConnectDB()
	router := gin.Default()
	router.Use(middlewares.JWTAuth)
	router.GET("/", Ping)
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.GenerateToken)

	router.Run("localhost:8000")
}
