package service

type TaskService interface {
	CreateTask()
	GetTask()
	UpdateTask()
	DeleteTask()
	GetTasks()
}
