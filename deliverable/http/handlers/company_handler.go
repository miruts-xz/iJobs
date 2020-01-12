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
	"time"
)

type CompanyHandler struct {
	tmpl    *template.Template
	jsSrv   jobseeker.JobseekerService
	cmpSrv  company.CompanyService
	ctgSrv  job.CategoryService
	addrSrv jobseeker.AddressService
	appSrv  application.IAppService
	sessSrv session.SessionService
	jobSrv  job.JobService
}

// NewCompanyHandler creates new CompanyHandler
func NewCompanyHandler(tmpl *template.Template, jsSrv jobseeker.JobseekerService, cmpSrv company.CompanyService, ctgSrv job.CategoryService, addrSrv jobseeker.AddressService, appSrv application.IAppService, sessSrv session.SessionService, jobSrv job.JobService) *CompanyHandler {
	return &CompanyHandler{
		tmpl:    tmpl,
		jsSrv:   jsSrv,
		cmpSrv:  cmpSrv,
		ctgSrv:  ctgSrv,
		addrSrv: addrSrv,
		appSrv:  appSrv,
		sessSrv: sessSrv,
		jobSrv:  jobSrv,
	}
}

// CompanyPostJodNeed stores information required for posting job
type CompanyPostJobNeed struct {
	Categories []entity.Category
	Company    entity.Company
}
type CompanyHomeNeed struct {
	Company      entity.Company
	Applications []entity.Application
	Jobs         []entity.Job
}

