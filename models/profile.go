package models

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	ProfileID      string     `gorm:"default:cuid()" json:"profile_id"`
	UserID        uint       `gorm:"unique"`
	PhoneNumber    string    `gorm:"default: NULL" json:"phone_number"`
	FullName       string    `gorm:"default: NULL" json:"full_name"`
	ProfilePicture string    `gorm:"default: NULL" json:"profile_picture"`
	City           string    `gorm:"default: NULL" json:"city"`
	Country        string    `gorm:"default: NULL" json:"country"`
}
