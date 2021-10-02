package main

import (
	"dispatcher/pkg/handler"
	"dispatcher/pkg/repository"
	"dispatcher/pkg/service"
	"dispatcher/pkg/settings"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Env variables loading error: %s", err)
	}

	if err := settings.InitConfig(); err != nil {
		log.Fatalf("Config initialization error: %s", err)
	}

	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatalf("DB initialization error: %s", err)
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(Server)
	if err := srv.Run(settings.Config.Port, handlers.InitRoutes()); err != nil {
		log.Fatalf("Server running error: %s", err)
	}


}
