package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iamtaufik/coursehub/config"
	"github.com/iamtaufik/coursehub/models"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
    var body struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Username: body.Username, Email: body.Email, Password: body.Password}
	result := config.DB.Create(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(200, result)	
}