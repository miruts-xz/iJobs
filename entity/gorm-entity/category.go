package entity

import "time"

type Category struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`

	Name  string `json:"name"`
	Image string `json:"image"`
	Desc  string `json:"desc"`
}
