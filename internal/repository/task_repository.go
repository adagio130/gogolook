package repository

import (
	"database/sql"
	"fmt"
	"tasks/domain/entities"
	"tasks/domain/models"
)

type taskRepository struct {
	conn *sql.DB
}

func NewTaskRepository(conn *sql.DB) TaskRepository {
	return &taskRepository{conn: conn}
}

func (t *taskRepository) Find(id string) (models.Task, error) {
	task := models.Task{}
	t.conn.QueryRow("SELECT id,name,status,version,created_at FROM tasks WHERE id = ?", id).Scan(&task.ID, &task.Name, &task.Status, &task.Version, &task.CreatedAt)
	return task, nil
}

func (t *taskRepository) List(param entities.TaskQueryParam) ([]entities.Task, error) {
	//rows, err := t.conn.Query("SELECT id,name,status,version,created_at FROM tasks SortBy ? OrderBy ? Limit ? Offset ?", param.SortBy, param.Order, param.Size, param.Offset)
	rows, err := t.conn.Query("SELECT id,name,status,version,created_at FROM tasks")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	result := make([]entities.Task, 0)
	for rows.Next() {
		task := entities.Task{}
		err = rows.Scan(&task.ID, &task.Name, &task.Status, &task.Version, &task.CreatedAt)
		fmt.Println("Task: ", task)
		if err != nil {
			return nil, err
		}
		result = append(result, task)
	}
	return result, nil
}

func (t *taskRepository) Create(task entities.Task) error {
	t.conn.Exec("INSERT INTO tasks (id, name, status, version, created_at) VALUES (?, ?, ?, ?, ?)", task.ID, task.Name, task.Status, task.Version, task.CreatedAt)
	record, err := t.Find(task.ID)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Task: ", record)
	return nil
}

func (t *taskRepository) Update(task entities.Task) error {
	record, err := t.Find(task.ID)
	if err != nil {
		return err
	}
	version := record.Version + 1
	_, err = t.conn.Exec("UPDATE tasks SET name = ?, status = ?, version = ? WHERE id = ? and version = ?", task.Name, task.Status, version, task.ID, record.Version)
	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) Delete(id string) error {
	_, err := t.conn.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
