package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/AmirHosein-Gharaati/user-management/internal/adapter/config"
	"github.com/AmirHosein-Gharaati/user-management/internal/adapter/handler"
	"github.com/AmirHosein-Gharaati/user-management/internal/adapter/storage/postgres"
	"github.com/AmirHosein-Gharaati/user-management/internal/adapter/storage/postgres/repository"
	"github.com/AmirHosein-Gharaati/user-management/internal/core/service"
	"github.com/go-chi/chi/v5"
)

func main() {
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		return
	}

	// init database
	database, err := postgres.New(config.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		return
	}
	defer database.DB.Close()

	// migrate
	database.Migrate()

	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router := chi.NewRouter()
	router.Post("/register", userHandler.Register)

	listenAddr := fmt.Sprintf(":%s", config.HTTP.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)

	server := &http.Server{
		Addr:    listenAddr,
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
	}
}
