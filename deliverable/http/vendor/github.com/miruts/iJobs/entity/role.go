package entity

import "github.com/jinzhu/gorm"

type Role struct {
	ID    uint
	Name  string `gorm:"type:varchar(255)"`
	Users []User
}
type User struct {
	gorm.Model
	RoleID uint
}
