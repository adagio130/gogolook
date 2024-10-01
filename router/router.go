package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"tasks/internal/handler"
)

type Attach interface {
	Attach(router *gin.Engine)
}

type baseRouter struct {
	rootPath    string
	middlewares []gin.HandlerFunc
}

func (r *baseRouter) Attach(router *gin.Engine) {
	group := router.Group(r.rootPath, r.middlewares...)
	group.GET("/", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})
	group.GET("/health", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})
}

type taskRouter struct {
	rootPath    string
	middlewares []gin.HandlerFunc
	handlers    handler.TaskHandler
}

func NewTaskRouter(taskHandler handler.TaskHandler, middleware []gin.HandlerFunc) Attach {
	return &taskRouter{
		rootPath:    "/tasks",
		middlewares: middleware,
		handlers:    taskHandler,
	}
}

func (r *taskRouter) Attach(router *gin.Engine) {
	group := router.Group(r.rootPath, r.middlewares...)
	group.GET("/", r.handlers.GetTasks)
	group.POST("/", r.handlers.CreateTask)
	//group.GET("/:id", r.handlers.GetTask)
	group.PUT("/:id", r.handlers.UpdateTask)
	group.DELETE("/:id", r.handlers.DeleteTask)
}

type swaggerRouter struct {
	rootPath string
	handler  gin.HandlerFunc
}

func NewSwaggerRouter() Attach {
	swaggerHandler := ginSwagger.WrapHandler(swaggerfiles.Handler)
	return &swaggerRouter{
		rootPath: "/swagger",
		handler:  swaggerHandler,
	}
}

func (s swaggerRouter) Attach(router *gin.Engine) {
	router.GET("/swagger/*any", s.handler)
}

func NewBaseRouter() Attach {
	return &baseRouter{
		rootPath: "/",
	}
}
