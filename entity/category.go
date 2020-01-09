package entity

type Category struct {
	ID    int64  `json:"id" gorm:"primary_key;"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Desc  string `json:"desc"`
}
