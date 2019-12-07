package entity

import "time"

type Job struct {
	ID, CompanyID, CategoryID, RequiredNum int
	Salary                                 float64
	Name, Description, JobTime             string
	Deadline                               time.Time
}
