package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/iJobs/deliverable/http/api"
	"github.com/miruts/iJobs/deliverable/http/handlers"
	"github.com/miruts/iJobs/entity"
	apprepo "github.com/miruts/iJobs/usecases/application/repository"
	appsrv "github.com/miruts/iJobs/usecases/application/service"
	cmprepo "github.com/miruts/iJobs/usecases/company/repository"
	cmpsrv "github.com/miruts/iJobs/usecases/company/service"
	jobrepo "github.com/miruts/iJobs/usecases/job/repository"
	jobsrv "github.com/miruts/iJobs/usecases/job/service"
	jsrepo "github.com/miruts/iJobs/usecases/jobseeker/repository"
	jssrv "github.com/miruts/iJobs/usecases/jobseeker/service"
	"github.com/miruts/iJobs/usecases/session/repository"
	"github.com/miruts/iJobs/usecases/session/service"

	apijobhandler "github.com/miruts/iJobs/deliverable/http/api"

	"html/template"
	"net/http"
)

const (
	domain    = "localhost"
	apiDomain = "api." + domain
)

var gormDB *gorm.DB
var err error
var errs error
var tmpl = template.Must(template.New("index").Funcs(funcMaps).ParseGlob("ui/template/*.html"))
var pqconnjs, pqconncmp *sql.DB

var funcMaps = template.FuncMap{"cmp": handlers.JobCmp, "appGetJob": handlers.AppJob, "appGetJs": handlers.AppJs, "appGetJobCatId": handlers.AppGetJobCatId, "appGetCmpLogo": handlers.AppGetCmpLogo, "appGetJobName": handlers.AppGetJobsName, "appGetCmpName": handlers.AppGetCmpName, "appGetLoc": handlers.AppGetLocation}

func init() {
	// Template

	//Company database connection
	// //pqconncmp, err = sql.Open("postgres", "user=company password=company database=ijobs sslmode=disable")
	// defer pqconncmp.Close()
	// //Jobseeker database connection
	// pqconnjs, err = sql.Open("postgres", "user=jobseeker password=jobseeker database=ijobs sslmode=disable")
	// defer pqconnjs.Close()
}

