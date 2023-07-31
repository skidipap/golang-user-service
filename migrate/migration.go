package main

import (
	"example/users-service/initializers"
	"example/users-service/models"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
