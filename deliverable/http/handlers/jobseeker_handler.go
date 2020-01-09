package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/application"
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
func NewJobseekerHandler(tmpl *template.Template, jss jobseeker.JobseekerService, jcs job.CategoryService, adds jobseeker.AddressService, apps application.IAppService, ss session.SessionService) *JobseekerHandler {
	return &JobseekerHandler{
		tmpl:    tmpl,
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
func (jsh *JobseekerHandler) JobseekerRegister(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "GET" {
		jobCtgs, err := jsh.ctgSrv.Categories()
		Regions := []string{entity.Tigray, entity.Amhara, entity.Sidama, entity.Afar, entity.Somalia, entity.Gambella, entity.Harare, entity.Snnpr, entity.Oromia, entity.Benshangul}
		Cities := []string{entity.Addis, entity.Mekele, entity.Hawassa, entity.Adamma, entity.Gonder}
		SubCities := []string{entity.Gulele, entity.Arada, entity.Yeka, entity.Bole, entity.Cherkos, entity.AddisKetema}
		var registerNeed RegisterNeed
		registerNeed.Categories = jobCtgs
		registerNeed.Regions = Regions
		registerNeed.Cities = Cities
		registerNeed.Subcities = SubCities
		if err != nil {
			fmt.Println(err)
			return
		}
		err = jsh.tmpl.ExecuteTemplate(w, "jobseeker.register.layout", registerNeed)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
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
}
func hasvalue(value interface{}) bool {
	if value != nil {
		return true
	}
	return false
}
func (jsh *JobseekerHandler) GetJobseekerHome(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	session, err := util.Authenticate(jsh.sessSrv, r)
	if err == nil {
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
			// Other http methods

		}
	} else {
		err := jsh.tmpl.ExecuteTemplate(w, "login.layout", nil)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
func (jsh *JobseekerHandler) JobseekerAppliedJobs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	session, err := util.Authenticate(jsh.sessSrv, r)
	if err == nil {
		if r.Method == "GET" {
			// Get http method
			jobappneeds := JobseekerAppliedNeed{}
			Appls, _ := jsh.appSrv.UserApplication(int(session.UserID))
			Ctgs, _ := jsh.ctgSrv.Categories()
			jobappneeds.Categories = Ctgs
			jobappneeds.Applications = Appls
			err := jsh.tmpl.ExecuteTemplate(w, "jobseeker.appliedjobs.layout", jobappneeds)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			// Other http methods

		}
	}
}
func (jsh *JobseekerHandler) JobseekerProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	session, err := util.Authenticate(jsh.sessSrv, r)
	if err == nil {
		if r.Method == "GET" {

		}
	}
}
