package unoptimized

import (
	
	"unicode"
	"github.com/aleedurrani/TimeComplexity/pkg/utils/fileHandling"
	"github.com/aleedurrani/TimeComplexity/pkg/utils/helperFunctions"
)

// These functions read the file multiple times, once for each counting operation

// CountWords counts the total number of words in the file
func CountWords() int {

    file := fileHandling.OpenFile()
    scanner := fileHandling.CreateRuneScanner(file)

    wordCount := 0
    inWord := false

    for scanner.Scan() {
        char := scanner.Text()
        if unicode.IsSpace([]rune(char)[0]) {
            if inWord {
                wordCount++
                inWord = false
            }
        } else {
            inWord = true
        }
    }

    // Handle the case where the file doesn't end with whitespace
    if inWord {
        wordCount++
    }
   
    return wordCount
}


// CountPunctuation counts the total number of punctuation marks in the file
func CountPunctuation() int {

    file := fileHandling.OpenFile()
    scanner := fileHandling.CreateRuneScanner(file)
    
    punctCount := 0
    for scanner.Scan() {
        char := scanner.Text()
        if helperFunctions.IsPunctuation(char) {
            punctCount++
        }
    }
    return punctCount
}


// CountVowels counts the total number of vowels in the file
func CountVowels() int {

	file := fileHandling.OpenFile()
    scanner := fileHandling.CreateRuneScanner(file)
    
    vowelCount := 0
    for scanner.Scan() {
        char := scanner.Text()
        if helperFunctions.IsVowel(char) {
            vowelCount++
        }
    }
    return vowelCount
}


// CountSentences counts the total number of sentences in the file
func CountSentences() int {
	
    file := fileHandling.OpenFile()
    scanner := fileHandling.CreateRuneScanner(file)

    sentenceCount := 0
    inSentence := false
    for scanner.Scan() {
        char := scanner.Text()
        if !inSentence && !unicode.IsSpace([]rune(char)[0]) {
            inSentence = true
        }
        if inSentence && helperFunctions.IsSentence(char) {
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


// CountParagraphs counts the total number of paragraphs in the file
func CountParagraphs() int {
	
    file := fileHandling.OpenFile()
    scanner := fileHandling.CreateRuneScanner(file)

    paragraphCount := 0
    for scanner.Scan() {
        char := scanner.Text()[0]

        if char == '\n' {
            paragraphCount++
        }
    }
    return paragraphCount
}

// CountDigits counts the total number of digits in the file
func CountDigits() int {
	
    file := fileHandling.OpenFile()
    scanner := fileHandling.CreateRuneScanner(file)

    digitCount := 0
    for scanner.Scan() {
        char := scanner.Text()
        if unicode.IsDigit([]rune(char)[0]) {
            digitCount++
        }
    }
    return digitCount
}
