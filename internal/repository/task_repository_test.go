package repository

import (
	"database/sql"
	"errors"
	"tasks/constants"
	"tasks/domain/entities"
	customError "tasks/errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func Test_taskRepository_Find(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	logger := zap.NewNop()

	repo := NewTaskRepository(db, logger)

	t.Run("successfully find task", func(t *testing.T) {
		columns := []string{"id", "name", "status", "version", "created_at"}
		mock.ExpectQuery("SELECT id,name,status,version,created_at FROM tasks WHERE id = ?").
			WithArgs("task-123").
			WillReturnRows(sqlmock.NewRows(columns).AddRow("task-123", "Test Task", 0, 1, time.Now()))

		task, err := repo.Find("task-123")
		assert.NoError(t, err)
		assert.Equal(t, "task-123", task.ID)
		assert.Equal(t, "Test Task", task.Name)

		mock.ExpectationsWereMet()
	})

	t.Run("task not found", func(t *testing.T) {
		mock.ExpectQuery("SELECT id,name,status,version,created_at FROM tasks WHERE id = ?").
			WithArgs("task-123").
			WillReturnError(sql.ErrNoRows)

		task, err := repo.Find("task-123")
		assert.Nil(t, task)
		assert.True(t, errors.Is(err, customError.TaskNotFound))

		mock.ExpectationsWereMet()
	})

	t.Run("find task error", func(t *testing.T) {
		mock.ExpectQuery("SELECT id,name,status,version,created_at FROM tasks WHERE id = ?").
			WithArgs("task-123").
			WillReturnError(errors.New("db error"))

		task, err := repo.Find("task-123")
		assert.Nil(t, task)
		assert.EqualError(t, err, "db error")

		mock.ExpectationsWereMet()
	})
}

func Test_taskRepository_List(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	logger := zap.NewNop() // 使用空的 logger
	repo := NewTaskRepository(db, logger)

	columns := []string{"id", "name", "status", "version"}

	t.Run("successfully list tasks", func(t *testing.T) {
		mock.ExpectQuery("SELECT id,name,status,version FROM tasks LIMIT ? OFFSET ?").
			WithArgs(10, 0).
			WillReturnRows(sqlmock.NewRows(columns).AddRow("task-123", "Test Task", 0, 1))

		param := entities.TaskQueryParam{
			Size:   10,
			Offset: 0,
		}

		tasks, err := repo.List(param)
		assert.NoError(t, err)
		assert.Len(t, tasks, 1)
		assert.Equal(t, "task-123", tasks[0].ID)

		mock.ExpectationsWereMet()
	})

	t.Run("list tasks error", func(t *testing.T) {
		mock.ExpectQuery("*").
			WithArgs(10, 0).
			WillReturnError(errors.New("db error"))

		param := entities.TaskQueryParam{
			Size:   10,
			Offset: 0,
		}

		tasks, err := repo.List(param)
		assert.Nil(t, tasks)
		assert.EqualError(t, err, "db error")

		mock.ExpectationsWereMet()
	})
}

func Test_taskRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	logger := zap.NewNop()
	repo := NewTaskRepository(db, logger)

	t.Run("successfully create task", func(t *testing.T) {
		mock.ExpectPrepare("INSERT INTO tasks").
			ExpectExec().
			WithArgs("task-123", "Test Task", 0, 0, sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))

		task := entities.Task{
			ID:        "task-123",
			Name:      "Test Task",
			Status:    constants.Incomplete,
			Version:   0,
			CreatedAt: time.Now().UTC(),
		}

		err := repo.Create(task)
		assert.NoError(t, err)

		mock.ExpectationsWereMet()
	})

	t.Run("create task error", func(t *testing.T) {
		mock.ExpectPrepare("INSERT INTO tasks").
			ExpectExec().
			WithArgs("task-123", "Test Task", 0, 0, sqlmock.AnyArg()).
			WillReturnError(errors.New("db error"))

		task := entities.Task{
			ID:        "task-123",
			Name:      "Test Task",
			Status:    constants.Incomplete,
			Version:   0,
			CreatedAt: time.Now().UTC(),
		}

		err := repo.Create(task)
		assert.EqualError(t, err, "db error")

		mock.ExpectationsWereMet()
	})
}

//func Test_taskRepository_Update(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	assert.NoError(t, err)
//	defer db.Close()
//
//	logger := zap.NewNop()
//	repo := NewTaskRepository(db, logger)
//
//	t.Run("successfully update task", func(t *testing.T) {
//		columns := []string{"id", "name", "status", "version", "created_at"}
//		mock.ExpectQuery("SELECT id,name,status,version,created_at FROM tasks WHERE id = ?").
//			WithArgs("task-123").
//			WillReturnRows(sqlmock.NewRows(columns).AddRow("task-123", "Test Task", 0, 1, time.Now()))
//
//		mock.ExpectPrepare("UPDATE tasks SET name = ?, status = ?, version = ? WHERE id = ? and version = ?").
//			ExpectExec().
//			WithArgs("Updated Task", 0, 2, "task-123", 1).
//			WillReturnResult(sqlmock.NewResult(1, 1))
//
//		task := entities.Task{
//			ID:     "task-123",
//			Name:   "Updated Task",
//			Status: 0,
//		}
//
//		err := repo.Update(task)
//		assert.NoError(t, err)
//
//		mock.ExpectationsWereMet()
//	})
//
//	t.Run("update task error", func(t *testing.T) {
//		mock.ExpectQuery("SELECT id,name,status,version,created_at FROM tasks WHERE id = ?").
//			WithArgs("task-123").
//			WillReturnError(errors.New("db error"))
//
//		task := entities.Task{
//			ID:     "task-123",
//			Name:   "Updated Task",
//			Status: 0,
//		}
//
//		err := repo.Update(task)
//		assert.EqualError(t, err, "db error")
//
//		mock.ExpectationsWereMet()
//	})
//}

func Test_taskRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	logger := zap.NewNop() // 使用空的 logger
	repo := NewTaskRepository(db, logger)

	t.Run("successfully delete task", func(t *testing.T) {
		mock.ExpectPrepare("DELETE FROM tasks WHERE id = ?").
			ExpectExec().
			WithArgs("task-123").
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Delete("task-123")
		assert.NoError(t, err)

		mock.ExpectationsWereMet()
	})

	t.Run("delete task error", func(t *testing.T) {
		mock.ExpectPrepare("DELETE FROM tasks WHERE id = ?").
			ExpectExec().
			WithArgs("task-123").
			WillReturnError(errors.New("db error"))

		err := repo.Delete("task-123")
		assert.EqualError(t, err, "db error")

		mock.ExpectationsWereMet()
	})
}
