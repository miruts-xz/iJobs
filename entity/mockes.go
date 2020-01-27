package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

var addressmock = Address{
	Model:     gorm.Model{},
	Region:    Oromia,
	City:      Addis,
	SubCity:   Gulele,
	LocalName: "Piassa",
}
var appliccatiomock = Application{
	Model:       gorm.Model{},
	JobID:       0,
	JobseekerID: 0,
	Response:    ACCEPTED,
	Status:      "unreviewed",
}
var categorymock = Category{
	Model: gorm.Model{},
	Jobs:  nil,
	Name:  "Tech",
	Image: "",
	Descr: "Technology related job category",
}
var companymock = Company{
	Jobs:        nil,
	Address:     nil,
	CompanyName: "ABC Software Inc",
	Password:    "123",
	Email:       "abc@software.com",
	Phone:       "094554545",
	Logo:        "",
	ShortDesc:   "",
	DetailInfo:  "",
}
var jobmock = Job{
	Model:        gorm.Model{},
	CompanyID:    0,
	Categories:   []Category{{Name: "Tech"}},
	Applications: nil,
	RequiredNum:  3,
	Salary:       6000,
	Name:         "Go Coder",
	Description:  "",
	JobTime:      "",
	Deadline:     time.Time{},
}
var jobseekermock = Jobseeker{
	Address:        nil,
	Applications:   nil,
	Categories:     []Category{{Name: "Tech"}},
	Age:            20,
	Phone:          "0987654321",
	WorkExperience: 5,
	Username:       "user",
	Fullname:       "user name",
	Password:       "user123",
	Email:          "user@gmail.com",
	Profile:        "",
	Portfolio:      "",
	CV:             "",
	Gender:         "M",
	EmpStatus:      "",
}
var sessionmock = Session{
	Model:  gorm.Model{},
	Uuid:   "sessionuuid",
	UserID: 0,
	Email:  "user@gmail.com",
}
