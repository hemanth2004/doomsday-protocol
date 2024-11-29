package debug

import (
	"log"
	"os"
	"path/filepath"
	"sync"
)

var (
	logger      *log.Logger
	file        *os.File
	initOnce    sync.Once
	mutex       sync.Mutex
	logDir      = "debug"
	logFile     = "DEBUG.txt"
	logFilePath string
)

// initLogger initializes the logger and ensures the file and directory exist.
func initLogger() {
	logFilePath = filepath.Join(logDir, logFile)

	// Create the directory if it doesn't exist
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	// Open the log file for appending
	file, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Create the logger
	logger = log.New(file, "", log.LstdFlags)
}

// Log writes a log message to debug/DEBUG.txt
func Log(message string) {
	initOnce.Do(initLogger) // Ensure the logger is initialized only once

	mutex.Lock()
	defer mutex.Unlock()

	// Write the message
	logger.Printf("%s", message)
}

// Close cleans up the resources when the program exits.
func Close() {
	if file != nil {
		file.Close()
	}
}
