package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

var Addressmock1 = Address{
	Model: gorm.Model{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	},
	Region:    Amhara,
	City:      Addis,
	SubCity:   Bole,
	LocalName: "Bole",
}
var Addressmock2 = Address{
	Model: gorm.Model{
		ID:        1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	},
	Region:    Oromia,
	City:      Addis,
	SubCity:   Arada,
	LocalName: "Piassa",
}
var Applicatiomock1 = Application{
	Model: gorm.Model{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	},
	JobID:       0,
	JobseekerID: 0,
	Response:    ACCEPTED,
	Status:      "reviewed",
}
var Applicatiomock2 = Application{
	Model: gorm.Model{
		ID:        1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	},
	JobID:       1,
	JobseekerID: 1,
	Response:    REJECTED,
	Status:      "reviewed",
}
var Categorymock1 = Category{
	Model: gorm.Model{
		ID: 0,
	},
	Jobs:  nil,
	Name:  "Tech",
	Image: "",
	Descr: "Technology related job category",
}
var Categorymock2 = Category{
	Model: gorm.Model{
		ID: 1,
	},
	Jobs:  nil,
	Name:  "General",
	Image: "",
	Descr: "General solution related job category",
}
var Companymock1 = Company{
	User:        Usermock1,
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
var Companymock2 = Company{
	User:        Usermock2,
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
var Jobmock1 = Job{
	Model:        gorm.Model{ID: 0},
	CompanyID:    0,
	Categories:   []Category{{Name: "Tech"}},
	Applications: nil,
	RequiredNum:  3,
	Salary:       6000,
	Name:         "Go Coder",
	Description:  "Go web developer",
	JobTime:      "fulltime",
	Deadline:     time.Time{},
}
var Jobmock2 = Job{
	Model:        gorm.Model{ID: 0},
	CompanyID:    0,
	Categories:   []Category{{Name: "Tech"}},
	Applications: nil,
	RequiredNum:  3,
	Salary:       10000,
	Name:         "Go Coder",
	Description:  "Go web developer",
	JobTime:      "fulltime",
	Deadline:     time.Time{},
}
var Jobseekermock1 = Jobseeker{
	User:           Usermock3,
	Address:        nil,
	Applications:   nil,
	Categories:     []Category{{Name: "Tech"}},
	Age:            20,
	Phone:          "0987654321",
	WorkExperience: 5,
	Username:       "user1",
	Fullname:       "user1 name",
	Password:       "user123",
	Email:          "user1@gmail.com",
	Profile:        "",
	Portfolio:      "",
	CV:             "",
	Gender:         "M",
	EmpStatus:      "",
}
var Jobseekermock2 = Jobseeker{
	User:           Usermock4,
	Address:        nil,
	Applications:   nil,
	Categories:     []Category{{Name: "Tech"}},
	Age:            20,
	Phone:          "0987654321",
	WorkExperience: 5,
	Username:       "user2",
	Fullname:       "user2 name",
	Password:       "user123",
	Email:          "user2@gmail.com",
	Profile:        "",
	Portfolio:      "",
	CV:             "",
	Gender:         "M",
	EmpStatus:      "",
}
var Sessionmock1 = Session{
	Uuid:       "sessionuuid1",
	UserID:     0,
	Email:      "user@gmail.com",
	SigningKey: []byte("sessionmock1"),
	Expires:    0,
}
var Sessionmock2 = Session{
	Uuid:       "sessionuuid2",
	UserID:     1,
	Email:      "user@gmail.com",
	SigningKey: []byte("sessionmock2"),
	Expires:    0,
}
var Usermock1 = User{
	Model: gorm.Model{
		ID: 0,
	},
	RoleID: 2,
}
var Usermock2 = User{
	Model: gorm.Model{
		ID: 1,
	},
	RoleID: 2,
}
var Usermock3 = User{
	Model: gorm.Model{
		ID: 2,
	},
	RoleID: 1,
}
var Usermock4 = User{
	Model: gorm.Model{
		ID: 3,
	},
	RoleID: 1,
}
var Rolemock1 = Role{
	ID:    1,
	Name:  "JOBSEEKER",
	Users: nil,
}
var Rolemock2 = Role{
	ID:    2,
	Name:  "COMPANY",
	Users: nil,
}
