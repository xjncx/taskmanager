package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type TaskState string

const (
	StatePending   TaskState = "pending"
	StateRunning   TaskState = "running"
	StateDone      TaskState = "done"
	StateError     TaskState = "error"
	StateCancelled TaskState = "cancelled"
)

type TaskStatus struct {
	CreatedAt   time.Time
	CompletedAt time.Time
	Duration    time.Duration
	Result      string
	State       TaskState
}

type Task struct {
	id     uuid.UUID
	ctx    context.Context
	cancel context.CancelFunc
	status TaskStatus
}
