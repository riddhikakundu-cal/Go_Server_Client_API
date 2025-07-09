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

// 9th july
// package controllers

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"movie-api/go-client/models"
// 	"net/http"
// 	"time"
// )

// const baseURL = "http://localhost:8080/api/movies"

// func SendBatchRequests() {
// 	fmt.Println("🎬 Sending batch of requests all at once to the server...")

// 	// Prepare 20 movies
// 	var movieBatch []models.Movie
// 	for i := 1; i <= 20; i++ {
// 		movie := models.Movie{
// 			ID:       fmt.Sprintf("%d", i),
// 			Title:    fmt.Sprintf("Movie %d", i),
// 			Director: fmt.Sprintf("Director %d", i),
// 		}
// 		movieBatch = append(movieBatch, movie)
// 	}

// 	// Send POST
// 	taskID := sendBatchPOST(movieBatch)
// 	if taskID == "" {
// 		fmt.Println("Task ID not received. Exiting.")
// 		return
// 	}

// 	// Poll status every 30 seconds until complete
// 	for {
// 		time.Sleep(30 * time.Second)
// 		fmt.Println("\nPolling task status...")

// 		status := checkStatus(taskID)
// 		if status == "Completed" {
// 			break
// 		}
// 	}
// }

// func sendBatchPOST(movies []models.Movie) string {
// 	bodyBytes, err := json.Marshal(movies)
// 	if err != nil {
// 		fmt.Println("Error marshaling movies:", err)
// 		return ""
// 	}

// 	req, err := http.NewRequest("POST", baseURL+"/batch", bytes.NewBuffer(bodyBytes))
// 	if err != nil {
// 		fmt.Println("Failed to create POST request:", err)
// 		return ""
// 	}
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{Timeout: 2 * time.Minute}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("POST failed:", err)
// 		return ""
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading POST response:", err)
// 		return ""
// 	}

// 	var result map[string]interface{}
// 	if err := json.Unmarshal(body, &result); err != nil {
// 		fmt.Println("Failed to unmarshal POST response:", err)
// 		return ""
// 	}

// 	pretty, _ := json.MarshalIndent(result, "", "  ")
// 	fmt.Println("POST response:\n", string(pretty))

// 	taskID, ok := result["taskId"].(string)
// 	if !ok || taskID == "" {
// 		fmt.Println("taskId not found in response")
// 		return ""
// 	}

// 	return taskID
// }

// func checkStatus(taskID string) string {
// 	url := baseURL + "/status/" + taskID
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("Status check failed:", err)
// 		return ""
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading status response:", err)
// 		return ""
// 	}

// 	var result map[string]interface{}
// 	if err := json.Unmarshal(body, &result); err != nil {
// 		fmt.Println("Failed to unmarshal status response:", err)
// 		return ""
// 	}

// 	pretty, _ := json.MarshalIndent(result, "", "  ")
// 	fmt.Println("Status response:\n", string(pretty))

// 	if result["status"] == "Completed" {
// 		return "Completed"
// 	}
// 	return "In Progress"
// }
// 9th july

// package controllers

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"movie-api/go-client/models"
// 	"net/http"
// 	"time"
// )

// const baseURL = "http://localhost:8080/api/movies"

// func SendBatchRequests() {
// 	fmt.Println("Sending batch of requests all at once to the server...")

// 	// Prepare 20 movies
// 	var movieBatch []models.Movie
// 	for i := 1; i <= 20; i++ {
// 		movie := models.Movie{
// 			ID:       fmt.Sprintf("%d", i),
// 			Title:    fmt.Sprintf("Movie %d", i),
// 			Director: fmt.Sprintf("Director %d", i),
// 		}
// 		movieBatch = append(movieBatch, movie)
// 	}

// 	// Send POST
// 	taskID := sendBatchPOST(movieBatch)

// 	// Poll status every 30 seconds until complete
// 	for {
// 		time.Sleep(30 * time.Second)
// 		fmt.Println("\nPolling task status...")

// 		status := checkStatus(taskID)
// 		if status == "Completed" {
// 			break
// 		}
// 	}
// }

// func sendBatchPOST(movies []models.Movie) string {
// 	bodyBytes, _ := json.Marshal(movies)
// 	req, _ := http.NewRequest("POST", baseURL+"/batch", bytes.NewBuffer(bodyBytes))
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{Timeout: 1 * time.Minute}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("POST failed:", err)
// 		return ""
// 	}
// 	defer resp.Body.Close()

// 	body, _ := io.ReadAll(resp.Body)

// 	var result map[string]string
// 	json.Unmarshal(body, &result)

// 	taskID := result["taskId"]
// 	pretty, _ := json.MarshalIndent(result, "", "  ")
// 	fmt.Println("POST response:\n", string(pretty))

// 	return taskID
// }

