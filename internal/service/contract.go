package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/xjncx/taskmanager/internal/model"
)

type TaskService interface {
	Create(ctx context.Context) (uuid.UUID, error)
	Get(ctx context.Context, id uuid.UUID) (*model.Task, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
