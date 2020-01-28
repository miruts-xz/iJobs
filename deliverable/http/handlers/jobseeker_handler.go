package handlers

import (
	"context"
	"fmt"
	"github.com/betsegawlemma/web-prog-go-sample/rtoken"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/role"
	"github.com/miruts/iJobs/security/form"
	"github.com/miruts/iJobs/security/permission"
	sess "github.com/miruts/iJobs/security/session"
	"github.com/miruts/iJobs/usecases/application"
	"github.com/miruts/iJobs/usecases/company"
	"github.com/miruts/iJobs/usecases/job"
	"github.com/miruts/iJobs/usecases/jobseeker"
	"github.com/miruts/iJobs/usecases/session"
	"github.com/miruts/iJobs/util"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
	jobSrv  job.JobService
	cmpSrv  company.CompanyService

	userSess     *entity.Session
	loggedInUser *entity.Jobseeker
	userRole     role.RoleService
	csrfSignKey  []byte
}
type contextKey string

var ctxUserSessionKey = contextKey("signed_in_user_session")

func NewJobseekerHandler(cmpSrv company.CompanyService, jobSrv job.JobService, tmpl *template.Template, jsSrv jobseeker.JobseekerService, ctgSrv job.CategoryService, addrSrv jobseeker.AddressService, appSrvcc application.IAppService, sessSrv session.SessionService, userSess *entity.Session, userRole role.RoleService, csrfSignKey []byte) *JobseekerHandler {
	jsSrvc = jsSrv
	ctgSrvc = ctgSrv
	appSrvc = appSrvcc
	jobSrvc = jobSrv
	cmpSrvc = cmpSrv
	return &JobseekerHandler{jobSrv: jobSrv, tmpl: tmpl, jsSrv: jsSrv, ctgSrv: ctgSrv, addrSrv: addrSrv, appSrv: appSrvcc, sessSrv: sessSrv, userSess: userSess, userRole: userRole, csrfSignKey: csrfSignKey}
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
type JobseekerProfileEditNeed struct {
	Jobseeker  entity.Jobseeker
	Categories []entity.Category
	Regions    []string
	Cities     []string
	SubCities  []string
	FName      string
	LName      string
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

// Authenticated checks if a user is authenticated to access a given route
func (uh *JobseekerHandler) Authenticated(next http.Handler) httprouter.Handle {
	fn := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ok := uh.loggedIn(r)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), ctxUserSessionKey, uh.userSess)
		//call next middleware with new context
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return fn
}

func (jsh *JobseekerHandler) JobseekerApply(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	jobid := r.URL.Query().Get("jobid")
	jobidint, err := strconv.Atoi(jobid)
	job, err := jsh.jobSrv.Job(jobidint)

	applyform := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}

	alreadyApplied := jsh.jsSrv.AlreadyApplied(jsh.loggedInUser.ID, job.ID)
	if alreadyApplied {
		applyform.VErrors.Add("generic", "Already Applied")
		jsh.tmpl.ExecuteTemplate(w, "jobseeker.appliedjobs.layout", applyform)
		return
	}
	appl := entity.Application{Status: "Unreviewed", JobID: job.ID, JobseekerID: jsh.loggedInUser.ID}
	app, err := jsh.appSrv.Store(&appl)
	jsh.loggedInUser.Applications = append(jsh.loggedInUser.Applications, *app)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/jobseeker/"+jsh.loggedInUser.Username+"/appliedjobs", http.StatusSeeOther)
	return
}
func contains(apps []entity.Application, job entity.Job) bool {
	for _, a := range apps {
		if a.JobID == job.ID {
			return true
		}
	}
	return false
}
func (jsh *JobseekerHandler) JobseekerHome(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		Ctgs, _ := jsh.ctgSrv.Categories()
		Apps, _ := jsh.appSrv.UserApplication(int(jsh.loggedInUser.ID))
		Suggs, _ := jsh.jsSrv.Suggestions(int(jsh.loggedInUser.ID))
		var sugges []entity.Job
		for _, s := range Suggs {
			if !contains(Apps, s) {
				sugges = append(sugges, s)
			}
		}
		jobseekerhomeinfo := struct {
			Categories   []entity.Category
			Applications []entity.Application
			Suggestions  []entity.Job
			Jobseeker    entity.Jobseeker
		}{
			Categories:   Ctgs,
			Applications: Apps,

			Suggestions: sugges,
			Jobseeker:   *jsh.loggedInUser,
		}
		err := jsh.tmpl.ExecuteTemplate(w, "jobseeker.layout", jobseekerhomeinfo)
		if err != nil {
			w.Header().Set("Status", http.StatusText(http.StatusInternalServerError))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		} else {
			w.Header().Set("Status", http.StatusText(http.StatusOK))
		}
	}
}

