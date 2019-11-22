package mainRequests

import (
	"html/template"
	"net/http"
)

var Tmpl = template.Must(template.ParseGlob("deliverable/template/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	_ = Tmpl.ExecuteTemplate(w, "index.layout", nil)
}
