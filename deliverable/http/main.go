package main

import (
	"database/sql"
	_ "database/sql"
	"errors"
	"fmt"
	_ "fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/iJobs/deliverable/http/api"
	entity "github.com/miruts/iJobs/entity/gorm-entity"
	jobrepo "github.com/miruts/iJobs/usecases/job/repository"
	jobservice "github.com/miruts/iJobs/usecases/job/service"
	"github.com/miruts/iJobs/usecases/jobseeker/repository"
	"github.com/miruts/iJobs/usecases/jobseeker/service"
	"html/template"
	"net/http"
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
	gormdb, err := gorm.Open("postgres", "user=postgres dbname=ijobs_gorm_db password=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer gormdb.Close()
	js := entity.Jobseeker{}
	Ctgs := entity.Category{
		Name:  "Software Development",
		Image: "software.jpg",
		Desc:  "Jobs related to Software design and development",
	}
	Addr := entity.Address{
		Region:    "Oromia",
		City:      "Mekelle",
		SubCity:   "Somewhere",
		LocalName: "localname",
	}
	js.Categories = []entity.Category{Ctgs}
	js.Address = []entity.Address{Addr}
	js.Username = "akayou"
	js.Fullname = "akayou adane"
	js.Gender = entity.MALE
	js.Profile = "akayou.png"
	js.WorkExperience = 2
	js.CV = "akayou.cv.pdf"
	js.Portfolio = "www.github.com/akuadane"
	js.Password = "akayou@password"
	js.Email = "akayou.adane@aait.edu.et"
	js.EmpStatus = entity.UNEMPLD
	js.Age = 21
	js.Phone = "251454545454"
	gormdb.Create(&js)

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
