package entity

import "github.com/jinzhu/gorm"

type Session struct {
	gorm.Model

	Uuid   string `json:"uuid" gorm:"unique;not null"`
	UserID uint   `json:"user_id" gorm:"not null"`
	Email  string `json:"email" gorm:"not null"`
}
