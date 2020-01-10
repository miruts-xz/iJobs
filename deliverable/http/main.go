package main

import (
	"database/sql"
	_ "database/sql"
	"errors"
	_ "fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

var templ = template.Must(template.ParseGlob("../../ui/template/*"))

func init() {

}
func main() {
	/**
	templates, global database connection and interfaces
	*/
	// Company database connection
	//pqconncmp, errcmp := sql.Open("postgres", "user=company password=company database=ijobs sslmode=disable")
	// Jobseeker database connection
	pqconnjs, errjs := sql.Open("postgres", "user=postgres password=akuadane database=ijobs sslmode=disable")

	//Job repoHandler
	//jobRepoHandler := repository.NewJobRepository(pqconnjs)

	// if errcmp != nil {
	// 	panic(errors.New("unable to connect with database with company account"))
	// }
	// if err := pqconncmp.Ping(); err != nil {
	// 	panic(err)
	// }
	if errjs != nil {
		panic(errors.New("unable to connect with database with jobseeker account"))
	}
	if err := pqconnjs.Ping(); err != nil {
		panic(err)
	}

	fs := http.FileServer(http.Dir("../../ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "welcome", nil)
}