// CreateTables creates gorm tables provided entity structs
func CreateTables(db *gorm.DB) {
	errs := db.CreateTable(&entity.Session{}, &entity.Address{}, &entity.Application{}, &entity.Category{}, &entity.Job{}, &entity.Company{}, entity.Jobseeker{}).GetErrors()
	if len(errs) > 0 {
		fmt.Println(errs[0])
		return
	}
}
func main() {
	// Gorm Database Connection
	gormDB, err = gorm.Open("postgres", "user=postgres dbname=ijobs_gorm_db password=tsedekeme sslmode=disable")
	if errs != nil {
		fmt.Println(err)
		return
	}
	gormDB.AutoMigrate(&entity.Category{})
	// Create Gorm Tables
	// Run Once
	//CreateTables(gormDB)

	// Data Repositories
	applicationRepo := apprepo.NewAppGormRepositoryImpl(gormDB)
	companyRepo := cmprepo.NewCompanyGormRepositoryImpl(gormDB)
	jobRepo := jobrepo.NewJobGormRepositoryImpl(gormDB)
	categoryRepo := jobrepo.NewCategoryGormRepositoryImpl(gormDB)
	jobseekerRepo := jsrepo.NewJobseekerGormRepositoryImpl(gormDB)
	addressRepo := jsrepo.NewAddressGormRepositoryImpl(gormDB)
	sessionRepo := repository.NewSessionGormRepositoryImpl(gormDB)

	// Services
	companySrv := cmpsrv.NewCompanyServiceImpl(companyRepo)

	categorySrv := jobsrv.NewCategoryServiceImpl(categoryRepo)
	jobSrv := jobsrv.NewJobServices(jobRepo, categorySrv)
	jobseekerSrv := jssrv.NewJobseekerServiceImpl(jobseekerRepo, jobSrv)
	addressSrv := jssrv.NewAddressServiceImpl(addressRepo)
	sessionSrv := service.NewSessionServiceImpl(sessionRepo)
	applicationSrv := appsrv.NewAppService(applicationRepo, jobseekerSrv, jobSrv, companySrv)
	// Handlers
	loginHandler := handlers.NewLoginHandler(tmpl, jobseekerSrv, companySrv, sessionSrv, categorySrv)
	logoutHandler := handlers.NewLogoutHandler(tmpl, jobseekerSrv, companySrv, sessionSrv)
	welcomeHandler := handlers.NewWelcomeHandler(tmpl, sessionSrv, jobseekerSrv, companySrv)
	jobseekerHandler := handlers.NewJobseekerHandler(tmpl, jobseekerSrv, categorySrv, addressSrv, applicationSrv, sessionSrv, jobSrv, companySrv)
	jobseekerAPIHandler := api.NewJobseekerHandler(jobseekerSrv)
	companyHandler := handlers.NewCompanyHandler(tmpl, jobseekerSrv, companySrv, categorySrv, addressSrv, applicationSrv, sessionSrv, jobSrv)
	//logoutHandler := handlers.NewLogoutHandler(tmpl, jobseekerSrv, companySrv, sessionSrv)
	//go util.ClearExpiredSessions(sessionSrv)

	//RESTApi Handlers
	apiJobHandler := apijobhandler.NewJobApiHandler(jobSrv)
	apiJobSkHandler := apijobhandler.NewJobseekerHandler(jobseekerSrv)

	//File Server
	//fs := http.FileServer(http.Dir("ui/asset"))
	router := httprouter.New()

	// Welcome SignIn/Up path registration
	router.GET("/", welcomeHandler.Welcome)
	router.GET("/signout", logoutHandler.Logout)
	router.GET("/login", loginHandler.GetLogin)
	router.POST("/login", loginHandler.PostLogin)
	router.GET("/logout", logoutHandler.Logout)
	router.POST("/signup/jobseeker", jobseekerHandler.JobseekerRegister)
	router.POST("/signup/company", companyHandler.CompanyRegister)

	router.GET("/company/:username/postjob", companyHandler.CompanyPostJob)
	router.POST("/company/:username/postjob", companyHandler.CompanyPostJob)

	router.GET("/company/:username", companyHandler.CompanyHome)

	// Jobseeker path registration
	router.GET("/jobseeker/:username", jobseekerHandler.JobseekerHome)
	router.POST("/jobseeker/:username", jobseekerHandler.JobseekerHome)
	router.GET("/jobseeker/:username/apply/:id", jobseekerHandler.JobseekerApply)
	router.GET("/jobseeker/:username/profile", jobseekerHandler.JobseekerProfile)
	router.GET("/jobseeker/:username/profile/edit", jobseekerHandler.ProfileEdit)
	router.POST("/jobseeker/:username/profile/edit", jobseekerHandler.ProfileEdit)
	router.GET("/jobseeker/:username/appliedjobs", jobseekerHandler.JobseekerAppliedJobs)
	router.GET("/jobseeker/:username/appliedjobs/:id", jobseekerHandler.JobseekerAppliedJobs)
	router.GET("/api/jobseekers", jobseekerAPIHandler.Jobseekers)

	//REST Api registration
	//Job Api Handlers
	router.GET("/api/jobs", apiJobHandler.Jobs)
	router.GET("/api/jobs/:id", apiJobHandler.Job)
	router.POST("/api/jobs/", apiJobHandler.AddJob)
	router.PUT("/api/jobs/:id", apiJobHandler.UpdateJob)
	router.DELETE("/api/jobs/:id", apiJobHandler.DeleteJob)

	//JobSeeker Api Handler

	router.GET("/api/jobseeker", apiJobSkHandler.Jobseekers)
	router.GET("/api/jobseekers/:id", apiJobSkHandler.Jobseeker)
	router.POST("/api/jobseekers/", apiJobSkHandler.AddJobseeker)
	router.PUT("/api/jobseekers/:id", apiJobSkHandler.UpdateJobseeker)
	router.DELETE("/api/jobseekers/:id", apiJobSkHandler.DeleteJobseeker)

	// Static file registration
	router.ServeFiles("/assets/*filepath", http.Dir("ui/asset"))
	// Start Serving
	err := http.ListenAndServe(":8181", router)
	if err != nil {
		fmt.Printf("server failed: %s", err)
	}
}
