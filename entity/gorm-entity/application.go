package entity

import (
	"time"
)

type Status string

const (
	ACCEPTED Status = "accepted"
	REJECTED Status = "rejected"
	FURTHER  Status = "further"
)

type Response string

const (
	REVIEWED   Status = "reviewed"
	UNREVIEWED string = "unreviewed"
)

type Application struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`

	Job       Job
	JobSeeker Jobseeker

	Response Status   `json:"response"`
	Status   Response `json:"status"`
}
