package entity

import "github.com/jinzhu/gorm"

const (
	Tigray     string = "Tigray"
	Amhara     string = "Amhara"
	Oromia     string = "Oromia"
	Afar       string = "Afar"
	Somalia    string = "Somalia"
	Benshangul string = "Benshangul Gumz"
	Harare     string = "Harare"
	Sidama     string = "Sidama"
	Gambella   string = "Gambella"
	Snnpr      string = "SNNPR"
)
const (
	Addis   string = "Addis Ababa"
	Adamma  string = "Adamma"
	Hawassa string = "Hawassa"
	Mekele  string = "Mekele"
	Gonder  string = "Gonder"
)
const (
	Gulele      string = "Gullele"
	Arada       string = "Arada"
	Akaki       string = "Akaki Kality"
	Bole        string = "Bole"
	Cherkos     string = "Cherkos"
	Yeka        string = "Yeka"
	AddisKetema string = "Addis Ketema"
)

type Address struct {
	gorm.Model

	Region    string `json:"region" gorm:"type:varchar(255)"`
	City      string `json:"city" gorm:"type:varchar(255)"`
	SubCity   string `json:"sub_city" gorm:"type:varchar(255)"`
	LocalName string `json:"local_name" gorm:"type:varchar(255);not null"`
}
