package job

import (
	"github.com/miruts/iJobs/entity"
)

type JobService interface {
	Jobs() ([]entity.Job, error)
	JobsOfCategory(cat_id int) ([]entity.Job, error)
	Job(id int) (entity.Job, error)
	UpdateJob(job entity.Job) error
	DeleteJob(id int) error
	StoreJob(job entity.Job) error
}

// CategoryService specifies food menu category services
type CategoryService interface {
	Categories() ([]entity.Category, error)
	Category(id int) (entity.Category, error)
	UpdateCategory(category entity.Category) error
	DeleteCategory(id int) error
	StoreCategory(category entity.Category) error
}
