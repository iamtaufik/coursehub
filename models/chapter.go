package models

import "gorm.io/gorm"

type Chapter struct {
	gorm.Model
	CourseID  uint     `json:"course_id"` 
	Name      string   `json:"name"`
	Modules   []Module `gorm:"constraint:OnDelete:CASCADE"`
}