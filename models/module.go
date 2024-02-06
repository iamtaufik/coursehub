package models

import (
	"gorm.io/gorm"
)

type Module struct {
	gorm.Model
	ChapterID 	uint	`json:"chapter_id"`
	Title     	string	`json:"title"`
	Duration  	int		`json:"duration"`	
	URL       	string	`json:"url"`
	IsTrailer 	bool     `gorm:"default:false"`
}