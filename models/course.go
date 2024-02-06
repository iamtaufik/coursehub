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
	Title          string		  	`json:"title"`
	Description    string	  	  	`json:"description"`
	Image          *string	  	  	`gorm:"default: NULL" json:"image"`
	TelegramGroup  *string		  	`gorm:"default: NULL" json:"telegram_group"`
	Requirements   string		  	`json:"requirements"`
	Level          Levels         	`gorm:"default:beginner" json:"level"`
	Price          int			  	`json:"price"`
	Author         string		  	`json:"author"`
	CategoryID     uint			  	`json:"category_id"`
	Chapters 	   []Chapter      	`gorm:"constraint:OnDelete:CASCADE" json:"chapters"`
	Users		   []*User         	`gorm:"many2many:user_courses;" json:"-"`
}