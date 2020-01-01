package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/jobseeker"
	"github.com/miruts/iJobs/util"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
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
func (jsh *JobseekerHandler) JobseekerRegister(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	err = r.ParseMultipartForm(1024)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	jobseeker := entity.JobSeeker{}
	uname := ps.ByName("uname")
	if !hasvalue(uname) {
	}
	jobseeker.Username = uname
	// todo process firstname and lastname
	fname := ps.ByName("fname")
	if !hasvalue(fname) {
	}
	lname := ps.ByName("lname")
	if !hasvalue(lname) {

	}
	jobseeker.Fullname = fname + " " + lname
	// todo process, make it secure and store user entered password
	pswd := ps.ByName("pswd")
	if !hasvalue(pswd) || len(pswd) < 4 {

	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(pswd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	hashedpwsd := string(hashed)
	fmt.Print(hashedpwsd)
	jobseeker.Password = hashedpwsd
	// todo process and store user entered work experience
	wrkexp := ps.ByName("wrkexp")
	wrkexpint, err := strconv.Atoi(wrkexp)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	if !hasvalue(wrkexp) || wrkexpint < 0 {
	}
	jobseeker.WorkExperience = wrkexpint

	// todo process and store user entered portfolio
	portf := r.FormValue("portf")
	if !hasvalue(portf) {

	}
	jobseeker.Portfolio = portf

	// todo process and store user entered profile picture
	propic, fh, err := r.FormFile("propic")
	if err != nil {

	}
	path, err := os.Getwd()
	fmt.Println(path)
	path = path[:len(path)-25]
	fmt.Println(path)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	path = filepath.Join(path, "ui", "assets", "jsdata", uname, "pp")
	err = os.MkdirAll(path, 0644)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	path = filepath.Join(path, fh.Filename)
	written := util.SaveFile(propic, path)
	if !written {

	}

	// todo process and store user entered cv
	cv, fh, err := r.FormFile("cv")
	if err != nil {

	}
	path, err = os.Getwd()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Println(path)
	path = path[:len(path)-25]
	fmt.Println(path)
	path = filepath.Join(path, "ui", "assets", "jsdata", uname, "cv")
	err = os.MkdirAll(path, 0644)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	path = filepath.Join(path, fh.Filename)
	cvWritten := util.SaveFile(cv, path)
	if !cvWritten {

	}

	err = jsh.jsSrv.StoreJobSeeker(jobseeker)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	js, err := jsh.jsSrv.JobSeekers()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	var jsid int64
	for _, v := range js {
		if v.Username == uname {
			jsid = int64(v.ID)
			break
		}
	}

	// todo process and store selected interested job categories
	intjobcat := r.Form["intjobcat"]
	if !hasvalue(intjobcat) {
	}
	for v := range intjobcat {
		jcid := v
		err = jsh.jsSrv.AddIntCategory(int(jsid), jcid)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
	}
}
func hasvalue(value interface{}) bool {
	if value != nil {
		return true
	}
	return false
}
