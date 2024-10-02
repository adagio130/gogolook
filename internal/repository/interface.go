package repository

import (
	"tasks/domain/entities"
	"tasks/domain/models"
)

type TaskRepository interface {
	Find(id string) (*models.Task, error)
	List(param entities.TaskQueryParam) ([]*models.Task, error)
	Create(task entities.Task) error
	Update(task entities.Task) error
	Delete(id string) error
}
