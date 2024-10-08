package router

import (
	"context"
	"errors"
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"tasks/config"
	"tasks/router/middleware"
	"time"
)

type Server struct {
	port       string
	router     *gin.Engine
	httpServer *http.Server
	logger     *zap.Logger
}

func NewServer(conf *config.Config, logger *zap.Logger) *Server {
	serverConf := conf.Server
	logger.Debug("Starting server", zap.String("port", serverConf.Port), zap.String("mode", serverConf.Mode))
	router := gin.New()
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(gin.Recovery())
	errorMiddleware := middleware.NewResponseMiddleware()
	router.Use(errorMiddleware.GetResponseHandler())
	return &Server{
		port:   serverConf.Port,
		router: router,
		logger: logger,
	}
}

func (s *Server) Run(ctx context.Context, cancel context.CancelFunc, finishChan chan struct{}, attaches ...Attach) {
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.port),
		Handler: s.router,
	}
	go func() {
		<-ctx.Done()
		s.Shutdown(httpServer, ctx)
		finishChan <- struct{}{}
	}()
	for _, a := range attaches {
		a.Attach(s.router)
	}
	s.logger.Info("run http server address success", zap.String("address", httpServer.Addr))
	err := httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.Error("run http server error", zap.Error(err))
		cancel()
	}
}

func (s *Server) Shutdown(httpServer *http.Server, ctx context.Context) {
	s.logger.Info("Shutting down server...")
	if err := httpServer.Shutdown(ctx); err != nil {
		s.logger.Error("http server shutdown error", zap.Error(err))
	}
	s.logger.Info("Server shutdown complete")
}
