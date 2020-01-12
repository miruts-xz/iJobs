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
func (cpr *CompanyGormRepositoryImpl) Companies() ([]entity.Company, error) {
	var companies []entity.Company
	errs := cpr.conn.Find(&companies).GetErrors()
	if len(errs) > 0 {
		return companies, errs[0]
	}
	return companies, nil
}
func (cpr *CompanyGormRepositoryImpl) Company(cid int) (entity.Company, error) {
	var company entity.Company
	errs := cpr.conn.First(&company, cid).GetErrors()
	if len(errs) > 0 {
		return company, errs[0]
	}
	return company, nil
}
func (cpr *CompanyGormRepositoryImpl) UpdateCompany(cmp *entity.Company) (*entity.Company, error) {
	company := cmp
	errs := cpr.conn.Save(&company).GetErrors()
	if len(errs) > 0 {
		return company, errs[0]
	}
	return company, nil
}
func (cpr *CompanyGormRepositoryImpl) DeleteCompany(cid int) (entity.Company, error) {
	var company entity.Company
	errs := cpr.conn.First(&company, cid).GetErrors()
	if len(errs) > 0 {
		return company, errs[0]
	}
	return company, nil
}
func (cpr *CompanyGormRepositoryImpl) StoreCompany(cmp *entity.Company) (*entity.Company, error) {
	company := cmp
	errs := cpr.conn.Create(&company).GetErrors()
	if len(errs) > 0 {
		return company, errs[0]
	}
	return company, nil
}
func (cpr *CompanyGormRepositoryImpl) PostedJobs(cid int) ([]entity.Job, error) {
	company, err := cpr.Company(cid)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var jobs []entity.Job
	errs := cpr.conn.Model(&company).Related(&jobs).GetErrors()
	if len(errs) > 0 {
		return jobs, errs[0]
	}
	return jobs, nil
}
func (cpr *CompanyGormRepositoryImpl) CompanyByEmail(email string) (entity.Company, error) {
	var company entity.Company
	errs := cpr.conn.Where("email = ?", email).First(&company).GetErrors()
	if len(errs) > 0 {
		return company, errs[0]
	}
	return company, nil
}
func (cpr *CompanyGormRepositoryImpl) CompanyAddress(id uint) (entity.Address, error) {
	address := entity.Address{}
	company, err := cpr.Company(int(id))
	if err != nil {
		return address, err
	}
	errs := cpr.conn.Model(&company).Related(&address, "Address").GetErrors()
	if len(errs) > 0 {
		return address, errs[0]
	}
	return address, nil
}
