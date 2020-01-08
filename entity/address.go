package entity

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
	AddisKetema        = "Addis Ketema"
)

type Address struct {
	Add_ID    int    `json:"add_id" gorm:"primary_key"`
	Region    string `json:"region"`
	City      string `json:"city"`
	SubCity   string `json:"sub_city"`
	LocalName string `json:"local_name"`
}
