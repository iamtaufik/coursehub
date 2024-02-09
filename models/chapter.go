package models

import (
	"time"

	"gorm.io/gorm"
)

type Chapter struct {
	ID        		uint 		`gorm:"primarykey" json:"id"`
	CreatedAt 		time.Time 	`json:"created_at"`
	UpdatedAt 		time.Time 	`json:"-"`
	DeletedAt 		gorm.DeletedAt `gorm:"index" json:"-"`
	CourseID  uint     `json:"-"` 
	Name      string   `json:"name"`
	Modules   []Module `gorm:"constraint:OnDelete:CASCADE" json:"modules"`
}