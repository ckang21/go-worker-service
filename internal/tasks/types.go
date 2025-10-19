package tasks

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

type Status string

const (
	StatusQueued   Status = "queued"
	StatusRunning  Status = "running"
	StatusDone     Status = "done"
	StatusFailed   Status = "failed"
)

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Job struct {
	TaskID string
}

func NewID() string {
	var b [8]byte
	_, _ = rand.Read(b[:])
	return hex.EncodeToString(b[:])
}
