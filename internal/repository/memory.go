package repository

import (
	"sync"

	"github.com/google/uuid"
)

type TaskManager struct {
	mu    sync.RWMutex
	tasks map[uuid.UUID]*Task
}
