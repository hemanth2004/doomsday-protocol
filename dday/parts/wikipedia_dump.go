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

// InitiateWikipediaDownload aligns with the OSM download format
func InitiateWikipediaDownload(folderPath string, logFunction func(string), downloadStruct *core.Resource) error {

	// Ensure the folder path is valid
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return fmt.Errorf("folder path does not exist: %s", folderPath)
	}

	// Create the full file path
	filePath := filepath.Join(folderPath, downloadStruct.FileName)

	client := &http.Client{
		Timeout: 10 * time.Minute,
	}

	// Start the request
	url := downloadStruct.UrlGetter.GetUrl()
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("failed to start download: %v", err)
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP %d error while fetching data", resp.StatusCode)
	}

	// Set up the file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	// Initialize download information
	downloadStruct.Info.Size = float64(resp.ContentLength)
	downloadStruct.Info.StartTime = time.Now()

	// Download the file in chunks
	buffer := make([]byte, 32*1024) // 32 KB buffer
	var doneBytes float64
	startTime := time.Now()

	for {
		n, readErr := resp.Body.Read(buffer)
		if n > 0 {
			_, writeErr := file.Write(buffer[:n])
			if writeErr != nil {
				return fmt.Errorf("error writing to file: %v", writeErr)
			}

			// Update progress information
			doneBytes += float64(n)
			downloadStruct.Info.Done = doneBytes

			elapsed := time.Since(startTime).Seconds()
			if elapsed > 0 {
				downloadStruct.Info.Bandwidth = doneBytes / elapsed
				downloadStruct.Info.ETA = uint64((downloadStruct.Info.Size - doneBytes) / downloadStruct.Info.Bandwidth)
			}
		}

		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			return fmt.Errorf("error reading data: %v", readErr)
		}
	}

	downloadStruct.Info.EndTime = time.Now()
	return nil
}
