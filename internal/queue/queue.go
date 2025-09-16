// queue interface and job structure

package queue

import (
	"time"
)

type JobStatus string // 5 states

const (
	StatusQueued     JobStatus = "queued"
	StatusRunning    JobStatus = "running"
	StatusFailed     JobStatus = "failed"
	StatusCompleted  JobStatus = "completed"
	StatusTerminated JobStatus = "dead" // in DLQ
)

type Job struct {
	ID           string
	Payload      []byte
	DAGID        string
	NodeID       string
	CronExpr     string
	Dependencies []string // job IDs that must complete before this runs
	Priority     int      // higher means higher priority
	RunAt        time.Time
	Attempts     int
	MaxAttempts  int
	Status       JobStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LastError    string
}

type Queue interface {
	Enqueue(job *Job) (string, error)
	Dequeue() (*Job, error)
	SetCompleted(jobID string) error
	SetFailed(jobID string, errMsg string) error
	GetJob(jobID string) (*Job, error)
	DLQ() ([]*Job, error)
}
