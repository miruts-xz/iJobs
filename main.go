package main

import (
	"github.com/miruts/iJobs/middleware"
	"html/template"
)

var Tmpl = template.Must(template.ParseGlob("deliverable/template/*"))

func main() {
	middleware.Run()
}
