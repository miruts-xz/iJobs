package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	apprepo "github.com/miruts/iJobs/usecases/application/repository"
	appsrv "github.com/miruts/iJobs/usecases/application/service"
	cmprepo "github.com/miruts/iJobs/usecases/company/repository"
	cmpsrv "github.com/miruts/iJobs/usecases/company/service"
	jobrepo "github.com/miruts/iJobs/usecases/job/repository"
	jobsrv "github.com/miruts/iJobs/usecases/job/service"
	jsrepo "github.com/miruts/iJobs/usecases/jobseeker/repository"
	jssrv "github.com/miruts/iJobs/usecases/jobseeker/service"
	"html/template"
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
func main() {
	// Gorm Database Connection
	gormDB, err = gorm.Open("postgres", "user=postgres dbname=ijobs_gorm_db password=postgres sslmode=disable")
	if errs != nil {
		fmt.Println(err)
		return
	}
	defer gormDB.Close()

	// Data Repositories
	applicationRepo := apprepo.NewAppGormRepo(gormDB)
	companyRepo := cmprepo.NewCompanyGormRepositoryImpl(gormDB)
	jobRepo := jobrepo.NewJobGormRepository(gormDB)
	categoryRepo := jobrepo.NewCategoryGormRepositoryImpl(gormDB)
	jobseekerRepo := jsrepo.NewJobseekerGormRepositoryImpl(gormDB)
	addressRepo := jsrepo.NewAddressGormRepositoryImpl(gormDB)

	// Services
	applicationSrv := appsrv.NewAppservice(applicationRepo)
	companySrv := cmpsrv.NewCompanyServiceImpl(companyRepo)
	jobSrv := jobsrv.NewJobService(jobRepo)
	categorySrv := jobsrv.NewCategoryServiceImpl(categoryRepo)
	jobseekerSrv := jssrv.NewJobseekerServiceImpl(jobseekerRepo, jobSrv)
	addressSrv := jssrv.NewAddressServiceImpl(addressRepo)

	// Handlers

}
