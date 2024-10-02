package service

import (
	"context"
	"tasks/constants"
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
	err := t.repo.Update(param)
	if err != nil {
		return err
	}
	return nil
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
	result.Tasks = make([]entities.Task, 0)
	for _, task := range tasks {
		result.Tasks = append(result.Tasks, entities.Task{
			ID:      task.ID,
			Name:    task.Name,
			Status:  constants.Status(task.Status),
			Version: task.Version,
		})

	}
	result.Page = param.Offset/param.Size + 1
	result.Size = len(tasks)
	return &result, nil

}
