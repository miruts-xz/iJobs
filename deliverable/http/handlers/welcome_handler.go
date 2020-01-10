package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/iJobs/usecases/company"
	"github.com/miruts/iJobs/usecases/jobseeker"
	"github.com/miruts/iJobs/usecases/session"
	"github.com/miruts/iJobs/util"
	"html/template"
	"net/http"
)

type WelcomeHandler struct {
	tmpl    *template.Template
	sessSrv session.SessionService
	jsSrv   jobseeker.JobseekerService
	cmpSrv  company.CompanyService
}

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
	ok, session := util.Authenticate(wh.sessSrv, r)
	if !ok {
		err := wh.tmpl.ExecuteTemplate(w, "welcome.layout", nil)
		if err != nil {
			fmt.Printf("Login Templating error: %s", err)
			return
		}
		return
	}
	util.DetectUser(&w, r, session, wh.jsSrv, wh.cmpSrv)
}
