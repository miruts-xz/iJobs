package entity

import "github.com/jinzhu/gorm"

type Company struct {
	gorm.Model

	Jobs    []Job     `json:"jobs" gorm:"foreignkey:CompanyID"`
	Address []Address `json:"address" gorm:"many2many:company_addresses"`

	CompanyName string `json:"company_name" gorm:"varchar(255);not null;"`
	Password    string `json:"password" gorm:"type:varchar(255);not null"`
	Email       string `json:"email" gorm:"type:varchar(255);not null;unique;"`
	Phone       string `json:"phone" gorm:"type:varchar(255);unique;"`
	Logo        string `json:"logo" gorm:"type:varchar(255)"`
	ShortDesc   string `json:"short_desc" gorm:"type:varchar(255)"`
	DetailInfo  string `json:"detail_info" gorm:"type:text"`
}
