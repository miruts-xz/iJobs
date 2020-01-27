package entity

import "github.com/jinzhu/gorm"

// Category represents job categories
type Category struct {
	gorm.Model

	Jobs  []Job  `json:"jobs" gorm:"many2many:job_categories"`
	Name  string `json:"name" gorm:"type:varchar(255);not null"`
	Image string `json:"image" gorm:"type:varchar(255)"`
	Descr string `json:"descr" gorm:"type:text"`
}
