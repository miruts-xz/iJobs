package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/company"
	"github.com/miruts/iJobs/usecases/job"
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
	ctgSrv  job.CategoryService
}

// NewLoginHandler creates new LoginHandler
func NewLoginHandler(tmpl *template.Template, jsSrv jobseeker.JobseekerService, cmpSrv company.CompanyService, sessSrv session.SessionService, ctgSrv job.CategoryService) *LoginHandler {
	return &LoginHandler{tmpl: tmpl, jsSrv: jsSrv, cmpSrv: cmpSrv, sessSrv: sessSrv, ctgSrv: ctgSrv}
}

/**
Handles GET request to localhost/login
*/
func (lh *LoginHandler) GetLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ok, sess := util.Authenticate(lh.sessSrv, r)
	if ok {
		err := util.DestroySession(&w, r)
		if err != nil {
			return
		}
		_, err = lh.sessSrv.DeleteSession(int(sess.ID))
		if err != nil {
			return
		}
		jobCtgs, err := lh.ctgSrv.Categories()
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
		err = lh.tmpl.ExecuteTemplate(w, "signInUp.layout", registerNeed)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		jobCtgs, err := lh.ctgSrv.Categories()
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
		err = lh.tmpl.ExecuteTemplate(w, "signInUp.layout", registerNeed)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
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
	fmt.Println(jobseeker.Email)
	fmt.Println(company.Email)
	fmt.Println(err1)
	fmt.Println(err2)
	if err2 != nil && err1 == nil {
		// Its jobseeker
		err = bcrypt.CompareHashAndPassword([]byte(jobseeker.Password), []byte(password))
		fmt.Println(err)
		if err == nil {
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
					fmt.Println(err)
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
		err = bcrypt.CompareHashAndPassword([]byte(jobseeker.Password), []byte(password))
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
