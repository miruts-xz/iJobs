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
	id                                        int64
	Uname, Fname, Lname, Email, Bio, password string
	Age, Phone                                int
	WrExprs                                   []Experience
	Gender                                    Gender
	Ctgrs                                     []Category
	EmpStatus                                 EmpStatus
	Portfolio                                 []Portfolio
	CvUrl                                     string
	Profile                                   string
	Address                                   Address
}

// getters and setters
func (js *JobSeeker) GetId() int64 {
	return js.id
}
func (js *JobSeeker) SetId(id int64) {
	js.id = id
}
func (js *JobSeeker) GetPass() string {
	return js.password
}
func (js *JobSeeker) SetPass(pass string) {
	js.password = pass
}
