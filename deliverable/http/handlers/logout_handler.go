package handlers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/iJobs/usecases/company"
	"github.com/miruts/iJobs/usecases/job"
	"github.com/miruts/iJobs/usecases/jobseeker"
	"github.com/miruts/iJobs/usecases/session"
	"github.com/miruts/iJobs/util"
	"html/template"
	"net/http"
)

// LogoutHandler specifies a logout handler
type LogoutHandler struct {
	tmpl    *template.Template
	jsSrv   jobseeker.JobseekerService
	cmpSrv  company.CompanyService
	sessSrv session.SessionService
	ctgSrvc job.CategoryService
}

// NewLogoutHandler create new LogoutHandler
func NewLogoutHandler(tmpl *template.Template, jsSrv jobseeker.JobseekerService, cmpSrv company.CompanyService, sessSrv session.SessionService) *LogoutHandler {
	return &LogoutHandler{tmpl: tmpl, jsSrv: jsSrv, cmpSrv: cmpSrv, sessSrv: sessSrv}
}

// Logout handles get request at /logout
func (lh *LogoutHandler) Logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ok, session := util.Authenticate(lh.sessSrv, r)
	if ok {
		if ok {
			err := util.DestroySession(&w, r)
			if err != nil {
				return
			}
			_, err = lh.sessSrv.DeleteSession(int(session.ID))
			if err != nil {
				return
			}
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
