package fileHandling

import (
	"log"
	"os"
	"bufio"
	"io"
)

// OpenFile opens the file and returns a pointer to the file
func OpenFile() *os.File {
    file, err := os.Open("../../assets/file.txt")
    if err != nil {
        log.Fatal(err)
    }
    return file
}

// CreateRuneScanner creates a rune scanner for the file
func CreateRuneScanner(file *os.File) *bufio.Scanner {
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanRunes)
    return scanner
}

// getFileSize returns the size of the given file
func GetFileSize(file *os.File) int64 {
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	return fileInfo.Size()
}

// ReadChunk reads a chunk of the file from the given start position
func ReadChunk(file *os.File, chunk []byte, start int64) error {
	_, err := file.ReadAt(chunk, start)
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}