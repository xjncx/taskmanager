package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/xjncx/taskmanager/internal/manager"
	"github.com/xjncx/taskmanager/internal/model"
)

type Service struct {
	tm *manager.TaskManager
}

func NewService(tm *manager.TaskManager) *Service {
	return &Service{tm: tm}
}

func (s *Service) Create(ctx context.Context) (uuid.UUID, error) {
	return s.tm.Create(ctx)
}

func (s *Service) Get(ctx context.Context, id uuid.UUID) (*model.Task, error) {
	return s.tm.Get(id)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.tm.Delete(id)
}
