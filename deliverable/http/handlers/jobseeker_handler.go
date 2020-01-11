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
type AppliedJobCatNeed struct {
	Categories []entity.Category
	Jobseeker  entity.Jobseeker
	Catid      int
}
type JobseekerHomeNeed struct {
	Applications []entity.Application
	Suggestions  []entity.Job
	Categories   []entity.Category
	Jobseeker    entity.Jobseeker
}
type JobseekerAppliedNeed struct {
	Applications []entity.Application
	Categories   []entity.Category
	Jobseeker    entity.Jobseeker
}
type JobseekerProfileNeed struct {
	Categories []entity.Category
	Jobseeker  entity.Jobseeker
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
		fmt.Printf("Error Parsing Form: Line %d", 106)
		return
	}
	jobseeker := entity.Jobseeker{}
	empstatus := r.FormValue("empstatus")
	jobseeker.EmpStatus = entity.EmpStatus(empstatus)
	uname := r.FormValue("uname")
	age := r.FormValue("age")
	gender := r.FormValue("gender")

	jobseeker.Gender = entity.Gender(gender)
	if age != "" {
		ageint, err := strconv.Atoi(age)
		if err != nil {
			fmt.Println("Invalid Age value ")
			jobseeker.Age = 20
		} else {
			if ageint > 13 {
				jobseeker.Age = uint(ageint)
			} else {
				jobseeker.Age = 20
			}
		}
	}
	phone := r.FormValue("phone")

	jobseeker.Phone = phone
	if uname == "" {
		fmt.Printf("Please provide Username: Line %d", 112)
		return
	}
	_, err = jsh.jsSrv.JobseekerByUsername(uname)
	if err == nil {
		fmt.Printf("username already taken : Line %d", 116)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	email := r.FormValue("email")

	if email == "" {
		fmt.Printf("Email is empty: Line %d", 124)
	}
	_, err = jsh.jsSrv.JobseekerByEmail(email)
	if err == nil {
		fmt.Printf("Email must be unique: Line %d", 128)
		return
	}
	jobseeker.Email = email

	jobseeker.Username = uname
	// todo process firstname and lastname
	fname := r.FormValue("fname")
	if fname == "" {
		fmt.Printf("First Name is required : Line %d", 137)
	}
	lname := r.FormValue("lname")
	if lname == "" {
		fmt.Printf("Last name is required : Line %d", 141)
	}
	jobseeker.Fullname = fname + " " + lname
	// todo process, make it secure and store user entered password
	pswd := r.FormValue("pswd")
	if len(pswd) < 3 {
		fmt.Printf("Empty or invalid password : Line %d", 147)
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(pswd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error Hashing : Line %d", 151)
		return
	}
	hashedpwsd := string(hashed)

	jobseeker.Password = hashedpwsd
	// todo process and store user entered work experience
	wrkexp := r.FormValue("wrkexp")
	if wrkexp == "" {
		wrkexp = string(0)
	}
	wrkexpint, err := strconv.Atoi(wrkexp)
	if err != nil {
		fmt.Printf("Invalid work experience : Line %d", 164)
	}
	jobseeker.WorkExperience = wrkexpint

	// todo process and store user entered portfolio
	portf := r.FormValue("portf")
	if portf == "" {
		fmt.Printf("Portfolio not provided %d", 171)
	} else {
		jobseeker.Portfolio = portf
	}
	// todo process and store user entered profile picture
	propic, fh, err := r.FormFile("propic")
	if err != nil {
		fmt.Printf("Profile picture not provided : Line %d", 178)
		path := "/assets/img/avatar"
		jobseeker.Profile = path

	} else {
		path, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		path = filepath.Join(path, "ui", "asset", "jsdata", uname, "pp")
		err = os.MkdirAll(path, 0644)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		path = filepath.Join(path, fh.Filename)
		written := util.SaveFile(propic, path)
		if !written {
			fmt.Printf("Not written profile picture : Line %d", 197)
		}
		jobseeker.Profile = fh.Filename
	}
	// todo process and store user entered cv
	cv, fh, err := r.FormFile("cv")
	if err != nil {
		return
	}
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("Portfolio is required : Line %d", 209)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	fmt.Println(path)
	path = filepath.Join(path, "ui", "asset", "jsdata", uname, "cv")
	err = os.MkdirAll(path, 0644)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	path = filepath.Join(path, fh.Filename)
	cvWritten := util.SaveFile(cv, path)
	if !cvWritten {
		fmt.Printf("Not written curriculum vitae : Line %d", 223)
	}
	jobseeker.CV = fh.Filename

	// todo process and store selected interested job categories
	intjobcat := r.Form["intjobcat"]
	if !hasvalue(intjobcat) {
		return
	}
	var categories []entity.Category
	for v := range intjobcat {
		ctg, err := jsh.ctgSrv.Category(v)
		if err != nil {
			fmt.Printf("Category not found")
			continue
		}
		categories = append(categories, ctg)
	}
	jobseeker.Categories = categories

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
	jobseeker.Address = []entity.Address{address}

	js, err := jsh.jsSrv.StoreJobSeeker(&jobseeker)
	if err != nil {
		fmt.Printf("Storing jobseeker failed : Line %d", 268)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	fmt.Println("Jobseeker registered successfully", js)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
func (jsh *JobseekerHandler) JobseekerHome(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ok, session := util.Authenticate(jsh.sessSrv, r)
	if ok == true {
		if r.Method == "GET" {
			// Get http method
			jsneeds := JobseekerHomeNeed{}
			jobseeker, err := jsh.jsSrv.JobSeeker(int(session.UserID))
			if err != nil {
				fmt.Printf("Couldn't find jobseeker info : Line %d", 271)
				return
			}
			Ctgs, _ := jsh.ctgSrv.Categories()
			Apps, _ := jsh.appSrv.UserApplication(int(session.UserID))
			Suggs, _ := jsh.jsSrv.Suggestions(int(session.UserID))
			jsneeds.Categories = Ctgs
			jsneeds.Applications = Apps
			jsneeds.Suggestions = Suggs
			jsneeds.Jobseeker = jobseeker
			err = jsh.tmpl.ExecuteTemplate(w, "jobseeker.layout", jsneeds)
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
			jobseeker, err := jsh.jsSrv.JobSeeker(int(session.UserID))
			if err != nil {
				fmt.Println("Unable to retrieve jobseeker: Line %", 328)
				http.Redirect(w, r, "/login", http.StatusSeeOther)
			}
			jobappneeds := JobseekerAppliedNeed{}
			Appls, _ := jsh.appSrv.UserApplication(int(session.UserID))
			Ctgs, _ := jsh.ctgSrv.Categories()
			jobappneeds.Jobseeker = jobseeker
			jobappneeds.Categories = Ctgs
			jobappneeds.Applications = Appls

			err = jsh.tmpl.ExecuteTemplate(w, "jobseeker.appliedJobs.layout", jobappneeds)
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

			fmt.Println(jobseeker)
			if err != nil {
				return
			}
			Ctgs, err := jsh.ctgSrv.Categories()
			if err != nil {
				return
			}
			jspneeds := JobseekerProfileNeed{}
			jspneeds.Categories = Ctgs
			jspneeds.Jobseeker = jobseeker

			fmt.Println("Employment Status", jobseeker.EmpStatus)
			err = jsh.tmpl.ExecuteTemplate(w, "jobseeker.profile.layout", jspneeds)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	} else {
		err := util.DestroySession(&w, r)
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
func (jsh *JobseekerHandler) AppliedJobCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ok, session := util.Authenticate(jsh.sessSrv, r)
	if ok {
		if r.Method == "GET" {
			id := ps.ByName("id")
			idint, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println(err)
				return
			}
			jobseeker, err := jsh.jsSrv.JobSeeker(int(session.UserID))

			fmt.Println(jobseeker)
			if err != nil {
				return
			}
			Ctgs, err := jsh.ctgSrv.Categories()
			if err != nil {
				return
			}
			appjobneeds := AppliedJobCatNeed{}
			appjobneeds.Categories = Ctgs
			appjobneeds.Jobseeker = jobseeker
			appjobneeds.Catid = idint

			err = jsh.tmpl.ExecuteTemplate(w, "jobseeker.appliedJobs.category.layout", appjobneeds)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	} else {
		err := util.DestroySession(&w, r)
		if err != nil {
			fmt.Println(err)
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
func AppGetCmpLogo(app entity.Application) (string, error) {
	job, err := jobSrvc.Job(int(app.JobID))
	if err != nil {
		return string(job.CompanyID), err
	}
	cmp, err := cmpSrvc.Company(int(job.CompanyID))
	if err != nil {
		return cmp.Logo, err
	}
	return cmp.Logo, nil
}
func AppGetJobCatId(app entity.Application) ([]int, error) {
	job, err := jobSrvc.Job(int(app.JobID))
	var catsId []int
	if err != nil {
		fmt.Println(err)
		return catsId, err
	}
	for id, _ := range job.Categories {
		catsId = append(catsId, id)
	}
	return catsId, nil
}
