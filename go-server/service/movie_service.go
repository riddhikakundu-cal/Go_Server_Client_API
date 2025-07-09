package service

import (
	"movie-api/go-server/models"
	"sync"
	"time"

	"github.com/google/uuid"
)

type taskInfo struct {
	Total     int
	Done      int
	DoneFlag  bool
	Movies    []models.Movie
	DoneMutex sync.Mutex
}

var tasks = make(map[string]*taskInfo)
var tasksMu sync.Mutex

func StartBatchProcessing(batch []models.Movie) string {
	id := uuid.New().String()
	ti := &taskInfo{Total: len(batch), Movies: []models.Movie{}}

	tasksMu.Lock()
	tasks[id] = ti
	tasksMu.Unlock()

	go func() {
		for _, m := range batch {
			time.Sleep(5 * time.Second)
			ti.DoneMutex.Lock()
			ti.Movies = append(ti.Movies, m)
			ti.Done++
			ti.DoneMutex.Unlock()
		}
		ti.DoneMutex.Lock()
		ti.DoneFlag = true
		ti.DoneMutex.Unlock()
	}()

	return id
}

func GetTaskStatus(taskID string) (string, bool, []models.Movie) {
	tasksMu.Lock()
	task, exists := tasks[taskID]
	tasksMu.Unlock()
	if !exists {
		return "not_found", false, nil
	}
	task.DoneMutex.Lock()
	defer task.DoneMutex.Unlock()

	if task.DoneFlag {
		return "completed", true, task.Movies
	}
	return "in_progress", false, nil
}
