package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/miruts/iJobs/entity"
	jssrv "github.com/miruts/iJobs/usecases/jobseeker"
)

type JobseekerHandler struct {
	jsSrv jssrv.JobseekerService
}

func NewJobseekerHandler(jsSrv jssrv.JobseekerService) *JobseekerHandler {
	return &JobseekerHandler{jsSrv: jsSrv}
}

//Jobseekers handles all JobSeekers JSon data
func (jsh *JobseekerHandler) Jobseekers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("I'm here")
	w.Header().Set("Content-Type", "application/json")
	jobseekers, err := jsh.jsSrv.JobSeekers()
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(jobseekers)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	_, err = w.Write(response)
}
func (jsh *JobseekerHandler) Jobseeker(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	jobseeker, err := jsh.jsSrv.JobSeeker(idint)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(jobseeker)
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
func (jsh *JobseekerHandler) UpdateJobseeker(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	_, err = jsh.jsSrv.JobSeeker(idint)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	var jobseeker entity.Jobseeker
	err = json.NewDecoder(r.Body).Decode(&jobseeker)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	js, err := jsh.jsSrv.UpdateJobSeeker(&jobseeker)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(js)
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
func (jsh *JobseekerHandler) DeleteJobseeker(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	_, err = jsh.jsSrv.JobSeeker(idint)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	js, err := jsh.jsSrv.DeleteJobSeeker(idint)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(js)
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
func (jsh *JobseekerHandler) AddJobseeker(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var jobseeker entity.Jobseeker
	err := json.NewDecoder(r.Body).Decode(&jobseeker)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	js, err := jsh.jsSrv.StoreJobSeeker(&jobseeker)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(js)
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
