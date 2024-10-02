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
// @Param sort_by query string false "sort_by"
// @Param order query string false "order"
// @Success 200 {array} Task
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /tasks [get]
func (h *taskHandler) GetTasks(ginCtx *gin.Context) {
	ctx := context.Background()
	query := formatQuery(ginCtx.Query("size"), ginCtx.Query("page"), ginCtx.Query("sort_by"), ginCtx.Query("order"))
	tasks, err := h.taskService.GetTasks(ctx, query)
	if err != nil {
		_ = ginCtx.Error(err)
		return
	}
	ginCtx.JSON(http.StatusOK, tasks)
}

func formatQuery(size, page, sort, order string) entities.TaskQueryParam {
	defaultSize := 10
	result := entities.TaskQueryParam{
		Size:   defaultSize,
		Offset: 0,
		SortBy: "created_at",
		Order:  "desc",
	}
	s, err := strconv.Atoi(size)
	if err == nil && s > 0 {
		result.Size = s
	}
	p, err := strconv.Atoi(page)
	if err == nil && p > 0 {
		result.Offset = (p - 1) * result.Size
	}
	if sort != "" {
		result.SortBy = sort
	}
	if order == "asc" || order == "desc" {
		result.Order = order
	}
	return result
}

// CreateTask godoc
// @Summary Create task
// @Description Create task
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body CreateTaskReq true "task"
// @Success 201
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
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
// @Produce json
// @Param id path string true "task id"
// @Param task body UpdateTaskReq true "task"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
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
		_ = ginCtx.Error(customError.InvalidRequest.New("should bind json error"))
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
// @Produce json
// @Param id path string true "task id"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
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
