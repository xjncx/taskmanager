package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/xjncx/taskmanager/internal/model"
	"github.com/xjncx/taskmanager/internal/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(ctx context.Context, duration time.Duration) (uuid.UUID, error) {

	if duration <= 0 || duration > 10 {
		return uuid.Nil, ErrInvalidDuration
	}

	id := uuid.New()
	ctx, cancel := context.WithCancel(ctx)
	start := time.Now()

	task := &model.Task{
		ID:     id,
		Ctx:    ctx,
		Cancel: cancel,
		Data: model.TaskData{

			CreatedAt: start,
			Duration:  duration,
			State:     model.StatePending,
		},
	}

	err := s.repo.Add(task)

	if err != nil {
		if errors.Is(err, repository.ErrTaskExists) {
			return uuid.Nil, err
		}

		return uuid.Nil, ErrInsertTask
	}

	go func(task *model.Task) {
		task.MarkRunning()

		select {
		case <-task.Ctx.Done():
			task.MarkCancelled()
		case <-time.After(task.Data.Duration):
			task.MarkDone()

		}
	}(task)

	return id, nil
}

func (s *Service) Get(ctx context.Context, id uuid.UUID) (*model.Task, error) {

	task, err := s.repo.Get(id)
	if err != nil {
		log.Printf("failed to get task: %v", err)
		return nil, err
	}

	return task, nil

}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {

	task, err := s.Get(ctx, id)
	if err != nil {
		return err
	}

	if !task.IsDone() && !task.IsCancelled() {
		task.MarkCancelled()
	}

	err = s.repo.Delete(id)
	if err != nil {
		log.Printf("failed to delete task: %v", err)
		return err
	}

	return nil
}
