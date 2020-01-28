package handlers

import (
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

func TestJobseekerHandler_JobseekerHome(test *testing.T) {
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

	jsmockhandler := NewJobseekerHandler(cmpmocksrv, jobmocksrv, tmpl, jsmocksrv, categmockserv, addmockserv, appmockserv, sessmockserv, &entity.Sessionmock1, rolemockserv, []byte("mysinging key"))
	jsmockhandler.loggedInUser = &entity.Jobseekermock1
	tests := []struct {
		Name   string
		jsmock *JobseekerHandler
		wanted string
	}{{Name: "Status Ok", jsmock: jsmockhandler, wanted: http.StatusText(http.StatusOK)},
		{Name: "Status Error", jsmock: jsmockhandler, wanted: http.StatusText(http.StatusInternalServerError)}}

	for _, tst := range tests {
		test.Run(tst.Name, func(t *testing.T) {
			httprec := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/jobseeker/{username}/home", nil)
			tst.jsmock.JobseekerHome(httprec, req)
			resp := httprec.Result()
			if http.StatusText(resp.StatusCode) != tst.wanted {
				t.Errorf("wanted %s got %s", tst.wanted, http.StatusText(resp.StatusCode))
			}
		})
	}

}
