package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/xjncx/taskmanager/internal/manager"
	"github.com/xjncx/taskmanager/internal/repository"
	"github.com/xjncx/taskmanager/internal/service"
)

type Handler struct {
	svc service.TaskService
}

func NewHandler(svc service.TaskService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	id, err := h.svc.Create(r.Context())
	if err != nil {
		handleError(w, err)
		return
	}

	resp := CreateTaskResponse{ID: id.String()}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		handleError(w, ErrInvalidUUID)
		return
	}

	task, err := h.svc.Get(r.Context(), id)
	if err != nil {
		handleError(w, err)
		return
	}

	resp := TaskResponse{
		ID:        task.ID.String(),
		State:     task.GetState().String(),
		CreatedAt: task.Data.CreatedAt,
		Duration:  task.CurrentDuration(),
		Result:    task.Data.Result.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		handleError(w, err)
		return
	}

	err = h.svc.Delete(r.Context(), id)
	if err != nil {
		handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func handleError(w http.ResponseWriter, err error) {
	log.Printf("[ERROR] %v", err)

	code := http.StatusInternalServerError
	msg := "internal error"

	switch {
	case errors.Is(err, manager.ErrInsertTask):
		code = http.StatusInternalServerError
		msg = "could not create task"

	case errors.Is(err, manager.ErrGetTaskID):
		code = http.StatusBadRequest
		msg = "task ID required"

	case errors.Is(err, repository.ErrTaskNotFound):
		code = http.StatusNotFound
		msg = "task not found"

	case errors.Is(err, repository.ErrTaskExists):
		code = http.StatusConflict
		msg = "task already exists"

	case errors.Is(err, ErrInvalidUUID):
		code = http.StatusBadRequest
		msg = "invalid task ID format"

	case errors.Is(err, repository.ErrStorageFailure):
		code = http.StatusBadGateway
		msg = "internal storage error"
	}

	writeJSON(w, code, map[string]string{"error": msg})
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
