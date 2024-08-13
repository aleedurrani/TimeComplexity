package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aleedurrani/TimeComplexity/internal/handlers/routes"
)

// main function
func main() {
	// Retry for database initialization
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		err := routes.InitDB()
		if err == nil {
			break
		}
		if i == maxRetries-1 {
			log.Fatalf("Error initializing database after %d attempts: %v", maxRetries, err)
		}
		log.Printf("Failed to initialize database. Retrying in %d seconds...", i+1)
		time.Sleep(time.Duration(i+1) * time.Second)
	}
	defer routes.CloseDB()
	
	http.HandleFunc("/analyze", routes.AnalyzeHandler)
	http.HandleFunc("/unoptimized", routes.UnoptimizedHandler)
	http.HandleFunc("/optimized", routes.OptimizedHandler)
	http.HandleFunc("/parallel", routes.ParallelHandler)
	http.HandleFunc("/parallelExtended", routes.ParallelExtendedHandler)

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}