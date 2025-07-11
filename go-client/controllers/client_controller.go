package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"movie-api/go-client/models"
	"net/http"
	"time"
)

const baseURL = "http://localhost:8080/api/movies"

func SendBatchRequests() {
	fmt.Println("Sending batch of requests all at once to the server...")

	// Prepare 20 movies
	var movieBatch []models.Movie
	for i := 1; i <= 20; i++ {
		movie := models.Movie{
			ID:       fmt.Sprintf("%d", i),
			Title:    fmt.Sprintf("Movie %d", i),
			Director: fmt.Sprintf("Director %d", i),
		}
		movieBatch = append(movieBatch, movie)
	}

	// Send POST
	taskID := sendBatchPOST(movieBatch)

	// Poll status every 30 seconds until complete
	for {
		time.Sleep(30 * time.Second)
		fmt.Println("\nPolling task status...")

		status := checkStatus(taskID)
		if status == "Completed" {
			break
		}
	}
}

func sendBatchPOST(movies []models.Movie) string {
	bodyBytes, _ := json.Marshal(movies)
	req, _ := http.NewRequest("POST", baseURL+"/batch", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 1 * time.Minute}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("POST failed:", err)
		return ""
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result map[string]string
	json.Unmarshal(body, &result)

	taskID := result["taskId"]
	pretty, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println("POST response:\n", string(pretty))

	return taskID
}

func checkStatus(taskID string) string {
	resp, err := http.Get(baseURL + "/status/" + taskID)
	if err != nil {
		fmt.Println("Status check failed:", err)
		return ""
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	pretty, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(pretty))

	if result["status"] == "completed" {
		fmt.Println("Batch completed. Movie data:")
		moviesJSON, _ := json.MarshalIndent(result["movies"], "", "  ")
		fmt.Println(string(moviesJSON))
		return "Completed"
	}

	return "In Progress"
}
