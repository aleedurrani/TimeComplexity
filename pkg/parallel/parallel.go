package parallel

import (
	"bufio"
	"log"
	"os"
	"sync"
	"unicode"
	"strings"
)

// This function uses goroutines to process the file in parallel

// ParallelCountAll counts words, punctuation, vowels, sentences, paragraphs, and digits using goroutines
func ParallelCountAll() (int, int, int, int, int, int) {
	file, err := os.Open("../../assets/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	var wg sync.WaitGroup
	wordChan := make(chan int)
	punctChan := make(chan int)
	vowelChan := make(chan int)
	sentenceChan := make(chan int)
	paragraphChan := make(chan int)
	digitChan := make(chan int)

	chunkSize := 1000000 
	chunk := make([]string, 0, chunkSize)

	// processChunk processes a chunk of the file and sends results through channels
	processChunk := func(chunk []string, isFirstChunk, isLastChunk bool) {
		defer wg.Done()
		wordCount, punctCount, vowelCount, sentenceCount, paragraphCount, digitCount := 0, 0, 0, 0, 0, 0
		inWord := false
		prevChar := ""

		for i, char := range chunk {
			if unicode.IsSpace([]rune(char)[0]) {
				if inWord {
					wordCount++
					inWord = false
				}
				if char == "\n" && prevChar != "\n" {
					paragraphCount++
				}
			} else {
				inWord = true
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

			// Handle the case where a word is split between chunks
			if !isFirstChunk && i == 0 && !unicode.IsSpace([]rune(char)[0]) && !inWord {
				wordCount++
			}

			prevChar = char
		}

		// Only count the last word if it's the last chunk and we're in a word
		if inWord && isLastChunk {
			wordCount++
		}

		// Only count the last paragraph if it's the last chunk and the last character isn't a newline
		if isLastChunk && prevChar != "\n" {
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

	chunkCount := 0
	// Read the file in chunks and process each chunk in a separate goroutine
	for scanner.Scan() {
		chunk = append(chunk, scanner.Text())
		if len(chunk) == chunkSize {
			wg.Add(1)
			go processChunk(chunk, chunkCount == 0, false)
			chunk = make([]string, 0, chunkSize)
			chunkCount++
		}
	}

	// Process the last chunk if it's not empty
	if len(chunk) > 0 {
		wg.Add(1)
		go processChunk(chunk, chunkCount == 0, true)
		chunkCount++
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

	for i := 0; i < chunkCount; i++ {
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