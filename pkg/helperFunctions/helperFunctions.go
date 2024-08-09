package helperFunctions

import (
	"strings"
	"unicode"
)

// Counts is a struct that holds the counts of words, punctuation, vowels, sentences, paragraphs, and digits
type Counts struct {
	Word      int
	Punct     int
	Vowel     int
	Sentence  int
	Paragraph int
	Digit     int
}

// CountChannels is a struct that holds the channels for the counts
type CountChannels struct {
	WordChan      chan int
	PunctChan     chan int
	VowelChan     chan int
	SentenceChan  chan int
	ParagraphChan chan int
	DigitChan     chan int
}

// IsPunctuation checks if a given character is a punctuation mark
func IsPunctuation(s string) bool {
    punctuations := ".,;:!?-()[]{}'\""
    return len(s) == 1 && strings.Contains(punctuations, s)
}

// IsVowel checks if a given character is a vowel
func IsVowel(char string) bool {
	vowels := "aeiouAEIOU"
	return len(char) == 1 && strings.Contains(vowels, char)
}

// IsSentence checks if a given character is a sentence-ending punctuation
func IsSentence(char string) bool {
	sentences := "."
	return len(char) == 1 && strings.Contains(sentences, char)
}

// ProcessChar processes a character and updates the counts
func ProcessChar(char byte, inWord *bool, counts *Counts) {
	if unicode.IsSpace(rune(char)) {
		if *inWord {
			counts.Word++
			*inWord = false
		}
		if char == '\n' {
			counts.Paragraph++
		}
	} else {
		*inWord = true
	}

	// Check for letters, digits, and punctuation
	if unicode.IsLetter(rune(char)) {
		if IsVowel(string(char)) {
			counts.Vowel++
		}
	} else if unicode.IsDigit(rune(char)) {
		counts.Digit++
	} else if IsPunctuation(string(char)) {
		counts.Punct++
	}

	// Check for sentence endings
	if IsSentence(string(char)) {
		counts.Sentence++
	}
}
