# Time Complexity Analysis: Text File Processing in GoLang


## Unoptimized Code (Reading the file multiple times using different functions)

### Overview
This Go program processes a text file and performs various counting operations. It calculates the number of words, punctuation marks, vowels, sentences, paragraphs, and digits in the file. The program is structured with separate functions for each counting operation.

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
- The `unoptimizedHandler()` function orchestrates the execution of all counting operations and measures the total execution time.

### Performance Analysis
Based on the output provided:

1. **Execution Time**: The program completed all operations in approximately 6.5 seconds.

2. **File Statistics**:
   - Words: 14,766,930
   - Punctuation marks: 4,266,002
   - Vowels: 31,995,015
   - Sentences: 1,804,847
   - Paragraphs: 164,076
   - Digits: 1,968,924

## Optimized Code (Reading the file once)

### Overview
The optimized version of the program uses a single function `optimizedCountAll()` to process the text file once, counting all the required elements simultaneously. This approach significantly reduces the number of file I/O operations and improves overall performance.

### Code Structure
- A single function `optimizedCountAll()` replaces the multiple counting functions from the unoptimized version.
- This function reads the file character by character and updates all counters in a single pass.
- The `optimizedHandler()` function now calls this optimized function and measures its execution time.

### Performance Analysis
Based on the output provided:

1. **Execution Time**: The optimized version completed all operations in approximately 3.4 seconds.

2. **File Statistics**:
   - Words: 14,766,930
   - Punctuation marks: 4,266,002
   - Vowels: 31,995,015
   - Sentences: 1,804,847
   - Paragraphs: 164,077
   - Digits: 1,968,924

3. **Performance Improvement**:
   - The optimized code achieved a 47-48% improvement in execution time compared to the unoptimized version.

## Further Optimization (Using goroutines)

### Overview
This version further optimizes the code by using goroutines to process chunks of the file in parallel.

### Code Structure
- The `parallelCountAll()` function processes the file using a similar approach to the optimized version, but with minimal parallelization.
- It may use a small number of goroutines or a simplified concurrent approach that doesn't fully utilize parallel processing.

### Performance Analysis
Based on the output provided:

1. **Execution Time**: The parallel version completed all operations in approximately 3.3 seconds.

2. **File Statistics**:
   - Words: 14,766,930
   - Punctuation marks: 4,266,002
   - Vowels: 31,995,015
   - Sentences: 1,804,847
   - Paragraphs: 164,077
   - Digits: 1,968,924

3. **Performance Comparison**:
   - The parallel version achieved similar performance to the optimized (single-pass) version, with only a marginal improvement of about 2-3%.
   - Compared to the original unoptimized version, the total improvement is approximately 49-50%.

## Parallel Extended Optimization

### Overview
This version further enhances the parallel processing by implementing improved chunk processing and experimenting with different numbers of goroutines.

### Code Structure
- The `parallelExtendedCountAll()` function in `parallel_extended.go` implements this optimization.
- It allows for adjusting the number of goroutines to find the optimal performance.

### Performance Analysis
Based on the provided data for different numbers of goroutines:

1. **2 Goroutines**:
   - Execution Time: 1.7638407s
   - Performance: 16.70% slower than the simple parallel version

2. **8 Goroutines**:
   - Execution Time: 912.2323ms
   - Performance: 40.19% improvement over the simple parallel version

3. **100 Goroutines**:
   - Execution Time: 914.7199ms
   - Performance: 41.06% improvement over the simple parallel version

4. **100,000 Goroutines**:
   - Execution Time: 1.1569884s
   - Performance: 26.01% improvement over the simple parallel version

5. **200,000 Goroutines**:
   - Execution Time: 1.8465987s
   - Performance: **15.63% slower** than the simple parallel version

