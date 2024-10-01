package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"tasks/config"
	"tasks/internal/handler"
	"tasks/router"
)

func main() {
	ctx := context.Background()
	svcCtx, cancel := context.WithCancel(ctx)

	v := viper.New()
	v.SetConfigFile("config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	var conf config.Config
	if err := v.Unmarshal(&conf); err != nil {
		panic(fmt.Errorf("umarshal config error: %s \n", err))
	}
	v.WatchConfig()
	logger, _ := zap.NewProduction()
	taskHandler := handler.NewTaskHandler()
	middleware := []gin.HandlerFunc{}
	attaches := []router.Attach{
		router.NewBaseRouter(),
		router.NewTaskRouter(taskHandler, middleware),
		router.NewSwaggerRouter(),
	}
	server := router.NewServer(&conf, logger)
	finishChan := make(chan struct{})
	defer func() {
		close(finishChan)
		logger.Info("Server shutdown")
	}()
	go func() {
		signalChan := make(chan os.Signal, 2)
		defer signal.Stop(signalChan)

		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		<-signalChan
		logger.Info("Received shutdown signal")
		cancel()
	}()
	go server.Run(svcCtx, cancel, finishChan, attaches...)
	<-finishChan
	return
}
