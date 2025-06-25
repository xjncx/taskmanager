package model

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Result string

func (r Result) String() string {
	return string(r)
}

const (
	ResultSuccess   Result = "success"
	ResultCancelled Result = "cancelled"
	ResultUnknown   Result = "unknown or in process"
)

type TaskState string

func (s TaskState) String() string {
	return string(s)
}

const (
	StatePending   TaskState = "pending"
	StateRunning   TaskState = "running"
	StateDone      TaskState = "done"
	StateError     TaskState = "error"
	StateCancelled TaskState = "cancelled"
)

type TaskData struct {
	CreatedAt   time.Time
	CompletedAt time.Time
	Duration    time.Duration
	Result      Result
	State       TaskState
}

type Task struct {
	ID     uuid.UUID
	Ctx    context.Context
	Cancel context.CancelFunc
	Data   TaskData
	mu     sync.RWMutex
}

func (t *Task) GetState() TaskState {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.Data.State
}

func (t *Task) CurrentDuration() time.Duration {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if t.Data.CompletedAt.IsZero() {
		return time.Since(t.Data.CreatedAt)
	}

	return t.Data.CompletedAt.Sub(t.Data.CreatedAt)
}

func (t *Task) MarkRunning() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Data.State = StateRunning
}

func (t *Task) MarkCancelled() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Data.State = StateCancelled
	t.Data.CompletedAt = time.Now()
	t.Data.Result = ResultCancelled
}

func (t *Task) MarkDone() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Data.State = StateDone
	t.Data.CompletedAt = time.Now()
	t.Data.Result = ResultSuccess
}

func (t *Task) IsRunning() bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.Data.State == StateRunning
}

func (t *Task) IsDone() bool {
	t.mu.RLock()
	defer t.mu.Unlock()
	return t.Data.State == StateDone
}

func (t *Task) IsCancelled() bool {
	t.mu.RLock()
	defer t.mu.Unlock()
	return t.Data.State == StateCancelled
}
