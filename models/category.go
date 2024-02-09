package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        		uint 		`gorm:"primarykey" json:"id"`
	CreatedAt 		time.Time 	`json:"created_at"`
	UpdatedAt 		time.Time 	`json:"-"`
	DeletedAt 		gorm.DeletedAt `gorm:"index" json:"-"`
	Name       string `json:"name"`
	Courses    []Course `gorm:"constraint:OnDelete:CASCADE"`
}