package main

import (
	"fmt"
	"time"
	"time_complexity/unoptimized"
	"time_complexity/optimized"
	"time_complexity/parallel"
	"time_complexity/parallelextended"
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

	optimizedWordCount, optimizedPunctCount, optimizedVowelCount, optimizedSentenceCount, optimizedParagraphCount, optimizedDigitCount := optimized.OptimizedCountAll()

	optimizedDuration := time.Since(start)

	fmt.Printf("Total word count: %d\n", optimizedWordCount)
	fmt.Printf("Total punctuation count: %d\n", optimizedPunctCount)
	fmt.Printf("Total vowel count: %d\n", optimizedVowelCount)
	fmt.Printf("Total sentence count: %d\n", optimizedSentenceCount)
	fmt.Printf("Total paragraph count: %d\n", optimizedParagraphCount)
	fmt.Printf("Total digit count: %d\n", optimizedDigitCount)
	fmt.Printf("Total execution time: %v\n", optimizedDuration)

	fmt.Printf("\nPerformance improvement: %.2f%%\n", (1 - float64(optimizedDuration)/float64(duration)) * 100)

	// Parallel version
	fmt.Println("\nFurther Optimized Code (Using goroutines)")
	start = time.Now()

	parallelWordCount, parallelPunctCount, parallelVowelCount, parallelSentenceCount, parallelParagraphCount, parallelDigitCount := parallel.ParallelCountAll()

	parallelDuration := time.Since(start)

	fmt.Printf("Total word count: %d\n", parallelWordCount)
	fmt.Printf("Total punctuation count: %d\n", parallelPunctCount)
	fmt.Printf("Total vowel count: %d\n", parallelVowelCount)
	fmt.Printf("Total sentence count: %d\n", parallelSentenceCount)
	fmt.Printf("Total paragraph count: %d\n", parallelParagraphCount)
	fmt.Printf("Total digit count: %d\n", parallelDigitCount)
	fmt.Printf("Total execution time: %v\n", parallelDuration)

	fmt.Printf("\nPerformance improvement: %.2f%%\n", (1 - float64(parallelDuration)/float64(optimizedDuration)) * 100)

	// Parallel Extended version
	fmt.Println("\nParallel Extended Code (Using goroutines with improved chunk processing)")
	start = time.Now()

	parallelExtendedWordCount, parallelExtendedPunctCount, parallelExtendedVowelCount, parallelExtendedSentenceCount, parallelExtendedParagraphCount, parallelExtendedDigitCount := parallelextended.ParallelCountAll()

	parallelExtendedDuration := time.Since(start)

	fmt.Printf("Total word count: %d\n", parallelExtendedWordCount)
	fmt.Printf("Total punctuation count: %d\n", parallelExtendedPunctCount)
	fmt.Printf("Total vowel count: %d\n", parallelExtendedVowelCount)
	fmt.Printf("Total sentence count: %d\n", parallelExtendedSentenceCount)
	fmt.Printf("Total paragraph count: %d\n", parallelExtendedParagraphCount)
	fmt.Printf("Total digit count: %d\n", parallelExtendedDigitCount)
	fmt.Printf("Total execution time: %v\n", parallelExtendedDuration)

	fmt.Printf("\nPerformance improvement %.2f%%\n", (1 - float64(parallelExtendedDuration)/float64(parallelDuration)) * 100)
	
}