//JobseekerAppliedJobs displays Job ith specified ID or AllJobs Posted in the Category
func (jsh *JobseekerHandler) JobseekerAppliedJobs(w http.ResponseWriter, r *http.Request) {
	Appls, _ := jsh.appSrv.UserApplication(int(jsh.loggedInUser.ID))
	Ctgs, _ := jsh.ctgSrv.Categories()

	appliedjobsinfo := struct {
		Jobseeker    entity.Jobseeker
		Categories   []entity.Category
		Applications []entity.Application
	}{
		Jobseeker:    *jsh.loggedInUser,
		Categories:   Ctgs,
		Applications: Appls,
	}

	err := jsh.tmpl.ExecuteTemplate(w, "jobseeker.appliedJobs.layout", appliedjobsinfo)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

//JobseekerProfile display the JobSeeker profile page
func (jsh *JobseekerHandler) JobseekerProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		Ctgs, err := jsh.ctgSrv.Categories()
		if err != nil {
			return
		}
		jsprofileinfo := struct {
			Categories []entity.Category
			Jobseeker  entity.Jobseeker
		}{
			Categories: Ctgs,
			Jobseeker:  *jsh.loggedInUser,
		}

		err = jsh.tmpl.ExecuteTemplate(w, "jobseeker.profile.layout", jsprofileinfo)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}

