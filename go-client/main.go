package main

import (
	"fmt"
	"movie-api/go-client/controllers"
)

func main() {
	fmt.Println("Go client started...")
	controllers.SendBatchRequests()
}
