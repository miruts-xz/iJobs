package job

import "github.com/akuadane/iJobs/entity"

type JobRepository interface {
	Jobs() ([]entity.Job, error)
	JobsOfCategory(cat_id int) ([]entity.Job, error)
	Job(id int) (entity.Job, error)
	UpdateJob(job entity.Job) error
	DeleteJob(id int) error
	StoreJob(job entity.Job) error
}
