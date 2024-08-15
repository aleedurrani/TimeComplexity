package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aleedurrani/TimeComplexity/internal/utils/common"
	"github.com/aleedurrani/TimeComplexity/pkg/optimized"
	"github.com/aleedurrani/TimeComplexity/pkg/parallel"
	"github.com/aleedurrani/TimeComplexity/pkg/parallelExtended"
	"github.com/aleedurrani/TimeComplexity/pkg/unoptimized"
	"github.com/aleedurrani/TimeComplexity/pkg/utils/helperFunctions"
	"github.com/aleedurrani/TimeComplexity/internal/dbConnection"
	_ "github.com/lib/pq"
)

// AnalyzeHandler handles the analyze endpoint
func AnalyzeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handlePostRequest(w, r)
	case http.MethodGet:
		handleGetRequest(w)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


// handlePostRequest handles the POST request for analyzing the file
func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	_, err := common.GetFileContent(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	numRoutines := common.GetNumRoutines(w, r)
	if numRoutines == 0 {
		return
	}

	response := make(map[string]interface{})

	// Unoptimized
	unoptimizedCounts, unoptimizedDuration := common.RunMethod(func() helperFunctions.Counts {
		return helperFunctions.Counts{
			Word:      unoptimized.CountWords(),
			Punct:     unoptimized.CountPunctuation(),
			Vowel:     unoptimized.CountVowels(),
			Sentence:  unoptimized.CountSentences(),
			Paragraph: unoptimized.CountParagraphs(),
			Digit:     unoptimized.CountDigits(),
		}
	})
	response["unoptimized"] = map[string]interface{}{
		"counts":   unoptimizedCounts,
		"duration": unoptimizedDuration.String(),
	}

	// Optimized
	optimizedCounts, optimizedDuration := common.RunMethod(optimized.OptimizedCountAll)
	response["optimized"] = map[string]interface{}{
		"counts":   optimizedCounts,
		"duration": optimizedDuration.String(),
	}

	// Parallel
	parallelCounts, parallelDuration := common.RunMethod(parallel.ParallelCountAll)
	response["parallel"] = map[string]interface{}{
		"counts":   parallelCounts,
		"duration": parallelDuration.String(),
	}

	// Parallel Extended
	parallelExtendedCounts, parallelExtendedDuration := common.RunMethod(func() helperFunctions.Counts {
		return parallelExtended.ParallelCountAll(numRoutines)
	})
	response["parallelExtended"] = map[string]interface{}{
		"counts":   parallelExtendedCounts,
		"duration": parallelExtendedDuration.String(),
	}

	err = dbConnection.StoreResponse("all",response)
	if err != nil {
		log.Printf("Error storing response: %v", err)
		http.Error(w, "Error storing response", http.StatusInternalServerError)
		return
	}

	response["message"] = "Results added to database"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handleGetRequest handles the GET request for retrieving the analysis results
func handleGetRequest(w http.ResponseWriter) {
	records, err := dbConnection.RetrieveRecords("all")
	if err != nil {
		if err.Error() == "no records found" {
			http.Error(w, "No records found", http.StatusNotFound)
		} else {
			log.Printf("Error retrieving records: %v", err)
			http.Error(w, "Error retrieving records", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}