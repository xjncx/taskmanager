package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/xjncx/taskmanager/internal/service"
)

type Handler struct {
	svc service.TaskService
}

func NewHandler(svc service.TaskService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	id, err := h.svc.CreateTask(r.Context(), time.Duration(req.Duration)*time.Second)
}

// func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {...}
// func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {...}
