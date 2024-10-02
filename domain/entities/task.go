package entities

import (
	"tasks/constants"
	"time"
)

type Task struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	Status    constants.Status `json:"status"`
	Version   int              `json:"-"`
	CreatedAt time.Time        `json:"-"`
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
	Size  int    `json:"size"`
	Page  int    `json:"page"`
}

type TaskQueryParam struct {
	Size   int    `json:"size"`
	Offset int    `json:"offset"`
	SortBy string `json:"order_by"`
	Order  string `json:"sort"`
}
