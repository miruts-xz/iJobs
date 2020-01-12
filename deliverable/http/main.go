package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
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

	apiHandler "github.com/miruts/iJobs/deliverable/http/api"

	"html/template"
	"net/http"
)

var gormDB *gorm.DB
var err error
var errs error
var tmpl = template.Must(template.New("index").Funcs(funcMaps).ParseGlob("../../ui/template/*.html"))
var pqconnjs, pqconncmp *sql.DB

var funcMaps = template.FuncMap{"appGetJobCatId": handlers.AppGetJobCatId, "appGetCmpLogo": handlers.AppGetCmpLogo, "appGetJobName": handlers.AppGetJobsName, "appGetCmpName": handlers.AppGetCmpName, "appGetLoc": handlers.AppGetLocation}

func init() {
	// Template

	//Company database connection
	// //pqconncmp, err = sql.Open("postgres", "user=company password=company database=ijobs sslmode=disable")
	// defer pqconncmp.Close()
	// //Jobseeker database connection
	// pqconnjs, err = sql.Open("postgres", "user=jobseeker password=jobseeker database=ijobs sslmode=disable")
	// defer pqconnjs.Close()
}

func CreateTables(db *gorm.DB) {
	errs := db.CreateTable(&entity.Session{}, &entity.Address{}, &entity.Application{}, &entity.Category{}, &entity.Job{}, &entity.Company{}, entity.Jobseeker{}).GetErrors()
	if len(errs) > 0 {
		fmt.Println(errs[0])
		return
	}
}
func main() {
	// Gorm Database Connection
	gormDB, err = gorm.Open("postgres", "user=postgres dbname=ijobs_gorm_db password=akuadane sslmode=disable")
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
	applicationSrv := appsrv.NewAppService(applicationRepo, jobseekerSrv, jobSrv)
	// Handlers
	loginHandler := handlers.NewLoginHandler(tmpl, jobseekerSrv, companySrv, sessionSrv, categorySrv)
	welcomeHandler := handlers.NewWelcomeHandler(tmpl, sessionSrv, jobseekerSrv, companySrv)
	jobseekerHandler := handlers.NewJobseekerHandler(tmpl, jobseekerSrv, categorySrv, addressSrv, applicationSrv, sessionSrv, jobSrv, companySrv)
	//go util.ClearExpiredSessions(sessionSrv)

	//RESTApi Handlers
	apiJobHandler := apiHandler.NewJobApiHandler(jobSrv)
	apiJobSkHandler := apiHandler.NewJobseekerHandler(jobseekerSrv)
	apiAppHandler := apiHandler.NewAppApiHandler(applicationSrv)
	apiCmpHandler := apiHandler.NewCompanyHandler(companySrv)

	//File Server
	//fs := http.FileServer(http.Dir("ui/asset"))
	router := httprouter.New()

	// Welcome SignIn/Up path registration
	router.GET("/", welcomeHandler.Welcome)
	router.GET("/login", loginHandler.GetLogin)
	router.POST("/login", loginHandler.PostLogin)
	router.POST("/signup/jobseeker", jobseekerHandler.JobseekerRegister)

	// Jobseeker path registration
	router.GET("/jobseeker/:username", jobseekerHandler.JobseekerHome)
	router.POST("/jobseeker/:username", jobseekerHandler.JobseekerHome)
	router.GET("/jobseeker/:username/profile", jobseekerHandler.JobseekerProfile)
	router.GET("/jobseeker/:username/profile/edit", jobseekerHandler.ProfileEdit)
	router.POST("/jobseeker/:username/profile/edit", jobseekerHandler.ProfileEdit)
	router.GET("/jobseeker/:username/appliedjobs", jobseekerHandler.JobseekerAppliedJobs)
	router.GET("/jobseeker/:username/appliedjobs/:id", jobseekerHandler.JobseekerAppliedJobs)

	//REST Api registration

	//Job Api Handlers
	router.GET("/api/job", apiJobHandler.Jobs)
	router.GET("/api/job/:id", apiJobHandler.Job)
	router.POST("/api/job/", apiJobHandler.AddJob)
	router.PUT("/api/job/:id", apiJobHandler.UpdateJob)
	router.DELETE("/api/job/:id", apiJobHandler.DeleteJob)

	//JobSeeker Api Handler
	router.GET("/api/jobseeker", apiJobSkHandler.Jobseeker)
	router.GET("/api/jobseeker/:id", apiJobSkHandler.Jobseekers)
	router.POST("/api/jobseeker/", apiJobSkHandler.AddJobseeker)
	router.PUT("/api/jobseeker/:id", apiJobSkHandler.UpdateJobseeker)
	router.DELETE("/api/jobseeker/:id", apiJobSkHandler.DeleteJobseeker)

	//Company Api Handler
	router.GET("/api/company", apiCmpHandler.Companies)
	router.GET("/api/company/:id", apiCmpHandler.Company)
	router.POST("/api/company/", apiCmpHandler.AddCompany)
	router.PUT("/api/company/:id", apiCmpHandler.UpdateCompany)
	router.DELETE("/api/company/:id", apiCmpHandler.DeleteCompany)

	//Application Api Handler
	//router.GET("/api/application/job/:jobId", apiAppHandler.ApplicationsOnJob)
	router.GET("/api/application/:id", apiAppHandler.Application)
	router.POST("/api/application/", apiAppHandler.AddApplication)
	router.DELETE("/api/application/:id", apiAppHandler.DeleteApp)

	// Static file registration
	router.ServeFiles("/assets/*filepath", http.Dir("../../ui/asset"))
	// Start Serving
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Printf("server failed: %s", err)
	}
}
