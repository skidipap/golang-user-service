package controllers

import (
	"example/users-service/helpers"
	"example/users-service/initializers"
	"example/users-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		c.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		c.Abort()
		return
	}
	if err := initializers.DB.Create(&user).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		c.Abort()
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"data": map[string]any{"created_at": user.CreatedAt, "name": user.Name, "email": user.Email}})
}

func GenerateToken(c *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		c.Abort()
		return
	}

	if err := initializers.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if err := user.CheckPassword(request.Password); err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	tokenString, err := helpers.GenerateJWT(user.Email, user.Name)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}
