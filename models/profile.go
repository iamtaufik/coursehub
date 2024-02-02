package models

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	UserID      uint   `json:"user_id"`
}