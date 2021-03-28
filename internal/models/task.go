package models

type Status int

const (
	New Status = iota
	Assigned
	InProgress
	Completed
	Failed
	Abandoned
)

type Task struct {
	Model
	Title     string `json:"title"`
	Story     string `json:"story"`
	Status    uint8  `json:"status" default:"0"`
}
