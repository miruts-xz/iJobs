package util

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

// SaveFile saves multipart file on the given path
func SaveFile(file multipart.File, path string) bool {
	var data []byte
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return false
	}
	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return false
	}
	return true
}
func Authenticate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
