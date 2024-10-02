package handler

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"tasks/constants"
	"tasks/domain/entities"
	"testing"
)

// MockTaskService 模擬 TaskService
type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) UpdateTask(ctx context.Context, param entities.Task) error {
	args := m.Called(ctx, param)
	return args.Error(0)
}

func (m *MockTaskService) DeleteTask(ctx context.Context, taskId string) error {
	args := m.Called(ctx, taskId)
	return args.Error(0)
}

func (m *MockTaskService) GetTasks(ctx context.Context, param entities.TaskQueryParam) (*entities.Tasks, error) {
	args := m.Called(ctx, param)
	return args.Get(0).(*entities.Tasks), args.Error(1)
}

func (m *MockTaskService) CreateTask(ctx context.Context, task entities.Task) error {
	args := m.Called(ctx, task)
	return args.Error(0)
}

func Test_taskHandler_GetTasks(t *testing.T) {

	tasks := &entities.Tasks{
		Tasks: []entities.Task{
			{
				ID:     "task-1",
				Name:   "Task 1",
				Status: constants.Incomplete,
			},
		},
		Page: 1,
		Size: 1,
	}
	t.Run("successful task retrieval", func(t *testing.T) {
		mockTaskService := new(MockTaskService)
		h := &taskHandler{
			taskService: mockTaskService,
		}
		mockTaskService.On("GetTasks", mock.Anything, mock.Anything).Return(tasks, nil)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest("GET", "/tasks?size=1&page=1", nil)
		c.Request = req

		h.GetTasks(c)

		assert.Equal(t, http.StatusOK, w.Code)
		expected := `{"tasks":[{"id":"task-1","name":"Task 1","status":0}],"page":1,"size":1}`
		assert.JSONEq(t, expected, w.Body.String())

		mockTaskService.AssertExpectations(t)
	})
}

func Test_taskHandler_CreateTask(t *testing.T) {
	t.Run("successful task creation", func(t *testing.T) {
		mockTaskService := new(MockTaskService)
		h := &taskHandler{
			taskService: mockTaskService,
		}
		mockTaskService.On("CreateTask", mock.Anything, mock.Anything).Return(nil)

		body := `{"name": "Test Task"}`
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		c.Request = req

		h.CreateTask(c)

		assert.Equal(t, http.StatusCreated, w.Code)

		mockTaskService.AssertExpectations(t)
	})
}

func Test_taskHandler_UpdateTask(t *testing.T) {
	t.Run("successful task update", func(t *testing.T) {
		mockTaskService := new(MockTaskService)
		h := &taskHandler{
			taskService: mockTaskService,
		}

		mockTaskService.On("UpdateTask", mock.Anything, mock.Anything).Return(nil)

		body := `{"name":"task-test","status": 1}`
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest("PUT", "/tasks/task-1", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		c.Request = req
		c.Params = gin.Params{gin.Param{Key: "id", Value: "task-1"}}

		h.UpdateTask(c)

		assert.Equal(t, http.StatusNoContent, w.Code)

		mockTaskService.AssertExpectations(t)
	})
}

func Test_taskHandler_DeleteTask(t *testing.T) {
	t.Run("successful task deletion", func(t *testing.T) {
		mockTaskService := new(MockTaskService)
		h := &taskHandler{
			taskService: mockTaskService,
		}
		mockTaskService.On("DeleteTask", mock.Anything, "task-1").Return(nil)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest("DELETE", "/tasks/task-1", nil)
		c.Request = req
		c.Params = gin.Params{gin.Param{Key: "id", Value: "task-1"}}

		h.DeleteTask(c)

		assert.Equal(t, http.StatusNoContent, w.Code)

		mockTaskService.AssertExpectations(t)
	})
}
