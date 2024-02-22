package controllers

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/iamtaufik/coursehub/config"
	"github.com/iamtaufik/coursehub/dto"
	"github.com/iamtaufik/coursehub/models"
)

func CreateCourse(c *gin.Context){
	var body dto.CreateCourseDto

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
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		result := config.DB.Create(&course)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
			return
		}

		for _, chapter := range body.Chapters {

			var newChapter models.Chapter
			newChapter.Name = chapter.Name
			newChapter.CourseID = course.ID
			if err := config.DB.Create(&newChapter).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}

			for _, module := range chapter.Modules {
				var newModule models.Module
				newModule.Title = module.Title
				newModule.Duration = module.Duration
				newModule.URL = module.URL
				newModule.IsTrailer = module.IsTrailer
				newModule.ChapterID = newChapter.ID
				if err := config.DB.Create(&newModule).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err})
					return
				}
			}
		}


		c.JSON(http.StatusCreated, gin.H{"message": "Course created successfully"})
	}()

	wg.Wait()
	c.Status(http.StatusCreated)
}

func GetCourses(c *gin.Context){
	if c.Query("category") != "" {
		GetCourseByCategory(c)
		return
	}
	var courses []models.Course
	
	var wg sync.WaitGroup
    wg.Add(1)
	go func() {
		defer wg.Done()

		if err := config.DB.Find(&courses).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": courses})
		}()
	wg.Wait()	
	c.Status(http.StatusOK)
}

func GetCourse(c *gin.Context){
	var course models.Course
	
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := config.DB.Preload("Chapters").Preload("Chapters.Modules").First(&course, "id = ?", c.Param("id")).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data kursus"})
            return
		}

		c.JSON(http.StatusOK, gin.H{"data": course})
	}()

	wg.Wait()
	c.Status(http.StatusOK)
}

func GetCourseByCategory(c *gin.Context){
	var courses []models.Course
	var category models.Category

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
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
	}()

	wg.Wait()
	c.Status(http.StatusOK)
}

func JoinCourse(c *gin.Context){
	var user models.User
	config.DB.First(&user, "id = ?", c.MustGet("user").(models.User).ID)

	
	var course models.Course
	config.DB.First(&course, "id = ?", c.Param("id"))
	
	// check if user is already enrolled in the course
	for _, userCourse := range user.Courses   {
		if userCourse.ID == course.ID {
			c.JSON(http.StatusBadRequest, gin.H{"message": "You are already enrolled in this course"})
			return
		}
	}
	if course.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Course not found"})
		return
	}

	config.DB.Model(&user).Association("Courses").Append(&course)

	c.JSON(http.StatusOK, gin.H{"message": "Course joined successfully"})
}

func MyCourses(c *gin.Context){
	var user models.User
	config.DB.Preload("Courses").First(&user, "id = ?", c.MustGet("user").(models.User).ID)

	c.JSON(http.StatusOK, gin.H{"data": user.Courses})
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