package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	entity2 "github.com/miruts/iJobs/entity"
)

type CompanyGormRepositoryImpl struct {
	conn *gorm.DB
}

func NewCompanyGormRepositoryImpl(db *gorm.DB) *CompanyGormRepositoryImpl {
	return &CompanyGormRepositoryImpl{conn: db}
}
func (cgr *CompanyGormRepositoryImpl) Companies() ([]entity2.Company, error) {
	var companies []entity2.Company
	errs := cgr.conn.Find(&companies).GetErrors()
	if len(errs) > 0 {
		return companies, errs[0]
	}
	return companies, nil
}
func (cgr *CompanyGormRepositoryImpl) Company(cid int) (entity2.Company, error) {
	var company entity2.Company
	errs := cgr.conn.First(&company, cid).GetErrors()
	if len(errs) > 0 {
		return company, errs[0]
	}
	return company, nil
}
func (cgr *CompanyGormRepositoryImpl) UpdateCompany(cmp *entity2.Company) (*entity2.Company, error) {
	company := cmp
	errs := cgr.conn.Save(&company).GetErrors()
	if len(errs) > 0 {
		return company, errs[0]
	}
	return company, nil
}
func (cgr *CompanyGormRepositoryImpl) DeleteCompany(cid int) (entity2.Company, error) {
	var company entity2.Company
	errs := cgr.conn.First(&company, cid).GetErrors()
	if len(errs) > 0 {
		return company, errs[0]
	}
	return company, nil
}
func (cgr *CompanyGormRepositoryImpl) StoreCompany(cmp *entity2.Company) (*entity2.Company, error) {
	company := cmp
	errs := cgr.conn.Create(&company).GetErrors()
	if len(errs) > 0 {
		return company, errs[0]
	}
	return company, nil
}
func (cgr *CompanyGormRepositoryImpl) PostedJobs(cid int) ([]entity2.Job, error) {
	company, err := cgr.Company(cid)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var jobs []entity2.Job
	errs := cgr.conn.Model(&company).Related(&jobs).GetErrors()
	if len(errs) > 0 {
		return jobs, errs[0]
	}
	return jobs, nil
}
