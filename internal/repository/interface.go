package repository

import "tasks/domain/entities"

type TaskRepository interface {
	Find(id string) (entities.Task, error)
	FindAll() ([]entities.Task, error)
	Create(task entities.Task) error
	Update(task entities.Task) error
	Delete(id string) error
}
