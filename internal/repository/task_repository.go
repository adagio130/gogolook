package repository

import (
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"tasks/domain/entities"
	"tasks/domain/models"
	customError "tasks/errors"
)

type taskRepository struct {
	conn   *sql.DB
	logger *zap.Logger
}

func NewTaskRepository(conn *sql.DB, logger *zap.Logger) TaskRepository {
	return &taskRepository{conn: conn, logger: logger}
}

func (t *taskRepository) Find(id string) (*models.Task, error) {
	task := models.Task{}
	row := t.conn.QueryRow("SELECT id,name,status,version,created_at FROM tasks WHERE id = ?", id)
	err := row.Scan(&task.ID, &task.Name, &task.Status, &task.Version, &task.CreatedAt)
	if err != nil {
		t.logger.Error("Find task error", zap.String("id", id), zap.Error(err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customError.TaskNotFound.Wrap(err, "task not found")
		}
		return nil, err
	}
	return &task, nil
}

func (t *taskRepository) List(param entities.TaskQueryParam) ([]*models.Task, error) {
	rows, err := t.conn.Query("SELECT id,name,status,version FROM tasks LIMIT ? OFFSET ? ", param.Size, param.Offset)
	defer rows.Close()
	if err != nil {
		t.logger.Error("List task error", zap.Any("param", param), zap.Error(err))
		return nil, err
	}
	result := make([]*models.Task, 0)
	for rows.Next() {
		task := models.Task{}
		err = rows.Scan(&task.ID, &task.Name, &task.Status, &task.Version)
		if err != nil {
			t.logger.Error("Scan task error", zap.Any("param", param), zap.Error(err))
			return nil, err
		}
		result = append(result, &task)
	}

	return result, nil
}

func (t *taskRepository) Create(task entities.Task) error {
	stmt, err := t.conn.Prepare("INSERT INTO tasks (id, name, status, version, created_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		t.logger.Error("Prepare insert stmt error", zap.Any("task", task), zap.Error(err))
		return err
	}
	_, err = stmt.Exec(task.ID, task.Name, task.Status, task.Version, task.CreatedAt)
	if err != nil {
		t.logger.Error("Execute insert stmt error", zap.Any("task", task), zap.Error(err))
		return err
	}
	return nil
}

func (t *taskRepository) Update(task entities.Task) error {
	record, err := t.Find(task.ID)
	if err != nil {
		return err
	}
	version := record.Version + 1
	stmt, err := t.conn.Prepare("UPDATE tasks SET name = ?, status = ?, version = ? WHERE id = ? and version = ?")
	if err != nil {
		t.logger.Error("Prepare update stmt error", zap.String("id", task.ID), zap.Error(err))
		return err
	}
	name := task.Name
	if task.Name == "" {
		name = record.Name
	}
	_, err = stmt.Exec(name, task.Status, version, task.ID, record.Version)
	if err != nil {
		t.logger.Error("Execute update stmt error", zap.String("id", task.ID), zap.Error(err))
		return err
	}
	return nil
}

func (t *taskRepository) Delete(id string) error {
	stmt, err := t.conn.Prepare("DELETE FROM tasks WHERE id = ?")
	if err != nil {
		return err
	}
	rows, err := stmt.Exec(id)
	if err != nil {
		t.logger.Error("Execute delete stmt error", zap.String("id", id), zap.Error(err))
		return err
	}
	effectRows, err := rows.RowsAffected()
	if err != nil {
		t.logger.Error("Execute delete stmt error", zap.String("id", id), zap.Error(err))
		return err
	}
	if effectRows == 0 {
		return customError.TaskNotFound.New("task not found")
	}
	return nil
}