// CompanyRegister handles company signup request.
func (ch *CompanyHandler) CompanyRegister(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	err = r.ParseMultipartForm(1024)
	if err != nil {
		fmt.Printf("Error Parsing Form: Line %d", 106)
		return
	}
	company := entity.Company{}
	name := r.FormValue("name")
	if name == "" {
		fmt.Println("Name field is require")
		return
	}
	company.CompanyName = name

	password := r.FormValue("pswd")
	if password == "" {
		fmt.Println("Password is require")
		return
	}
	confirm := r.FormValue("confirm")
	if password != confirm {
		fmt.Println("Password not maching")
		return
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedpassword := string(hashed)
	company.Password = hashedpassword

	email := r.FormValue("email")
	if email == "" {
		fmt.Println("Email is required")
		return
	}
	_, err1 := ch.jsSrv.JobseekerByEmail(email)
	_, err2 := ch.cmpSrv.CompanyByEmail(email)
	if err1 == nil || err2 == nil {
		fmt.Println("Email already taken ")
		return
	}
	company.Email = email

	phone := r.FormValue("phone")
	company.Phone = phone

	detailinfo := r.FormValue("detailinfo")
	shortdescr := r.FormValue("shortdesc")

	if shortdescr == "" {
		fmt.Println("Short Description is required")
		return
	}
	company.DetailInfo = detailinfo
	company.ShortDesc = shortdescr

	logo, fh, err := r.FormFile("logo")
	if err != nil {
		fmt.Println(err)
		return
	}
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	path = filepath.Join(path, "ui", "asset", "cmpdata", name, "logo")
	err = os.MkdirAll(path, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	path = filepath.Join(path, fh.Filename)
	logoWritten := util.SaveFile(logo, path)
	if !logoWritten {
		fmt.Println("Not written company logo")
	}
	company.Logo = fh.Filename

	region := r.FormValue("region")
	city := r.FormValue("city")
	subcity := r.FormValue("subcity")
	localname := r.FormValue("localname")
	address := entity.Address{}
	address.Region = region
	address.City = city
	address.SubCity = subcity
	address.LocalName = localname
	company.Address = []entity.Address{address}
	cmp, err := ch.cmpSrv.StoreCompany(&company)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	fmt.Println("Company registered successfully", cmp)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return
}
func (ch *CompanyHandler) CompanyHome(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ok, session := util.Authenticate(ch.sessSrv, r)
	if ok == true {
		if r.Method == "GET" {
			cmpneeds := CompanyHomeNeed{}
			company, err := ch.cmpSrv.Company(int(session.UserID))
			if err != nil {
				fmt.Println("Couldn't find company info")
				return
			}
			Jobs, err := ch.jobSrv.CompanyJobs(ch.cmpSrv, int(session.UserID))
			Applications, err := ch.appSrv.ApplicationForCompany(int(session.UserID))
			cmpneeds.Company = company
			cmpneeds.Jobs = Jobs
			cmpneeds.Applications = Applications

			err = ch.tmpl.ExecuteTemplate(w, "company.home.layout", cmpneeds)
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
func (ch *CompanyHandler) CompanyPostJob(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ok, session := util.Authenticate(ch.sessSrv, r)
	if ok {
		company, err := ch.cmpSrv.Company(int(session.UserID))
		if r.Method == "GET" {
			// Get http method
			if err != nil {
				fmt.Println("Unable to retrieve Company: Line ")
				http.Redirect(w, r, "/login", http.StatusSeeOther)
			}
			postneeds := CompanyPostJobNeed{}
			Ctgs, _ := ch.ctgSrv.Categories()
			postneeds.Categories = Ctgs
			postneeds.Company = company
			err = ch.tmpl.ExecuteTemplate(w, "company.postjob.layout", postneeds)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			err := r.ParseForm()
			err = r.ParseMultipartForm(1024)
			if err != nil {
				fmt.Println(err)
				return
			}
			var job entity.Job
			ctgs := r.Form["ctgs"]
			jbtitle := r.FormValue("jbtitle")
			description := r.FormValue("description")
			reqnum := r.FormValue("reqnum")
			salary := r.FormValue("salary")
			jbtime := r.FormValue("jbtime")
			deadline := r.FormValue("deadline")
			fmt.Println(ctgs)
			var categories []entity.Category
			for i, _ := range ctgs {
				intid, err := strconv.Atoi(ctgs[i])
				fmt.Println(intid)
				if err == nil {
					category, err := ch.ctgSrv.Category(intid)
					if err == nil {
						categories = append(categories, category)
					}
				}
			}
			fmt.Println(categories)
			job.Categories = categories

			job.Name = jbtitle
			job.Description = description
			req, _ := strconv.Atoi(reqnum)
			job.RequiredNum = uint(req)
			sal, _ := strconv.ParseFloat(salary, 64)
			job.Salary = sal
			job.JobTime = jbtime
			job.Deadline, _ = time.Parse("2006-01-02", deadline)
			if jbtitle == "" {
				fmt.Println("Job title is required")
				return
			}
			if len(ctgs) == 0 {
				fmt.Println("Please choose atleast one category")
				return
			}
			if reqnum == "" {
				job.RequiredNum = 1
			}
			if salary == "" {
				fmt.Println("Please enter salary")
				return
			}
			if jbtime == "" {
				job.JobTime = "fulltime"
			}
			if deadline == "" {
				job.Deadline = time.Now().Add(24 * 14 * time.Hour)
			}
			company.Jobs = []entity.Job{job}
			cmp, err := ch.cmpSrv.UpdateCompany(&company)
			j, err := ch.jobSrv.UpdateJob(&job)
			fmt.Println(job.Categories)
			if err != nil {
				fmt.Println("Unable to post a job", cmp, j)
				return
			}
			fmt.Println("Job posted successfully")
			http.Redirect(w, r, "/company/"+company.CompanyName+"jobs", http.StatusSeeOther)
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
func AppJs(app entity.Application) (entity.Jobseeker, error) {
	jsid := app.JobseekerID
	jobseeker, err := jsSrvc.ApplicationJobseeker(int(jsid))
	if err != nil {
		fmt.Println(err)
		return jobseeker, err
	}
	return jobseeker, nil
}
func AppJob(app entity.Application) (entity.Job, error) {
	jid := app.JobID
	job, err := jobSrvc.Job(int(jid))
	if err != nil {
		fmt.Println(err)
		return job, err
	}
	return job, nil
}
func JobCmp(job entity.Job) (entity.Company, error) {
	cmid := job.CompanyID
	fmt.Println(cmid)
	company, err := cmpSrvc.Company(int(cmid))
	if err != nil {
		fmt.Println(err)
		return company, err
	}
	return company, nil
}
