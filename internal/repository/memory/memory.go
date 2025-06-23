package repository

import (
	"sync"

	"github.com/google/uuid"
	"github.com/xjncx/taskmanager/internal/model"
	"github.com/xjncx/taskmanager/internal/repository"
)

type InMemoryRepo struct {
	mu    sync.RWMutex
	tasks map[uuid.UUID]*model.Task
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		tasks: make(map[uuid.UUID]*model.Task),
	}
}

func (r *InMemoryRepo) Add(task *model.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.tasks[task.ID]; ok {
		return repository.ErrTaskExists
	}

	r.tasks[task.ID] = task
	return nil
}

func (r *InMemoryRepo) Get(id uuid.UUID) (*model.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	t, ok := r.tasks[id]
	if !ok {
		return nil, repository.ErrTaskNotFound
	}
	return t, nil
}

func (r *InMemoryRepo) Delete(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	t, ok := r.tasks[id]
	if !ok {
		return repository.ErrTaskNotFound
	}

	t.Cancel()
	delete(r.tasks, id)

	return nil
}
