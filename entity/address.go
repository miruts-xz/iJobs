package entity

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
	AddisKetema         = "Addis Ketema"
)

type Address struct {
	Add_ID    int     `json:"add_id" gorm:"primary_key"`
	Region    Region  `json:"region"`
	City      City    `json:"city"`
	SubCity   SubCity `json:"sub_city"`
	LocalName string  `json:"local_name"`
}
