package fileHandling

import (
	"bufio"
	"bytes"
	"io"
	"sync"
)

var (
	uploadedFile []byte
	fileMutex    sync.RWMutex
)

// SetUploadedFile sets the uploaded file content
func SetUploadedFile(content []byte) {
	fileMutex.Lock()
	defer fileMutex.Unlock()
	uploadedFile = content
}

// OpenFile returns a bytes.Reader for the uploaded file
func OpenFile() *bytes.Reader {
	fileMutex.RLock()
	defer fileMutex.RUnlock()
	return bytes.NewReader(uploadedFile)
}

// CreateRuneScanner creates a rune scanner for the file
func CreateRuneScanner(file *bytes.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	return scanner
}

// GetFileSize returns the size of the uploaded file
func GetFileSize(file *bytes.Reader) int64 {
	return int64(file.Len())
}

// ReadChunk reads a chunk of the file from the given start position
func ReadChunk(file *bytes.Reader, chunk []byte, start int64) error {
	_, err := file.ReadAt(chunk, start)
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}