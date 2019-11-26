package mainRequests

import (
	"github.com/miruts/iJobs/middleware"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	_ = middleware.Tmpl.ExecuteTemplate(w, "index.layout", nil)
}
