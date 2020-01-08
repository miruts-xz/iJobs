package entity

import "time"

type Region string
type City string
type SubCity string

const (
	Tigray     Region = "Tigray"
	Amhara     Region = "Amhara"
	Oromia     Region = "Oromia"
	Afar       Region = "Afar"
	Somalia    Region = "Somalia"
	Benshangul Region = "Benshangul Gumz"
	Harare     Region = "Harare"
	Sidama     Region = "Sidama"
	Gambella   Region = "Gambella"
	Snnpr      Region = "SNNPR"
)
const (
	Addis   City = "Addis Ababa"
	Adamma  City = "Adamma"
	Hawassa City = "Hawassa"
	Mekele  City = "Mekele"
	Gonder  City = "Gonder"
)
const (
	Gulele      SubCity = "Gullele"
	Arada       SubCity = "Arada"
	Akaki       SubCity = "Akaki Kality"
	Bole        SubCity = "Bole"
	Cherkos     SubCity = "Cherkos"
	Yeka        SubCity = "Yeka"
	AddisKetema SubCity = "Addis Ketema"
)

type Address struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`

	Region    Region  `json:"region"`
	City      City    `json:"city"`
	SubCity   SubCity `json:"sub_city"`
	LocalName string  `json:"local_name"`
}
