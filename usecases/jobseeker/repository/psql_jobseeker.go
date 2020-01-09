package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/miruts/iJobs/entity"
)

// JobseekerGormRepositoryIMpl implements JobseekerRepository interface
type JobseekerRepositoryImpl struct {
	conn *sql.DB
}

// NewJobseekerRepositoryImpl returns new JobseekerGormRepositoryIMpl
func NewJobseekerRepositoryImpl(jsr *sql.DB) *JobseekerRepositoryImpl {
	return &JobseekerRepositoryImpl{conn: jsr}
}

// JobSeekers retrieves and returns all jobseekers
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

// JobSeeker return a jobseeker with given id
func (jsr *JobseekerRepositoryImpl) JobSeeker(id int) (entity.JobSeeker, error) {
	query := "select * from jobseekers where id = $1"
	var jobSeeker entity.JobSeeker
	err := jsr.conn.QueryRow(query, id).Scan(&jobSeeker.ID, &jobSeeker.Username, &jobSeeker.Fullname, &jobSeeker.Email, &jobSeeker.Phone, &jobSeeker.Password, &jobSeeker.Profile, &jobSeeker.WorkExperience, &jobSeeker.CV, &jobSeeker.Portfolio, &jobSeeker.EmpStatus, &jobSeeker.Gender, &jobSeeker.Age)
	if err != nil {
		return jobSeeker, errors.New("unable to retrieve jobseeker")
	}
	return jobSeeker, nil
}

// UpdateJobSeeker updates a given jobseeker
func (jsr *JobseekerRepositoryImpl) UpdateJobSeeker(js *entity.JobSeeker) (*entity.JobSeeker, error) {
	query := "update jobseekers set id=$1, username=$2, fullname=$3, email=$4, phone=$5, password=$6, profile=$7, work_exp=$8, cv=$9, portfolio=$10, emp_status=$11, gender=$12, age=$13"
	_, err := jsr.conn.Exec(query, js.ID, js.Username, js.Fullname, js.Email, js.Phone, js.Password, js.Profile, js.WorkExperience, js.CV, js.Portfolio, js.EmpStatus, js.Gender, js.Age)
	if err != nil {
		return js, errors.New("unable to update jobseeker")
	}
	return js, nil
}

// DeleteJobSeeker deletes a jobseeker with a given id
func (jsr *JobseekerRepositoryImpl) DeleteJobSeeker(id int) (entity.JobSeeker, error) {
	js, err := jsr.JobSeeker(id)
	if err != nil {
		return js, nil
	}
	query := "delete from jobseekers where id=$1"
	_, err = jsr.conn.Exec(query, id)
	if err != nil {
		return js, errors.New("unable to delete jobseeker")
	}
	return js, nil
}

// JsCategories return all interested job categories of jobseeker with a given jobseeker id
func (jsr *JobseekerRepositoryImpl) JsCategories(id int) ([]entity.Category, error) {
	query := "select cat_id from jobseeker_categories where js_id = $1"
	rows, err := jsr.conn.Query(query, id)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	var category entity.Category
	var categories []entity.Category
	categquery := "select * from job_categories where id = $1"
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return categories, err
		}
		ctgrows, err := jsr.conn.Query(categquery, id)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return categories, err
		}
		for ctgrows.Next() {
			if err := ctgrows.Scan(&category.ID, &category.Name, &category.Desc, &category.Image); err != nil {
				return categories, nil
			}
			categories = append(categories, category)
		}
	}
	return categories, nil

}

// StoreJobSeeker stores new jobseeker
func (jsr *JobseekerRepositoryImpl) StoreJobSeeker(js *entity.JobSeeker) (*entity.JobSeeker, error) {
	query := "insert into jobseekers (username, fullname, email, phone, password, profile, work_exp, cv, portfolio, emp_status, gender, age) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"
	_, err := jsr.conn.Exec(query, js.ID, js.Username, js.Fullname, js.Email, js.Phone, js.Password, js.Profile, js.WorkExperience, js.CV, js.Portfolio, js.EmpStatus, js.Gender, js.Age)
	if err != nil {
		return js, errors.New("unable to store jobseeker")
	}
	return js, nil
}

// AddIntCategory adds new Interested category list given jobseeker and category id
func (jss *JobseekerRepositoryImpl) AddIntCategory(jsid, jcid int) error {
	query := "insert into jobseeker_categories (js_id, cat_id) values ($1, $2)"
	_, err := jss.conn.Exec(query, jsid, jcid)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return err
	}
	return nil
}

// RemoveIntCategory removes category from interested list of categories given category and jobseeker id
func (jss *JobseekerRepositoryImpl) RemoveIntCategory(jsid, jcid int) error {
	query := "delete from jobseeker_categories where js_id = $1 and cat_id = $2"
	_, err := jss.conn.Exec(query, jsid, jcid)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return err
	}
	return nil
}
