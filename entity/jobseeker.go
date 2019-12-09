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

type JobSeeker struct {
	ID, Age, Phone, WorkExperience                              int64
	Username, Email, Fullname, Password, Profile, Portfolio, CV string
	Gender                                                      Gender
	Categories                                                  []Category
	EmpStatus                                                   EmpStatus
	Address                                                     Address
}
