package main

import (
	"dispatcher/pkg/handler"
	"dispatcher/pkg/repository"
	"dispatcher/pkg/service"
	"dispatcher/pkg/settings"
	"fmt"
	"log"
)

func main() {

	if err := settings.InitConfig(); err != nil {
		log.Fatalf("Config initialization error: %s", err)
	}
	fmt.Println("port ", settings.Config.Port)
	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(Server)
	if err := srv.Run(settings.Config.Port, handlers.InitRoutes()); err != nil {
		log.Fatalf("Server running error: %s", err)
	}
}
