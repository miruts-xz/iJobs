package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/iJobs/usecases/company"
	"github.com/miruts/iJobs/usecases/jobseeker"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

// LoginHandler specifies a login handler
type LoginHandler struct {
	tmpl   *template.Template
	jsSrv  jobseeker.JobseekerService
	cmpSrv company.CompanyService
}

// NewLoginHandler creates new LoginHandler
func NewLoginHandler(tmpl *template.Template, jsSrv jobseeker.JobseekerService, cmpSrv company.CompanyService) *LoginHandler {
	return &LoginHandler{tmpl: tmpl, jsSrv: jsSrv, cmpSrv: cmpSrv}
}

/**
Handles GET request to localhost/login
*/
func (lh *LoginHandler) GetLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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
		return
	}
	err = r.ParseMultipartForm(1024)
	if err != nil {
		fmt.Printf("Error Parsing Multipart Data: %s", err)
	}
	email := r.FormValue("email")
	password := r.FormValue("password")
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

}

/**
Handles Request to localhost/authenticate
*/
