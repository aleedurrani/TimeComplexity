package optimized

import (
	"github.com/aleedurrani/TimeComplexity/pkg/helperFunctions"
	"github.com/aleedurrani/TimeComplexity/pkg/fileHandling"
)


// OptimizedCountAll counts words, punctuation, vowels, sentences, paragraphs, and digits in a single pass
func OptimizedCountAll() (helperFunctions.Counts) {

	file := fileHandling.OpenFile()
	defer file.Close()
	scanner := fileHandling.CreateRuneScanner(file)

	counts := helperFunctions.Counts{}
	inWord := false
	prevChar := ""

	for scanner.Scan() {
		char := scanner.Text()[0]
		helperFunctions.ProcessChar(char, &inWord, &counts)
		prevChar = string(char)
	}

	// Handle the last word if the file doesn't end with whitespace
	if inWord {
		counts.Word++
	}

	// Handle the last paragraph if the file doesn't end with an empty line
	if prevChar != "\n" {
		counts.Paragraph++
	}

	return counts
}

