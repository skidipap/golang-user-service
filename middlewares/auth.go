package middlewares

import (
	"example/users-service/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	if err := helpers.ValidateToken(tokenString); err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.Next()
}
