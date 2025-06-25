package manager

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/xjncx/taskmanager/internal/model"
	"github.com/xjncx/taskmanager/internal/repository"
)

type TaskManager struct {
	mu      sync.RWMutex
	repo    repository.Repository
	running map[uuid.UUID]*model.Task
}

func NewTaskManager(repo repository.Repository) *TaskManager {
	return &TaskManager{
		repo:    repo,
		running: make(map[uuid.UUID]*model.Task),
	}
}

func (tm *TaskManager) Create(ctx context.Context) (uuid.UUID, error) {

	min := 3 * time.Minute
	max := 5 * time.Minute
	delta := rand.Int63n(int64(max - min))
	duration := min + time.Duration(delta)

	id := uuid.New()
	taskCtx, cancel := context.WithCancel(context.Background())
	start := time.Now()

	task := &model.Task{
		ID:     id,
		Ctx:    taskCtx,
		Cancel: cancel,
		Data: model.TaskData{

			CreatedAt: start,
			Duration:  duration,
			State:     model.StatePending,
		},
	}

	if err := tm.repo.Add(task); err != nil {
		if errors.Is(err, repository.ErrTaskExists) {
			return uuid.Nil, fmt.Errorf("%w: %v", ErrInsertTask, err)
		}

		return uuid.Nil, fmt.Errorf("failed to add task: %w", err)
	}

	go tm.run(task)

	return task.ID, nil
}

func (tm *TaskManager) run(task *model.Task) {
	task.MarkRunning()

	tm.mu.Lock()
	tm.running[task.ID] = task
	tm.mu.Unlock()

	select {
	case <-task.Ctx.Done():
		task.MarkCancelled()
	case <-time.After(task.Data.Duration):
		task.MarkDone()
	}

	tm.mu.Lock()
	delete(tm.running, task.ID)
	tm.mu.Unlock()
}

func (tm *TaskManager) Get(id uuid.UUID) (*model.Task, error) {
	return tm.repo.Get(id)
}

func (tm *TaskManager) Delete(id uuid.UUID) error {

	task, err := tm.repo.Get(id)
	if err != nil {
		return fmt.Errorf("failed to get task: %w", err)
	}

	if task.IsRunning() {
		log.Printf("TaskManager cancelling running task %s", id)
		task.Cancel()
	}

	tm.mu.Lock()
	delete(tm.running, id)
	tm.mu.Unlock()
	return tm.repo.Delete(id)
}

func (tm *TaskManager) Shutdown() {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	for _, task := range tm.running {
		log.Printf("Shutting down task: %s", task.ID)
		task.Cancel()
	}
}
