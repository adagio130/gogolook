package entities

import (
	"tasks/constants"
	"time"
)

type Task struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	Status    constants.Status `json:"status"`
	Version   int              `json:"version"`
	CreatedAt time.Time        `json:"created_at"`
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
}

type TaskQueryParam struct {
	Size   int    `json:"size"`
	Offset int    `json:"offset"`
	SortBy string `json:"order_by"`
	Order  string `json:"sort"`
}
