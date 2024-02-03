package models

import "gorm.io/gorm"

type Chapter struct {
	gorm.Model
	ChapterID string `gorm:"default:cuid()" json:"chapter_id"`
	Name      string `json:"name"`
	CourseID  uint   `json:"course_id"`
	Courses   Course `gorm:"foreignKey:CourseID"`
	Modules   []Module `gorm:"foreignKey:ChapterID" json:"modules"`
}