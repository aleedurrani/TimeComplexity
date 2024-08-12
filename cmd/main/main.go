package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aleedurrani/TimeComplexity/internal/handlers/routes"
)

func main() {
	http.HandleFunc("/analyze", routes.AnalyzeHandler)
	http.HandleFunc("/unoptimized", routes.UnoptimizedHandler)
	http.HandleFunc("/optimized", routes.OptimizedHandler)
	http.HandleFunc("/parallel", routes.ParallelHandler)
	http.HandleFunc("/parallelExtended", routes.ParallelExtendedHandler)

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}