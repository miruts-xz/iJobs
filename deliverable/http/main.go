package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/iJobs/deliverable/http/handlers"
	"github.com/miruts/iJobs/entity"
	repository2 "github.com/miruts/iJobs/role/repository"
	service2 "github.com/miruts/iJobs/role/service"
	"github.com/miruts/iJobs/security/rndtoken"
	apprepo "github.com/miruts/iJobs/usecases/application/repository"
	appsrv "github.com/miruts/iJobs/usecases/application/service"
	cmprepo "github.com/miruts/iJobs/usecases/company/repository"
	cmpsrv "github.com/miruts/iJobs/usecases/company/service"
	jobrepo "github.com/miruts/iJobs/usecases/job/repository"
	jobsrv "github.com/miruts/iJobs/usecases/job/service"
	jsrepo "github.com/miruts/iJobs/usecases/jobseeker/repository"
	jssrv "github.com/miruts/iJobs/usecases/jobseeker/service"
	"time"

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

var funcMaps = template.FuncMap{"cmp": handlers.JobCmp, "appGetJob": handlers.AppGetJob, "appGetJs": handlers.AppJs, "appGetJobCatId": handlers.AppGetJobCatId, "appGetCmp": handlers.AppGetCmp, "appGetLoc": handlers.AppGetLocation}

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
	errs := db.CreateTable(&entity.Session{}, &entity.Address{}, &entity.Application{}, &entity.Category{}, &entity.Job{}, &entity.Company{}, &entity.Jobseeker{}, &entity.Role{}, &entity.User{}).GetErrors()
	if len(errs) > 0 {
		fmt.Println(errs[0])
		return
	}
}
func main() {
	// Gorm Database Connection
	gormDB, err = gorm.Open("postgres", "user=postgres dbname=ijobs_gorm_db_2 password=postgres sslmode=disable")
	if errs != nil {
		fmt.Println(err)
		return
	}

	// Create Gorm Tables
	// Run Once

	gormDB.Set("gorm:insert_option", "ON DUPLICATE KEY UPDATE")
	//gormDB.AutoMigrate(&entity.Session{}, &entity.Address{}, &entity.Application{}, &entity.Category{}, &entity.Job{}, &entity.Company{}, &entity.Jobseeker{}, &entity.Role{})
	//CreateTables(gormDB)
	// Data Repositories

	sess := configSess()
	csrfSignKey := []byte(rndtoken.GenerateRandomID(32))
	applicationRepo := apprepo.NewAppGormRepositoryImpl(gormDB)
	companyRepo := cmprepo.NewCompanyGormRepositoryImpl(gormDB)
	jobRepo := jobrepo.NewJobGormRepositoryImpl(gormDB)
	categoryRepo := jobrepo.NewCategoryGormRepositoryImpl(gormDB)
	jobseekerRepo := jsrepo.NewJobseekerGormRepositoryImpl(gormDB)
	addressRepo := jsrepo.NewAddressGormRepositoryImpl(gormDB)
	sessionRepo := repository2.NewSessionGormRepo(gormDB)
	roleRepo := repository2.NewRoleGormRepo(gormDB)

	// Services
	companySrv := cmpsrv.NewCompanyServiceImpl(companyRepo)

	categorySrv := jobsrv.NewCategoryServiceImpl(categoryRepo)
	jobSrv := jobsrv.NewJobServices(jobRepo, categorySrv)
	jobseekerSrv := jssrv.NewJobseekerServiceImpl(jobseekerRepo, jobSrv)
	addressSrv := jssrv.NewAddressServiceImpl(addressRepo)
	sessionSrv := service2.NewSessionService(sessionRepo)
	applicationSrv := appsrv.NewAppService(applicationRepo, jobseekerSrv, jobSrv, companySrv)
	roleSrv := service2.NewRoleService(roleRepo)
	// Handlers
	welcomeHandler := handlers.NewWelcomeHandler(tmpl, sessionSrv, jobseekerSrv, companySrv)
	jh := handlers.NewJobseekerHandler(companySrv, jobSrv, tmpl, jobseekerSrv, categorySrv, addressSrv, applicationSrv, sessionSrv, sess, roleSrv, csrfSignKey)
	ch := handlers.NewCompanyHandler(tmpl, jobseekerSrv, companySrv, categorySrv, addressSrv, applicationSrv, sessionSrv, jobSrv, sess, roleSrv, csrfSignKey)

	//go util.ClearExpiredSessions(sessionSrv)

	//RESTApi Handlers
	apiJobHandler := apijobhandler.NewJobApiHandler(jobSrv)
	apiJobSkHandler := apijobhandler.NewJobseekerHandler(jobseekerSrv)

	//File Server
	//fs := http.FileServer(http.Dir("ui/asset"))
	router := httprouter.New()

	// Welcome SignIn/Up path registration
	router.GET("/", welcomeHandler.Welcome)
	router.GET("/login", jh.Signup)
	router.POST("/login/jobseeker", jh.Login)
	router.POST("/login/company", ch.Login)
	router.POST("/signup/jobseeker", jh.Signup)
	router.POST("/signup/company", ch.Signup)

	router.GET("/company/:username/postjob", ch.Authenticated(ch.Authorized(http.HandlerFunc(ch.CompanyPostJob))))
	router.POST("/company/:username/postjob", ch.Authenticated(ch.Authorized(http.HandlerFunc(ch.CompanyPostJob))))
	router.GET("/company/:username/jobs", ch.Authenticated(ch.Authorized(http.HandlerFunc(ch.CompanyJobs))))
	router.GET("/company/:username/home", ch.Authenticated(ch.Authorized(http.HandlerFunc(ch.CompanyHome))))

	// Jobseeker path registration
	router.GET("/jobseeker/:username/home", jh.Authenticated(jh.Authorized(http.HandlerFunc(jh.JobseekerHome))))
	router.GET("/jobseeker/:username/apply", jh.Authenticated(jh.Authorized(http.HandlerFunc(jh.JobseekerApply))))
	router.GET("/jobseeker/:username/profile", jh.Authenticated(jh.Authorized(http.HandlerFunc(jh.JobseekerProfile))))
	router.GET("/jobseeker/:username/profile/edit", jh.Authenticated(jh.Authorized(http.HandlerFunc(jh.ProfileEdit))))
	router.POST("/jobseeker/:username/profile/edit", jh.Authenticated(jh.Authorized(http.HandlerFunc(jh.ProfileEdit))))
	router.GET("/jobseeker/:username/appliedjobs", jh.Authenticated(jh.Authorized(http.HandlerFunc(jh.JobseekerAppliedJobs))))

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
func configSess() *entity.Session {
	tokenExpires := time.Now().Add(time.Minute * 30).Unix()
	sessionID := rndtoken.GenerateRandomID(32)
	signingString, err := rndtoken.GenerateRandomString(32)
	if err != nil {
		panic(err)
	}
	signingKey := []byte(signingString)

	return &entity.Session{
		Expires:    tokenExpires,
		SigningKey: signingKey,
		Uuid:       sessionID,
	}
}
