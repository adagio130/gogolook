package entities

import "tasks/constants"

type Task struct {
	ID      int              `json:"id"`
	Name    string           `json:"name"`
	Status  constants.Status `json:"status"`
	Version int              `json:"version"`
}
