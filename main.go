package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
	"unicode"
)

// Unoptimized Code Section
// These functions read the file multiple times, once for each counting operation

// countWords counts the total number of words in the file
func countWords() int {
    file, err := os.Open("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)
    wordCount := 0
    for scanner.Scan() {
        wordCount++
    }
    return wordCount
}

// countPunctuation counts the total number of punctuation marks in the file
func countPunctuation() int {
	file, err := os.Open("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanRunes)
    punctCount := 0
    for scanner.Scan() {
        char := scanner.Text()
        if isPunctuation(char) {
            punctCount++
        }
    }
    return punctCount
}

// isPunctuation checks if a given character is a punctuation mark
func isPunctuation(s string) bool {
    punctuations := ".,;:!?-()[]{}'\""
    return len(s) == 1 && strings.Contains(punctuations, s)
}

// countVowels counts the total number of vowels in the file
func countVowels() int {
	file, err := os.Open("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanRunes)
    vowelCount := 0
    for scanner.Scan() {
        char := scanner.Text()
        if isVowel(char) {
            vowelCount++
        }
    }
    return vowelCount
}

// isVowel checks if a given character is a vowel
func isVowel(char string) bool {
	vowels := "aeiouAEIOU"
	return len(char) == 1 && strings.Contains(vowels, char)
}

// countSentences counts the total number of sentences in the file
func countSentences() int {
	file, err := os.Open("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanRunes)
    sentenceCount := 0
    for scanner.Scan() {
        char := scanner.Text()
        if isSentence(char) {
            sentenceCount++
        }
    }
    return sentenceCount
}

// isSentence checks if a given character is a sentence-ending punctuation
func isSentence(char string) bool {
	sentences := ".!?"
	return len(char) == 1 && strings.Contains(sentences, char)
}

// countParagraphs counts the total number of paragraphs in the file
func countParagraphs() int {
	file, err := os.Open("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    paragraphCount := 0
    emptyLine := false
    for scanner.Scan() {
        line := scanner.Text()
        if len(strings.TrimSpace(line)) == 0 {
            if !emptyLine {
                paragraphCount++
                emptyLine = true
            }
        } else {
            emptyLine = false
        }
    }
    if !emptyLine {
        paragraphCount++
    }
    return paragraphCount
}

// countDigits counts the total number of digits in the file
func countDigits() int {
	file, err := os.Open("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanRunes)
    digitCount := 0
    for scanner.Scan() {
        char := scanner.Text()
        if isDigit(char) {
            digitCount++
        }
    }
    return digitCount
}

// isDigit checks if a given character is a digit
func isDigit(char string) bool {
	digits := "0123456789"
	return len(char) == 1 && strings.Contains(digits, char)
}

// Optimized Code Section
// This function reads the file only once and performs all counting operations simultaneously

// optimizedCountAll counts words, punctuation, vowels, sentences, paragraphs, and digits in a single pass
func optimizedCountAll() (int, int, int, int, int, int) {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	wordCount, punctCount, vowelCount, sentenceCount, paragraphCount, digitCount := 0, 0, 0, 0, 0, 0
	inWord := false
	emptyLine := true
	consecutiveNewlines := 0

	for scanner.Scan() {
		char := scanner.Text()

		// Check for word boundaries
		if unicode.IsSpace([]rune(char)[0]) {
			if inWord {
				wordCount++
				inWord = false
			}
			if char == "\n" {
				consecutiveNewlines++
				if consecutiveNewlines == 2 {
					paragraphCount++
					emptyLine = true
					consecutiveNewlines = 1
				}
			} else {
				consecutiveNewlines = 0
			}
		} else {
			inWord = true
			emptyLine = false
			consecutiveNewlines = 0
		}

		// Check for letters, digits, and punctuation
		if unicode.IsLetter([]rune(char)[0]) {
			if isVowel(char) {
				vowelCount++
			}
		} else if unicode.IsDigit([]rune(char)[0]) {
			digitCount++
		} else if isPunctuation(char) {
			punctCount++
		}

		// Check for sentence endings
		if isSentence(char) {
			sentenceCount++
		}
	}

	// Handle the last word if the file doesn't end with whitespace
	if inWord {
		wordCount++
	}

	// Handle the last paragraph if the file doesn't end with an empty line
	if !emptyLine {
		paragraphCount++
	}

	return wordCount, punctCount, vowelCount, sentenceCount, paragraphCount, digitCount
}

// Further Optimization Section
// This function uses goroutines to process the file in parallel

// parallelCountAll counts words, punctuation, vowels, sentences, paragraphs, and digits using goroutines
func parallelCountAll() (int, int, int, int, int, int) {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	var wg sync.WaitGroup
	wordChan := make(chan int, 8)
	punctChan := make(chan int, 8)
	vowelChan := make(chan int, 8)
	sentenceChan := make(chan int, 8)
	paragraphChan := make(chan int, 8)
	digitChan := make(chan int, 8)

	chunkSize := 1000000 
	chunk := make([]string, 0, chunkSize)

	// processChunk processes a chunk of the file and sends results through channels
	processChunk := func(chunk []string) {
		defer wg.Done()
		wordCount, punctCount, vowelCount, sentenceCount, paragraphCount, digitCount := 0, 0, 0, 0, 0, 0
		inWord := false
		emptyLine := true
		consecutiveNewlines := 0

		for _, char := range chunk {
			// Similar logic to optimizedCountAll, but for a chunk of the file
			if unicode.IsSpace([]rune(char)[0]) {
				if inWord {
					wordCount++
					inWord = false
				}
				if char == "\n" {
					consecutiveNewlines++
					if consecutiveNewlines == 2 {
						paragraphCount++
						emptyLine = true
						consecutiveNewlines = 1
					}
				} else {
					consecutiveNewlines = 0
				}
			} else {
				inWord = true
				emptyLine = false
				consecutiveNewlines = 0
			}

			if unicode.IsLetter([]rune(char)[0]) {
				if isVowel(char) {
					vowelCount++
				}
			} else if unicode.IsDigit([]rune(char)[0]) {
				digitCount++
			} else if isPunctuation(char) {
				punctCount++
			}

			if isSentence(char) {
				sentenceCount++
			}
		}

		if inWord {
			wordCount++
		}

		if !emptyLine {
			paragraphCount++
		}

		// Send results through channels
		wordChan <- wordCount
		punctChan <- punctCount
		vowelChan <- vowelCount
		sentenceChan <- sentenceCount
		paragraphChan <- paragraphCount
		digitChan <- digitCount
	}

	// Read the file in chunks and process each chunk in a separate goroutine
	for scanner.Scan() {
		chunk = append(chunk, scanner.Text())
		if len(chunk) == chunkSize {
			wg.Add(1)
			go processChunk(chunk)
			chunk = make([]string, 0, chunkSize)
		}
	}

	// Process the last chunk if it's not empty
	if len(chunk) > 0 {
		wg.Add(1)
		go processChunk(chunk)
	}

	// Close channels when all goroutines are done
	go func() {
		wg.Wait()
		close(wordChan)
		close(punctChan)
		close(vowelChan)
		close(sentenceChan)
		close(paragraphChan)
		close(digitChan)
	}()

	// Sum up the results from all goroutines
	totalWordCount, totalPunctCount, totalVowelCount, totalSentenceCount, totalParagraphCount, totalDigitCount := 0, 0, 0, 0, 0, 0

	for wordCount := range wordChan {
		totalWordCount += wordCount
	}
	for punctCount := range punctChan {
		totalPunctCount += punctCount
	}
	for vowelCount := range vowelChan {
		totalVowelCount += vowelCount
	}
	for sentenceCount := range sentenceChan {
		totalSentenceCount += sentenceCount
	}
	for paragraphCount := range paragraphChan {
		totalParagraphCount += paragraphCount
	}
	for digitCount := range digitChan {
		totalDigitCount += digitCount
	}

	return totalWordCount, totalPunctCount, totalVowelCount, totalSentenceCount, totalParagraphCount, totalDigitCount
}

func main() {
	// Unoptimized version
	start := time.Now()

	wordCount := countWords()
	punctCount := countPunctuation()
	vowelCount := countVowels()
	sentenceCount := countSentences()
	paragraphCount := countParagraphs()
	digitCount := countDigits()

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

	optimizedWordCount, optimizedPunctCount, optimizedVowelCount, optimizedSentenceCount, optimizedParagraphCount, optimizedDigitCount := optimizedCountAll()

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
	fmt.Println("\nOptimized Code (Using goroutines)")
	start = time.Now()

	parallelWordCount, parallelPunctCount, parallelVowelCount, parallelSentenceCount, parallelParagraphCount, parallelDigitCount := parallelCountAll()

	parallelDuration := time.Since(start)

	fmt.Printf("Total word count: %d\n", parallelWordCount)
	fmt.Printf("Total punctuation count: %d\n", parallelPunctCount)
	fmt.Printf("Total vowel count: %d\n", parallelVowelCount)
	fmt.Printf("Total sentence count: %d\n", parallelSentenceCount)
	fmt.Printf("Total paragraph count: %d\n", parallelParagraphCount)
	fmt.Printf("Total digit count: %d\n", parallelDigitCount)
	fmt.Printf("Total execution time: %v\n", parallelDuration)

   
	fmt.Printf("\nPerformance improvement: %.2f%%\n", (1 - float64(parallelDuration)/float64(optimizedDuration)) * 100)
}