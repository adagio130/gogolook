package entities

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status Status `json:"status"`
}

type Status int

const (
	Incomplete Status = iota
	Complete
)
