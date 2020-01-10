package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/akuadane/iJobs/usecases/application/service"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/iJobs/entity"
)

type ApplicationApiHandler struct {
	appService service.AppService
}

func NewJobApiHandler(appSrv service.AppService) *ApplicationApiHandler {
	return &ApplicationApiHandler{appService: appSrv}
}

func (appHandler *ApplicationApiHandler) ApplicationsOnJob(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")
	idint, err := strconv.Atoi(id)
	apps, err := appHandler.appService.ApplicationsOnJob(idint)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(apps)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	_, err = w.Write(response)

}

func (appHandler *ApplicationApiHandler) ApplicationsOnJob(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	app, err := jobHander.jobService.Job(idint)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(job)
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

func (jobHander *JobApiHandler) UpdateJob(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	_, err = jobHander.jobService.Job(idint)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	job := entity.Job{}
	err = json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	err = jobHander.jobService.UpdateJob(job)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(job)
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

func (jobHander *JobApiHandler) DeleteJob(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	_, err = jobHander.jobService.Job(idint)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	err = jobHander.jobService.DeleteJob(idint)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	return
}

func (jobHander *JobApiHandler) AddJob(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var job entity.Job
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	err = jobHander.jobService.StoreJob(job)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(job)
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
