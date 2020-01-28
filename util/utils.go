package util

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
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
