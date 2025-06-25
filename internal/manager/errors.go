package manager

import "errors"

var (
	ErrInsertTask = errors.New("failed to insert task into repository")
	ErrGetTaskID  = errors.New("failed to get Task ID. ID cannot be empty")
)
