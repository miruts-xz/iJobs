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

	userSess     *entity.Session
	loggedInUser *entity.Company
	userRole     role.RoleService
	csrfSignKey  []byte
}

// NewCompanyHandler creates new CompanyHandler
func NewCompanyHandler(tmpl *template.Template, jsSrv jobseeker.JobseekerService, cmpSrv company.CompanyService, ctgSrv job.CategoryService, addrSrv jobseeker.AddressService, appSrv application.IAppService, sessSrv session.SessionService, jobSrv job.JobService, userSess *entity.Session, userRole role.RoleService, csrfSignKey []byte) *CompanyHandler {
	return &CompanyHandler{tmpl: tmpl, jsSrv: jsSrv, cmpSrv: cmpSrv, ctgSrv: ctgSrv, addrSrv: addrSrv, appSrv: appSrv, sessSrv: sessSrv, jobSrv: jobSrv, userSess: userSess, userRole: userRole, csrfSignKey: csrfSignKey}
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

// Authenticated checks if a user is authenticated to access a given route
func (uh *CompanyHandler) Authenticated(next http.Handler) httprouter.Handle {
	fn := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ok := uh.loggedIn(r)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), ctxUserSessionKey, uh.userSess)
		ctx = context.WithValue(r.Context(), "params", ps)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return fn
}

// Authorized checks if a user has proper authority to access a give route
func (uh *CompanyHandler) Authorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if uh.loggedInUser == nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		roles, errs := uh.cmpSrv.UserRoles(uh.loggedInUser)
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

