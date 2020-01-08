package entity

import (
	"time"
)

type Company struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`

	Jobs    []Job   `json:"jobs"`
	Address Address `json:"address"`

	CompanyName string `json:"company_name"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Logo        string `json:"logo"`
	ShortDesc   string `json:"short_desc"`
	DetailInfo  string `json:"detail_info"`
}
