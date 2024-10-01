package repository

import (
	"tasks/domain/entities"
	"tasks/domain/models"
)

type TaskRepository interface {
	Find(id string) (entities.Task, error)
	List() ([]entities.Task, error)
	Create(task models.Task) error
	Update(task entities.Task) error
	Delete(id string) error
}
