package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"tasks/config"
	"tasks/internal/handler"
	"tasks/internal/repository"
	"tasks/internal/service"
	"tasks/router"
)

func main() {
	ctx := context.Background()
	svcCtx, cancel := context.WithCancel(ctx)
	logger, _ := zap.NewProduction()
	conf := initConfig()
	db := initStorage(conf)
	finishChan := make(chan struct{})
	attaches, server := initServer(db, logger, conf)
	defer func() {
		db.Close()
		close(finishChan)
		logger.Info("Server shutdown")
	}()
	go func() {
		signalChan := make(chan os.Signal, 1)
		defer signal.Stop(signalChan)

		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		<-signalChan
		logger.Info("Received shutdown signal")
		cancel()
	}()
	server.Run(svcCtx, cancel, finishChan, attaches...)
	<-finishChan
	return
}

func initConfig() config.Config {
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
	return conf
}

func initServer(db *sql.DB, logger *zap.Logger, conf config.Config) ([]router.Attach, *router.Server) {
	taskRepo := repository.NewTaskRepository(db, logger)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)
	attaches := []router.Attach{
		router.NewBaseRouter(),
		router.NewTaskRouter(taskHandler, []gin.HandlerFunc{}),
	}
	server := router.NewServer(&conf, logger)

	return attaches, server
}

func initStorage(conf config.Config) *sql.DB {
	db, err := sql.Open(conf.DB.Driver, conf.DB.Dsn)
	if err != nil {
		panic(fmt.Errorf("Fatal error db file: %s \n", err))
	}
	db.SetMaxOpenConns(conf.DB.MaxOpen)
	if err = db.Ping(); err != nil {
		panic(fmt.Errorf("ping db error: %s \n", err))
	}
	if err = checkTables(db); err != nil {
		panic(fmt.Errorf("check tables error: %s \n", err))
	}
	return db
}

func checkTables(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS tasks 
		(id TEXT PRIMARY KEY NOT NULL, 
		name TEXT NOT NULL, 
		status INTEGER NOT NULL, 
		version INTEGER, 
		created_at TEXT
		)
	`)
	if err != nil {
		return err
	}
	return nil
}
