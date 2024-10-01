package views

type GetTaskReq struct {
	ID string `json:"id" uri:"id" binding:"required"`
}

type GetTasksReq struct {
	Size   int    `json:"size" form:"size"`
	Page   int    `json:"page" form:"page"`
	SortBy string `json:"sort" form:"sort"`
	Order  string `json:"order" form:"order"`
}

type CreateTaskReq struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTaskReq struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type DeleteTaskReq struct {
	ID string `json:"id" uri:"id" binding:"required"`
}
