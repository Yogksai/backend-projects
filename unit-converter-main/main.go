package main

import (
	"log/slog"

	"github.com/fentezi/unit-converter/handler"
	"github.com/fentezi/unit-converter/router"
	"github.com/fentezi/unit-converter/service"
)

func main() {
	// Init logger
	log := slog.Default()

	// Init service
	service := service.NewService(log)

	//Init handler
	handler := handler.NewHandler(log, service)

	// Init router
	router := router.NewRouter(log, handler)

	router.StartServer()

}
