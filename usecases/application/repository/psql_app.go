package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/miruts/iJobs/entity"
)

type AppRepository struct {
	conn *sql.DB
}

func NewAppRepo(conn *sql.DB) *AppRepository {
	return &AppRepository{conn: conn}
}

func (appRepo *AppRepository) Store(app *entity.Application) error {

	query := "INSERT INTO applications (job_id,jobseeker_id,status,response) values ($1,$2,$3,$4);"
	_, err := appRepo.conn.Exec(query, app.ID, app.JobID, app.JobseekerID, app.Status, app.Response)

	if err != nil {
		return errors.New("Unable to insert application")
	}

	return nil

}
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
