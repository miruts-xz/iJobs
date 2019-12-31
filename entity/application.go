package entity

type Response string

const (
	ACCEPTED Response = "accepted"
	REJECTED Response = "rejected"
	FURTHER  Response = "further"
)

type Status string

const (
	REVIEWED   Status = "reviewed"
	UNREVIEWED Status = "unreviewed"
)

type Application struct {
	ID, JobID, JsID int
	Response        string // will be changed to type Response type after testing
	Status          string // will be changed to type Stattus type after testing
}
