package main

import (
	"database/sql"
	_ "database/sql"
	"errors"
	"fmt"
	_ "fmt"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/iJobs/deliverable/http/api"
	"github.com/miruts/iJobs/entity"
	jobrepo "github.com/miruts/iJobs/usecases/job/repository"
	"github.com/miruts/iJobs/usecases/jobseeker/repository"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
	jobservice "github.com/miruts/iJobs/usecases/job/service"
	"github.com/miruts/iJobs/usecases/jobseeker/service"
)

func init() {

}
func main() {
	/**
	templates, global database connection and interfaces
	*/
	_ = template.Must(template.ParseGlob("ui/template/*.html"))
	// Company database connection
	pqconncmp, errcmp := sql.Open("postgres", "user=company password=company database=ijobs sslmode=disable")
	// Jobseeker database connection
	pqconnjs, errjs := sql.Open("postgres", "user=jobseeker password=jobseeker database=ijobs sslmode=disable")
	//Jobseeker gorm database connection
	gormdb, err := gorm.Open("postgres", "user=postgres dbname=apidatabase password=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer gormdb.Close()
	gormdb.CreateTable(&entity.JobSeeker{}, &entity.Address{}, &entity.Category{})
	if errcmp != nil {
		panic(errors.New("unable to connect with database with company account"))
	}
	if err := pqconncmp.Ping(); err != nil {
		panic(err)
	}
	if errjs != nil {
		panic(errors.New("unable to connect with database with jobseeker account"))
	}
	if err := pqconnjs.Ping(); err != nil {
		panic(err)
	}

	// Job Service Infrastructure
	jobRepo := jobrepo.NewJobRepository(pqconnjs)
	jobSrv := jobservice.NewJobService(jobRepo)

	// Jobseeker API Infrastructure
	jsRepo := repository.NewJobseekerGormRepositoryImpl(gormdb)
	jsSrv := service.NewJobseekerServiceImpl(jsRepo, jobSrv)

	// JobSeeker API Handler
	jsHandler := api.NewJobseekerHandler(jsSrv)

	router := httprouter.New()
	router.GET("/jobseekers", jsHandler.Jobseekers)
	router.GET("/jobseekers/:id", jsHandler.Jobseeker)
	router.POST("/jobseekers", jsHandler.AddJobseeker)
	router.PUT("/jobseekers/:id", jsHandler.UpdateJobseeker)
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
		return
	}
}
