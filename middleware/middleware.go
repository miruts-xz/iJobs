package middleware

import (
	"github.com/miruts/iJobs/deliverable/http/apiRequest"
	"github.com/miruts/iJobs/deliverable/http/mainRequest"
	"net/http"
)

const (
	domain string = "localhost"
	apisd  string = "api"
	authsd string = "auth"
)

func Run() {
	fs := http.FileServer(http.Dir("deliverable/asset"))
	http.HandleFunc(apisd+"."+domain+"/users", apiRequest.GetUsers)
	http.Handle("/deliverable/asset/", http.StripPrefix("/deliverable/asset/", fs))
	http.HandleFunc("/", mainRequest.Index)
	_ = http.ListenAndServe(":8080", nil)
}
