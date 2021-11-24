package main

import (
	"context"
	"dispatcher/pkg/handler"
	"dispatcher/pkg/repository"
	"dispatcher/pkg/service"
	"dispatcher/pkg/settings"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

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

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

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
	//if err = client.Disconnect(ctx); err != nil {
	//	logrus.Fatalf("DB connection close error: %s", err)
	//}

}
