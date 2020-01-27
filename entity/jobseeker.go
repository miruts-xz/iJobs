package entity

type Gender string
type EmpStatus string

const (
	EMPLD   EmpStatus = "employed"
	UNEMPLD EmpStatus = "unemployed"
)
const (
	MALE   Gender = "male"
	FEMALE Gender = "female"
	OTHER  Gender = "other"
)

// Jobseeker represents users seeking job
type Jobseeker struct {
	User
	Address      []Address     `json:"address" gorm:"many2many:jobseeker_addresses"`
	Applications []Application `json:"applications" gorm:"foreignkey:JobseekerID"`
	Categories   []Category    `json:"categories" gorm:"many2many:jobseeker_categories"`

	Age            uint   `json:"age"`
	Phone          string `json:"phone"`
	WorkExperience int    `json:"work_experience"`
	Username       string `json:"username" gorm:"unique;not null"`
	Fullname       string `json:"fullname"`
	Password       string `json:"password"`
	Email          string `json:"email" gorm:"not null;unique"`
	Profile        string `json:"profile"`
	Portfolio      string `json:"portfolio"`
	CV             string `json:"cv" gorm:"not null"`
	Gender         string `json:"gender"`
	EmpStatus      string `json:"emp_status" gorm:"not null"`
}

func (js *Jobseeker) Addressable() {

}
