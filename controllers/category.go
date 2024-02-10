package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamtaufik/coursehub/config"
	"github.com/iamtaufik/coursehub/models"
)

type APICategory struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func GetCategories(c *gin.Context){

	var categories []models.Category

	result := config.DB.Omit("Courses").Find(&categories)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": categories})
}