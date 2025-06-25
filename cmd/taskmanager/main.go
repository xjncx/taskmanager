package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/xjncx/taskmanager/internal/api"
	"github.com/xjncx/taskmanager/internal/manager"
	repository "github.com/xjncx/taskmanager/internal/repository/memory"
	service "github.com/xjncx/taskmanager/internal/service"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("env file not found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	repo := repository.NewInMemoryRepo()
	tm := manager.NewTaskManager(repo)
	taskService := service.NewService(tm)
	handler := api.NewHandler(taskService)
	router := api.NewRouter(handler)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Println("Server running on http://localhost:8080")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}

}
