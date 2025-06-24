package main

import (
	"log"
	"net/http"

	"github.com/xjncx/taskmanager/internal/api"
	repository "github.com/xjncx/taskmanager/internal/repository/memory"
	service "github.com/xjncx/taskmanager/internal/service"
)

func main() {
	repo := repository.NewInMemoryRepo()
	taskService := service.NewService(repo)
	handler := api.NewHandler(taskService)
	router := api.NewRouter(handler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server running on http://localhost:8080")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}

}
