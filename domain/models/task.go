package models

import "tasks/constants"

type Task struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	Status    constants.Status `json:"status"`
	Version   int              `json:"version"`
	CreatedAt string           `json:"created_at"`
}

type TaskQueryParam struct {
}
