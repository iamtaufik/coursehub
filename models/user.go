package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uint     `gorm:"primaryKey"`
	UserID     string   `gorm:"default:cuid()" json:"user_id"`
	Username   string   `json:"username"`
	Email      string   `gorm:"unique" json:"email"`
	Password   string   `json:"password"`
	Courses    []Course `gorm:"many2many:user_courses"`
	Profile    Profile 	`gorm:"foreignKey:UserID"`
	IsVerified bool      `gorm:"default:false"`
}
