package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/miruts/iJobs/entity"
	appsrv "github.com/miruts/iJobs/usecases/application/service"
)

// ApplicationApiHandler represents application api handler
type ApplicationApiHandler struct {
	appService *appsrv.AppService
}

// NewAppApiHandler creates new ApplicationApiHandler
func NewAppApiHandler(appSrv *appsrv.AppService) *ApplicationApiHandler {
	return &ApplicationApiHandler{appService: appSrv}
}

// ApplicationsOnJob retrieves all Application on a given job
func (appHandler *ApplicationApiHandler) ApplicationsOnJob(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("jobId")
	idint, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}

	app, err := appHandler.appService.ApplicationsOnJob(idint)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(app)
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

// Application finds application by id
func (appHandler *ApplicationApiHandler) Application(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	app, err := appHandler.appService.Application(idint)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(app)
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

// UserApplication finds all application given jobseeker id and service
func (appHandler *ApplicationApiHandler) ApplicationsOfJs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	app, err := appHandler.appService.UserApplication(idint)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(app)
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

// Delete Application deletes application with given id
func (appHandler *ApplicationApiHandler) DeleteApp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	app, err := appHandler.appService.DeleteApplication(idint)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(app)
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

// ApplicationForCompany retrieves all job-applications for a given company
func (appHandler *ApplicationApiHandler) AddApplication(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var app entity.Application
	err := json.NewDecoder(r.Body).Decode(&app)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	err = appHandler.appService.Store(&app)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(app)
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
