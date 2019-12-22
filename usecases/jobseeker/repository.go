package jobseeker

import "github.com/miruts/iJobs/entity"

// JobseekerRepository interface defines all jobseeker related data/database operations
type JobseekerRepository interface {
	JobSeekers() ([]entity.JobSeeker, error)
	JobSeeker(id int) (entity.JobSeeker, error)
	UpdateJobSeeker(js entity.JobSeeker) error
	DeleteJobSeeker(id int) error
	StoreJobSeeker(js entity.JobSeeker) error
	JsCategories(id int) ([]entity.Category, error)
}
