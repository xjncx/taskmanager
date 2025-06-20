package service

import "errors"

var (
	ErrInvalidDuration = errors.New("invalid task duration (must be > 0)")
	ErrTooLongDuration = errors.New("duration exceeds allowed max")
)
