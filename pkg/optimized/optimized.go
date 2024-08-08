package optimized

import (
	"bufio"
	"log"
	"os"
	"unicode"
	"strings"
)

// optimizedCountAll counts words, punctuation, vowels, sentences, paragraphs, and digits in a single pass
func OptimizedCountAll() (int, int, int, int, int, int) {
	file, err := os.Open("../../assets/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	wordCount, punctCount, vowelCount, sentenceCount, paragraphCount, digitCount := 0, 0, 0, 0, 0, 0
	inWord := false
	prevChar := ""

	for scanner.Scan() {
		char := scanner.Text()

		// Check for word boundaries
		if unicode.IsSpace([]rune(char)[0]) {
			if inWord {
				wordCount++
				inWord = false
			}
			if char == "\n" {
				paragraphCount++
			}
		} else {
			inWord = true
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

		prevChar = char
	}

	// Handle the last word if the file doesn't end with whitespace
	if inWord {
		wordCount++
	}

	// Handle the last paragraph if the file doesn't end with an empty line
	if prevChar != "\n" {
		paragraphCount++
	}

	return wordCount, punctCount, vowelCount, sentenceCount, paragraphCount, digitCount
}

// isPunctuation checks if a given character is a punctuation mark
func isPunctuation(s string) bool {
    punctuations := ".,;:!?-()[]{}'\""
    return len(s) == 1 && strings.Contains(punctuations, s)
}

// isVowel checks if a given character is a vowel
func isVowel(char string) bool {
	vowels := "aeiouAEIOU"
	return len(char) == 1 && strings.Contains(vowels, char)
}

// isSentence checks if a given character is a sentence-ending punctuation
func isSentence(char string) bool {
	sentences := "."
	return len(char) == 1 && strings.Contains(sentences, char)
}