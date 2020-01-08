package entity

import (
	"time"
)

type Gender string
type EmpStatus string

const (
	EMPLD   EmpStatus = "employed"
	UNEMPLD EmpStatus = "unemployed"
)
const (
	MALE   Gender = "male"
	FEMALE Gender = "female"
	OTHER  Gender = "other"
)

type Jobseeker struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`

	Applications []Application
	Categories   []Category `json:"categories" gorm:"many2many:jobseeker_categories;"`
	Address      Address

	Age            uint      `json:"age"`
	Phone          int       `json:"phone"`
	WorkExperience int       `json:"work_experience"`
	Username       string    `json:"username"`
	Fullname       string    `json:"fullname"`
	Password       string    `json:"password"`
	Email          string    `json:"email" gorm:"unique"`
	Profile        string    `json:"profile"`
	Portfolio      string    `json:"portfolio"`
	CV             string    `json:"cv"`
	Gender         Gender    `json:"gender"`
	EmpStatus      EmpStatus `json:"emp_status"`
}
