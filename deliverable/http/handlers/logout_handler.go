package handlers

import (
	"github.com/miruts/iJobs/usecases/company"
	"github.com/miruts/iJobs/usecases/jobseeker"
	"html/template"
)

// LogoutHandler specifies a logout handler
type LogoutHandler struct {
	tmpl   *template.Template
	jsSrv  jobseeker.JobseekerService
	cmpSrv company.CompanyService
}

// NewLogoutHandler create new LogoutHandler
func NewLogoutHandler(tmpl *template.Template, jsSrv jobseeker.JobseekerService, cmpSrv company.CompanyService) *LogoutHandler {
	return &LogoutHandler{tmpl: tmpl, jsSrv: jsSrv, cmpSrv: cmpSrv}
}
