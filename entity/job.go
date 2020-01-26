package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Job represents a job posted by company
type Job struct {
	gorm.Model

	CompanyID    uint          `json:"company_id" gorm:"not null"`
	Categories   []Category    `json:"categories" gorm:"many2many:job_categories"`
	Applications []Application `json:"applications" gorm:"foreign_key:JobID"`

	RequiredNum uint      `json:"required_num" gorm:"not null"`
	Salary      float64   `json:"salary"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description" gorm:"type:text"`
	JobTime     string    `json:"job_time" gorm:"type:varchar(255);not null"`
	Deadline    time.Time `json:"deadline" gorm:"not null"`
}
