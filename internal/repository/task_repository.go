package repository

import (
	"database/sql"
	"tasks/domain/entities"
	"tasks/domain/models"
)

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (t *taskRepository) Find(id string) (entities.Task, error) {
	return entities.Task{}, nil
}

func (t *taskRepository) List() ([]entities.Task, error) {
	return []entities.Task{}, nil
}

func (t *taskRepository) Create(task models.Task) error {
	return nil
}

func (t *taskRepository) Update(task entities.Task) error {
	return nil
}

func (t *taskRepository) Delete(id string) error {
	return nil
}

func (t *taskRepository) Migrate() error {
	return nil
}
