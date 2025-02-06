package main

import (
	"log"
	"net/http"
	"telegram-bot/internal/config"
	"telegram-bot/internal/handler"
	"telegram-bot/internal/repository"
	"telegram-bot/internal/service"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig()

	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handlers.InitRoutes(),
	}

	log.Println("Starting server on :8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
}
