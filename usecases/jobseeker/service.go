package jobseeker

import "github.com/miruts/iJobs/entity"

type JobseekerService interface {
	JobSeekers() ([]entity.JobSeeker, error)
	JobSeeker(id int) (entity.JobSeeker, error)
	UpdateJobSeeker(js entity.JobSeeker) error
	DeleteJobSeeker(id int) error
	StoreJobSeeker(js entity.JobSeeker) error
	Suggestions(id int) ([]entity.Job, error)
}
