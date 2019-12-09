package repository

import (
	"database/sql"
	"errors"
	"github.com/miruts/iJobs/entity"
)

type JobseekerRepositoryImpl struct {
	conn *sql.DB
}

func NewJobseekerRepositoryImpl(jsr *sql.DB) *JobseekerRepositoryImpl {
	return &JobseekerRepositoryImpl{conn: jsr}
}
func (jsr *JobseekerRepositoryImpl) JobSeekers() ([]entity.JobSeeker, error) {
	query := "select * from jobseekers"
	rows, err := jsr.conn.Query(query)
	if err != nil {
		return nil, errors.New("unable to retrieve jobseekers")
	}
	var jobSeekers []entity.JobSeeker
	var jobSeeker entity.JobSeeker
	for rows.Next() {
		if err := rows.Scan(&jobSeeker.ID, &jobSeeker.Username, &jobSeeker.Fullname, &jobSeeker.Email, &jobSeeker.Phone, &jobSeeker.Password, &jobSeeker.Profile, &jobSeeker.WorkExperience, &jobSeeker.CV, &jobSeeker.Portfolio, &jobSeeker.EmpStatus, &jobSeeker.Gender, &jobSeeker.Age); err != nil {
			return nil, errors.New("unable to retrieve jobseekers")
		}
		jobSeekers = append(jobSeekers, jobSeeker)
	}
	return jobSeekers, nil
}
func (jsr *JobseekerRepositoryImpl) JobSeeker(id int) (entity.JobSeeker, error) {
	query := "select * from jobseekers where id = $1"
	var jobSeeker entity.JobSeeker
	err := jsr.conn.QueryRow(query, id).Scan(&jobSeeker.ID, &jobSeeker.Username, &jobSeeker.Fullname, &jobSeeker.Email, &jobSeeker.Phone, &jobSeeker.Password, &jobSeeker.Profile, &jobSeeker.WorkExperience, &jobSeeker.CV, &jobSeeker.Portfolio, &jobSeeker.EmpStatus, &jobSeeker.Gender, &jobSeeker.Age)
	if err != nil {
		return jobSeeker, errors.New("unable to retrieve jobseeker")
	}
	return jobSeeker, nil
}
func (jsr *JobseekerRepositoryImpl) UpdateJobSeeker(js entity.JobSeeker) error {
	query := "update jobseekers set id=$1, username=$2, fullname=$3, email=$4, phone=$5, password=$6, profile=$7, work_exp=$8, cv=$9, portfolio=$10, emp_status=$11, gender=$12, age=$13"
	_, err := jsr.conn.Exec(query, js.ID, js.Username, js.Fullname, js.Email, js.Phone, js.Password, js.Profile, js.WorkExperience, js.CV, js.Portfolio, js.EmpStatus, js.Gender, js.Age)
	if err != nil {
		return errors.New("unable to update jobseeker")
	}
	return nil
}
func (jsr *JobseekerRepositoryImpl) DeleteJobSeeker(id int) error {
	query := "delete from jobseekers where id=$1"
	_, err := jsr.conn.Exec(query, id)
	if err != nil {
		return errors.New("unable to delete jobseeker")
	}
	return nil
}
func (jsr *JobseekerRepositoryImpl) StoreJobSeeker(js entity.JobSeeker) error {
	query := "insert into jobseekers (username, fullname, email, phone, password, profile, work_exp, cv, portfolio, emp_status, gender, age) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"
	_, err := jsr.conn.Exec(query, js.ID, js.Username, js.Fullname, js.Email, js.Phone, js.Password, js.Profile, js.WorkExperience, js.CV, js.Portfolio, js.EmpStatus, js.Gender, js.Age)
	if err != nil {
		return errors.New("unable to store jobseeker")
	}
	return nil
}
