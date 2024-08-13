package common

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/aleedurrani/TimeComplexity/pkg/utils/fileHandling"
	"github.com/aleedurrani/TimeComplexity/pkg/utils/helperFunctions"
)


// RunMethod runs the method and returns the counts and duration
func RunMethod(countFunc func() helperFunctions.Counts) (helperFunctions.Counts, time.Duration) {
	start := time.Now()
	counts := countFunc()
	duration := time.Since(start)
	return counts, duration
}

// GetFileContent gets the file content from the request
func GetFileContent(r *http.Request) ([]byte, error) {
	file, err := ParseAndGetFile(r)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	fileHandling.SetUploadedFile(content)
	return content, nil
}

// ParseAndGetFile parses the request and gets the file
func ParseAndGetFile(r *http.Request) (multipart.File, error) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit on In-Memory Parsing
	if err != nil {
		return nil, fmt.Errorf("unable to parse form: %v", err)
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("error retrieving file: %v", err)
	}

	return file, nil
}

// GetNumRoutines gets the number of routines from the request
func GetNumRoutines(w http.ResponseWriter, r *http.Request) int {
	routines := r.URL.Query().Get("routines")
	numRoutines := 8
	if routines != "" {
		parsedRoutines, err := strconv.Atoi(routines)
		if err != nil {
			http.Error(w, "Invalid routines parameter", http.StatusBadRequest)
			return 0
		}
		if parsedRoutines == 0 {
			http.Error(w, "Number of routines cannot be zero", http.StatusBadRequest)
			return 0
		}
		if parsedRoutines > 0 {
			numRoutines = parsedRoutines
		}
	}
	return numRoutines
}