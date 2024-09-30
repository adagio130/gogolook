package handler

import "context"

type TaskHandler interface {
	GetTasks(ctx context.Context, size int, page int) (interface{}, error)
}
