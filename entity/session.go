package entity

// Session represents active user
type Session struct {
	ID     uint
	Uuid   string `json:"uuid" gorm:"unique;not null"`
	UserID uint   `json:"user_id" gorm:"not null"`
	Email  string `json:"email" gorm:"not null"`

	Expires    int64  `gorm:"type:varchar(255);not null"`
	SigningKey []byte `gorm:"type:varchar(255);not null"`
}
