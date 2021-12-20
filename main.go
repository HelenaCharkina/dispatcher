package main

import (
	"context"
	dispatcher2 "dispatcher/pkg/dispatcher"
	"dispatcher/pkg/handler"
	"dispatcher/pkg/repository"
	"dispatcher/pkg/service"
	"dispatcher/pkg/settings"
	"dispatcher/types"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var wsChan chan types.WsChanMessage

func main() {
	ctx := context.TODO()
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Env variables loading error: %s", err)
	}

	if err := settings.InitConfig(); err != nil {
		logrus.Fatalf("Config initialization error: %s", err)
	}

	db, err := repository.NewClickhouseDB()
	if err != nil {
		logrus.Fatalf("DB initialization error: %s", err)
	}

	cache, err := repository.NewRedisCache(ctx)
	if err != nil {
		logrus.Fatalf("Cache initialization error: %s", err)
	}

	wsChan = make(chan types.WsChanMessage, settings.Config.QueueCap)

	repo := repository.NewRepository(db, cache)
	services := service.NewService(repo)
	dispatcher := dispatcher2.NewDispatcher(services, wsChan)
	if err = dispatcher.Run(); err != nil {
		logrus.Fatalf("Dispatcher running error: %s", err)
	}
	handlers := handler.NewHandler(services, wsChan, dispatcher.CmdChan)

	srv := new(Server)
	go func() {
		if err := srv.Run(settings.Config.Port, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Server running error: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err = srv.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server shutting down error: %s", err)
	}
	if err = db.Close(); err != nil {
		logrus.Fatalf("DB connection close error: %s", err)
	}
	if err = cache.Close(); err != nil {
		logrus.Fatalf("Cache connection close error: %s", err)
	}
}
