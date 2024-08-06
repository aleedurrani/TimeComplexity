package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
)

//Unoptimized Code




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

func isPunctuation(s string) bool {
    punctuations := ".,;:!?-()[]{}'\""
    return len(s) == 1 && strings.Contains(punctuations, s)
}


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

func isVowel(char string) bool {
	vowels := "aeiouAEIOU"
	return len(char) == 1 && strings.Contains(vowels, char)
}

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

func isSentence(char string) bool {
	sentences := ".!?"
	return len(char) == 1 && strings.Contains(sentences, char)
}

func countParagraphs() int {
	file, err := os.Open("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanRunes)
    paragraphCount := 0
    for scanner.Scan() {
        char := scanner.Text()
        if isParagraph(char) {
            paragraphCount++
        }
    }
    return paragraphCount
}

func isParagraph(char string) bool {
	paragraphs := "\n"
	return len(char) == 1 && strings.Contains(paragraphs, char)
}

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

func isDigit(char string) bool {
	digits := "0123456789"
	return len(char) == 1 && strings.Contains(digits, char)
}


//Optimized Code (Reading the file once)

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

	for scanner.Scan() {
		char := scanner.Text()

		if unicode.IsLetter([]rune(char)[0]) {
			if !inWord {
				wordCount++
				inWord = true
			}
			if isVowel(char) {
				vowelCount++
			}
		} else {
			inWord = false
			if isPunctuation(char) {
				punctCount++
			}
			if isSentence(char) {
				sentenceCount++
			}
			if isParagraph(char) {
				paragraphCount++
			}
			if isDigit(char) {
				digitCount++
			}
		}
	}

	return wordCount, punctCount, vowelCount, sentenceCount, paragraphCount, digitCount
}




func main() {
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
}