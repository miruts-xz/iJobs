package entity

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model

	Name  string `json:"name" gorm:"type:varchar(255);not null"`
	Image string `json:"image" gorm:"type:varchar(255)"`
	Desc  string `json:"desc" gorm:"type:text"`
}
