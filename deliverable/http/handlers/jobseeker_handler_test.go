package handlers

import (
	"fmt"
	"github.com/miruts/iJobs/util"
	"os"
	"testing"
)

func TestWriteProfileToLocal(t *testing.T) {
	file, err := os.Open("miruts.jpg")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	var data []byte
	_, err = file.Read(data)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	//expected := trued
	real := util.SaveFile(file, "miruts.jpg")
	fmt.Print("data", real)
}
func TestWriteCvToLocal(t *testing.T) {
	file, err := os.Open("miruts.pdf")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	var data []byte
	_, err = file.Read(data)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	real := util.SaveFile(file, "miruts.pdf")
	fmt.Print("data", real)
}
func TestNewJobseekerHandler(t *testing.T) {

}
