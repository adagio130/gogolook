package service

import (
	"context"
	"tasks/domain/entities"
	"tasks/internal/repository"
	"testing"
)

func TestTaskServiceCreateTask(t1 *testing.T) {
	type fields struct {
		repo repository.TaskRepository
	}
	type args struct {
		ctx   context.Context
		param entities.Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &taskService{
				repo: tt.fields.repo,
			}
			if err := t.CreateTask(tt.args.ctx, tt.args.param); (err != nil) != tt.wantErr {
				t1.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
