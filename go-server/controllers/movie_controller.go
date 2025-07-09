package controllers

import (
	"fmt"
	"movie-api/go-server/models"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskStatus struct {
	StartTime time.Time
	Duration  time.Duration
	Completed bool
	Movies    []models.Movie
}

var tasks = make(map[string]*TaskStatus)
var tasksMu sync.RWMutex

func CreateMovieBatch(c *gin.Context) {
	var movies []models.Movie
	if err := c.ShouldBindJSON(&movies); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	taskID := uuid.New().String()
	duration := 10 * time.Minute

	tasksMu.Lock()
	tasks[taskID] = &TaskStatus{
		StartTime: time.Now(),
		Duration:  duration,
		Completed: false,
		Movies:    movies, // ✅ Store movies
	}
	tasksMu.Unlock()

	// Process in background
	go func(id string) {
		time.Sleep(duration)
		tasksMu.Lock()
		tasks[id].Completed = true
		tasksMu.Unlock()
	}(taskID)

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Batch processing started",
		"taskId":  taskID,
	})
}

func CheckTaskStatus(c *gin.Context) {
	taskID := c.Param("taskId")

	tasksMu.RLock()
	task, exists := tasks[taskID]
	tasksMu.RUnlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task ID not found"})
		return
	}

	elapsed := time.Since(task.StartTime)
	progress := float64(elapsed) / float64(task.Duration)
	if progress > 1 {
		progress = 1
	}
	remaining := task.Duration - elapsed
	if remaining < 0 {
		remaining = 0
	}

	if task.Completed {
		c.JSON(http.StatusOK, gin.H{
			"taskId":   taskID,
			"status":   "completed",
			"message":  "POST request fully processed.",
			"progress": "100%",
			"movies":   task.Movies, // ✅ Respond with batch
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"taskId":           taskID,
		"status":           "in_progress",
		"message":          "POST request is still being processed...",
		"progress":         fmt.Sprintf("%.0f%%", progress*100),
		"remainingPercent": fmt.Sprintf("%.0f%%", (1-progress)*100),
		"remainingTime":    remaining.Truncate(time.Second).String(),
	})
}

//9th july
// package controllers

// import (
// 	"fmt"
// 	"movie-api/go-server/models"
// 	"net/http"
// 	"sync"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// type TaskStatus struct {
// 	StartTime time.Time
// 	Duration  time.Duration
// 	Completed bool
// }

// var tasks = make(map[string]*TaskStatus)
// var tasksMu sync.RWMutex

// func CreateMovieBatch(c *gin.Context) {
// 	var movies []models.Movie
// 	if err := c.ShouldBindJSON(&movies); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
// 		return
// 	}

// 	taskID := uuid.New().String()
// 	duration := 10 * time.Minute

// 	tasksMu.Lock()
// 	tasks[taskID] = &TaskStatus{
// 		StartTime: time.Now(),
// 		Duration:  duration,
// 		Completed: false,
// 	}
// 	tasksMu.Unlock()

// 	go func(id string, mv []models.Movie) {
// 		time.Sleep(duration) // Simulate long processing

// 		tasksMu.Lock()
// 		tasks[id].Completed = true
// 		tasksMu.Unlock()
// 	}(taskID, movies)

// 	c.JSON(http.StatusAccepted, gin.H{
// 		"message": "Batch processing started",
// 		"taskId":  taskID, // ✅ camelCase for consistency
// 	})
// }

// func CheckTaskStatus(c *gin.Context) {
// 	taskID := c.Param("taskId")

// 	tasksMu.RLock()
// 	task, exists := tasks[taskID]
// 	tasksMu.RUnlock()

// 	if !exists {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Task ID not found"})
// 		return
// 	}

// 	elapsed := time.Since(task.StartTime)
// 	progress := float64(elapsed) / float64(task.Duration)
// 	if progress > 1 {
// 		progress = 1
// 	}
// 	remaining := task.Duration - elapsed
// 	if remaining < 0 {
// 		remaining = 0
// 	}

