package mainRequest

import (
	"encoding/json"
	"fmt"
	"github.com/miruts/iJobs/entity"
	"html/template"
	"io/ioutil"
	"net/http"
)

var Tmpl = template.Must(template.ParseGlob("deliverable/template/*.html"))

// Passing data to template using the rest api provided
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(Tmpl)
	resp, err := http.Get("http://api.localhost:8080/users")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	var user entity.JobSeeker
	_ = json.Unmarshal(body, &user)

	_ = Tmpl.ExecuteTemplate(w, "index.layout", user)
}