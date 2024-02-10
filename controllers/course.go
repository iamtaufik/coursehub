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
	if c.Query("category") != "" {
		GetCourseByCategory(c)
		return
	}
	var courses []models.Course
	config.DB.Find(&courses)

	c.JSON(http.StatusOK, gin.H{"data": courses})
}

func GetCourse(c *gin.Context){
	var course models.Course
	config.DB.Preload("Chapters").Preload("Chapters.Modules").First(&course, "id = ?", c.Param("id"))

	c.JSON(http.StatusOK, gin.H{"data": course})
}

func GetCourseByCategory(c *gin.Context){
	var courses []models.Course
	var category models.Category
		//  search category ny name use LIKE
	config.DB.Where("name LIKE ?", "%" + c.Query("category") + "%").First(&category)
	
	if category.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Category not found"})
		return
	}

	config.DB.Where("category_id = ?",  category.ID).Find(&courses)
	
	if len(courses) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No courses found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": courses})
}

func UpdateCourse(c *gin.Context){
	var body struct {
		Title 		string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
		Image 		*string `json:"image"`
		TelegramGroup *string `json:"telegram_group"`
		Requirements string `json:"requirements" binding:"required"`
		Level 		string `json:"level" binding:"required"`
		Price 		int `json:"price" binding:"required"`
		Author 		string `json:"author" binding:"required"`
		CategoryID 	uint `json:"category_id" binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	var course models.Course
	config.DB.First(&course, "id = ?", c.Param("id"))

	if course.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Course not found"})
		return
	}

	course.Title = body.Title
	course.Description = body.Description
	course.Image = body.Image
	course.TelegramGroup = body.TelegramGroup
	course.Requirements = body.Requirements
	course.Level = models.Levels(body.Level)
	course.Price = body.Price
	course.Author = body.Author
	course.CategoryID = body.CategoryID

	config.DB.Save(&course)

	c.JSON(http.StatusOK, gin.H{"message": "Course updated successfully"})
}

func DeleteCourse(c *gin.Context){
	var course models.Course
	config.DB.First(&course, "id = ?", c.Param("id"))

	if course.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Course not found"})
		return
	}

	config.DB.Delete(&course)

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}