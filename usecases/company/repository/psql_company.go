package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/miruts/iJobs/entity"
)

// CompanyRepositoryImpl implements CompanyRepository interface
type CompanyRepositoryImpl struct {
	conn *sql.DB
}

// NewCompanyRepositoryImpl returns new CompanyRepositoryImpl
func NewCompanyRepositoryImpl(cpr *sql.DB) *CompanyRepositoryImpl {
	return &CompanyRepositoryImpl{conn: cpr}
}

// Companies retrieves and returns all companies
func (cpr *CompanyRepositoryImpl) Companies() ([]entity.Company, error) {
	query := "select * from companies"
	rows, err := cpr.conn.Query(query)
	if err != nil {
		return nil, errors.New("unable to retrieve companys")
	}
	var companys []entity.Company
	var company entity.Company
	for rows.Next() {
		if err := rows.Scan(&company.ID, &company.CompanyName, &company.Password, &company.Email, &company.Phone, &company.Logo, &company.ShortDesc, &company.DetailInfo, &company.Address); err != nil {
			return nil, errors.New("unable to retrieve companys")
		}
		companys = append(companys, company)
	}
	return companys, nil
}

// Company return a Company with given id
func (cpr *CompanyRepositoryImpl) Company(id int) (entity.Company, error) {
	query := "select * from companies where id = $1"
	var company entity.Company
	err := cpr.conn.QueryRow(query, id).Scan(company.ID, &company.CompanyName, &company.Password, &company.Email, &company.Phone, &company.Logo, &company.ShortDesc, &company.DetailInfo, &company.Address)
	if err != nil {
		return company, errors.New("unable to retrieve company")
	}
	return company, nil
}

// UpdateCompany updates a given company
func (cpr *CompanyRepositoryImpl) UpdateCompany(cp *entity.Company) (*entity.Company, error) {
	query := "update companies set id=$1, company_name=$2, password=$3, email=$4, phone=$5, logo=$7, short_desc=$9, detail_info=$10"
	_, err := cpr.conn.Exec(query, cp.ID, cp.CompanyName, cp.Password, cp.Email, cp.Phone, cp.Logo, cp.ShortDesc, cp.DetailInfo, cp.Address)
	if err != nil {
		return cp, errors.New("unable to update company")
	}
	return cp, nil
}

// DeleteCompany deletes a company with a given id
func (cpr *CompanyRepositoryImpl) DeleteCompany(id int) (entity.Company, error) {
	company, err := cpr.Company(id)
	if err != nil {
		fmt.Println(err)
		return company, err
	}
	query := "delete from companies where id=$1"
	_, err = cpr.conn.Exec(query, id)
	if err != nil {
		return company, errors.New("unable to delete company")
	}
	return company, nil
}

// StoreCompany stores new company
func (cpr *CompanyRepositoryImpl) StoreCompany(cp *entity.Company) (*entity.Company, error) {
	query := "insert into companies (ID, company_name, Password, Email, Phone, Logo, Short_desc, Detail_info) values ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err := cpr.conn.Exec(query, cp.ID, cp.CompanyName, cp.Password, cp.Email, cp.Phone, cp.Logo, cp.ShortDesc, cp.DetailInfo)
	if err != nil {
		return cp, errors.New("unable to store company")
	}
	return cp, nil
}

// Posted jobs retrieves jobs jobs posted by company
func (cpr *CompanyRepositoryImpl) PostedJobs(cid int) ([]entity.Job, error) {
	var jobs []entity.Job
	return jobs, errors.New("un implemented method error")
}

// CompanyByEmail retrieves company given email
func (cpr *CompanyRepositoryImpl) CompanyByEmail(email string) (entity.Company, error) {
	var company entity.Company
	return company, errors.New("un implemented method error")
}

// CompanyAddress retrieves address of a company given company id
func (cpr *CompanyRepositoryImpl) CompanyAddress(id uint) (entity.Address, error) {
	var address entity.Address
	return address, errors.New("un implemented method error")
}
