package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/iJobs/usecases/company"
	"github.com/miruts/iJobs/usecases/jobseeker"
	"github.com/miruts/iJobs/usecases/session"
	"html/template"
	"net/http"
	"time"
)

// LogoutHandler specifies a logout handler
type LogoutHandler struct {
	tmpl    *template.Template
	jsSrv   jobseeker.JobseekerService
	cmpSrv  company.CompanyService
	sessSrv session.SessionService
}

// NewLogoutHandler create new LogoutHandler
func NewLogoutHandler(tmpl *template.Template, jsSrv jobseeker.JobseekerService, cmpSrv company.CompanyService, sessSrv session.SessionService) *LogoutHandler {
	return &LogoutHandler{tmpl: tmpl, jsSrv: jsSrv, cmpSrv: cmpSrv, sessSrv: sessSrv}
}

func (l *LogoutHandler) Logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cookie, _ := r.Cookie("_cookie")
	cookie = &http.Cookie{
		Name:     "_cookie",
		Value:    "",
		Path:     "/",
		MaxAge:   0,
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	sessId, _ := l.sessSrv.SessionByValue(cookie.Value)
	_, errors := l.sessSrv.DeleteSession(int(sessId.ID))
	if errors != nil {
		fmt.Println(errors)
	}
}
