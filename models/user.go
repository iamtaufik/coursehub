package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	Admin    Role = "admin"
	Student  Role = "student"
)

type User struct {
	ID        		uint 		`gorm:"primarykey" json:"id"`
	CreatedAt 		time.Time 	`json:"created_at"`
	UpdatedAt 		time.Time 	`json:"-"`
	DeletedAt 		gorm.DeletedAt `gorm:"index" json:"-"`
	Username   		string   	`json:"username"`
	Email      		string   	`gorm:"unique" json:"email"`
	Password   		string   	`json:"password"`
	Role 	 		Role     	`gorm:"default:student" json:"role"`
	PhoneNumber    	*string    	`gorm:"default: NULL" json:"phone_number"`
	FullName       	*string    	`gorm:"default: NULL" json:"full_name"`
	ProfilePicture 	*string    	`gorm:"default: NULL" json:"profile_picture"`
	Address        	*string    	`gorm:"default: NULL" json:"address"`
	Courses 		[]Course   	`gorm:"many2many:user_courses;" json:"courses"`
	IsVerified 		bool     	`gorm:"default:false" json:"is_verified"`
}
