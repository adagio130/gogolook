package service

import (
	"context"
	"tasks/domain/entities"
	"tasks/internal/repository"
)

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (t *taskService) CreateTask(ctx context.Context, param entities.Task) error {
	return t.repo.Create(param)
}

func (t *taskService) UpdateTask(ctx context.Context, param entities.Task) error {
	return t.repo.Update(param)
}

func (t *taskService) DeleteTask(ctx context.Context, taskId string) error {
	return t.repo.Delete(taskId)
}

func (t *taskService) GetTasks(ctx context.Context, param entities.TaskQueryParam) (*entities.Tasks, error) {
	var result entities.Tasks
	tasks, err := t.repo.List(param)
	if err != nil {
		return nil, err
	}
	result.Tasks = tasks
	result.Limit = len(tasks)
	return &result, nil

}
