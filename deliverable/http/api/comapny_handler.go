package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/miruts/iJobs/entity"
	cmpsrv "github.com/miruts/iJobs/usecases/company/service"
)

// CompanyHandler handles company related http requests
type CompanyHandler struct {
	cmpSrv *cmpsrv.CompanyServiceImpl
}

// NewCompanyHandler returns new CompanyHandler object
func NewCompanyHandler(cmpSrv *cmpsrv.CompanyServiceImpl) *CompanyHandler {
	return &CompanyHandler{cmpSrv: cmpSrv}
}

// Companies retrieves and returns all companies
func (cmph *CompanyHandler) Companies(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	companies, err := cmph.cmpSrv.Companies()
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(companies)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	_, err = w.Write(response)
}

// Company return a Company with given id
func (cmph *CompanyHandler) Company(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	company, err := cmph.cmpSrv.Company(idint)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(company)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	_, err = w.Write(response)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
}

// UpdateCompany updates a given company
func (cmph *CompanyHandler) UpdateCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	_, err = cmph.cmpSrv.Company(idint)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	var company entity.Company
	err = json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	cmp, err := cmph.cmpSrv.UpdateCompany(&company)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(cmp)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	_, err = w.Write(response)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
}

// DeleteCompany deletes a company with a given id
func (cmph *CompanyHandler) DeleteCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	_, err = cmph.cmpSrv.Company(idint)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	cmp, err := cmph.cmpSrv.DeleteCompany(idint)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(cmp)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
}

// StoreCompany stores new company
func (cmph *CompanyHandler) AddCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var company entity.Company
	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	cmp, err := cmph.cmpSrv.StoreCompany(&company)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(cmp)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(response)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
}
