package parallelextended

import (
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"unicode"
)

// This function uses goroutines to process the file in parallel

// parallelCountAll counts words, punctuation, vowels, sentences, paragraphs, and digits using goroutines
func ParallelCountAll() (int, int, int, int, int, int) {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Get file size
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fileSize := fileInfo.Size()

	// Define number of routines
	numRoutines := 8
	chunkSize := fileSize / int64(numRoutines)

	var wg sync.WaitGroup
	wordChan := make(chan int, numRoutines)
	punctChan := make(chan int, numRoutines)
	vowelChan := make(chan int, numRoutines)
	sentenceChan := make(chan int, numRoutines)
	paragraphChan := make(chan int, numRoutines)
	digitChan := make(chan int, numRoutines)

	// processChunk processes a chunk of the file and sends results through channels
	processChunk := func(start, end int64) {
		defer wg.Done()
		wordCount, punctCount, vowelCount, sentenceCount, paragraphCount, digitCount := 0, 0, 0, 0, 0, 0
		inWord := false
		prevChar := ""

		chunk := make([]byte, end-start)
		_, err := file.ReadAt(chunk, start)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		for i, char := range string(chunk) {
			if unicode.IsSpace(rune(char)) {
				if inWord {
					wordCount++
					inWord = false
				}
				if char == '\n' && prevChar != "\n" {
					paragraphCount++
				}
			} else {
				inWord = true
			}

			if unicode.IsLetter(rune(char)) {
				if isVowel(string(char)) {
					vowelCount++
				}
			} else if unicode.IsDigit(rune(char)) {
				digitCount++
			} else if isPunctuation(string(char)) {
				punctCount++
			}

			if isSentence(string(char)) {
				sentenceCount++
			}

			// Handle the case where a word is split between chunks
			if start > 0 && i == 0 && !unicode.IsSpace(rune(char)) && !inWord {
				wordCount++
			}

			prevChar = string(char)
		}

		// Only count the last word if it's the last chunk and we're in a word
		if end >= fileSize && inWord {
			wordCount++
		}

		// Only count the last paragraph if it's the last chunk and the last character isn't a newline
		if end >= fileSize && prevChar != "\n" {
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

	for i := 0; i < numRoutines; i++ {
		totalWordCount += <-wordChan
		totalPunctCount += <-punctChan
		totalVowelCount += <-vowelChan
		totalSentenceCount += <-sentenceChan
		totalParagraphCount += <-paragraphChan
		totalDigitCount += <-digitChan
	}

	return totalWordCount, totalPunctCount, totalVowelCount, totalSentenceCount, totalParagraphCount, totalDigitCount
}

// isSentence checks if a given character is a sentence-ending punctuation
func isSentence(char string) bool {
	sentences := "."
	return len(char) == 1 && strings.Contains(sentences, char)
}

// isVowel checks if a given character is a vowel
func isVowel(char string) bool {
	vowels := "aeiouAEIOU"
	return len(char) == 1 && strings.Contains(vowels, char)
}

// isPunctuation checks if a given character is a punctuation mark
func isPunctuation(s string) bool {
    punctuations := ".,;:!?-()[]{}'\""
    return len(s) == 1 && strings.Contains(punctuations, s)
}