package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/iJobs/deliverable/http/handlers"
	"github.com/miruts/iJobs/entity"
	cmprepo "github.com/miruts/iJobs/usecases/company/repository"
	cmpsrv "github.com/miruts/iJobs/usecases/company/service"
	jobrepo "github.com/miruts/iJobs/usecases/job/repository"
	jobsrv "github.com/miruts/iJobs/usecases/job/service"
	jsrepo "github.com/miruts/iJobs/usecases/jobseeker/repository"
	jssrv "github.com/miruts/iJobs/usecases/jobseeker/service"
	"github.com/miruts/iJobs/usecases/session/repository"
	"github.com/miruts/iJobs/usecases/session/service"
	"html/template"
	"net/http"
)

var gormDB *gorm.DB
var err error
var errs error
var tmpl *template.Template
var pqconnjs, pqconncmp *sql.DB

func init() {
	// Template
	tmpl = template.Must(template.ParseGlob("ui/template/*.html"))
	//Company database connection
	pqconncmp, err = sql.Open("postgres", "user=company password=company database=ijobs sslmode=disable")

	//Jobseeker database connection
	pqconnjs, err = sql.Open("postgres", "user=jobseeker password=jobseeker database=ijobs sslmode=disable")

}
func CreateTables(db *gorm.DB) {
	errs := db.CreateTable(&entity.Session{}, &entity.Address{}, &entity.Category{}, &entity.Application{}, &entity.Job{}, &entity.Company{}, entity.Jobseeker{}).GetErrors()
	if len(errs) > 0 {
		fmt.Println(errs[0])
		return
	}
}
func main() {
	// Gorm Database Connection
	gormDB, err = gorm.Open("postgres", "user=postgres dbname=ijobs_gorm_db password=postgres sslmode=disable")
	if errs != nil {
		fmt.Println(err)
		return
	}
	defer gormDB.Close()

	// Create Gorm Tables
	// Run Once

	//CreateTables(gormDB)

	// Data Repositories
	//applicationRepo := apprepo.NewAppGormRepo(gormDB)
	companyRepo := cmprepo.NewCompanyGormRepositoryImpl(gormDB)
	jobRepo := jobrepo.NewJobRepository(pqconnjs)
	//categoryRepo := jobrepo.NewCategoryGormRepositoryImpl(gormDB)
	jobseekerRepo := jsrepo.NewJobseekerGormRepositoryImpl(gormDB)
	//addressRepo := jsrepo.NewAddressGormRepositoryImpl(gormDB)
	sessionRepo := repository.NewSessionGormRepositoryImpl(gormDB)
	// Services
	//applicationSrv := appsrv.NewAppservice(applicationRepo)
	companySrv := cmpsrv.NewCompanyServiceImpl(companyRepo)
	jobSrv := jobsrv.NewJobService(jobRepo)
	//categorySrv := jobsrv.NewCategoryServiceImpl(categoryRepo)
	jobseekerSrv := jssrv.NewJobseekerServiceImpl(jobseekerRepo, jobSrv)
	//addressSrv := jssrv.NewAddressServiceImpl(addressRepo)
	sessionSrv := service.NewSessionServiceImpl(sessionRepo)

	// Handlers
	loginHandler := handlers.NewLoginHandler(tmpl, jobseekerSrv, companySrv, sessionSrv)
	welcomeHandler := handlers.NewWelcomeHandler(tmpl, sessionSrv, jobseekerSrv, companySrv)
	//go util.ClearExpiredSessions(sessionSrv)

	//File Server
	//fs := http.FileServer(http.Dir("ui/asset"))
	router := httprouter.New()

	// Login path registration

	router.GET("/", welcomeHandler.Welcome)
	router.GET("/login", loginHandler.GetLogin)
	router.POST("/login", loginHandler.PostLogin)
	router.GET("/jobseeker/home")
	router.ServeFiles("/assets/*filepath", http.Dir("ui/asset"))
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Printf("server failed: %s", err)
	}
}
