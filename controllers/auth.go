package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"github.com/iamtaufik/coursehub/config"
	"github.com/iamtaufik/coursehub/dto"
	"github.com/iamtaufik/coursehub/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var body dto.RegisterDto

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)


	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	user := models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: string(hash),	
	}

	result := config.DB.Create(&user)


	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
	var body dto.LoginDto

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	var user models.User
	config.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0{
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Invalid email or password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 3).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Invalid create token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenStr,
	})
}

func RegisterAdmin(c *gin.Context){
	var body dto.RegisterDto

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)


	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	user := models.User{
		Username: body.Username,
		Email:    body.Email,
		Role: models.Admin,
		Password: string(hash),	
	}

	result := config.DB.Create(&user)


	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func LoginAdmin(c *gin.Context) {
	var body dto.LoginDto

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	var user models.User
	config.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Invalid email or password",
		})
		return
	}

	if user.Role != models.Admin {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Invalid email or password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 3).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Invalid create token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenStr,
	})
}

func Whoami(c *gin.Context){
	user := c.MustGet("user")

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}