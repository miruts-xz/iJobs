package entity

import "github.com/jinzhu/gorm"

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
	gorm.Model

	Address      []Address     `json:"address" gorm:"many2many:jobseeker_addresses"`
	Applications []Application `json:"applications" gorm:"foreignkey:JobseekerID"`
	Categories   []Category    `json:"categories" gorm:"many2many:jobseeker_categories"`

	Age            uint      `json:"age"`
	Phone          string    `json:"phone" gorm:"unique;"`
	WorkExperience int       `json:"work_experience"`
	Username       string    `json:"username" gorm:"unique;not null"`
	Fullname       string    `json:"fullname"`
	Password       string    `json:"password"`
	Email          string    `json:"email" gorm:"unique"`
	Profile        string    `json:"profile"`
	Portfolio      string    `json:"portfolio"`
	CV             string    `json:"cv" gorm:"not null;unique"`
	Gender         Gender    `json:"gender"`
	EmpStatus      EmpStatus `json:"emp_status" gorm:"not null"`
}
