package unoptimized

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

// Unoptimized Code Section
// These functions read the file multiple times, once for each counting operation

// CountWords counts the total number of words in the file
func CountWords() int {
    file, err := os.Open("../../assets/file.txt")
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

// CountPunctuation counts the total number of punctuation marks in the file
func CountPunctuation() int {
	file, err := os.Open("../../assets/file.txt")
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

// CountVowels counts the total number of vowels in the file
func CountVowels() int {
	file, err := os.Open("../../assets/file.txt")
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

// CountSentences counts the total number of sentences in the file
func CountSentences() int {
	file, err := os.Open("../../assets/file.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanRunes)
    sentenceCount := 0
    inSentence := false
    for scanner.Scan() {
        char := scanner.Text()
        if !inSentence && !unicode.IsSpace([]rune(char)[0]) {
            inSentence = true
        }
        if inSentence && isSentence(char) {
            sentenceCount++
            inSentence = false
        }
    }
    // Handle the case where the file doesn't end with a sentence-ending punctuation
    if inSentence {
        sentenceCount++
    }
    return sentenceCount
}

// isSentence checks if a given character is a sentence-ending punctuation
func isSentence(char string) bool {
	sentences := "."
	return len(char) == 1 && strings.Contains(sentences, char)
}

// CountParagraphs counts the total number of paragraphs in the file
func CountParagraphs() int {
	file, err := os.Open("../../assets/file.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanRunes)
    paragraphCount := 0
    for scanner.Scan() {
        char := scanner.Text()
        if char == "\n" {
            paragraphCount++
        }
    }
    return paragraphCount
}

// CountDigits counts the total number of digits in the file
func CountDigits() int {
	file, err := os.Open("../../assets/file.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanRunes)
    digitCount := 0
    for scanner.Scan() {
        char := scanner.Text()
        if unicode.IsDigit([]rune(char)[0]) {
            digitCount++
        }
    }
    return digitCount
}
