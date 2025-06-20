package service

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Service struct{}

func (s *Service) CreateTask(ctx context.Context, duration time.Duration) (uuid.UUID, error) {

	if duration <= 0 || duration > 10 {
		return uuid.Nil, ErrInvalidDuration
	}

	id := uuid.New()
	ctx, cancel := context.WithCancel()
	start = time.Now()

	*Task := Model.Task{}

}
