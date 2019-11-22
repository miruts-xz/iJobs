package middleware

import (
	"github.com/miruts/iJobs/handler/mainRequests"
	"net/http"
)

func Run() {
	fs := http.FileServer(http.Dir("deliverable/asset"))
	http.Handle("/deliverable/asset/", http.StripPrefix("/deliverable/asset/", fs))
	http.HandleFunc("/", mainRequests.Index)
	_ = http.ListenAndServe(":9090", nil)
}