//ProfileEdit display and edit JobSeekers Profile
func (jsh *JobseekerHandler) ProfileEdit(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(jsh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	jobCtgs, err := jsh.ctgSrv.Categories()
	Regions := []string{entity.Tigray, entity.Amhara, entity.Sidama, entity.Afar, entity.Somalia, entity.Gambella, entity.Harare, entity.Snnpr, entity.Oromia, entity.Benshangul}
	Cities := []string{entity.Addis, entity.Mekele, entity.Hawassa, entity.Adamma, entity.Gonder}
	SubCities := []string{entity.Gulele, entity.Arada, entity.Yeka, entity.Bole, entity.Cherkos, entity.AddisKetema}
	names := strings.Split(jsh.loggedInUser.Fullname, " ")
	FName := names[0]
	var LName string
	if len(names) > 1 {
		LName = names[1]
	}
	profileinfo := struct {
		Inputs form.Input

		Jobseeker  entity.Jobseeker
		FName      string
		LName      string
		Regions    []string
		Categories []entity.Category
		SubCities  []string
		Cities     []string
	}{

		Inputs: struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{Values: r.PostForm, VErrors: form.ValidationErrors{}, CSRF: token},
		Jobseeker:  *jsh.loggedInUser,
		FName:      FName,
		LName:      LName,
		Regions:    Regions,
		Categories: jobCtgs,
		SubCities:  SubCities,
		Cities:     Cities,
	}
	if r.Method == "GET" {
		err = jsh.tmpl.ExecuteTemplate(w, "jobseeker.profile.edit.layout", profileinfo)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		_ = r.ParseMultipartForm(1024)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		profileinfo.Inputs.Required("fname", "lname", "wrkexpr", "email", "uname", "age", "empstatus", "localname")
		profileinfo.Inputs.MatchesPattern("email", form.EmailRX)
		profileinfo.Inputs.MatchesPattern("phone", form.PhoneRX)

		profileinfo.Inputs.CSRF = token

		// If there are any errors, redisplay the signup form.
		if !profileinfo.Inputs.Valid() {
			jsh.tmpl.ExecuteTemplate(w, "jobseeker.profile.edit.layout", profileinfo)
			return
		}

		if jsh.loggedInUser.Username != r.FormValue("uname") {
			uExists := jsh.jsSrv.UsernameExists(r.FormValue("uname"))
			if uExists {
				profileinfo.Inputs.VErrors.Add("uname", "Username Already Exists")
				jsh.tmpl.ExecuteTemplate(w, "jobseeker.profile.edit.layout", profileinfo)
				return
			}
		}
		if jsh.loggedInUser.Email != r.FormValue("email") {
			eExists := jsh.jsSrv.EmailExists(r.FormValue("email"))
			if eExists {
				profileinfo.Inputs.VErrors.Add("email", "Email Already Exists")
				jsh.tmpl.ExecuteTemplate(w, "jobseeker.profile.edit.layout", profileinfo)
				return
			}
		}

		role, errs := jsh.userRole.RoleByName("JOBSEEKER")

		if len(errs) > 0 {
			profileinfo.Inputs.VErrors.Add("role", "could not assign role to the user")
			jsh.tmpl.ExecuteTemplate(w, "jobseeker.profile.edit.layout", profileinfo)
			return
		}
		addr := &entity.Address{
			Region:    r.FormValue("region"),
			City:      r.FormValue("city"),
			SubCity:   r.FormValue("subcity"),
			LocalName: r.FormValue("localname"),
		}
		intjobcat := r.Form["intjobcat"]
		var categories []entity.Category
		for _, v := range intjobcat {
			fmt.Println(v)
			catidint, _ := strconv.Atoi(v)
			ctg, _ := jsh.ctgSrv.Category(catidint)
			categories = append(categories, ctg)
		}
		var password = ""
		if r.FormValue("newpswd") != "" {
			profileinfo.Inputs.MinLength("newpswd", 8)
			profileinfo.Inputs.PasswordMatches("newpswd", "pswdconfirm")
			err = bcrypt.CompareHashAndPassword([]byte(jsh.loggedInUser.Password), []byte(r.FormValue("currpswd")))
			if err != nil {
				profileinfo.Inputs.VErrors.Add("currpswd", "Password incorrect")
				jsh.tmpl.ExecuteTemplate(w, "jobseeker.profile.edit.layout", profileinfo)
				return
			}
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("newpswd")), 12)
			if err != nil {
				profileinfo.Inputs.VErrors.Add("newpswd", "Password Could not be stored")
				jsh.tmpl.ExecuteTemplate(w, "jobseeker.profile.edit.layout", profileinfo)
				return
			}
			password = string(hashedPassword)
		}
		ageint, err := strconv.Atoi(r.FormValue("age"))
		wrkint, err := strconv.Atoi(r.FormValue("wrkexpr"))

		js := &entity.Jobseeker{
			Address:        []entity.Address{*addr},
			Applications:   jsh.loggedInUser.Applications,
			Categories:     categories,
			Age:            uint(ageint),
			Phone:          r.FormValue("phone"),
			WorkExperience: wrkint,
			Username:       r.FormValue("uname"),
			Fullname:       r.FormValue("fname") + " " + r.FormValue("lname"),
			Email:          r.FormValue("email"),
			Portfolio:      r.FormValue("portf"),
			Gender:         r.FormValue("gender"),
			EmpStatus:      r.FormValue("empstatus"),
		}
		js.ID = jsh.loggedInUser.ID
		js.RoleID = role.ID
		if password == "" {
			js.Password = jsh.loggedInUser.Password
		} else {
			js.Password = password
		}
		// todo process and store user entered profile picture
		propic, fh, err := r.FormFile("propic")
		if err == nil {
			path, err := os.Getwd()
			if err != nil {
				fmt.Printf("Error: %v", err)
				return
			}
			path = filepath.Join(path, "ui", "asset", "jsdata", js.Username, "pp")
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
			js.Profile = fh.Filename
		} else {
			js.Profile = jsh.loggedInUser.Profile
		}
		// todo process and store user entered cv
		cv, fh, err := r.FormFile("cv")
		if err == nil {
			path, err := os.Getwd()
			fmt.Println(path)
			path = filepath.Join(path, "ui", "asset", "jsdata", js.Username, "cv")
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
			js.CV = fh.Filename
		} else {
			js.CV = jsh.loggedInUser.CV
		}
		jsh.loggedInUser, err = jsh.jsSrv.UpdateJobSeeker(js)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/jobseeker/"+jsh.loggedInUser.Username+"/home", http.StatusSeeOther)

	}
}
func (jsh *JobseekerHandler) AppliedJobCategory(w http.ResponseWriter, r *http.Request) {
	ps := r.Context().Value("params").(httprouter.Params)
	if r.Method == "GET" {
		id := ps.ByName("id")
		idint, err := strconv.Atoi(id)
		Ctgs, _ := jsh.ctgSrv.Categories()

		appliedjobcatinfo := struct {
			Categories []entity.Category
			Catid      int
		}{
			Categories: Ctgs,
			Catid:      idint,
		}
		err = jsh.tmpl.ExecuteTemplate(w, "jobseeker.appliedJobs.category.layout", appliedjobcatinfo)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}
func AppGetJob(app entity.Application) (entity.Job, error) {
	fmt.Println(app.JobID)
	job, err := jobSrvc.Job(int(app.JobID))
	if err != nil {
		return job, err
	}
	return job, nil
}
func AppGetCmp(app entity.Application) (entity.Company, error) {
	jb, err := jobSrvc.Job(int(app.JobID))
	var cmp entity.Company
	if err == nil {
		fmt.Println(jb.CompanyID)
		if err != nil {
			return cmp, err
		}
		cmp, err = cmpSrvc.Company(int(jb.CompanyID))
		if err != nil {
			return cmp, err
		}
		return cmp, nil
	}
	return cmp, nil

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
func (uh *JobseekerHandler) loggedIn(r *http.Request) bool {
	if uh.userSess == nil {
		return false
	}
	userSess := uh.userSess
	c, err := r.Cookie(userSess.Uuid)
	if err != nil {
		return false
	}
	ok, err := sess.Valid(c.Value, userSess.SigningKey)
	if !ok || (err != nil) {
		return false
	}
	return true
}

// Authorized checks if a user has proper authority to access a give route
func (uh *JobseekerHandler) Authorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if uh.loggedInUser == nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		roles, errs := uh.jsSrv.UserRoles(uh.loggedInUser)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		for _, role := range roles {
			permitted := permission.HasPermission(r.URL.Path, role.Name, r.Method)
			if !permitted {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}
		if r.Method == http.MethodPost {
			ok, err := rtoken.ValidCSRF(r.FormValue("_csrf"), uh.csrfSignKey)
			if !ok || (err != nil) {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

// Login hanldes the GET/POST /login requests
func (uh *JobseekerHandler) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		loginForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		err := uh.tmpl.ExecuteTemplate(w, "signInUp.layout", loginForm)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		jobCtgs, _ := uh.ctgSrv.Categories()
		Regions := []string{entity.Tigray, entity.Amhara, entity.Sidama, entity.Afar, entity.Somalia, entity.Gambella, entity.Harare, entity.Snnpr, entity.Oromia, entity.Benshangul}
		Cities := []string{entity.Addis, entity.Mekele, entity.Hawassa, entity.Adamma, entity.Gonder}
		SubCities := []string{entity.Gulele, entity.Arada, entity.Yeka, entity.Bole, entity.Cherkos, entity.AddisKetema}
		signUpForm := struct {
			Inputs     form.Input
			Regions    []string
			Cities     []string
			Subcities  []string
			Categories []entity.Category
		}{
			Inputs: struct {
				Values  url.Values
				VErrors form.ValidationErrors
				CSRF    string
			}{Values: r.PostForm, VErrors: form.ValidationErrors{}, CSRF: token},
			Regions:    Regions,
			Cities:     Cities,
			Subcities:  SubCities,
			Categories: jobCtgs,
		}
		usr, errs := uh.jsSrv.JobseekerByEmail(r.FormValue("email"))
		if errs != nil {
			signUpForm.Inputs.VErrors.Add("generic", "Your email address and/or password is wrong")
			w.Header().Set("Location", "/login")
			uh.tmpl.ExecuteTemplate(w, "signInUp.layout", signUpForm)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(r.FormValue("password")))
		if err != nil {
			signUpForm.Inputs.VErrors.Add("generic", "Your email address or password is wrong")
			uh.tmpl.ExecuteTemplate(w, "signInUp.layout", signUpForm)
			return
		}
		uh.loggedInUser = &usr
		claims := rtoken.Claims(usr.Email, uh.userSess.Expires)
		sess.Create(claims, uh.userSess.Uuid, uh.userSess.SigningKey, w)
		newSess, er := uh.sessSrv.StoreSession(uh.userSess)
		if len(er) > 0 {
			signUpForm.Inputs.VErrors.Add("generic", "Failed to store session")
			uh.tmpl.ExecuteTemplate(w, "signInUp.layout", signUpForm)
			return
		}
		uh.userSess = newSess
		http.Redirect(w, r, "/jobseeker/"+usr.Username+"/home", http.StatusSeeOther)
	}
}

// Logout hanldes the POST /logout requests
func (uh *JobseekerHandler) Logout(w http.ResponseWriter, r *http.Request) {
	userSess, _ := r.Context().Value(ctxUserSessionKey).(*entity.Session)
	sess.Remove(userSess.Uuid, w)
	uh.sessSrv.DeleteSession(userSess.Uuid)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Signup hanldes the GET/POST /signup requests
func (uh *JobseekerHandler) Signup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		jobCtgs, _ := uh.ctgSrv.Categories()
		Regions := []string{entity.Tigray, entity.Amhara, entity.Sidama, entity.Afar, entity.Somalia, entity.Gambella, entity.Harare, entity.Snnpr, entity.Oromia, entity.Benshangul}
		Cities := []string{entity.Addis, entity.Mekele, entity.Hawassa, entity.Adamma, entity.Gonder}
		SubCities := []string{entity.Gulele, entity.Arada, entity.Yeka, entity.Bole, entity.Cherkos, entity.AddisKetema}
		signUpForm := struct {
			Inputs     form.Input
			Regions    []string
			Cities     []string
			Subcities  []string
			Categories []entity.Category
		}{
			Inputs: struct {
				Values  url.Values
				VErrors form.ValidationErrors
				CSRF    string
			}{Values: r.PostForm, VErrors: form.ValidationErrors{}, CSRF: token},
			Regions:    Regions,
			Cities:     Cities,
			Subcities:  SubCities,
			Categories: jobCtgs,
		}
		_ = uh.tmpl.ExecuteTemplate(w, "signInUp.layout", signUpForm)

	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		_ = r.ParseMultipartForm(1024)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		jobCtgs, _ := uh.ctgSrv.Categories()
		Regions := []string{entity.Tigray, entity.Amhara, entity.Sidama, entity.Afar, entity.Somalia, entity.Gambella, entity.Harare, entity.Snnpr, entity.Oromia, entity.Benshangul}
		Cities := []string{entity.Addis, entity.Mekele, entity.Hawassa, entity.Adamma, entity.Gonder}
		SubCities := []string{entity.Gulele, entity.Arada, entity.Yeka, entity.Bole, entity.Cherkos, entity.AddisKetema}
		signUpForm := struct {
			Inputs     form.Input
			Regions    []string
			Cities     []string
			Subcities  []string
			Categories []entity.Category
		}{
			Inputs: struct {
				Values  url.Values
				VErrors form.ValidationErrors
				CSRF    string
			}{Values: r.PostForm, VErrors: form.ValidationErrors{}, CSRF: token},
			Regions:    Regions,
			Cities:     Cities,
			Subcities:  SubCities,
			Categories: jobCtgs,
		}
		signUpForm.Inputs.Required("fname", "lname", "email", "pswd", "pswdconfirm", "phone", "uname", "empstatus", "localname")
		signUpForm.Inputs.MatchesPattern("email", form.EmailRX)
		signUpForm.Inputs.MatchesPattern("phone", form.PhoneRX)
		signUpForm.Inputs.MinLength("pswd", 8)
		signUpForm.Inputs.PasswordMatches("pswd", "pswdconfirm")
		signUpForm.Inputs.CSRF = token

		// If there are any errors, redisplay the signup form.
		if !signUpForm.Inputs.Valid() {
			uh.tmpl.ExecuteTemplate(w, "signInUp.layout", signUpForm)
			return
		}

		uExists := uh.jsSrv.UsernameExists(r.FormValue("phone"))
		if uExists {
			signUpForm.Inputs.VErrors.Add("phone", "Username Already Exists")
			uh.tmpl.ExecuteTemplate(w, "signInUp.layout", signUpForm)
			return
		}
		eExists := uh.jsSrv.EmailExists(r.FormValue("email"))
		if eExists {
			signUpForm.Inputs.VErrors.Add("email", "Email Already Exists")
			uh.tmpl.ExecuteTemplate(w, "signInUp.layout", signUpForm)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("pswd")), 12)
		if err != nil {
			signUpForm.Inputs.VErrors.Add("password", "Password Could not be stored")
			uh.tmpl.ExecuteTemplate(w, "signInUp.layout", signUpForm)
			return
		}

		role, errs := uh.userRole.RoleByName("JOBSEEKER")

		if len(errs) > 0 {
			signUpForm.Inputs.VErrors.Add("role", "could not assign role to the user")
			uh.tmpl.ExecuteTemplate(w, "signInUp.layout", signUpForm)
			return
		}
		addr := &entity.Address{
			Region:    r.FormValue("region"),
			City:      r.FormValue("city"),
			SubCity:   r.FormValue("subcity"),
			LocalName: r.FormValue("localname"),
		}
		intjobcat := r.Form["intjobcat"]
		var categories []entity.Category
		for _, v := range intjobcat {
			fmt.Println(v)
			catidint, _ := strconv.Atoi(v)
			ctg, _ := uh.ctgSrv.Category(catidint)
			categories = append(categories, ctg)
		}
		ageint, err := strconv.Atoi(r.FormValue("age"))
		wrkint, err := strconv.Atoi(r.FormValue("wrkexp"))
		js := &entity.Jobseeker{
			Address:        []entity.Address{*addr},
			Applications:   nil,
			Categories:     categories,
			Age:            uint(ageint),
			Phone:          r.FormValue("phone"),
			WorkExperience: wrkint,
			Username:       r.FormValue("uname"),
			Fullname:       r.FormValue("fname") + " " + r.FormValue("lname"),
			Password:       string(hashedPassword),
			Email:          r.FormValue("email"),
			Profile:        r.FormValue("propic"),
			Portfolio:      r.FormValue("portf"),
			CV:             r.FormValue("cv"),
			Gender:         r.FormValue("gender"),
			EmpStatus:      r.FormValue("empstatus"),
		}
		js.RoleID = role.ID
		// todo process and store user entered profile picture
		propic, fh, err := r.FormFile("propic")
		if err == nil {
			path, err := os.Getwd()
			if err != nil {
				fmt.Printf("Error: %v", err)
				return
			}
			path = filepath.Join(path, "ui", "asset", "jsdata", js.Username, "pp")
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
			js.Profile = fh.Filename
		} else {
			js.Profile = "Avatar.ico"
		}
		// todo process and store user entered cv
		cv, fh, err := r.FormFile("cv")
		if err == nil {
			path, err := os.Getwd()
			fmt.Println(path)
			path = filepath.Join(path, "ui", "asset", "jsdata", js.Username, "cv")
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
			js.CV = fh.Filename
		} else {
			js.CV = "sample.cv.txt"
		}
		_, err = uh.jsSrv.StoreJobSeeker(js)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
