package models

import "time"

type TaskStatus string
type TaskPriority string

const (
	StatusBacklog   TaskStatus = "Backlog"
	StatusWIP       TaskStatus = "WIP"
	StatusCompleted TaskStatus = "Completed"

	PriorityLow    TaskPriority = "Low"
	PriorityMedium TaskPriority = "Medium"
	PriorityHigh   TaskPriority = "High"
	PriorityUnset  TaskPriority = "Unset"
)

type Task struct {
	ID          int          `json:"id"`
	UserID      int          `json:"userId"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Status      TaskStatus   `json:"status"`
	Priority    TaskPriority `json:"priority"`
	StartTime   time.Time    `json:"startTime"`
	EndTime     time.Time    `json:"endTime"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
}
