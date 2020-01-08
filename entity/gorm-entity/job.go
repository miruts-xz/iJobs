package entity

import (
	"time"
)

type Job struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`

	Company  Company
	Category []Category `gorm:"many2many:job_categories"`

	RequiredNum int       `json:"required_num"`
	Salary      float64   `json:"salary"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	JobTime     string    `json:"job_time"`
	Deadline    time.Time `json:"deadline"`
}
