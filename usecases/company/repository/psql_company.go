package repository

import (
	"database/sql"
	"errors"
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

// Companys retrieves and returns all companys
func (cpr *CompanyRepositoryImpl) Companys() ([]entity.Company, error) {
	query := "select * from company"
	rows, err := cpr.conn.Query(query)
	if err != nil {
		return nil, errors.New("unable to retrieve companys")
	}
	var companys []entity.Company
	var company entity.Company
	for rows.Next() {
		if err := rows.Scan(&company.ID, &company.CompanyName, &company.Password, &company.Email, &company.Phone, &company.Logo, &company.Short_desc, &company.Detail_info, &company.Address); err != nil {
			return nil, errors.New("unable to retrieve companys")
		}
		companys = append(companys, company)
	}
	return companys, nil
}

// Company return a Company with given id
func (cpr *CompanyRepositoryImpl) Company(id int) (entity.Company, error) {
	query := "select * from companys where id = $1"
	var company entity.Company
	err := cpr.conn.QueryRow(query, id).Scan(company.ID, &company.CompanyName, &company.Password, &company.Email, &company.Phone, &company.Logo, &company.Short_desc, &company.Detail_info, &company.Address)
	if err != nil {
		return company, errors.New("unable to retrieve company")
	}
	return company, nil
}

// UpdateCompany updates a given company
func (cpr *CompanyRepositoryImpl) UpdateCompany(cp entity.Company) error {
	query := "update companys set id=$1, companyname=$2, password=$3, email=$4, phone=$5, logo=$7, shortdesc=$9, detailinfo=$10, address=$11"
	_, err := cpr.conn.Exec(query, cp.ID, cp.CompanyName, cp.Password, cp.Email, cp.Phone, cp.Logo, cp.Short_desc, cp.Detail_info, cp.Address)
	if err != nil {
		return errors.New("unable to update company")
	}
	return nil
}

// DeleteCompany deletes a company with a given id
func (cpr *CompanyRepositoryImpl) DeleteCompany(id int) error {
	query := "delete from companys where id=$1"
	_, err := cpr.conn.Exec(query, id)
	if err != nil {
		return errors.New("unable to delete company")
	}
	return nil
}

// StoreCompany stores new company
func (cpr *CompanyRepositoryImpl) StoreCompany(cp entity.Company) error {
	query := "insert into companys (ID, CompanyName, Password, Email, Phone, Logo, Short_desc, Detail_info, Address) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	_, err := cpr.conn.Exec(query, cp.ID, cp.CompanyName, cp.Password, cp.Email, cp.Phone, cp.Logo, cp.Short_desc, cp.Detail_info, cp.Address)
	if err != nil {
		return errors.New("unable to store company")
	}
	return nil
}
