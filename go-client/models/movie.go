package models

type Movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
}

// type BatchRequest struct {
// 	Movies []Movie `json:"movies"`
// }

// type StatusResponse struct {
// 	TaskID  string `json:"task_id"`
// 	Message string `json:"message"`
// }
