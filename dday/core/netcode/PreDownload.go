package netcode

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/hemanth2004/doomsday-protocol/dday/core"
)

func prepareFS(folderPath string, logFunction func(string), downloadStruct *core.Resource) error {
	downloadStruct.Status = core.StatusIdle

	// Ensure the folder path exists or create it
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		errMsg := fmt.Sprintf("Folder path does not exist, creating: %s", folderPath)
		logFunction(errMsg)
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			errMsg := fmt.Sprintf("Failed to create folder path: %s", folderPath)
			logFunction(errMsg)
			return errors.New(errMsg)
		}
		logFunction(fmt.Sprintf("Folder created: %s", folderPath))
	}

	filePath := filepath.Join(folderPath, downloadStruct.FileName)
	downloadStruct.Location = filePath

	return nil
}

func startHTTPRequest(downloadStruct *core.Resource, logFunction func(string)) (*http.Response, error) {
	client := &http.Client{
		Timeout: 10 * time.Minute,
	}

	url := downloadStruct.UrlGetter.GetUrl()
	resp, err := client.Get(url)
	if err != nil {
		downloadStruct.Status = core.StatusFailed
		errMsg := fmt.Sprintf("Failed to start download: %v", err)
		logFunction(errMsg)
		return nil, errors.New(errMsg)
	}

	if resp.StatusCode != http.StatusOK {
		downloadStruct.Status = core.StatusFailed
		errMsg := fmt.Sprintf("HTTP %d error while fetching data", resp.StatusCode)
		logFunction(errMsg)
		return nil, errors.New(errMsg)
	}

	return resp, nil

}

func createOutputFile(downloadStruct *core.Resource, logFunction func(string)) (*os.File, error) {
	file, err := os.Create(downloadStruct.Location)
	if err != nil {
		downloadStruct.Status = core.StatusFailed
		errMsg := fmt.Sprintf("Error creating file: %v", err)
		logFunction(errMsg)
		return nil, errors.New(errMsg)
	}

	return file, nil

}

func initializeDownloadInfo(resp *http.Response, downloadStruct *core.Resource) {
	downloadStruct.Info.Size = float64(resp.ContentLength)
	downloadStruct.Info.StartTime = time.Now()
}
