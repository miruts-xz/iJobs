package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/jobseeker"
	"html/template"
	"net/http"
)

// JobseekerHandler handles jobseeker related http requests
type JobseekerHandler struct {
	tmpl  *template.Template
	jsSrv jobseeker.JobseekerService
}

// NewJobseekerHandler creates new JobseekerHandler
func NewJobseekerHandler(tmpl *template.Template, jss jobseeker.JobseekerService) *JobseekerHandler {
	return &JobseekerHandler{
		tmpl:  tmpl,
		jsSrv: jss,
	}
}

/**
 JobseekerRegister handles /jobseeker/register requests
	query parameter mapping
	form name 		data
	uname 		- username
	fname 		- first name
	lname 		- last name
	propic 		- profile picture
	pswd   		- password
	intjobcat 	- interested job categories
	wrkexp 		- work experience
	portf		- portfolio
	cv 			- cv
	phone 		- phone
	email 		- email
	gender 		- gender
*/
func JobseekerRegister(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	err = r.ParseMultipartForm(1024)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	uname := ps.ByName("uname")
	if !hasvalue(uname) {

	}
	// todo process firstname and lastname
	fname := ps.ByName("fname")
	if !hasvalue(fname) {

	}
	lname := ps.ByName("lname")
	if !hasvalue(lname) {

	}
	// todo process, make it secure and store user entered password
	pswd := ps.ByName("pswd")
	if !hasvalue(pswd) {

	}
	// todo process and store user entered work exprience
	wrkexp := ps.ByName("wrkexp")
	if !hasvalue(wrkexp) {

	}
	// todo process and store selected interested job categories
	intjobcat := r.Form["intjobcat"]
	if !hasvalue(intjobcat) {

	}
	// todo process and store user entered
	portf := r.Form["portf"]
	if !hasvalue(portf) {

	}
	// todo process and store user entered profile picture
	propic, fh, err := r.FormFile("propic")
	if err != nil {

	}
	// todo process and store user entered cv
	cv, fh, err := r.FormFile("cv")
	if err != nil {

	}

	jobseeker := entity.JobSeeker{}

}
func hasvalue(value interface{}) bool {
	if value != nil {
		return true
	}
	return false
}
