package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/application"
	"github.com/miruts/iJobs/usecases/company"
	"github.com/miruts/iJobs/usecases/job"
	"github.com/miruts/iJobs/usecases/jobseeker"
	"github.com/miruts/iJobs/usecases/session"
	"github.com/miruts/iJobs/util"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// JobseekerHandler handles jobseeker related http requests
var tmpl *template.Template
var jsSrvc jobseeker.JobseekerService
var appSrvc application.IAppService
var ctgSrvc job.CategoryService
var jobSrvc job.JobService
var cmpSrvc company.CompanyService

type JobseekerHandler struct {
	tmpl    *template.Template
	jsSrv   jobseeker.JobseekerService
	ctgSrv  job.CategoryService
	addrSrv jobseeker.AddressService
	appSrv  application.IAppService
	sessSrv session.SessionService
}
type RegisterNeed struct {
	Categories []entity.Category
	Regions    []string
	Cities     []string
	Subcities  []string
}
type JobseekerHomeNeed struct {
	Applications []entity.Application
	Suggestions  []entity.Job
	Categories   []entity.Category
}
type JobseekerAppliedNeed struct {
	Applications []entity.Application
	Categories   []entity.Category
}
type JobseekerProfileNeed struct {
	Categories []entity.Category
	jobseeker  entity.Jobseeker
}

