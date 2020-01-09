package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/miruts/iJobs/entity"
)

type CompanyGormRepositoryImpl struct {
	conn *gorm.DB
}

func NewCompanyGormRepositoryImpl(db *gorm.DB) *CompanyGormRepositoryImpl {
	return &CompanyGormRepositoryImpl{conn: db}
}
func (cgr *CompanyGormRepositoryImpl) Companies() ([]entity.Company, error) {
	var companies []entity.Company
	errs := cgr.conn.Find(&companies).GetErrors()
	if len(errs) > 0 {
		return companies, errs[0]
	}
	return companies, nil
}
func (cgr *CompanyGormRepositoryImpl) Company(cid int) (entity.Company, error) {
	var company entity.Company
	errs := cgr.conn.First(&company, cid).GetErrors()
	if len(errs) > 0 {
		return company, errs[0]
	}
	return company, nil
}
func (cgr *CompanyGormRepositoryImpl) UpdateCompany(cmp *entity.Company) (*entity.Company, error) {
	company := cmp
	errs := cgr.conn.Save(&company).GetErrors()
	if len(errs) > 0 {
		return company, errs[0]
	}
	return company, nil
}
func (cgr *CompanyGormRepositoryImpl) DeleteCompany(cid int) (entity.Company, error) {
	var company entity.Company
	errs := cgr.conn.First(&company, cid).GetErrors()
	if len(errs) > 0 {
		return company, errs[0]
	}
	return company, nil
}
func (cgr *CompanyGormRepositoryImpl) StoreCompany(cmp *entity.Company) (*entity.Company, error) {
	company := cmp
	errs := cgr.conn.Create(&company).GetErrors()
	if len(errs) > 0 {
		return company, errs[0]
	}
	return company, nil
}
func (cgr *CompanyGormRepositoryImpl) PostedJobs(cid int) ([]entity.Job, error) {
	company, err := cgr.Company(cid)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var jobs []entity.Job
	errs := cgr.conn.Model(&company).Related(&jobs).GetErrors()
	if len(errs) > 0 {
		return jobs, errs[0]
	}
	return jobs, nil
}
func (cgr *CompanyGormRepositoryImpl) CompanyByEmail(email string) (entity.Company, error) {
	var company entity.Company
	errs := cgr.conn.Where("email = ?", email).First(&company).GetErrors()
	if len(errs) > 0 {
		return company, errs[0]
	}
	return company, nil
}
func (cgr *CompanyGormRepositoryImpl) CompanyAddress(id uint) (entity.Address, error) {
	address := entity.Address{}
	company, err := cgr.Company(int(id))
	if err != nil {
		return address, err
	}
	errs := cgr.conn.Model(&company).Related(&address).GetErrors()
	if len(errs) > 0 {
		return address, errs[0]
	}
	return address, nil
}
