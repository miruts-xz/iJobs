package repository

import (
	"database/sql"
	"errors"
	"github.com/miruts/iJobs/usecases/company"

	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/job"
)

// JobRepository implements JobRepository interface
type JobRepository struct {
	conn *sql.DB
}

//Creates a new JobRepository Object
func NewJobRepository(conn *sql.DB) *JobRepository {
	return &JobRepository{conn: conn}
}

//Returns all the jobs that have been posted so far
func (jobRepo *JobRepository) Jobs() ([]entity.Job, error) {

	query := "SELECT * FROM jobs;"

	records, err := jobRepo.conn.Query(query)

	if err != nil {
		return nil, errors.New("Unable to retrieve jobs")
	}
	defer records.Close()

	jobs := []entity.Job{}

	for records.Next() {
		aJob := entity.Job{}
		records.Scan(&aJob.ID, &aJob.Name, &aJob.CompanyID, &aJob.Salary,
			&aJob.RequiredNum, &aJob.Categories, &aJob.Deadline, &aJob.Description, &aJob.JobTime)

		jobs = append(jobs, aJob)
	}

	return jobs, nil
}

//Returns all jobs under a specific category
func (jobRepo *JobRepository) JobsOfCategory(ctgSrv job.CategoryService, cat_id int) ([]entity.Job, error) {

	query := "SELECT * FROM jobs where id =$1;"

	records, err := jobRepo.conn.Query(query, cat_id)

	if err != nil {
		return nil, errors.New("Unable to retrieve jobs")
	}
	defer records.Close()

	jobs := []entity.Job{}

	for records.Next() {
		aJob := entity.Job{}
		records.Scan(&aJob.ID, &aJob.Name, &aJob.CompanyID, &aJob.Salary,
			&aJob.RequiredNum, &aJob.Categories, &aJob.Deadline, &aJob.Description, &aJob.JobTime)

		jobs = append(jobs, aJob)
	}

	return jobs, nil
}

//Returns a job given an its id
func (jobRepo *JobRepository) Job(id int) (entity.Job, error) {

	query := "SELECT * FROM jobs where id=$1;"
	record := jobRepo.conn.QueryRow(query, id)

	aJob := entity.Job{}

	err := record.Scan(&aJob.ID, &aJob.Name, &aJob.CompanyID, &aJob.Salary,
		&aJob.RequiredNum, &aJob.Categories, &aJob.Deadline, &aJob.Description, &aJob.JobTime)

	if err != nil {
		return aJob, errors.New("Unable to retrieve job")
	}
	return aJob, nil
}

//Updates a job given the udpated job object
func (jobRepo *JobRepository) UpdateJob(job *entity.Job) (*entity.Job, error) {

	query := "UPDATE jobs SET name=$1,salary=$2,required_num=$3,id=$4,deadline=$5,description=$6,job_time=$7 WHERE id=$8"
	_, err := jobRepo.conn.Exec(query, job.Name, job.Salary, job.RequiredNum, job.Categories, job.Deadline, job.Description, job.JobTime, job.ID)

	if err != nil {
		return job, errors.New("Unable to update job")
	}
	return job, nil

}

//Deletes a job given its id
func (jobRepo *JobRepository) DeleteJob(id int) (entity.Job, error) {
	var job entity.Job
	query := "DELETE FROM jobs WHERE id=$1"
	_, err := jobRepo.conn.Exec(query, id)

	if err != nil {
		return job, errors.New("Unable to delete job")
	}

	return job, nil
}

//Adds a job to the database
func (jobRepo *JobRepository) StoreJob(job *entity.Job) (*entity.Job, error) {

	query := "INSERT INTO jobs (name,id,salary,required_num,id,deadline,description,job_time) values ($1,$2,$3,$4,$5,$6,$7,$8);"
	_, err := jobRepo.conn.Exec(query, job.Name, job.CompanyID, job.Salary, job.RequiredNum, job.Categories, job.Deadline, job.Description, job.JobTime)

	if err != nil {
		return &entity.Job{}, errors.New("Unable to create job")
	}
	return &entity.Job{}, nil
}
func (jobRepo *JobRepository) CompanyJobs(cmpSrv company.CompanyService, cm_id int) ([]entity.Job, error) {
	return []entity.Job{entity.Jobmock1, entity.Jobmock2}, nil
}
