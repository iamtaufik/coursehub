package models

import (
	"gorm.io/gorm"
)

type Levels string

const (
	Beginner     Levels = "beginner"
	Intermediate Levels = "intermediate"
	Advanced     Levels = "advanced"
)

type Course struct {
	gorm.Model
	ID             uint           	`gorm:"primaryKey" json:"id"`
	CourseID       string         	`gorm:"default:cuid()" json:"course_id"`
	Title          string		  	`json:"title"`
	Description    string	  	  	`json:"description"`
	Image          string	  	  	`gorm:"default: NULL" json:"image"`
	TelegramGroup  string		  	`json:"telegram_group"`
	Requirements   []string		  	`json:"requirements"`
	Level          Levels         	`gorm:"default:beginner" json:"level"`
	Price          int			  	`json:"price"`
	Author         string		  	`json:"author"`
	Chapters       []Chapter      	`gorm:"foreignKey:CourseID" json:"chapters"`
	CategoryID     uint			  	`json:"category_id"`
	Category       Category     	`gorm:"foreignKey:CategoryID"`
	Users          []User        	`gorm:"many2many:user_courses"`
	IsDeleted      bool           	`gorm:"default:false"`
}