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

func TestCompanyHandler_CompanyHome(test *testing.T) {
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
	_ = NewJobseekerHandler(cmpmocksrv, jobmocksrv, tmplfake, jsmocksrv, categmockserv, addmockserv, appmockserv, sessmockserv, &entity.Sessionmock1, rolemockserv, []byte("mysigning key"))
	cmpmockhandler := NewCompanyHandler(tmpl, jsmocksrv, cmpmocksrv, categmockserv, addmockserv, appmockserv, sessmockserv, jobmocksrv, &entity.Sessionmock1, rolemockserv, []byte("mysigning key"))
	cmpmockhandlerfake := NewCompanyHandler(tmplfake, jsmocksrv, cmpmocksrv, categmockserv, addmockserv, appmockserv, sessmockserv, jobmocksrv, &entity.Sessionmock1, rolemockserv, []byte("mysigning key"))

	jsmockhandler.loggedInUser = &entity.Jobseekermock1
	cmpmockhandler.loggedInUser = &entity.Companymock1
	cmpmockhandlerfake.loggedInUser = &entity.Companymock2

	tests := []struct {
		Name   string
		jsmock *CompanyHandler
		method string
		path   string
		wanted string
	}{{Name: "CompanyHomeGET", jsmock: cmpmockhandlerfake, wanted: http.StatusText(http.StatusInternalServerError), method: "GET", path: "/company/{name}/home"},
		{Name: "CompanyHomeGET", jsmock: cmpmockhandler, wanted: http.StatusText(http.StatusOK), method: "GET", path: "/company/{name}/home"},
		{Name: "CompanyHomePOST", jsmock: cmpmockhandlerfake, wanted: http.StatusText(http.StatusOK), method: "POST", path: "/company/{name}/home"},
		{Name: "CompanyHomePOST", jsmock: cmpmockhandler, wanted: http.StatusText(http.StatusOK), method: "POST", path: "/company/{name}/home"},
		{Name: "CompanyPostJobGET", jsmock: cmpmockhandlerfake, wanted: http.StatusText(http.StatusInternalServerError), method: "GET", path: "/company/{name}/home"},
		{Name: "CompanyPostJobGET", jsmock: cmpmockhandler, wanted: http.StatusText(http.StatusOK), method: "GET", path: "/company/{name}/home"},
		{Name: "CompanyPostJobPOST", jsmock: cmpmockhandlerfake, wanted: http.StatusText(http.StatusSeeOther), method: "POST", path: "/company/{name}/home"},
		{Name: "CompanyPostJobPOST", jsmock: cmpmockhandler, wanted: http.StatusText(http.StatusOK), method: "POST", path: "/company/{name}/home"},
		{Name: "CompanyJobsGET", jsmock: cmpmockhandlerfake, wanted: http.StatusText(http.StatusBadRequest), method: "GET", path: "/company/{name}/name"},
		{Name: "CompanyJobsGET", jsmock: cmpmockhandler, wanted: http.StatusText(http.StatusOK), method: "GET", path: "/company/{name}/name"},
		{Name: "CompanyLoginGET", jsmock: cmpmockhandlerfake, wanted: http.StatusText(http.StatusInternalServerError), method: "GET", path: "/login"},
		{Name: "CompanyLoginGET", jsmock: cmpmockhandler, wanted: http.StatusText(http.StatusOK), method: "GET", path: "/login"},
		{Name: "CompanyLoginPOST", jsmock: cmpmockhandlerfake, wanted: http.StatusText(http.StatusOK), method: "POST", path: "/company/login"},
		{Name: "CompanyLoginPOST", jsmock: cmpmockhandler, wanted: http.StatusText(http.StatusOK), method: "POST", path: "/company/login"},
		{Name: "CompanyLogout", jsmock: cmpmockhandlerfake, wanted: http.StatusText(http.StatusOK), method: "GET", path: "/logout/compnay"},
		{Name: "CompanyLogout", jsmock: cmpmockhandler, wanted: http.StatusText(http.StatusOK), method: "GET", path: "/logout/compnay"},
		{Name: "CompanySignupGET", jsmock: cmpmockhandlerfake, wanted: http.StatusText(http.StatusOK), method: "GET", path: "/signup/company"},
		{Name: "CompanySignupGET", jsmock: cmpmockhandler, wanted: http.StatusText(http.StatusOK), method: "GET", path: "/signup/company"},
		{Name: "CompanySignupPOST", jsmock: cmpmockhandlerfake, wanted: http.StatusText(http.StatusOK), method: "POST", path: "/signup/company"},
		{Name: "CompanySignupPOST", jsmock: cmpmockhandler, wanted: http.StatusText(http.StatusOK), method: "POST", path: "/signup/company"}}
	for _, tst := range tests {
		test.Run(tst.Name, func(t *testing.T) {
			httprec := httptest.NewRecorder()

			switch tst.Name {
			case "CompanyHomeGET":
				req := httptest.NewRequest(tst.method, tst.path, nil)
				tst.jsmock.CompanyHome(httprec, req)
				break
			case "CompanyHomePOST":
				req := httptest.NewRequest(tst.method, tst.path, nil)
				tst.jsmock.CompanyHome(httprec, req)
				break
			case "CompanyPostJobGET":
				req := httptest.NewRequest(tst.method, tst.path, nil)
				tst.jsmock.CompanyPostJob(httprec, req)
				break
			case "CompanyPostJobPOST":
				req := httptest.NewRequest(tst.method, tst.path, nil)
				tst.jsmock.CompanyPostJob(httprec, req)
				break
			case "CompanyJobsGET":
				req := httptest.NewRequest(tst.method, tst.path, nil)
				tst.jsmock.CompanyJobs(httprec, req)
				break
			case "CompanyLoginGET":
				req := httptest.NewRequest(tst.method, tst.path, nil)
				tst.jsmock.Login(httprec, req, httprouter.Params{})
				break
			case "CompanyLoginPOST":
				req := httptest.NewRequest(tst.method, tst.path, nil)
				tst.jsmock.Login(httprec, req, httprouter.Params{})
				break
			case "CompanyLogout":
				req := httptest.NewRequest(tst.method, tst.path, nil)
				tst.jsmock.Logout(httprec, req)
				break
			case "CompanySignupGET":
				req := httptest.NewRequest(tst.method, tst.path, nil)
				tst.jsmock.Signup(httprec, req, httprouter.Params{})
				break
			case "CompanySignupPOST":
				req := httptest.NewRequest(tst.method, tst.path, nil)
				tst.jsmock.Signup(httprec, req, nil)
				break
			}
			resp := httprec.Result()
			if http.StatusText(resp.StatusCode) != tst.wanted {
				t.Errorf("wanted %s got %s", tst.wanted, http.StatusText(resp.StatusCode))
			}
		})
	}
}
