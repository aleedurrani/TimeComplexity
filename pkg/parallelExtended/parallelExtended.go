package parallelExtended

import (
	"log"
	"sync"
	"github.com/aleedurrani/TimeComplexity/pkg/utils/helperFunctions"
	"github.com/aleedurrani/TimeComplexity/pkg/utils/fileHandling"
)


// This function uses goroutines to process the file in parallel with chunking based on defined number of go routines
// The number of go routines is defined by the user
func ParallelCountAll(numRoutines int) helperFunctions.Counts {
	file := fileHandling.OpenFile()

	fileSize := fileHandling.GetFileSize(file)

	chunkSize := fileSize / int64(numRoutines)

	var wg sync.WaitGroup
	countsChan := make(chan helperFunctions.Counts, numRoutines)

	// processChunk processes a chunk of the file and sends results through the channel
	processChunk := func(start, end int64) {
		defer wg.Done()
		counts := helperFunctions.Counts{}
		inWord := false
		prevChar := ""

		chunk := make([]byte, end-start)

		if err := fileHandling.ReadChunk(file, chunk, start); err != nil {
			log.Fatal(err)
		}

		for _, char := range string(chunk) {
			helperFunctions.ProcessChar(byte(char), &inWord, &counts)
			prevChar = string(char)
		}

		// Only count the last word if it's the last chunk and we're in a word
		if end >= fileSize && inWord {
			counts.Word++
		}

		// Only count the last paragraph if it's the last chunk and the last character isn't a newline
		if end >= fileSize && prevChar != "\n" {
			counts.Paragraph++
		}

		// Send results through the channel
		countsChan <- counts
	}

	// Start goroutines
	for i := 0; i < numRoutines; i++ {
		start := int64(i) * chunkSize
		end := start + chunkSize
		if i == numRoutines-1 {
			end = fileSize
		}
		wg.Add(1)
		go processChunk(start, end)
	}

	// Close channel when all goroutines are done
	go func() {
		wg.Wait()
		close(countsChan)
	}()

	// Sum up the results from all goroutines
	totalCounts := helperFunctions.SumCounts(countsChan)

	return totalCounts
}