// NewJobseekerHandler creates new JobseekerHandler
func NewJobseekerHandler(tmplt *template.Template, jss jobseeker.JobseekerService, jcs job.CategoryService, adds jobseeker.AddressService, apps application.IAppService, ss session.SessionService, jssr job.JobService, cmpsrv company.CompanyService) *JobseekerHandler {
	tmpl = tmplt
	appSrvc = apps
	jsSrvc = jss
	ctgSrvc = jcs
	jobSrvc = jssr
	cmpSrvc = cmpsrv
	return &JobseekerHandler{
		tmpl:    tmplt,
		jsSrv:   jss,
		ctgSrv:  jcs,
		addrSrv: adds,
		appSrv:  apps,
		sessSrv: ss,
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
	region		-region
	city		-city
	subcity		-subcity
	localname	-localname
*/
func hasvalue(value interface{}) bool {
	if value != nil {
		return true
	}
	return false
}
func (jsh *JobseekerHandler) JobseekerRegister(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	err = r.ParseMultipartForm(1024)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	jobseeker := entity.Jobseeker{}
	uname := ps.ByName("uname")
	if !hasvalue(uname) {
		return
	}
	jobseeker.Username = uname
	// todo process firstname and lastname
	fname := ps.ByName("fname")
	if !hasvalue(fname) {
		return
	}
	lname := ps.ByName("lname")
	if !hasvalue(lname) {
		return
	}
	jobseeker.Fullname = fname + " " + lname
	// todo process, make it secure and store user entered password
	pswd := ps.ByName("pswd")
	if !hasvalue(pswd) || len(pswd) < 4 {
		return
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
		jobseeker.WorkExperience = 0
	}
	jobseeker.WorkExperience = wrkexpint

	// todo process and store user entered portfolio
	portf := r.FormValue("portf")
	if !hasvalue(portf) {
		return
	}
	jobseeker.Portfolio = portf

	// todo process and store user entered profile picture
	propic, fh, err := r.FormFile("propic")
	if err != nil {
		path := "/assets/img/avatar"
		jobseeker.Profile = path

	} else {
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
		imageUri := filepath.Join("/assets", "jsdata", uname, "pp", fh.Filename)
		fmt.Println(imageUri)
		jobseeker.Profile = imageUri
	}
	// todo process and store user entered cv
	cv, fh, err := r.FormFile("cv")
	if err != nil {
		return
	}
	path, err := os.Getwd()
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
	cvUri := filepath.Join("/assets", "jsdata", uname, "cv", fh.Filename)
	fmt.Println(cvUri)
	jobseeker.CV = cvUri
	_, err = jsh.jsSrv.StoreJobSeeker(&jobseeker)
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
		return
	}
	for v := range intjobcat {
		jcid := v
		err = jsh.jsSrv.AddIntCategory(int(jsid), jcid)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
	}
	// todo process and store addresses
	region := r.FormValue("region")
	city := r.FormValue("city")
	subcity := r.FormValue("subcity")
	localname := r.FormValue("localname")
	address := entity.Address{}
	address.Region = region
	address.City = city
	address.SubCity = subcity
	address.LocalName = localname
	adr, err := jsh.addrSrv.StoreAddress(&address)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = jsh.jsSrv.SetAddress(int(jobseeker.ID), int(adr.ID))
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (jsh *JobseekerHandler) JobseekerHome(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ok, session := util.Authenticate(jsh.sessSrv, r)
	if ok == true {
		if r.Method == "GET" {
			// Get http method
			jsneeds := JobseekerHomeNeed{}

			Ctgs, _ := jsh.ctgSrv.Categories()
			Apps, _ := jsh.appSrv.UserApplication(int(session.UserID))
			Suggs, _ := jsh.jsSrv.Suggestions(int(session.UserID))
			jsneeds.Categories = Ctgs
			jsneeds.Applications = Apps
			jsneeds.Suggestions = Suggs
			err := jsh.tmpl.ExecuteTemplate(w, "jobseeker.layout", jsneeds)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {

		}
	} else {
		err := util.DestroySession(&w, r)
		fmt.Println("Destroying Session")
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
func (jsh *JobseekerHandler) JobseekerAppliedJobs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ok, session := util.Authenticate(jsh.sessSrv, r)
	if ok {
		if r.Method == "GET" {
			// Get http method
			jobappneeds := JobseekerAppliedNeed{}
			Appls, _ := jsh.appSrv.UserApplication(int(session.UserID))
			Ctgs, _ := jsh.ctgSrv.Categories()
			jobappneeds.Categories = Ctgs
			jobappneeds.Applications = Appls
			err := jsh.tmpl.ExecuteTemplate(w, "jobseeker.appliedJobs.layout", jobappneeds)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			// Other http methods

		}
	} else {
		err := util.DestroySession(&w, r)
		if err != nil {
			fmt.Println(err)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
func (jsh *JobseekerHandler) JobseekerProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ok, session := util.Authenticate(jsh.sessSrv, r)
	if ok {
		if r.Method == "GET" {
			jobseeker, err := jsh.jsSrv.JobSeeker(int(session.UserID))
			if err != nil {
				return
			}
			Ctgs, err := jsh.ctgSrv.Categories()
			if err != nil {
				return
			}
			jspneeds := JobseekerProfileNeed{}
			jspneeds.Categories = Ctgs
			jspneeds.jobseeker = jobseeker
			err = jsh.tmpl.ExecuteTemplate(w, "jobseeker.profile.layout", jspneeds)
			if err != nil {
				return
			}
		}
	} else {
		err := util.DestroySession(&w, r)
		if err != nil {
			fmt.Println(err)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
func AppGetJobsName(app entity.Application) (string, error) {
	job, err := jobSrvc.Job(int(app.JobID))
	if err != nil {
		return job.Name, err
	}
	return job.Name, nil
}
func AppGetCmpName(app entity.Application) (string, error) {
	job, err := jobSrvc.Job(int(app.JobID))
	if err != nil {
		return string(job.CompanyID), err
	}
	cmp, err := cmpSrvc.Company(int(job.CompanyID))
	if err != nil {
		return cmp.CompanyName, err
	}
	return cmp.CompanyName, nil
}
func AppGetLocation(app entity.Application) (entity.Address, error) {
	var addr entity.Address
	job, err := jobSrvc.Job(int(app.JobID))
	if err != nil {
		return addr, err
	}
	cmp, err := cmpSrvc.Company(int(job.CompanyID))
	if err != nil {
		return addr, err
	}
	addr, err = cmpSrvc.CompanyAddress(cmp.ID)
	if err != nil {
		return addr, err
	}
	return addr, nil
}
