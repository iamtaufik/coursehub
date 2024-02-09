package models

import (
	"time"

	"gorm.io/gorm"
)

type Module struct {
	ID        		uint 		`gorm:"primarykey" json:"id"`
	CreatedAt 		time.Time 	`json:"created_at"`
	UpdatedAt 		time.Time 	`json:"-"`
	DeletedAt 		gorm.DeletedAt `gorm:"index" json:"-"`
	ChapterID 		uint		`json:"-"`
	Title     		string		`json:"title"`
	Duration  		int			`json:"duration"`	
	URL       		string		`json:"url"`
	IsTrailer 		bool     	`gorm:"default:false"`
}