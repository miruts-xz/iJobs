package entity

import "github.com/jinzhu/gorm"

// represents status of job-application
type Status string

const (
	ACCEPTED Status = "accepted"
	REJECTED Status = "rejected"
	FURTHER  Status = "further"
)

// represents a job-application response
type Response string

const (
	REVIEWED   Status = "reviewed"
	UNREVIEWED string = "unreviewed"
)

// Application represents a job-application
type Application struct {
	gorm.Model

	JobID       uint     `gorm:"not null"`
	JobseekerID uint     `gorm:"not null"`
	Response    Status   `json:"response" gorm:"varchar(255)"`
	Status      Response `json:"status" gorm:"varchar(255)"`
}