// 	status := "in_progress"
// 	message := "POST request is still being processed..."
// 	if task.Completed {
// 		status = "completed"
// 		message = "POST request has been fully processed."
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"taskId":           taskID,
// 		"status":           status,
// 		"message":          message,
// 		"progress":         fmt.Sprintf("%.0f%%", progress*100),
// 		"remainingPercent": fmt.Sprintf("%.0f%%", (1-progress)*100),
// 		"remainingTime":    remaining.Truncate(time.Second).String(),
// 	})
// }
// 9th july

// package controllers

// import (
// 	"fmt"
// 	"movie-api/go-server/models"
// 	"net/http"
// 	"sync"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// type TaskStatus struct {
// 	StartTime time.Time
// 	Duration  time.Duration
// 	Completed bool
// }

// var tasks = make(map[string]*TaskStatus)
// var tasksMu sync.RWMutex

// func CreateMovieBatch(c *gin.Context) {
// 	var movies []models.Movie
// 	if err := c.ShouldBindJSON(&movies); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
// 		return
// 	}

// 	taskID := uuid.New().String()
// 	duration := 10 * time.Minute

// 	tasksMu.Lock()
// 	tasks[taskID] = &TaskStatus{
// 		StartTime: time.Now(),
// 		Duration:  duration,
// 		Completed: false,
// 	}
// 	tasksMu.Unlock()

// 	go func(id string, mv []models.Movie) {
// 		// Simulate long task
// 		time.Sleep(duration)
// 		tasksMu.Lock()
// 		tasks[id].Completed = true
// 		tasksMu.Unlock()
// 	}(taskID, movies)

// 	c.JSON(http.StatusAccepted, gin.H{
// 		"message": "Batch processing started",
// 		"task_id": taskID,
// 	})
// }

// func CheckTaskStatus(c *gin.Context) {
// 	taskID := c.Param("taskId")
// 	tasksMu.RLock()
// 	task, exists := tasks[taskID]
// 	tasksMu.RUnlock()

// 	if !exists {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Task ID not found"})
// 		return
// 	}

// 	elapsed := time.Since(task.StartTime)
// 	progress := float64(elapsed) / float64(task.Duration)
// 	if progress > 1 {
// 		progress = 1
// 	}

// 	remainingDuration := task.Duration - elapsed
// 	if remainingDuration < 0 {
// 		remainingDuration = 0
// 	}

// 	message := "POST request is still being processed..."
// 	status := "in_progress"
// 	if task.Completed {
// 		message = "POST request has been fully processed."
// 		status = "completed"
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"task_id":              taskID,
// 		"status":               status,
// 		"message":              message,
// 		"progress":             fmt.Sprintf("%.0f%% completed", progress*100),
// 		"remaining_percentage": fmt.Sprintf("%.0f%%", (1-progress)*100),
// 		"remaining_time":       remainingDuration.Truncate(time.Second).String(),
// 	})
// }

// package controllers

// import (
// 	"movie-api/go-server/models"
// 	"movie-api/go-server/service"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func HandleBatchMovies(c *gin.Context) {
// 	var movies []models.Movie
// 	if err := c.ShouldBindJSON(&movies); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
// 		return
// 	}

// 	taskID := service.StartBatchProcessing(movies)
// 	c.JSON(http.StatusAccepted, gin.H{
// 		"task_id": taskID,
// 		"message": "Batch movie processing started",
// 	})
// }

// func CheckTaskStatus(c *gin.Context) {
// 	taskID := c.Param("taskId")
// 	status, done, result := service.GetTaskStatus(taskID)
// 	if !done {
// 		c.JSON(http.StatusOK, gin.H{
// 			"task_id": taskID,
// 			"status":  status,
// 			"message": "POST task is still in progress...",
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"task_id":          taskID,
// 		"status":           status,
// 		"message":          "POST task completed successfully.",
// 		"processed_movies": result,
// 	})
// }
