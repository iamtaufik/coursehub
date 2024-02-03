package models

import (
	"gorm.io/gorm"
)

type Module struct {
	gorm.Model
	ID        	uint 	`gorm:"primaryKey"`
	ModuleID  	string 	`gorm:"default:cuid()" json:"module_id"`
	Title     	string	`json:"title"`
	Duration  	int		`json:"duration"`	
	URL       	string	`json:"url"`
	ChapterID 	uint	`json:"chapter_id"`
	IsTrailer 	bool     `gorm:"default:false"`
	Chapters  	Chapter  `gorm:"foreignKey:ChapterID"`
}