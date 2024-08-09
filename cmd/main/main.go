package main

import (
	"fmt"
	"time"
	"github.com/aleedurrani/TimeComplexity/pkg/unoptimized"
	"github.com/aleedurrani/TimeComplexity/pkg/optimized"
	"github.com/aleedurrani/TimeComplexity/pkg/parallel"
	"github.com/aleedurrani/TimeComplexity/pkg/parallelExtended"
)



func main() {
	// Unoptimized version
	start := time.Now()

	wordCount := unoptimized.CountWords()
	punctCount := unoptimized.CountPunctuation()
	vowelCount := unoptimized.CountVowels()
	sentenceCount := unoptimized.CountSentences()
	paragraphCount := unoptimized.CountParagraphs()
	digitCount := unoptimized.CountDigits()

	duration := time.Since(start)

	fmt.Println("Unoptimized Code")
	fmt.Printf("Total word count: %d\n", wordCount)
	fmt.Printf("Total punctuation count: %d\n", punctCount)
	fmt.Printf("Total vowel count: %d\n", vowelCount)
	fmt.Printf("Total sentence count: %d\n", sentenceCount)
	fmt.Printf("Total paragraph count: %d\n", paragraphCount)
	fmt.Printf("Total digit count: %d\n", digitCount)
	fmt.Printf("Total execution time: %v\n", duration)

	// Optimized version (single pass)
	fmt.Println("\nOptimized Code (Reading the file once)")
	start = time.Now()

    optimizedCounts := optimized.OptimizedCountAll()

	optimizedDuration := time.Since(start)

	fmt.Printf("Total word count: %d\n", optimizedCounts.Word)
	fmt.Printf("Total punctuation count: %d\n", optimizedCounts.Punct)
	fmt.Printf("Total vowel count: %d\n", optimizedCounts.Vowel)
	fmt.Printf("Total sentence count: %d\n", optimizedCounts.Sentence)
	fmt.Printf("Total paragraph count: %d\n", optimizedCounts.Paragraph)
	fmt.Printf("Total digit count: %d\n", optimizedCounts.Digit)
	fmt.Printf("Total execution time: %v\n", optimizedDuration)

	fmt.Printf("\nPerformance improvement: %.2f%%\n", (1 - float64(optimizedDuration)/float64(duration)) * 100)

	// Parallel version
	fmt.Println("\nFurther Optimized Code (Using goroutines)")
	start = time.Now()

	parallelCounts := parallel.ParallelCountAll()

	parallelDuration := time.Since(start)

	fmt.Printf("Total word count: %d\n", parallelCounts.Word)
	fmt.Printf("Total punctuation count: %d\n", parallelCounts.Punct)
	fmt.Printf("Total vowel count: %d\n", parallelCounts.Vowel)
	fmt.Printf("Total sentence count: %d\n", parallelCounts.Sentence)
	fmt.Printf("Total paragraph count: %d\n", parallelCounts.Paragraph)
	fmt.Printf("Total digit count: %d\n", parallelCounts.Digit)
	fmt.Printf("Total execution time: %v\n", parallelDuration)

	fmt.Printf("\nPerformance improvement: %.2f%%\n", (1 - float64(parallelDuration)/float64(optimizedDuration)) * 100)

	// Parallel Extended version
	fmt.Println("\nParallel Extended Code (Using goroutines with improved chunk processing)")
	start = time.Now()

	parallelExtendedCounts := parallelExtended.ParallelCountAll()

	parallelExtendedDuration := time.Since(start)

	fmt.Printf("Total word count: %d\n", parallelExtendedCounts.Word)
	fmt.Printf("Total punctuation count: %d\n", parallelExtendedCounts.Punct)
	fmt.Printf("Total vowel count: %d\n", parallelExtendedCounts.Vowel)
	fmt.Printf("Total sentence count: %d\n", parallelExtendedCounts.Sentence)
	fmt.Printf("Total paragraph count: %d\n", parallelExtendedCounts.Paragraph)
	fmt.Printf("Total digit count: %d\n", parallelExtendedCounts.Digit)
	fmt.Printf("Total execution time: %v\n", parallelExtendedDuration)

	fmt.Printf("\nPerformance improvement %.2f%%\n", (1 - float64(parallelExtendedDuration)/float64(parallelDuration)) * 100)
	
}