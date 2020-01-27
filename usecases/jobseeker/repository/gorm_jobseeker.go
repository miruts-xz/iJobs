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
	var addresses []entity.Address
	var categories []entity.Category
	var applications []entity.Application
	errs := jsr.conn.Find(&jobseekers).GetErrors()
	if len(errs) > 0 {
		fmt.Printf("Error: %v", errs)
		return jobseekers, errs[0]
	}
	for i, _ := range jobseekers {
		_ = jsr.conn.Model(&jobseekers[i]).Related(&addresses, "Address").GetErrors()
		_ = jsr.conn.Model(&jobseekers[i]).Related(&categories, "Categories").GetErrors()
		_ = jsr.conn.Model(&jobseekers[i]).Related(&applications, "Applications").GetErrors()
		fmt.Println(addresses)
		fmt.Println(categories)
		jobseekers[i].Address = addresses
		jobseekers[i].Categories = categories
		jobseekers[i].Applications = applications
	}
	return jobseekers, nil
}

// JobSeeker return a jobseeker with given id
func (jsr *JobseekerGormRepositoryIMpl) JobSeeker(id int) (entity.Jobseeker, error) {
	var jobseeker entity.Jobseeker
	var addresses []entity.Address
	var categories []entity.Category
	var applications []entity.Application
	errs := jsr.conn.First(&jobseeker, id).GetErrors()
	_ = jsr.conn.Model(&jobseeker).Related(&addresses, "Address").GetErrors()
	_ = jsr.conn.Model(&jobseeker).Related(&categories, "Categories").GetErrors()
	_ = jsr.conn.Model(&jobseeker).Related(&applications, "Applications").GetErrors()
	jobseeker.Address = addresses
	jobseeker.Categories = categories
	jobseeker.Applications = applications
	if len(errs) > 0 {
		return jobseeker, errs[0]
	}
	return jobseeker, nil
}

// UpdateJobSeeker updates a given jobseeker
func (jsr *JobseekerGormRepositoryIMpl) UpdateJobSeeker(js *entity.Jobseeker) (*entity.Jobseeker, error) {
	jobseeker := js
	errs := jsr.conn.Save(jobseeker).GetErrors()
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
	jobseeker, err := jsr.JobSeeker(id)
	var categories []entity.Category
	if err != nil {
		return categories, err
	}
	errs := jsr.conn.Model(&jobseeker).Related(&categories, "Categories").GetErrors()
	if len(errs) > 0 {
		return categories, errs[0]
	}
	return categories, nil
}

// StoreJobSeeker stores new jobseeker
func (jsr *JobseekerGormRepositoryIMpl) StoreJobSeeker(js *entity.Jobseeker) (*entity.Jobseeker, error) {
	jobseeker := js
	fmt.Println(jobseeker.Email)
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
func (jss *JobseekerGormRepositoryIMpl) JobseekerByUsername(uname string) (entity.Jobseeker, error) {
	var jobseeker entity.Jobseeker
	errs := jss.conn.Where("username = ?", uname).First(&jobseeker).GetErrors()
	if len(errs) > 0 {
		return jobseeker, errs[0]
	}
	return jobseeker, nil
}
func (jss *JobseekerGormRepositoryIMpl) ApplicationJobseeker(id int) (entity.Jobseeker, error) {
	var jobseeker entity.Jobseeker
	errs := jss.conn.First(&jobseeker, id).GetErrors()
	if len(errs) > 0 {
		return jobseeker, errs[0]
	}
	return jobseeker, nil
}
func (jss *JobseekerGormRepositoryIMpl) UserRoles(user *entity.Jobseeker) ([]entity.Role, []error) {
	userRoles := []entity.Role{}
	errs := jss.conn.Model(user).Related(&userRoles).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return userRoles, errs
}

// PhoneExists check if a given phone number is found
func (userRepo *JobseekerGormRepositoryIMpl) PhoneExists(phone string) bool {
	user := entity.Jobseeker{}
	errs := userRepo.conn.Find(&user, "phone=?", phone).GetErrors()
	if len(errs) > 0 {
		return false
	}
	return true
}
func (jss *JobseekerGormRepositoryIMpl) UsernameExists(email string) bool {
	user := entity.Jobseeker{}
	errs := jss.conn.Find(&user, "email=?", email).GetErrors()
	if len(errs) > 0 {
		return false
	}
	return true
}

// EmailExists check if a given email is found
func (jss *JobseekerGormRepositoryIMpl) EmailExists(email string) bool {
	user := entity.Jobseeker{}
	errs := jss.conn.Find(&user, "email=?", email).GetErrors()
	if len(errs) > 0 {
		return false
	}
	return true
}
func (jss *JobseekerGormRepositoryIMpl) AlreadyApplied(id uint, id2 uint) bool {
	application := entity.Application{}
	errs := jss.conn.First(&application, "jobseeker_id=? and job_id=?", id, id2).GetErrors()
	if len(errs) > 0 {
		return false
	}
	return true
}
