package handler

import "github.com/gin-gonic/gin"

type taskHandler struct {
}

func NewTaskHandler() TaskHandler {
	return &taskHandler{}
}

func (h *taskHandler) GetTasks(ginCtx *gin.Context) {
	panic("implement me")
}

func (h *taskHandler) GetTask(ginCtx *gin.Context) {
	panic("implement me")
}

func (h *taskHandler) CreateTask(ginCtx *gin.Context) {
	panic("implement me")
}

func (h *taskHandler) UpdateTask(ginCtx *gin.Context) {
	panic("implement me")
}

func (h *taskHandler) DeleteTask(ginCtx *gin.Context) {
	panic("implement me")
}
