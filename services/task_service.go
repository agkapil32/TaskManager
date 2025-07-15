package services

import (
	"TaskManager/models"
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	taskList = []models.Task{}
	taskID   = 1
	lock     sync.Mutex
)

func CreateTask(task models.Task) models.Task {
	lock.Lock()
	defer lock.Unlock()

	task.ID = taskID
	task.Status = models.StatusBacklog
	if task.Priority == "" {
		task.Priority = models.PriorityUnset
	}
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	taskList = append(taskList, task)
	taskID++
	return task
}

func GetTaskByID(id int) (*models.Task, error) {
	for _, t := range taskList {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, errors.New("task not found")
}

func UpdateTask(id int, updated models.Task) (*models.Task, error) {
	lock.Lock()
	defer lock.Unlock()

	for i, t := range taskList {
		if t.ID == id {
			updated.ID = id
			updated.CreatedAt = t.CreatedAt
			updated.UpdatedAt = time.Now()
			taskList[i] = updated
			return &updated, nil
		}
	}
	return nil, errors.New("task not found")
}

func DeleteTask(id int) error {
	lock.Lock()
	defer lock.Unlock()

	for i, t := range taskList {
		if t.ID == id {
			taskList = append(taskList[:i], taskList[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

func ListTasks(filters map[string]string, page, pageSize int, sortBy string) []models.Task {
	filtered := []models.Task{}

	// Filter by userId and status
	for _, t := range taskList {
		match := true
		if v, ok := filters["userId"]; ok && v != "" && toString(t.UserID) != v {
			match = false
		}
		if v, ok := filters["status"]; ok && v != "" && !strings.EqualFold(string(t.Status), v) {
			match = false
		}
		if match {
			filtered = append(filtered, t)
		}
	}

	// Sorting (default by ID)
	switch sortBy {
	case "startTime":
		sort.Slice(filtered, func(i, j int) bool {
			return filtered[i].StartTime.Before(filtered[j].StartTime)
		})
	case "createdAt":
		sort.Slice(filtered, func(i, j int) bool {
			return filtered[i].CreatedAt.Before(filtered[j].CreatedAt)
		})
	default:
		sort.Slice(filtered, func(i, j int) bool {
			return filtered[i].ID < filtered[j].ID
		})
	}

	// Pagination
	start := (page - 1) * pageSize
	end := start + pageSize
	if start > len(filtered) {
		return []models.Task{}
	}
	if end > len(filtered) {
		end = len(filtered)
	}
	return filtered[start:end]
}

func toString(i int) string {
	return fmt.Sprintf("%d", i)
}
