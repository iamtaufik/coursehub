package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamtaufik/coursehub/config"
	"github.com/iamtaufik/coursehub/models"
)

func CreateCourse(c *gin.Context){
	var body struct {
		Title 		string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
		Image 		*string `json:"image"`
		TelegramGroup *string `json:"telegram_group"`
		Requirements string `json:"requirements" binding:"required"`
		Level 		string `json:"level" binding:"required"`
		Price 		int `json:"price" binding:"required"`
		Author 		string `json:"author" binding:"required"`
		Chapters    []struct {
			Name string `json:"name" binding:"required"`
			Modules []struct {
				Title string `json:"title" binding:"required"`
				Duration int `json:"duration" binding:"required"`
				URL string `json:"url" binding:"required"`
				IsTrailer bool `json:"is_trailer" binding:"required"`
			} `json:"modules" binding:"required"`
		} `json:"chapters" binding:"required"`
		CategoryID 	uint `json:"category_id" binding:"required"`

	}

	if c.Bind(&body) != nil {
//  show details of the error
		c.JSON(http.StatusBadRequest, gin.H{"error": c.Errors})
		return
	}

	course := models.Course{
		Title: body.Title,
		Description: body.Description,
		TelegramGroup: body.TelegramGroup,
		Requirements: body.Requirements,
		Level: models.Levels(body.Level),
		Price: body.Price,
		CategoryID: body.CategoryID,
	}
	
	result := config.DB.Create(&course)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	for _, chapter := range body.Chapters {

		var newChapter models.Chapter
		newChapter.Name = chapter.Name
		newChapter.CourseID = course.ID
		config.DB.Create(&newChapter)

		for _, module := range chapter.Modules {
			var newModule models.Module
			newModule.Title = module.Title
			newModule.Duration = module.Duration
			newModule.URL = module.URL
			newModule.IsTrailer = module.IsTrailer
			newModule.ChapterID = newChapter.ID
			config.DB.Create(&newModule)
		}
	}


	c.JSON(http.StatusCreated, gin.H{"message": "Course created successfully"})
}

func GetCourses(c *gin.Context){
	var courses []models.Course
	config.DB.Preload("Chapters").Preload("Chapters.Modules").Find(&courses)

	c.JSON(http.StatusOK, gin.H{"data": courses})
}

func GetCourse(c *gin.Context){
	var course models.Course
	config.DB.Preload("Chapters").Preload("Chapters.Modules").First(&course, "id = ?", c.Param("id"))

	c.JSON(http.StatusOK, gin.H{"data": course})
}