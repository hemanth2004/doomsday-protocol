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
	logDir      = "logs"
	logFile     = "logs.txt"
	logFilePath string
)

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

	logger = log.New(file, "", log.LstdFlags)
}

func Log(message string) {
	initOnce.Do(initLogger)

	mutex.Lock()
	defer mutex.Unlock()

	logger.Printf("%s", message)
}

func Close() {
	if file != nil {
		file.Close()
	}
}