func (ch *CompanyHandler) CompanyHome(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		aid := r.URL.Query().Get("aid")
		accept := r.URL.Query().Get("accept")
		if aid != "" {
			aidint, err := strconv.Atoi(aid)
			application, err := ch.appSrv.Application(aidint)
			switch accept {
			case "true":

				break
			case "false":
				break
			case "further":
				break
			}
		}

		cmpneeds := CompanyHomeNeed{}
		Jobs, err := ch.jobSrv.CompanyJobs(ch.cmpSrv, int(ch.loggedInUser.ID))
		Applications, err := ch.appSrv.ApplicationForCompany(int(ch.loggedInUser.ID))
		cmpneeds.Company = *ch.loggedInUser
		cmpneeds.Jobs = Jobs
		cmpneeds.Applications = Applications

		err = ch.tmpl.ExecuteTemplate(w, "company.home.layout", cmpneeds)
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}
func (ch *CompanyHandler) CompanyPostJob(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(ch.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == "GET" {
		jobCtgs, _ := ch.ctgSrv.Categories()
		postjobform := struct {
			Categories []entity.Category
			Company    entity.Company
			Inputs     form.Input
		}{
			Categories: jobCtgs,
			Company:    *ch.loggedInUser,
			Inputs: struct {
				Values  url.Values
				VErrors form.ValidationErrors
				CSRF    string
			}{Values: nil, VErrors: nil, CSRF: token},
		}
		err := ch.tmpl.ExecuteTemplate(w, "company.postjob.layout", postjobform)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		postjobform := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		postjobform.Required("reqnum", "jbtitle", "jbtime")

		if !postjobform.Valid() {
			ch.tmpl.ExecuteTemplate(w, "company.postjob.layout", postjobform)
		}
		jExists := ch.cmpSrv.JobExists(int(ch.loggedInUser.ID), r.FormValue("jbtitle"))
		if jExists {
			postjobform.VErrors.Add("name", "Name Already Exists")
			ch.tmpl.ExecuteTemplate(w, "company.postjob.layout", postjobform)
			return
		}
		intjobcat := r.Form["ctgs"]
		var categories []entity.Category
		for _, v := range intjobcat {
			catidint, _ := strconv.Atoi(v)
			ctg, _ := ch.ctgSrv.Category(catidint)
			categories = append(categories, ctg)
		}
		reqint, err := strconv.Atoi(r.FormValue("reqnum"))
		salint, err := strconv.Atoi(r.FormValue("salary"))
		deadline, errd := time.Parse("2006-01-02", r.FormValue("deadline"))

		job := entity.Job{
			CompanyID:    ch.loggedInUser.ID,
			Categories:   categories,
			Applications: nil,
			RequiredNum:  uint(reqint),
			Salary:       float64(salint),
			Name:         r.FormValue("jbtitle"),
			Description:  r.FormValue("description"),
			JobTime:      r.FormValue("jbtime"),
		}
		if errd != nil {
			job.Deadline = time.Now().Add(7 * 24 * time.Hour)
		} else {
			job.Deadline = deadline
		}
		updatejb, err := ch.jobSrv.StoreJob(&job)
		ch.loggedInUser.Jobs = append(ch.loggedInUser.Jobs, *updatejb)
		ch.loggedInUser, err = ch.cmpSrv.UpdateCompany(ch.loggedInUser)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/company/"+ch.loggedInUser.CompanyName+"/jobs", http.StatusSeeOther)
	}
}
func (ch *CompanyHandler) CompanyJobs(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		id := request.URL.Query().Get("id")
		if id != "" {
			idint, _ := strconv.Atoi(id)
			_, _ = ch.jobSrv.DeleteJob(idint)
			usr, _ := ch.cmpSrv.CompanyByEmail(ch.loggedInUser.Email)
			ch.loggedInUser = &usr
		}
		companyjobsinfo := struct {
			Company entity.Company
			Jobs    []entity.Job
		}{
			Company: *ch.loggedInUser,
			Jobs:    ch.loggedInUser.Jobs,
		}
		err := ch.tmpl.ExecuteTemplate(writer, "company.jobs.layout", companyjobsinfo)
		if err != nil {
			fmt.Println(err)
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
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
func (uh *CompanyHandler) loggedIn(r *http.Request) bool {
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

// Login hanldes the GET/POST /login requests
func (uh *CompanyHandler) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
			return
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
		usr, errs := uh.cmpSrv.CompanyByEmail(r.FormValue("email"))
		if errs != nil {
			signUpForm.Inputs.VErrors.Add("generic", "Your email address and/or password is wrong")
			uh.tmpl.ExecuteTemplate(w, "signInUp.layout", signUpForm)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(r.FormValue("password")))
		if err == bcrypt.ErrMismatchedHashAndPassword {
			signUpForm.Inputs.VErrors.Add("generic", "Your email address and/or password is wrong")
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
		http.Redirect(w, r, "/company/"+usr.CompanyName+"/home", http.StatusSeeOther)
	}
}

// Logout hanldes the POST /logout requests
func (uh *CompanyHandler) Logout(w http.ResponseWriter, r *http.Request) {
	userSess, _ := r.Context().Value(ctxUserSessionKey).(*entity.Session)
	sess.Remove(userSess.Uuid, w)
	uh.sessSrv.DeleteSession(userSess.Uuid)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Signup hanldes the GET/POST /signup requests
func (uh *CompanyHandler) Signup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		signUpForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		uh.tmpl.ExecuteTemplate(w, "signInUp.layout", signUpForm)
		return
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
		// Validate the form contents
		signUpForm.Inputs.Required("name", "phone", "pswd", "email", "confirm", "shortdesc")
		signUpForm.Inputs.MatchesPattern("email", form.EmailRX)
		signUpForm.Inputs.MatchesPattern("phone", form.PhoneRX)
		signUpForm.Inputs.MinLength("password", 8)
		signUpForm.Inputs.PasswordMatches("password", "confirmpassword")
		signUpForm.Inputs.CSRF = token

		// If there are any errors, redisplay the signup form.
		if !signUpForm.Inputs.Valid() {
			uh.tmpl.ExecuteTemplate(w, "signInUp.layout", signUpForm)
			return
		}

		uExists := uh.cmpSrv.UsernameExists(r.FormValue("username"))
		if uExists {
			signUpForm.Inputs.VErrors.Add("name", "Name Already Exists")
			uh.tmpl.ExecuteTemplate(w, "signInUp.layout", signUpForm.Inputs)
			return
		}
		eExists := uh.jsSrv.EmailExists(r.FormValue("email"))
		if eExists {
			signUpForm.Inputs.VErrors.Add("email", "Email Already Exists")
			uh.tmpl.ExecuteTemplate(w, "signInUp.layout", signUpForm.Inputs)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("pswd")), 12)
		if err != nil {
			signUpForm.Inputs.VErrors.Add("pswd", "Password Could not be stored")
			uh.tmpl.ExecuteTemplate(w, "signInUp.layout", signUpForm)
			return
		}

		role, errs := uh.userRole.RoleByName("COMPANY")

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
		cmp := &entity.Company{
			Address:     []entity.Address{*addr},
			CompanyName: r.FormValue("name"),
			Password:    string(hashedPassword),
			Email:       r.FormValue("email"),
			Phone:       r.FormValue("phone"),
			Logo:        r.FormValue("logo"),
			ShortDesc:   r.FormValue("shortDescr"),
			DetailInfo:  r.FormValue("detailDescr"),
		}
		cmp.RoleID = role.ID
		// todo process and store user entered profile picture
		propic, fh, err := r.FormFile("logo")
		if err == nil {
			path, err := os.Getwd()
			if err != nil {
				fmt.Printf("Error: %v", err)
				return
			}
			path = filepath.Join(path, "ui", "asset", "cmpdata", cmp.CompanyName, "logo")
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
			cmp.Logo = fh.Filename
		} else {
			cmp.Logo = "company.ico"
		}
		_, err = uh.cmpSrv.StoreCompany(cmp)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
