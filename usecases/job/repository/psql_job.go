package repository

import (
	"database/sql"

	"github.com/akuadane/iJobs/entity"
)

type JobRepository struct {
	conn *sql.DB
}

func NewJobRepository(conn *sql.DB) *JobRepository {
	return &JobRepository{conn: conn}
}

func (jobRepo *JobRepository) Jobs() ([]entity.Job, error) {

	query := "SELECT * FROM jobs;"

	records, err := jobRepo.conn.Query(query)

	if err != nil {
		return nil, err
	}
	defer records.Close()

	jobs := []entity.Job{}

	for records.Next() {
		aJob := entity.Job{}
		records.Scan(&aJob.ID, &aJob.Name, &aJob.CompanyID, &aJob.Salary,
			&aJob.RequiredNum, &aJob.CategoryID, &aJob.Deadline, &aJob.Description, &aJob.JobTime)

		jobs = append(jobs, aJob)
	}

	return jobs, nil
}

func (jobRepo *JobRepository) Job(id int) (entity.Job, error) {

	query := "SELECT * FROM jobs where id=$1;"
	

}
func (jobRepo *JobRepository) UpdateJob(job entity.Job) error {

}
func (jobRepo *JobRepository) DeleteJob(id int) error {

}
func (jobRepo *JobRepository) StoreJob(job entity.Job) error {

}