// func checkStatus(taskID string) string {
// 	resp, err := http.Get(baseURL + "/status/" + taskID)
// 	if err != nil {
// 		fmt.Println("Status check failed:", err)
// 		return ""
// 	}
// 	defer resp.Body.Close()

// 	body, _ := io.ReadAll(resp.Body)

// 	var result map[string]interface{}
// 	json.Unmarshal(body, &result)

// 	pretty, _ := json.MarshalIndent(result, "", "  ")
// 	fmt.Println(string(pretty))

// 	if result["status"] == "Completed" {
// 		return "Completed"
// 	}
// 	return "In Progress"
// }

// package controllers

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"movie-api/go-client/models"
// 	"net/http"
// 	"time"
// )

// const baseURL = "http://localhost:8080/api/movies"

// func SendBatchRequest() {
// 	movies := generateMovies()

// 	fmt.Println("Sending batch of requests all at once to the server...")

// 	// POST batch
// 	taskID := sendBatchPOST(movies)

// 	// Poll every 30s
// 	fmt.Println("Waiting for server to process...")
// 	for {
// 		time.Sleep(30 * time.Second)
// 		status := getTaskStatus(taskID)
// 		fmt.Println("\nStatus Check:")
// 		fmt.Println(status)

// 		if isCompleted(status) {
// 			break
// 		}
// 	}
// }

// func sendBatchPOST(movies []models.Movie) string {
// 	jsonData, _ := json.Marshal(movies)
// 	resp, _ := http.Post(baseURL+"/batch", "application/json", bytes.NewBuffer(jsonData))
// 	defer resp.Body.Close()

// 	body, _ := io.ReadAll(resp.Body)

// 	var response map[string]string
// 	json.Unmarshal(body, &response)
// 	fmt.Println("POST Response:")
// 	prettyPrint(body)

// 	return response["task_id"]
// }

// func getTaskStatus(taskID string) string {
// 	resp, _ := http.Get(baseURL + "/status/" + taskID)
// 	defer resp.Body.Close()
// 	body, _ := io.ReadAll(resp.Body)
// 	return string(body)
// }

// func isCompleted(body string) bool {
// 	var data map[string]interface{}
// 	json.Unmarshal([]byte(body), &data)
// 	return data["status"] == "completed"
// }

// func prettyPrint(body []byte) {
// 	var pretty bytes.Buffer
// 	json.Indent(&pretty, body, "", "  ")
// 	fmt.Println(pretty.String())
// }

// func generateMovies() []models.Movie {
// 	var movies []models.Movie
// 	for i := 1; i <= 20; i++ {
// 		movies = append(movies, models.Movie{
// 			ID:       fmt.Sprintf("%d", 100+i),
// 			Title:    fmt.Sprintf("Movie %d", i),
// 			Director: "Director Name",
// 		})
// 	}
// 	return movies
// }

// package controllers

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"movie-api/go-client/models"
// 	"net/http"
// 	"time"
// )

// const baseURL = "http://localhost:8080/api/movies"

// func SendBatchPostAndPoll() {
// 	batch := generateMovieBatch()
// 	body, _ := json.Marshal(batch)
// 	fmt.Println("📦 Sending POST batch request...")

// 	req, _ := http.NewRequest("POST", baseURL+"/batch", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	client := &http.Client{Timeout: 2 * time.Minute}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("❌ Failed to send batch POST:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	var result map[string]interface{}
// 	json.NewDecoder(resp.Body).Decode(&result)

// 	taskID, ok := result["task_id"].(string)
// 	if !ok {
// 		fmt.Println("❌ Task ID missing in response")
// 		return
// 	}

// 	fmt.Printf("✅ Task submitted. Task ID: %s\n", taskID)
// 	pollStatus(taskID)
// }

// func pollStatus(taskID string) {
// 	for {
// 		time.Sleep(30 * time.Second)
// 		fmt.Println("📡 Polling task status...")
// 		url := fmt.Sprintf("%s/status/%s", baseURL, taskID)

// 		resp, err := http.Get(url)
// 		if err != nil {
// 			fmt.Println("Poll error:", err)
// 			return
// 		}
// 		defer resp.Body.Close()

// 		body, _ := io.ReadAll(resp.Body)
// 		var pretty bytes.Buffer
// 		json.Indent(&pretty, body, "", "  ")
// 		fmt.Println(string(pretty.Bytes()))

// 		var status map[string]interface{}
// 		json.Unmarshal(body, &status)
// 		if status["status"] == "completed" {
// 			break
// 		}
// 	}
// }

// func generateMovieBatch() []models.Movie {
// 	var batch []models.Movie
// 	for i := 101; i <= 120; i++ {
// 		batch = append(batch, models.Movie{
// 			ID:       fmt.Sprintf("%d", i),
// 			Title:    fmt.Sprintf("Movie %d", i),
// 			Director: "Director X",
// 		})
// 	}
// 	return batch
// }
