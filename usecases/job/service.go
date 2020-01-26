package job

import (
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/company"
)

// JobService specifies all job related services
type JobService interface {
	Jobs() ([]entity.Job, error)
	JobsOfCategory(cat_id int) ([]entity.Job, error)
	Job(id int) (entity.Job, error)
	UpdateJob(job *entity.Job) (*entity.Job, error)
	DeleteJob(id int) (entity.Job, error)
	StoreJob(job *entity.Job) (*entity.Job, error)
	CompanyJobs(cmpSrv company.CompanyService, cmid int) ([]entity.Job, error)
}

// CategoryService specifies food menu category services
type CategoryService interface {
	Categories() ([]entity.Category, error)
	Category(id int) (entity.Category, error)
	UpdateCategory(category *entity.Category) (*entity.Category, error)
	DeleteCategory(id int) (entity.Category, error)
	StoreCategory(category *entity.Category) (*entity.Category, error)
}
