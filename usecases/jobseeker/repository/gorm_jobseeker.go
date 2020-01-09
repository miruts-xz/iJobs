package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/miruts/iJobs/entity"
)

// JobseekerGormRepositoryIMpl implements JobseekerRepository interface
type JobseekerGormRepositoryIMpl struct {
	conn *gorm.DB
}

// NewJobseekerRepositoryImpl returns new JobseekerGormRepositoryIMpl
func NewJobseekerGormRepositoryImpl(jsr *gorm.DB) *JobseekerGormRepositoryIMpl {
	return &JobseekerGormRepositoryIMpl{conn: jsr}
}

// JobSeekers retrieves and returns all jobseekers
func (jsr *JobseekerGormRepositoryIMpl) JobSeekers() ([]entity.Jobseeker, error) {
	var jobseekers []entity.Jobseeker
	errs := jsr.conn.Find(&jobseekers).GetErrors()
	if len(errs) > 0 {
		fmt.Printf("Error: %v", errs)
		return jobseekers, errs[0]
	}
	return jobseekers, nil
}

// JobSeeker return a jobseeker with given id
func (jsr *JobseekerGormRepositoryIMpl) JobSeeker(id int) (entity.Jobseeker, error) {
	var jobseeker entity.Jobseeker
	errs := jsr.conn.First(&jobseeker, id).GetErrors()
	if len(errs) > 0 {
		return jobseeker, errs[0]
	}
	return jobseeker, nil
}

// UpdateJobSeeker updates a given jobseeker
func (jsr *JobseekerGormRepositoryIMpl) UpdateJobSeeker(js *entity.Jobseeker) (*entity.Jobseeker, error) {
	jobseeker := js
	errs := jsr.conn.Save(&jobseeker).GetErrors()
	if len(errs) > 0 {
		return jobseeker, errs[0]
	}
	return jobseeker, nil
}

// DeleteJobSeeker deletes a jobseeker with a given id
func (jsr *JobseekerGormRepositoryIMpl) DeleteJobSeeker(id int) (entity.Jobseeker, error) {
	jobseeker, err := jsr.JobSeeker(id)
	if err != nil {
		return jobseeker, err
	}
	errs := jsr.conn.Delete(jobseeker, id).GetErrors()
	if len(errs) > 0 {
		return jobseeker, errs[0]
	}
	return jobseeker, nil
}

// JsCategories return all interested job categories of jobseeker with a given jobseeker id
func (jsr *JobseekerGormRepositoryIMpl) JsCategories(id int) ([]entity.Category, error) {
	var categories []entity.Category
	errs := jsr.conn.Where("id in (?)", jsr.conn.Table("jobseeker_categories").Select("cat_id").Where("js_id = ?", id)).Find(&categories).GetErrors()
	if len(errs) > 0 {
		return categories, errs[0]
	}
	return categories, nil
}

// StoreJobSeeker stores new jobseeker
func (jsr *JobseekerGormRepositoryIMpl) StoreJobSeeker(js *entity.Jobseeker) (*entity.Jobseeker, error) {
	jobseeker := js
	errs := jsr.conn.Create(jobseeker).GetErrors()
	if len(errs) > 0 {
		return jobseeker, errs[0]
	}
	return jobseeker, nil
}

// AddIntCategory adds new Interested category list given jobseeker and category id
func (jss *JobseekerGormRepositoryIMpl) AddIntCategory(jsid, cat_id int) error {
	errs := jss.conn.Raw("insert into jobseeker_categories (js_id, cat_id) values ($1, $2)", jsid, cat_id).GetErrors()
	if len(errs) > 0 {
		return errs[0]
	}
	return nil
}

// RemoveIntCategory removes category from interested list of categories given category and jobseeker id
func (jss *JobseekerGormRepositoryIMpl) RemoveIntCategory(jsid, jcid int) error {
	errs := jss.conn.Raw("delete from jobseeker_categories where js_id = $1 and cat_id = $2", jsid, jcid).GetErrors()
	if len(errs) > 0 {
		return errs[0]
	}
	return nil
}
func (jss *JobseekerGormRepositoryIMpl) SetAddress(jsid, addid int) error {
	errs := jss.conn.Raw("insert into jobseeker_addresses (js_id, add_id) values(jsid, addid").GetErrors()
	if len(errs) > 0 {
		return errs[0]
	}
	return nil
}
func (jss *JobseekerGormRepositoryIMpl) JobseekerByEmail(email string) (entity.Jobseeker, error) {
	var jobseeker entity.Jobseeker
	errs := jss.conn.Where("email = ?", email).First(&jobseeker).GetErrors()
	if len(errs) > 0 {
		return jobseeker, errs[0]
	}
	return jobseeker, nil
}
