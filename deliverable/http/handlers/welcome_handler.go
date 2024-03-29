package handlers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/iJobs/usecases/company"
	"github.com/miruts/iJobs/usecases/jobseeker"
	"github.com/miruts/iJobs/usecases/session"
	"html/template"
	"net/http"
)

// WelcomeHandler represents welcome page handler
type WelcomeHandler struct {
	tmpl    *template.Template
	sessSrv session.SessionService
	jsSrv   jobseeker.JobseekerService
	cmpSrv  company.CompanyService
}

// NewWelcomeHanlder creates new WelcomeHandler
func NewWelcomeHandler(tmpl *template.Template, sessSrv session.SessionService, jsSrv jobseeker.JobseekerService, cmpSrv company.CompanyService) *WelcomeHandler {
	return &WelcomeHandler{tmpl: tmpl, sessSrv: sessSrv, jsSrv: jsSrv, cmpSrv: cmpSrv}
}

/*func NewWelcomeHandler(tmpl *template.Template, ss session.SessionService) *WelcomeHandler{
	return &WelcomeHandler{tmpl:tmpl, sessSrv:ss}
}
*/
func (wh *WelcomeHandler) Welcome(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.URL.Path != "/" {
		wh.tmpl.ExecuteTemplate(w, "error.layout", http.StatusSeeOther)
		return
	}
	err := wh.tmpl.ExecuteTemplate(w, "welcome.layout", nil)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
