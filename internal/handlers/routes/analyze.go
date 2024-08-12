package routes

import (
	"encoding/json"
	"net/http"

	"github.com/aleedurrani/TimeComplexity/internal/utils/commonFunctions"
	"github.com/aleedurrani/TimeComplexity/pkg/optimized"
	"github.com/aleedurrani/TimeComplexity/pkg/parallel"
	"github.com/aleedurrani/TimeComplexity/pkg/parallelExtended"
	"github.com/aleedurrani/TimeComplexity/pkg/unoptimized"
	"github.com/aleedurrani/TimeComplexity/pkg/utils/helperFunctions"
)

func AnalyzeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	_, err := commonFunctions.GetFileContent(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	numRoutines := commonFunctions.GetNumRoutines(w, r)
	if numRoutines == 0 {
		return
	}

	response := make(map[string]interface{})

	// Unoptimized
	unoptimizedCounts, unoptimizedDuration := commonFunctions.RunMethod(func() helperFunctions.Counts {
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
	optimizedCounts, optimizedDuration := commonFunctions.RunMethod(optimized.OptimizedCountAll)
	response["optimized"] = map[string]interface{}{
		"counts":   optimizedCounts,
		"duration": optimizedDuration.String(),
	}

	// Parallel
	parallelCounts, parallelDuration := commonFunctions.RunMethod(parallel.ParallelCountAll)
	response["parallel"] = map[string]interface{}{
		"counts":   parallelCounts,
		"duration": parallelDuration.String(),
	}

	// Parallel Extended
	parallelExtendedCounts, parallelExtendedDuration := commonFunctions.RunMethod(func() helperFunctions.Counts {
		return parallelExtended.ParallelCountAll(numRoutines)
	})
	response["parallelExtended"] = map[string]interface{}{
		"counts":   parallelExtendedCounts,
		"duration": parallelExtendedDuration.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}