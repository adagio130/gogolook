package views

import "tasks/constants"

type GetTasksReq struct {
	Size   *int    `json:"size"`
	Page   *int    `json:"page" `
	SortBy *string `json:"sort"`
	Order  *string `json:"order"`
}

type CreateTaskReq struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTaskReq struct {
	ID     string           `json:"id" uri:"id"`
	Name   *string          `json:"name"`
	Status constants.Status `json:"status" enum:"0,1"`
}

type DeleteTaskReq struct {
	ID string `json:"id" uri:"id" binding:"required"`
}
