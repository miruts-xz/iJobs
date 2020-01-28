package repository

import (
	"github.com/miruts/iJobs/entity"
	"testing"
)

func TestJobseekerMockRepository(t *testing.T) {

	jsmockrepo := NewJobseekerMockRepository()
	test := []struct {
		Name      string
		parameter interface{}
		output    interface{}
	}{{Name: "Jobseekers", output: []entity.Jobseeker{entity.Jobseekermock1, entity.Jobseekermock2}},
		{Name: "Jobseeker", parameter: 1, output: entity.Jobseekermock1},
		{Name: "UpdateJobSeeker", parameter: &entity.Jobseekermock2, output: &entity.Jobseekermock2},
		{Name: "DeleteJobSeeker", parameter: 1, output: entity.Jobseekermock1},
		{Name: "StoreJobSeeker", parameter: &entity.Jobseekermock1, output: &entity.Applicatiomock1},
		{Name: "JsCategories", parameter: 1, output: []entity.Category{entity.Categorymock1, entity.Categorymock2}},
		{Name: "JobseekerByEmail", parameter: 1, output: entity.Jobseekermock1},
		{Name: "ApplicationByUsername", parameter: "user1", output: entity.Jobseekermock1},
		{Name: "ApplicationJobseeker", parameter: 1, output: entity.Jobseekermock1}}

	for _, tst := range test {
		t.Run(tst.Name, func(test *testing.T) {
			switch tst.Name {
			case "Jobseekers":
				_, err := jsmockrepo.JobSeekers()
				if err == nil {

				}
				break
			case "Jobseeker":
				op, _ := jsmockrepo.JobSeeker(1)
				if op.Username != entity.Jobseekermock1.Username {

				}
				break
			case "UpdateJobSeeker":
				op, _ := jsmockrepo.UpdateJobSeeker(&entity.Jobseekermock2)
				if op.Username != entity.Jobseekermock2.Username {
					t.Errorf("Error")
				}
				break
			case "DeleteJobSeeker":
				op, _ := jsmockrepo.DeleteJobSeeker(1)
				if op.Username != entity.Jobseekermock1.Username {
					t.Errorf("Error")
				}
				break
			case "StoreJobSeeker":
				op, _ := jsmockrepo.StoreJobSeeker(&entity.Jobseekermock1)
				if op.Username != entity.Jobseekermock1.Username {
					t.Errorf("Error")
				}
				break
			case "JsCategories":
				op, _ := jsmockrepo.JsCategories(1)
				if len(op) != 2 {
					t.Errorf("Error")
				}
				break
			case "JobseekerByEmail":
				op, _ := jsmockrepo.JobseekerByEmail(entity.Jobseekermock1.Username)
				if op.Username != entity.Jobseekermock1.Username {

				}
				break
			case "ApplicationByUsername":
				op, _ := jsmockrepo.JobseekerByUsername("user1")
				if op.Username != entity.Jobseekermock1.Username {
					t.Errorf("Error")
				}
				break
			case "ApplicationJobseeker":
				op, _ := jsmockrepo.ApplicationJobseeker(1)
				if op.Username != entity.Jobseekermock1.Username {
					t.Errorf("Error")
				}
				break
			}
		})
	}
}