### Analysis of Goroutine Impact
- The performance improves significantly as the number of goroutines increases from 2 to 8.
- The optimal performance is achieved with 8 to 100 goroutines, with the best results around 8 goroutines.
- Beyond 100,000 goroutines, the performance starts to degrade.
- At 200,000 goroutines, the performance becomes worse than the simple parallel version.

## Conclusion
The optimizations demonstrate significant improvements in processing time for the file:
1. The single-pass optimization reduced execution time by about 47-48%.
2. The simple parallel processing further reduced the time by an additional 54-56%.
3. The parallel extended version with **optimal goroutine count (8-100)** achieved the best performance, improving by about 40% over the simple parallel version.

These optimizations showcase:
- The importance of efficient I/O operations.
- The potential of concurrent processing in Go for handling large datasets.
- The critical role of finding the right balance in parallelism, as excessive goroutines can lead to diminishing returns and even performance degradation.

The parallel extended version demonstrates that while parallelism can significantly improve performance, there's an optimal point beyond which adding more goroutines becomes counterproductive due to increased overhead in goroutine management and synchronization.

## How to Run

To run this project using Docker, follow these steps:

1. Build the Docker image:
   ```
   docker compose build
   ```

2. Run the Docker container:
   ```
   docker compose up
   ```

This will start the API server on port 8080 of your local machine.


## API Endpoints

The project provides several API endpoints to analyze text files using different processing methods:

1. **Analyze (All Stats)**
   - Endpoint: `/analyze`
   - Methods: POST, GET
   - Query Parameter: `routines` (optional, for parallelExtended version)
   - POST Description: This endpoint processes the uploaded file using all available methods (unoptimized, optimized, parallel, and parallel extended) and returns the results for each. For the parallelExtended version, you can specify the number of goroutines using the `routines` query parameter (default is 8). The results are stored in the PostgreSQL database.
   - GET Description: Retrieves all stored analysis results from the database.
   - POST Response: JSON object containing counts and execution times for all methods.
   - GET Response: JSON array of all stored analysis results.
   - Example: `/analyze?routines=8` (uses 8 goroutines for parallelExtended)

2. **Unoptimized**
   - Endpoint: `/unoptimized`
   - Methods: POST, GET
   - POST Description: Processes the file using the unoptimized method, reading the file multiple times. Stores the results in the database.
   - GET Description: Retrieves all stored unoptimized analysis results from the database.
   - POST Response: JSON object with counts and execution time.
   - GET Response: JSON array of all stored unoptimized analysis results.

3. **Optimized**
   - Endpoint: `/optimized`
   - Methods: POST, GET
   - POST Description: Analyzes the file using the optimized single-pass method. Stores the results in the database.
   - GET Description: Retrieves all stored optimized analysis results from the database.
   - POST Response: JSON object with counts and execution time.
   - GET Response: JSON array of all stored optimized analysis results.

4. **Parallel**
   - Endpoint: `/parallel`
   - Methods: POST, GET
   - POST Description: Processes the file using simple parallel processing. Stores the results in the database.
   - GET Description: Retrieves all stored parallel analysis results from the database.
   - POST Response: JSON object with counts and execution time.
   - GET Response: JSON array of all stored parallel analysis results.

5. **Parallel Extended**
   - Endpoint: `/parallelExtended`
   - Methods: POST, GET
   - Query Parameter: `routines` (optional, default is 8)
   - POST Description: Analyzes the file using the extended parallel processing method with a customizable number of goroutines. Stores the results in the database.
   - GET Description: Retrieves all stored parallel extended analysis results from the database.
   - POST Response: JSON object with counts and execution time.
   - GET Response: JSON array of all stored parallel extended analysis results.
   - Example: `/parallelExtended?routines=16` (uses 16 goroutines for processing)

To use these endpoints:
1. Ensure the API server is running (follow the "How to Run" instructions above).
2. Use a tool like cURL or Postman to send POST requests to `http://localhost:8080/{endpoint}`.
3. Include the text file (file in the assets folder can be used) you want to analyze in the request body (type form-data) (with the key `file`).





