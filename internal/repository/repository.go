package repository

import (
	"github.com/google/uuid"
	"github.com/xjncx/taskmanager/internal/model"
)

type Repository interface {
	Add(task *model.Task) error
	Get(id uuid.UUID) (*model.Task, error)
	Delete(id uuid.UUID) error
}
