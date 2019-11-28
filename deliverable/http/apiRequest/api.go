package apiRequest

import (
	"encoding/json"
	"github.com/miruts/iJobs/entity"
	"net/http"
)

// Rest api to returns users
func GetUsers(w http.ResponseWriter, r *http.Request) {

	data, _ := json.Marshal(entity.JobSeeker{
		Uname:     "miruts",
		Fname:     "miruts",
		Lname:     "hadush",
		Email:     "miruts.hadush@aait.edu.et",
		Bio:       "Student at Addis Ababa Institute of Technology",
		Age:       21,
		Phone:     945964841,
		WrExprs:   nil,
		Gender:    entity.MALE,
		Ctgrs:     nil,
		EmpStatus: entity.UNEMPLD,
		Portfolio: nil,
		CvUrl:     "",
		Profile:   "/store/users/miruts/profile.pgn",
		Address: entity.Address{
			Ctr:     "Ethiopia",
			Rgn:     "Addis Ababa",
			City:    "Addis Ababa",
			SbCty:   "Gullele",
			LclName: "Addisu Gebeya",
		},
	})
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Write(data)

}
