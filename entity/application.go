package entity

type Status string

const (
	ACCEPTED Status = "accepted"
	REJECTED Status = "rejected"
	FURTHER  Status = "further"
)

type Response string

const (
	REVIEWED   Status = "reviewed"
	UNREVIEWED string = "unreviewed"
)

type Application struct {
	ID, JobID, JsID int
	Response        Status
	Status          Response
}
