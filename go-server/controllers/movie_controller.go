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

// CreateMovieBatch godoc
// @Summary      Start async batch movie processing
// @Description  Accepts a batch of movies and returns a task ID for progress polling
// @Tags         movies
// @Accept       json
// @Produce      json
// @Param        movies  body      []models.Movie  true  "List of movies"
// @Success      202     {object}  map[string]string
// @Failure      400     {object}  map[string]string
// @Router       /movies/batch [post]
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
		Movies:    movies,
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

// CheckTaskStatus godoc
// @Summary      Check task progress
// @Description  Returns progress and remaining time of the task
// @Tags         movies
// @Produce      json
// @Param        taskId  path      string  true  "Task ID"
// @Success      200     {object}  map[string]interface{}
// @Failure      404     {object}  map[string]string
// @Router       /movies/status/{taskId} [get]
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
			"movies":   task.Movies,
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
// 	Movies    []models.Movie
// }

// var tasks = make(map[string]*TaskStatus)
// var tasksMu sync.RWMutex

// // CreateMovieBatch godoc
// // @Summary      Start async batch movie processing
// // @Description  Accepts a batch of movies and returns a task ID for progress polling
// // @Tags         movies
// // @Accept       json
// // @Produce      json
// // @Param        movies  body      []models.Movie  true  "List of movies"
// // @Success      202     {object}  map[string]string
// // @Failure      400     {object}  map[string]string
// // @Router       /movies/batch [post]

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
// 		Movies:    movies, // Store movies
// 	}
// 	tasksMu.Unlock()

// 	// Process in background
// 	go func(id string) {
// 		time.Sleep(duration)
// 		tasksMu.Lock()
// 		tasks[id].Completed = true
// 		tasksMu.Unlock()
// 	}(taskID)

// 	c.JSON(http.StatusAccepted, gin.H{
// 		"message": "Batch processing started",
// 		"taskId":  taskID,
// 	})
// }

// // CheckTaskStatus godoc
// // @Summary      Check task progress
// // @Description  Returns progress and remaining time of the task
// // @Tags         movies
// // @Produce      json
// // @Param        taskId  path      string  true  "Task ID"
// // @Success      200     {object}  map[string]interface{}
// // @Failure      404     {object}  map[string]string
// // @Router       /movies/status/{taskId} [get]

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

// 	if task.Completed {
// 		c.JSON(http.StatusOK, gin.H{
// 			"taskId":   taskID,
// 			"status":   "completed",
// 			"message":  "POST request fully processed.",
// 			"progress": "100%",
// 			"movies":   task.Movies, // Respond with batch
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"taskId":           taskID,
// 		"status":           "in_progress",
// 		"message":          "POST request is still being processed...",
// 		"progress":         fmt.Sprintf("%.0f%%", progress*100),
// 		"remainingPercent": fmt.Sprintf("%.0f%%", (1-progress)*100),
// 		"remainingTime":    remaining.Truncate(time.Second).String(),
// 	})
// }
