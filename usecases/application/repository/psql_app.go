package repository

import (
	"database/sql"
	"errors"

	"github.com/miruts/iJobs/entity"
)

type AppRepository struct {
	conn *sql.DB
}

func NewAppRepo(conn *sql.DB) *AppRepository {
	return &AppRepository{conn: conn}
}

func (appRepo *AppRepository) Store(app entity.Application) error {

	query := "INSERT INTO applications (id,job_id,js_id,status,response) values ($1,$2,$3,$4);"
	_, err := appRepo.conn.Exec(query, app.ID, app.JobID, app.JobseekerID, app.Status, app.Response)

	if err != nil {
		return errors.New("Unable to insert application")
	}

	return nil

}

func (appRepo *AppRepository) Application(appId int) ([]entity.Application, error) {

	query := "SELECT * FROM applications WHERE id=$1"
	records, err := appRepo.conn.Query(query, appId)

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

func (appRepo *AppRepository) UserApplication(JsId int) ([]entity.Application, error) {

	query := "SELECT * FROM applications WHERE js_id=$1"
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

func (appRepo *AppRepository) DeleteApplication(id int) error {
	query := "DELETE FROM applications WHERE id=$1"
	_, err := appRepo.conn.Exec(query, id)

	if err != nil {
		return errors.New("Unable to delete application")
	}

	return nil
}
