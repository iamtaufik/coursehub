package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey"`
	CategoryID string `gorm:"default:cuid()" json:"category_id"`
	Name       string `json:"name"`
	Courses    []Course
}