package netcode

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/hemanth2004/doomsday-protocol/dday/core"
)

const CHUNK_SIZE = 4096 // bytes in buffer

// using http ranges
// first check if file exists
// if yes, check if complete
// if yes, finish up and set downloadStruct status
// if not, check if partially downloaded
// and if so, resume download

func InitiateHTTPDownload(folderPath string, logFunction func(string), downloadStruct *core.Resource) error {
	if err := prepareFS(folderPath, logFunction, downloadStruct); err != nil {
		return err
	}

	// Get HTTP response
	resp, err := startHTTPRequest(downloadStruct, logFunction)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := createOutputFile(downloadStruct, logFunction)
	if err != nil {
		return err
	}
	defer file.Close()

	initializeDownloadInfo(resp, downloadStruct)

	if err := performDownload(resp, file, downloadStruct, logFunction); err != nil {
		return err
	}

	// Finalize the download
	finalizeDownload(downloadStruct, logFunction)
	return nil
}

func performDownload(resp *http.Response, file *os.File, downloadStruct *core.Resource, logFunction func(string)) error {

	logFunction(fmt.Sprintf("Starting download: %s -> %s", downloadStruct.Name, downloadStruct.UrlGetter.GetUrl()))

	buffer := make([]byte, CHUNK_SIZE)
	var doneBytes float64
	startTime := time.Now()
	isPaused := false

	downloadStruct.Status = core.StatusDownloading

	// Download in CHUNK_SIZE chunks
	for {
		// Check control channel for commands
		// As this download op is a seperate goroutine
		select {
		case cmd := <-downloadStruct.ControlChannel:
			switch cmd {
			case core.Pause:
				isPaused = true
				downloadStruct.Info.Bandwidth = 0.0
				downloadStruct.Status = core.StatusPaused
				logFunction(fmt.Sprintf("Download paused: %s", downloadStruct.Name))
			case core.Start:
				isPaused = false
				downloadStruct.Status = core.StatusDownloading
				logFunction(fmt.Sprintf("Download resumed: %s", downloadStruct.Name))
			case core.Cancel:
				downloadStruct.Status = core.StatusIdle
				logFunction(fmt.Sprintf("Download cancelled: %s", downloadStruct.Name))
				return nil
			}
		default:
			// Continue with download if not paused
			if !isPaused {
				n, readErr := resp.Body.Read(buffer)
				if n > 0 {
					_, writeErr := file.Write(buffer[:n])
					if writeErr != nil {
						downloadStruct.Status = core.StatusFailed
						errMsg := fmt.Sprintf("Error writing to file: %v", writeErr)
						logFunction(errMsg)
						return errors.New(errMsg)
					}

					// Update progress on downloadStruct
					doneBytes += float64(n)
					downloadStruct.Info.Done = doneBytes

					elapsed := time.Since(startTime).Seconds()
					if elapsed > 0 {
						downloadStruct.Info.Bandwidth = doneBytes / elapsed
						downloadStruct.Info.ETA = uint64((downloadStruct.Info.Size - doneBytes) / downloadStruct.Info.Bandwidth)
					}
				}

				if readErr == io.EOF {
					downloadStruct.Status = core.StatusCompleted
					return nil
				}
				if readErr != nil {
					downloadStruct.Status = core.StatusFailed
					errMsg := fmt.Sprintf("Error reading data: %v", readErr)
					logFunction(errMsg)
					return errors.New(errMsg)
				}
			}
		}
	}
}

func finalizeDownload(downloadStruct *core.Resource, logFunction func(string)) {
	if downloadStruct.Status == core.StatusCompleted {
		downloadStruct.Info.EndTime = time.Now()
		downloadStruct.Info.Size = downloadStruct.Info.Done
		downloadStruct.Info.Bandwidth = 0.0
		downloadStruct.Info.ETA = 0
		logFunction(fmt.Sprintf("Download completed: %s, Time taken: %.2fs",
			downloadStruct.Location,
			time.Since(downloadStruct.Info.StartTime).Seconds()))
	} else if downloadStruct.Status == core.StatusFailed {
		logFunction(fmt.Sprintf("Download failed: %s", downloadStruct.Name))
	}
}
