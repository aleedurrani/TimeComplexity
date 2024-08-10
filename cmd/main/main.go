package main

import (
	"fmt"
	"time"
	"github.com/aleedurrani/TimeComplexity/pkg/unoptimized"
	"github.com/aleedurrani/TimeComplexity/pkg/optimized"
	"github.com/aleedurrani/TimeComplexity/pkg/parallel"
	"github.com/aleedurrani/TimeComplexity/pkg/parallelExtended"
	"github.com/aleedurrani/TimeComplexity/pkg/utils/helperFunctions"
)



func main() {
	// Unoptimized version
	start := time.Now()

	counts := helperFunctions.Counts{
		Word:      unoptimized.CountWords(),
		Punct:     unoptimized.CountPunctuation(),
		Vowel:     unoptimized.CountVowels(),
		Sentence:  unoptimized.CountSentences(),
		Paragraph: unoptimized.CountParagraphs(),
		Digit:     unoptimized.CountDigits(),
	}

	duration := time.Since(start)

	fmt.Printf(
		"Unoptimized Code\nTotal word count: %d\nTotal punctuation count: %d\nTotal vowel count: %d\n" +
        "Total sentence count: %d\nTotal paragraph count: %d\nTotal digit count: %d\nTotal execution time: %v\n",
        counts.Word, counts.Punct, counts.Vowel, counts.Sentence, counts.Paragraph, counts.Digit, duration,
	)


	// Optimized version (single pass)
	fmt.Println("\nOptimized Code (Reading the file once)")
	start = time.Now()

    optimizedCounts := optimized.OptimizedCountAll()

	optimizedDuration := time.Since(start)

	fmt.Printf(
		"Total word count: %d\n" + "Total punctuation count: %d\n" +"Total vowel count: %d\n" +"Total sentence count: %d\n" +
		"Total paragraph count: %d\n" + "Total digit count: %d\n" + "Total execution time: %v\n",
		optimizedCounts.Word, optimizedCounts.Punct, optimizedCounts.Vowel, optimizedCounts.Sentence, optimizedCounts.Paragraph,
        optimizedCounts.Digit,optimizedDuration,
	)
	fmt.Printf("\nPerformance improvement: %.2f%%\n", (1 - float64(optimizedDuration)/float64(duration)) * 100)


	// Parallel version
	fmt.Println("\nFurther Optimized Code (Using goroutines)")
	start = time.Now()

	parallelCounts := parallel.ParallelCountAll()

	parallelDuration := time.Since(start)

	fmt.Printf(
		"Total word count: %d\n" + "Total punctuation count: %d\n" + "Total vowel count: %d\n" + "Total sentence count: %d\n" +
		"Total paragraph count: %d\n" + "Total digit count: %d\n" + "Total execution time: %v\n",
		parallelCounts.Word, parallelCounts.Punct, parallelCounts.Vowel, parallelCounts.Sentence, parallelCounts.Paragraph,
		parallelCounts.Digit, parallelDuration,
	)
	fmt.Printf("\nPerformance improvement: %.2f%%\n", (1 - float64(parallelDuration)/float64(optimizedDuration)) * 100)


	// Parallel Extended version
	fmt.Println("\nParallel Extended Code (Using goroutines with improved chunk processing)")
	start = time.Now()

	parallelExtendedCounts := parallelExtended.ParallelCountAll()

	parallelExtendedDuration := time.Since(start)

	fmt.Printf(
		"Total word count: %d\n" + "Total punctuation count: %d\n" + "Total vowel count: %d\n" + "Total sentence count: %d\n" +
		"Total paragraph count: %d\n" + "Total digit count: %d\n" + "Total execution time: %v\n",
		parallelExtendedCounts.Word, parallelExtendedCounts.Punct, parallelExtendedCounts.Vowel, parallelExtendedCounts.Sentence,
		parallelExtendedCounts.Paragraph, parallelExtendedCounts.Digit, parallelExtendedDuration,
	)
	fmt.Printf("\nPerformance improvement %.2f%%\n", (1 - float64(parallelExtendedDuration)/float64(parallelDuration)) * 100)
	
}