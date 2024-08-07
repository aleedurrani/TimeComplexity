# Time Complexity Analysis: Text File Processing in GoLang

## Unoptimized Code (Reading the file multiple times using different functions)

### Overview
This Go program processes a text file (`file.txt`) and performs various counting operations. It calculates the number of words, punctuation marks, vowels, sentences, paragraphs, and digits in the file. The program is structured with separate functions for each counting operation.

### Code Structure
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
  - `isDigit()`
- The `main()` function orchestrates the execution of all counting operations and measures the total execution time.

### Performance Analysis
Based on the output provided:

1. **Execution Time**: The program completed all operations in 412.3734ms.

2. **File Statistics**:
   - Words: 1,048,500
   - Punctuation marks: 115,305
   - Vowels: 2,416,411
   - Sentences: 62,714
   - Paragraphs: 10,485
   - Digits: 407,936

## Optimized Code (Reading the file once)

### Overview
The optimized version of the program uses a single function `optimizedCountAll()` to process the text file once, counting all the required elements simultaneously. This approach significantly reduces the number of file I/O operations and improves overall performance.

### Code Structure
- A single function `optimizedCountAll()` replaces the multiple counting functions from the unoptimized version.
- This function reads the file character by character and updates all counters in a single pass.
- The `main()` function now calls this optimized function and measures its execution time.

### Performance Analysis
Based on the output provided:

1. **Execution Time**: The optimized version completed all operations in 242.6145ms.

2. **File Statistics**:
   - Words: 1,048,500
   - Punctuation marks: 115,305
   - Vowels: 2,416,411
   - Sentences: 62,714
   - Paragraphs: 10,485
   - Digits: 407,936

3. **Performance Improvement**:
   - The optimized code achieved a 41.17% improvement in execution time compared to the unoptimized version.

## Further Optimization (Using goroutines)

### Overview
This version further optimizes the code by using goroutines to process chunks of the file in parallel.

### Code Structure
- The `parallelCountAll()` function divides the file into chunks and processes them concurrently using goroutines.
- It uses channels to collect results from each goroutine and combines them for the final count.

### Performance Analysis
Based on the output provided:

1. **Execution Time**: The parallel version completed all operations in 166.0842ms.

2. **File Statistics**:
   - Words: 1,048,506
   - Punctuation marks: 115,305
   - Vowels: 2,416,411
   - Sentences: 62,714
   - Paragraphs: 10,492
   - Digits: 407,936

3. **Performance Improvement**:
   - The parallel version achieved a 31.54% improvement in execution time compared to the optimized (single-pass) version.
   - Compared to the original unoptimized version, the total improvement is approximately 59.73%.


## Conclusion
The optimizations demonstrate significant improvements in processing time:
1. The single-pass optimization reduced execution time by **41.17%**.
2. The parallel processing further reduced the time by an additional **31.54%**.
These optimizations showcase the importance of efficient I/O operations and the potential of concurrent processing in Go for handling large datasets.
