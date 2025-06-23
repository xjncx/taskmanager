package api

import "time"

type CreateTaskRequest struct {
	Duration int `json:"duration"`
}

type CreateTaskResponse struct {
	ID        string        `json:"id"`
	State     string        `json:"status"`
	CreatedAt time.Time     `json:"createdAt"`
	Duration  time.Duration `json:"duration"`
	Result    string        `json:"result"`
}
