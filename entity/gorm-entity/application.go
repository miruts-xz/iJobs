package entity

import "github.com/jinzhu/gorm"

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
	gorm.Model

	JobID       uint     `gorm:"not null"`
	JobseekerID uint     `gorm:"not null"`
	Response    Status   `json:"response" gorm:"varchar(255)"`
	Status      Response `json:"status" gorm:"varchar(255)"`
}
