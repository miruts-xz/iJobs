package handlers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/iJobs/entity"
	repository4 "github.com/miruts/iJobs/role/repository"
	service4 "github.com/miruts/iJobs/role/service"
	repository6 "github.com/miruts/iJobs/usecases/application/repository"
	appserv "github.com/miruts/iJobs/usecases/application/service"
	repository3 "github.com/miruts/iJobs/usecases/company/repository"
	service3 "github.com/miruts/iJobs/usecases/company/service"
	repository2 "github.com/miruts/iJobs/usecases/job/repository"
	service2 "github.com/miruts/iJobs/usecases/job/service"
	"github.com/miruts/iJobs/usecases/jobseeker/repository"
	"github.com/miruts/iJobs/usecases/jobseeker/service"
	repository5 "github.com/miruts/iJobs/usecases/session/repository"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJobseekerHandler_CompanyHome(test *testing.T) {
	jsmockrepo := repository.NewJobseekerMockRepository()
	jobmockrepo := repository2.NewJobMockRepository()
	cmpmockrepo := repository3.NewCompanyMockRepository()
	categmockrepo := repository2.NewCategoryMockRepository()
	rolemockrepo := repository4.NewRoleMockRepository()
	addmockrepo := repository.NewAddressMockRepository()
	sessmockrepo := repository5.NewSessionMockRepository()
	appmockrepo := repository6.NewApplicationMockRepository()

	sessmockserv := service4.NewSessionService(sessmockrepo)
	addmockserv := service.NewAddressServiceImpl(addmockrepo)
	rolemockserv := service4.NewRoleService(rolemockrepo)
	categsrv := service2.NewCategoryServiceImpl(categmockrepo)
	jobmocksrv := service2.NewJobServices(jobmockrepo, categsrv)
	categmockserv := service2.NewCategoryServiceImpl(categmockrepo)

	cmpmocksrv := service3.NewCompanyServiceImpl(cmpmockrepo)
	jsmocksrv := service.NewJobseekerServiceImpl(jsmockrepo, jobmocksrv)
	appmockserv := appserv.NewAppService(appmockrepo, jsmocksrv, jobmocksrv, cmpmockrepo)

	funcMaps := template.FuncMap{"cmp": JobCmp, "appGetJob": AppGetJob, "appGetJs": AppJs, "appGetJobCatId": AppGetJobCatId, "appGetCmp": AppGetCmp, "appGetLoc": AppGetLocation}

	tmpl = template.Must(template.New("index").Funcs(funcMaps).ParseGlob("../../../ui/template/*.html"))
	tmplfake := template.Must(template.New("fake").Funcs(funcMaps).ParseGlob("*.html"))

	jsmockhandler := NewJobseekerHandler(cmpmocksrv, jobmocksrv, tmpl, jsmocksrv, categmockserv, addmockserv, appmockserv, sessmockserv, &entity.Sessionmock1, rolemockserv, []byte("mysigning key"))
	jsmockhandlerfake := NewJobseekerHandler(cmpmocksrv, jobmocksrv, tmplfake, jsmocksrv, categmockserv, addmockserv, appmockserv, sessmockserv, &entity.Sessionmock1, rolemockserv, []byte("mysigning key"))
	cmpmockhandler := NewCompanyHandler(tmpl, jsmocksrv, cmpmocksrv, categmockserv, addmockserv, appmockserv, sessmockserv, jobmocksrv, &entity.Sessionmock1, rolemockserv, []byte("mysigning key"))
	cmpmockhandlerfake := NewCompanyHandler(tmplfake, jsmocksrv, cmpmocksrv, categmockserv, addmockserv, appmockserv, sessmockserv, jobmocksrv, &entity.Sessionmock1, rolemockserv, []byte("mysigning key"))

	jsmockhandler.loggedInUser = &entity.Jobseekermock1
	cmpmockhandler.loggedInUser = &entity.Companymock1
	cmpmockhandlerfake.loggedInUser = &entity.Companymock2
	jsmockhandlerfake.loggedInUser = &entity.Jobseekermock2

	tests := []struct {
		Name   string
		jsmock *JobseekerHandler
		method string
		path   string
		wanted string
	}{{Name: "JobseekerApplyGET", jsmock: jsmockhandlerfake, wanted: http.StatusText(http.StatusSeeOther), method: "GET", path: "/jobseeker/{name}/apply"},
		{Name: "JobseekerApplyGET", jsmock: jsmockhandler, wanted: http.StatusText(http.StatusSeeOther), method: "GET", path: "/jobseeker/{name}/apply"},
		{Name: "JobseekerHomeGET", jsmock: jsmockhandlerfake, wanted: http.StatusText(http.StatusInternalServerError), method: "GET", path: "/jobseeker/{name}/home"},
		{Name: "JobseekerHomeGET", jsmock: jsmockhandler, wanted: http.StatusText(http.StatusOK), method: "GET", path: "/jobseeker/{name}/home"},
		{Name: "JobseekerHomePOST", jsmock: jsmockhandlerfake, wanted: http.StatusText(http.StatusBadGateway), method: "POST", path: "/jobseeker/{name}/home"},
		{Name: "JobseekerHomePOST", jsmock: jsmockhandler, wanted: http.StatusText(http.StatusBadGateway), method: "POST", path: "/jobseeker/{name}/home"},
		{Name: "JobseekerAppliedJobsGET", jsmock: jsmockhandlerfake, wanted: http.StatusText(http.StatusInternalServerError), method: "GET", path: "/jobseeker/{name}/appliedjbos"},
		{Name: "JobseekerAppliedJobsGET", jsmock: jsmockhandler, wanted: http.StatusText(http.StatusOK), method: "GET", path: "/jobseeker/{name}/appliedjobs"},
		{Name: "JobseekerProfileGET", jsmock: jsmockhandlerfake, wanted: http.StatusText(http.StatusInternalServerError), method: "GET", path: "/jobseeker/{name}/profile"},
		{Name: "JobseekerProfileGET", jsmock: jsmockhandler, wanted: http.StatusText(http.StatusOK), method: "GET", path: "/jobseeker/{name}/profile"},

		{Name: "JobseekerProfilePOST", jsmock: jsmockhandlerfake, wanted: http.StatusText(http.StatusOK), method: "POST", path: "/jobseeker/{name}/profile/edit"},
		{Name: "JobseekerProfilePOST", jsmock: jsmockhandler, wanted: http.StatusText(http.StatusOK), method: "POST", path: "/jobseeker/{name}/profile/edit"},
		{Name: "JobseekerLogout", jsmock: jsmockhandler, wanted: http.StatusText(http.StatusOK), method: "GET", path: "/logout/jobseeker"},
		{Name: "JobseekerSignupGET", jsmock: jsmockhandlerfake, wanted: http.StatusText(http.StatusOK), method: "GET", path: "/signup/company"},
		{Name: "JobseekerSignupGET", jsmock: jsmockhandler, wanted: http.StatusText(http.StatusOK), method: "GET", path: "/signup/company"},
		{Name: "JobseekerSignupPOST", jsmock: jsmockhandlerfake, wanted: http.StatusText(http.StatusOK), method: "POST", path: "/signup/company"},
		{Name: "JobseekerSignupPOST", jsmock: jsmockhandler, wanted: http.StatusText(http.StatusOK), method: "POST", path: "/signup/company"}}
	for _, tst := range tests {
		test.Run(tst.Name, func(t *testing.T) {
			httprec := httptest.NewRecorder()

			switch tst.Name {
			case "JobseekerApplyGET":
				req := httptest.NewRequest(tst.method, tst.path, nil)
				tst.jsmock.JobseekerApply(httprec, req)
				break
			case "JobseekerHomeGET":
				req := httptest.NewRequest(tst.method, tst.path, nil)
				tst.jsmock.JobseekerHome(httprec, req)
				break
			case "JobseekerHomePOST":
				req := httptest.NewRequest(tst.method, tst.path, nil)
				tst.jsmock.JobseekerHome(httprec, req)
				break
			case "JobseekerAppliedJobsGET":
				req := httptest.NewRequest(tst.method, tst.path, nil)
				tst.jsmock.JobseekerAppliedJobs(httprec, req)
				break
			case "JobseekerProfileGET":
				req := httptest.NewRequest(tst.method, tst.path, nil)
				tst.jsmock.JobseekerProfile(httprec, req)
				break
			case "JobseekerProfilePOST":
				req := httptest.NewRequest(tst.method, tst.path, nil)
				tst.jsmock.JobseekerProfile(httprec, req)
				break
			case "JobseekerLogout":
				req := httptest.NewRequest(tst.method, tst.path, nil)
				tst.jsmock.Logout(httprec, req)
				break
			case "JobseekerSignupGET":
				req := httptest.NewRequest(tst.method, tst.path, nil)
				tst.jsmock.Signup(httprec, req, httprouter.Params{})
				break
			case "JobseekerSignupPOST":
				req := httptest.NewRequest(tst.method, tst.path, nil)
				tst.jsmock.Signup(httprec, req, httprouter.Params{})
				break
			}
			resp := httprec.Result()
			if http.StatusText(resp.StatusCode) != tst.wanted {
				t.Errorf("wanted %s got %s", tst.wanted, http.StatusText(resp.StatusCode))
			}
		})
	}
}
