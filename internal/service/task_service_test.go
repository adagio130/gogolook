package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"tasks/domain/entities"
	"tasks/domain/models"
	"testing"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Find(id string) (*models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockTaskRepository) Create(task entities.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) Update(task entities.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) Delete(taskID string) error {
	args := m.Called(taskID)
	return args.Error(0)
}

func (m *MockTaskRepository) List(param entities.TaskQueryParam) ([]*models.Task, error) {
	args := m.Called(param)
	return args.Get(0).([]*models.Task), args.Error(1)
}

func Test_taskService_CreateTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewTaskService(mockRepo)

	task := entities.Task{
		ID:   "task-123",
		Name: "Test Task",
	}

	t.Run("successfully create task", func(t *testing.T) {
		mockRepo.On("Create", task).Return(nil)

		err := service.CreateTask(context.Background(), task)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func Test_taskService_UpdateTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewTaskService(mockRepo)

	task := entities.Task{
		ID:   "task-123",
		Name: "Test Task",
	}

	t.Run("successfully update task", func(t *testing.T) {
		mockRepo.On("Update", task).Return(nil)

		err := service.UpdateTask(context.Background(), task)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func Test_taskService_DeleteTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewTaskService(mockRepo)

	taskID := "task-123"

	t.Run("successfully delete task", func(t *testing.T) {
		mockRepo.On("Delete", taskID).Return(nil)

		err := service.DeleteTask(context.Background(), taskID)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func Test_taskService_GetTasks(t *testing.T) {
	param := entities.TaskQueryParam{
		Offset: 0,
		Size:   10,
	}

	mockTasks := []*models.Task{
		{
			ID:     "task-123",
			Name:   "Test Task 1",
			Status: 0,
		},
		{
			ID:     "task-124",
			Name:   "Test Task 2",
			Status: 1,
		},
	}

	t.Run("successfully get tasks", func(t *testing.T) {
		mockRepo := new(MockTaskRepository)
		service := NewTaskService(mockRepo)
		mockRepo.On("List", param).Return(mockTasks, nil)

		tasks, err := service.GetTasks(context.Background(), param)

		assert.NoError(t, err)
		assert.Equal(t, 2, len(tasks.Tasks))
		mockRepo.AssertExpectations(t)
	})

	t.Run("fail to get tasks", func(t *testing.T) {
		mockRepo := new(MockTaskRepository)
		service := NewTaskService(mockRepo)
		mockRepo.On("List", param).Return([]*models.Task{}, errors.New("failed to list tasks"))

		tasks, err := service.GetTasks(context.Background(), param)

		assert.Error(t, err)
		assert.Nil(t, tasks)
		mockRepo.AssertExpectations(t)
	})
}
