package main

import (
	_ "database/sql"

	_ "fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func init() {

}

var tmpl = template.Must(template.ParseGlob("ui/template/*.html"))

func main() {

	fs := http.FileServer(http.Dir("ui/asset"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", Welcome)
	http.HandleFunc("/signInUp", signInUp)
	http.HandleFunc("/jobseeker/home", jsHome)
	http.HandleFunc("/jobseeker/appliedJobs", jsAppliedJobs)
	http.HandleFunc("/company/home", compHome)
	http.HandleFunc("/company/postjob", compPostJob)

	http.ListenAndServe(":8181", nil)

}

func Welcome(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "welcome.layout", nil)
}

func signInUp(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "signInUp.layout", nil)
}

func jsHome(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "jobseeker.home.layout", nil)
}

func jsAppliedJobs(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "jobseeker.appliedJobs.layout", nil)
}

func compHome(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "company.home.layout", nil)
}
func compPostJob(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "company.postjob.layout", nil)
}
