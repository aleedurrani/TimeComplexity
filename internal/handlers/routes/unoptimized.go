package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aleedurrani/TimeComplexity/internal/utils/common"
	"github.com/aleedurrani/TimeComplexity/pkg/unoptimized"
	"github.com/aleedurrani/TimeComplexity/pkg/utils/helperFunctions"
	"github.com/aleedurrani/TimeComplexity/internal/dbConnection"
	
)

// UnoptimizedHandler handles the unoptimized endpoint
func UnoptimizedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		_, err := common.GetFileContent(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		counts, duration := common.RunMethod(func() helperFunctions.Counts {
			return helperFunctions.Counts{
				Word:      unoptimized.CountWords(),
				Punct:     unoptimized.CountPunctuation(),
				Vowel:     unoptimized.CountVowels(),
				Sentence:  unoptimized.CountSentences(),
				Paragraph: unoptimized.CountParagraphs(),
				Digit:     unoptimized.CountDigits(),
			}
		})

		response := map[string]interface{}{
			"counts":   counts,
			"duration": duration.String(),
		}

		err = dbConnection.StoreResponse("unoptimized", response)
		if err != nil {
			log.Printf("Error storing response: %v", err)
			http.Error(w, "Error storing response", http.StatusInternalServerError)
			return
		}

		response["message"] = "Results added to database"
		json.NewEncoder(w).Encode(response)

	} else if r.Method == http.MethodGet {
		records, err := dbConnection.RetrieveRecords("unoptimized")
		if err != nil {
			if err.Error() == "no records found" {
				http.Error(w, "No records found", http.StatusNotFound)
			} else {
				log.Printf("Error retrieving records: %v", err)
				http.Error(w, "Error retrieving records", http.StatusInternalServerError)
			}
			return
		}

		json.NewEncoder(w).Encode(records)

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}



