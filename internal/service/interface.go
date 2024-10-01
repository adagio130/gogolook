package service

import (
	"context"
	"tasks/domain/entities"
)

type TaskService interface {
	CreateTask(ctx context.Context, param entities.Task) error
	UpdateTask(ctx context.Context, param entities.Task) error
	DeleteTask(ctx context.Context, taskId string) error
	GetTasks(ctx context.Context, param entities.TaskQueryParam) (*entities.Tasks, error)
}
