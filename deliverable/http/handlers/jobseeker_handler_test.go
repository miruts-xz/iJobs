package handlers

import (
	"fmt"
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
	//expected := true
	real := WriteProfileToLocal(file, "miruts.jpg", "miruts")
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
	real := WriteCvToLocal(file, "miruts.pdf", "miruts")
	fmt.Print("data", real)
}
func TestNewJobseekerHandler(t *testing.T) {

}
