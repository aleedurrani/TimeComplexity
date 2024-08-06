# Time Complexity Analysis: Text File Processing in GoLang

## Unoptimized Code

## Overview
This Go program processes a text file (`file.txt`) and performs various counting operations. It calculates the number of words, punctuation marks, vowels, sentences, paragraphs, and digits in the file. The program is structured with separate functions for each counting operation.

## Code Structure
- The main functionality is divided into several functions:
  - `countWords()`
  - `countPunctuation()`
  - `countVowels()`
  - `countSentences()`
  - `countParagraphs()`
  - `countDigits()`
- Helper functions are used to identify specific characters:
  - `isPunctuation()`
  - `isVowel()`
  - `isSentence()`
  - `isParagraph()`
  - `isDigit()`
- The `main()` function orchestrates the execution of all counting operations and measures the total execution time.

## Performance Analysis
Based on the output provided:

1. **Execution Time**: The program completed all operations in 522.5169ms, which is relatively fast for processing a large text file.

2. **File Statistics**:
   - Words: 1,048,500
   - Punctuation marks: 115,305
   - Vowels: 2,416,411
   - Sentences: 62,714
   - Paragraphs: 20,968
   - Digits: 407,936

3. **Observations**:
   - The text contains a high number of words, indicating a substantial document.
   - There's a significant number of vowels, as expected in most languages.
   - The ratio of sentences to paragraphs (about 3:1) suggests relatively short paragraphs on average.
   - The presence of over 400,000 digits indicates that the text likely includes numerical data or references.

