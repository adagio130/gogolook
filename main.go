package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"tasks/config"
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
	server := router.NewServer(&conf, logger)
	//server.SetupRouter(routeHandler *router.Router)
	go server.Run(svcCtx, cancel)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	logger.Info("Received shutdown signal")
	cancel()
	defer func() {
		signal.Stop(signalChan)
	}()
	logger.Info("Finish")
}
