package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"tasks/constants"
	"tasks/domain/entities"
	"tasks/domain/views"
	customError "tasks/errors"
	"tasks/internal/service"
	"time"
)

type taskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) TaskHandler {
	return &taskHandler{
		taskService: taskService,
	}
}

// GetTasks godoc
// @Summary Get tasks
// @Description Get tasks
// @Tags tasks
// @Accept json
// @Produce json
// @Param size query int false "size"
// @Param offset query int false "offset"
// @Success 200 {object} entities.Tasks
// @Router /tasks [get]
func (h *taskHandler) GetTasks(ginCtx *gin.Context) {
	ctx := context.Background()
	query := formatQuery(ginCtx.Query("size"), ginCtx.Query("page"))
	tasks, err := h.taskService.GetTasks(ctx, query)
	if err != nil {
		_ = ginCtx.Error(err)
		return
	}
	ginCtx.JSON(http.StatusOK, tasks)
}

func formatQuery(size, page string) entities.TaskQueryParam {
	defaultSize := 10
	result := entities.TaskQueryParam{
		Size:   defaultSize,
		Offset: 0,
	}
	s, err := strconv.Atoi(size)
	if err == nil && s > 0 {
		result.Size = s
	}
	p, err := strconv.Atoi(page)
	if err == nil && p > 0 {
		result.Offset = (p - 1) * result.Size
	}
	return result
}

// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task with a name and status
// @Tags tasks
// @Accept json
// @Param task body views.CreateTaskReq true "Task information"
// @Success 201
// @Failure 400 {object} error "request is invalid"
// @Failure 500 {object} error "server internal error"
// @Router /tasks [post]
func (h *taskHandler) CreateTask(ginCtx *gin.Context) {
	var req views.CreateTaskReq
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		_ = ginCtx.Error(customError.InvalidRequest.Wrap(err, "should bind json error"))
		return
	}
	ctx := context.Background()
	task := entities.Task{
		ID:        uuid.New().String(),
		Name:      req.Name,
		Status:    constants.Incomplete,
		Version:   0,
		CreatedAt: time.Now().UTC(),
	}
	err := h.taskService.CreateTask(ctx, task)
	if err != nil {
		_ = ginCtx.Error(err)
		return
	}
	ginCtx.AbortWithStatus(http.StatusCreated)
}

// UpdateTask godoc
// @Summary Update task
// @Description Update task
// @Tags tasks
// @Accept json
// @Param id path string true "task id"
// @Param task body views.UpdateTaskReq true "task"
// @Success 204
// @Failure 400 {object} error "request is invalid"
// @Failure 404 {object} error "task not found"
// @Failure 500 {object} error "server internal error"
// @Router /tasks/{id} [put]
func (h *taskHandler) UpdateTask(ginCtx *gin.Context) {
	taskId := ginCtx.Param("id")
	if taskId == "" {
		_ = ginCtx.Error(customError.InvalidRequest.New("task id is required"))
		return
	}
	var req views.UpdateTaskReq
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		_ = ginCtx.Error(customError.InvalidRequest.Wrap(err, "should bind json error"))
		return
	}
	if req.Status != constants.Complete && req.Status != constants.Incomplete {
		_ = ginCtx.Error(customError.InvalidRequest.New("status not supported"))
		return
	}
	ctx := context.Background()

	task := entities.Task{
		ID:     taskId,
		Status: req.Status,
	}
	if req.Name != nil {
		task.Name = *req.Name
	}
	err := h.taskService.UpdateTask(ctx, task)
	if err != nil {
		_ = ginCtx.Error(err)
		return
	}
	ginCtx.AbortWithStatus(http.StatusNoContent)
}

// DeleteTask godoc
// @Summary Delete task
// @Description Delete task
// @Tags tasks
// @Accept json
// @Param id path string true "task id"
// @Success 204
// @Failure 400 {object} error "request is invalid"
// @Failure 404 {object} error "task not found"
// @Failure 500 {object} error "server internal error"
// @Router /tasks/{id} [delete]
func (h *taskHandler) DeleteTask(ginCtx *gin.Context) {
	taskId := ginCtx.Param("id")
	if taskId == "" {
		_ = ginCtx.Error(customError.InvalidRequest.New("task id is required"))
		return
	}
	ctx := context.Background()
	err := h.taskService.DeleteTask(ctx, taskId)

	if err != nil {
		_ = ginCtx.Error(err)
		return
	}
	ginCtx.AbortWithStatus(http.StatusNoContent)
}
