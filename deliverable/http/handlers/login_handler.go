package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/company"
	"github.com/miruts/iJobs/usecases/jobseeker"
	"github.com/miruts/iJobs/usecases/session"
	"github.com/miruts/iJobs/util"
	"github.com/satori/uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

// LoginHandler specifies a login handler
type LoginHandler struct {
	tmpl    *template.Template
	jsSrv   jobseeker.JobseekerService
	cmpSrv  company.CompanyService
	sessSrv session.SessionService
}

// NewLoginHandler creates new LoginHandler
func NewLoginHandler(tmpl *template.Template, jsSrv jobseeker.JobseekerService, cmpSrv company.CompanyService, ss session.SessionService) *LoginHandler {
	return &LoginHandler{tmpl: tmpl, jsSrv: jsSrv, cmpSrv: cmpSrv, sessSrv: ss}
}

/**
Handles GET request to localhost/login
*/
func (lh *LoginHandler) GetLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess, err := util.Authenticate(lh.sessSrv, r)
	if err != nil {
		err := lh.tmpl.ExecuteTemplate(w, "login.layout", nil)
		if err != nil {
			fmt.Printf("Login Templating error: %s", err)
			return
		}
		return
	}
	util.DetectUser(&w, r, sess, lh.jsSrv, lh.cmpSrv)
}

/**
PostLogin Handles POST request to localhost/login

form field mapping
email - email
password - password
*/
func (lh *LoginHandler) PostLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Error Parsing Login Data: %s", err)
	}
	err = r.ParseMultipartForm(1024)
	if err != nil {
		fmt.Printf("Error Parsing Multipart Data: %s", err)
	}
	email := r.FormValue("email")
	fmt.Println(email)
	password := r.FormValue("password")
	fmt.Println(password)
	jobseeker, err1 := lh.jsSrv.JobseekerByEmail(email)
	company, err2 := lh.cmpSrv.CompanyByEmail(email)
	fmt.Println(jobseeker)
	fmt.Println(company)
	if err2 != nil && err1 == nil {
		// Its jobseeker
		err = bcrypt.CompareHashAndPassword([]byte(password), []byte(jobseeker.Password))
		if jobseeker.Password == password {
			sess := entity.Session{}
			uuid, err := uuid.NewV4()
			if err != nil {

			}
			fmt.Printf("I'm IN")
			sess.Uuid = uuid.String()
			sess.UserID = jobseeker.ID
			sess.Email = jobseeker.Email
			_, err = lh.sessSrv.StoreSession(&sess)
			if err == nil {
				err = util.CreateSession(&w, &sess)
				if err != nil {
					return
				}
			}
			http.Redirect(w, r, "/jobseeker/home", http.StatusSeeOther)
		} else {
			fmt.Printf("I'm IN below first")

			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	} else if err1 != nil && err2 == nil {
		fmt.Printf("I'm IN Second")

		// Its Company
		err = bcrypt.CompareHashAndPassword([]byte(password), []byte(company.Password))
		if err == nil {
			sess := entity.Session{}
			uuid, err := uuid.NewV4()
			if err != nil {

			}
			sess.Uuid = uuid.String()
			sess.UserID = company.ID
			sess.Email = company.Email
			err = util.CreateSession(&w, &sess)
			if err != nil {
				return
			}
			_, err = lh.sessSrv.StoreSession(&sess)
			if err != nil {
				fmt.Printf("storing session failed: %s", err)
			}
			http.Redirect(w, r, "/company/home", http.StatusSeeOther)
		} else {
			fmt.Printf("I'm IN Login Redirect")

			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}

	} else {
		fmt.Printf("I'm also in Login Redirect")

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
