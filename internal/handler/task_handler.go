package handler

import (
	"github.com/gin-gonic/gin"
	"tasks/internal/service"
)

type taskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) TaskHandler {
	return &taskHandler{
		taskService: taskService,
	}
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
