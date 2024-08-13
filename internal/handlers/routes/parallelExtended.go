package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aleedurrani/TimeComplexity/internal/utils/common"
	"github.com/aleedurrani/TimeComplexity/pkg/parallelExtended"
	"github.com/aleedurrani/TimeComplexity/pkg/utils/helperFunctions"
)

// ParallelExtendedHandler handles the parallelExtended endpoint
func ParallelExtendedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		numRoutines := common.GetNumRoutines(w, r)
		if numRoutines == 0 {
			return
		}

		_, err := common.GetFileContent(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		counts, duration := common.RunMethod(func() helperFunctions.Counts {
			return parallelExtended.ParallelCountAll(numRoutines)
		})

		response := map[string]interface{}{
			"counts":   counts,
			"duration": duration.String(),
		}

		err = StoreResponse("parallelExtended", response)
		if err != nil {
			log.Printf("Error storing response: %v", err)
			http.Error(w, "Error storing response", http.StatusInternalServerError)
			return
		}

		response["message"] = "Results added to database"
		json.NewEncoder(w).Encode(response)

	case http.MethodGet:
		records, err := RetrieveRecords("parallelExtended")
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

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

