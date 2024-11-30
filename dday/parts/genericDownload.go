package parts

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/hemanth2004/doomsday-protocol/dday/core"
)

// InitiateFileDownload handles downloading a file to a specified folder while updating core.Resource and logging progress.
func InitiateFileDownload(folderPath string, logFunction func(string), downloadStruct *core.Resource) error {
	downloadStruct.Status = core.StatusQueued

	// Ensure the folder path exists or create it
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		errMsg := fmt.Sprintf("Folder path does not exist, creating: %s", folderPath)
		logFunction(errMsg)
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			errMsg := fmt.Sprintf("Failed to create folder path: %s", folderPath)
			logFunction(errMsg)
			return fmt.Errorf(errMsg)
		}
		logFunction(fmt.Sprintf("Folder created: %s", folderPath))
	}

	// Create the full file path
	filePath := filepath.Join(folderPath, downloadStruct.FileName)

	client := &http.Client{
		Timeout: 10 * time.Minute,
	}

	// Start the download request
	url := downloadStruct.UrlGetter.GetUrl()
	resp, err := client.Get(url)
	if err != nil {
		downloadStruct.Status = core.StatusFailed
		errMsg := fmt.Sprintf("Failed to start download: %v", err)
		logFunction(errMsg)
		return fmt.Errorf(errMsg)
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		downloadStruct.Status = core.StatusFailed
		errMsg := fmt.Sprintf("HTTP %d error while fetching data", resp.StatusCode)
		logFunction(errMsg)
		return fmt.Errorf(errMsg)
	}

	// Create the output file
	file, err := os.Create(filePath)
	if err != nil {
		downloadStruct.Status = core.StatusFailed
		errMsg := fmt.Sprintf("Error creating file: %v", err)
		logFunction(errMsg)
		return fmt.Errorf(errMsg)
	}
	defer file.Close()

	// Initialize download information
	downloadStruct.Info.Size = float64(resp.ContentLength)
	downloadStruct.Info.StartTime = time.Now()

	// Log start of download
	logFunction(fmt.Sprintf("Starting download: %s -> %s", downloadStruct.Name, url))

	// Buffer for reading and progress tracking
	buffer := make([]byte, 4096)
	var doneBytes float64
	startTime := time.Now()

	downloadStruct.Status = core.StatusDownloading

	// Download in chunks
	for {
		n, readErr := resp.Body.Read(buffer)
		if n > 0 {
			_, writeErr := file.Write(buffer[:n])
			if writeErr != nil {
				downloadStruct.Status = core.StatusFailed
				errMsg := fmt.Sprintf("Error writing to file: %v", writeErr)
				logFunction(errMsg)
				return fmt.Errorf(errMsg)
			}

			// Update progress
			doneBytes += float64(n)
			downloadStruct.Info.Done = doneBytes

			elapsed := time.Since(startTime).Seconds()
			if elapsed > 0 {
				downloadStruct.Info.Bandwidth = doneBytes / elapsed
				downloadStruct.Info.ETA = uint64((downloadStruct.Info.Size - doneBytes) / downloadStruct.Info.Bandwidth)
			}
			downloadStruct.Status = core.StatusDownloading
		}

		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			downloadStruct.Status = core.StatusFailed
			errMsg := fmt.Sprintf("Error reading data: %v", readErr)
			logFunction(errMsg)
			return fmt.Errorf(errMsg)
		}
	}

	// Finalize download
	downloadStruct.Info.EndTime = time.Now()
	downloadStruct.Status = core.StatusCompleted
	logFunction(fmt.Sprintf("Download completed: %s, Time taken: %.2fs", filePath, time.Since(downloadStruct.Info.StartTime).Seconds()))

	return nil
}
