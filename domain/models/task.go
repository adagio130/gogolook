package models

import "tasks/constants"

type Task struct {
	ID     string           `json:"id"`
	Name   string           `json:"name"`
	Status constants.Status `json:"status"`
}
