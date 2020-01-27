package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/miruts/iJobs/entity"
)

// AppRepository respresents psql IAppRepository Implementation
type AppRepository struct {
	conn *sql.DB
}

// NewAppRepo creates new AppRepository
func NewAppRepo(conn *sql.DB) *AppRepository {
	return &AppRepository{conn: conn}
}

// Store stores application
func (appRepo *AppRepository) Store(app *entity.Application) error {

	query := "INSERT INTO applications (job_id,jobseeker_id,status,response) values ($1,$2,$3,$4);"
	_, err := appRepo.conn.Exec(query, app.ID, app.JobID, app.JobseekerID, app.Status, app.Response)

	if err != nil {
		return errors.New("Unable to insert application")
	}

	return nil

}

// Application finds application by id
func (appRepo *AppRepository) Application(id int) (entity.Application, error) {
	query := "select * from applications where id = $1"
	var application entity.Application
	err := appRepo.conn.QueryRow(query, id).Scan(application.ID, application.CreatedAt, application.UpdatedAt, application.DeletedAt, application.JobID, application.JobseekerID, application.Response, application.Status)
	if err != nil {
		fmt.Println(err)
		return application, err
	}
	return application, nil
}

// UserApplication finds all application given jobseeker id and service
func (appRepo *AppRepository) UserApplication(JsId int) ([]entity.Application, error) {

	query := "SELECT * FROM applications WHERE jobseeker_id=$1"
	records, err := appRepo.conn.Query(query, JsId)

	if err != nil {
		return nil, errors.New("Unable to fetch application")
	}

	apps := []entity.Application{}

	for records.Next() {
		app := entity.Application{}

		records.Scan(&app.ID, &app.JobID, &app.JobseekerID, &app.Status, &app.Response)

		apps = append(apps, app)
	}
	return apps, nil

}

// ApplicationsOnJob retrieves all Application on a given job
func (appRepo *AppRepository) ApplicationsOnJob(jobId int) ([]entity.Application, error) {

	query := "SELECT * FROM applications WHERE job_id=$1"
	records, err := appRepo.conn.Query(query, jobId)

	if err != nil {
		return nil, errors.New("Unable to fetch application")
	}

	apps := []entity.Application{}

	for records.Next() {
		app := entity.Application{}

		records.Scan(&app.ID, &app.JobID, &app.JobseekerID, &app.Status, &app.Response)
		apps = append(apps, app)
	}

	return apps, nil

}

// Delete Application deletes application with given id
func (appRepo *AppRepository) DeleteApplication(id int) (entity.Application, error) {
	query := "DELETE FROM applications WHERE id=$1"
	application, err := appRepo.Application(id)
	if err != nil {
		fmt.Println(err)
		return application, err
	}
	_, err = appRepo.conn.Exec(query, id)
	if err != nil {
		return application, errors.New("Unable to delete application")
	}

	return application, nil
}

// ApplicationForCompany retrieves all job-applications for a given company
func (agr *AppRepository) ApplicationForCompany(cmid int) ([]entity.Application, error) {
	var applications []entity.Application

	return applications, errors.New("unimplemented method")
}
