package main

import (
	"net/http"

	"github.com/xjncx/taskmanager/internal/api"
)

func main() {
	handler := api.NewHandler(service)
	router := api.NewRouter(handler)

	srv := &http.Server{
		Addr:    "8080",
		handler: router,
	}
}
