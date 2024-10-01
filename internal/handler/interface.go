package handler

import "github.com/gin-gonic/gin"

type TaskHandler interface {
	GetTasks(ginCtx *gin.Context)
	CreateTask(ginCtx *gin.Context)
	GetTask(ginCtx *gin.Context)
	UpdateTask(ginCtx *gin.Context)
	DeleteTask(ginCtx *gin.Context)
}
