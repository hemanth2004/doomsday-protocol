package core

import (
	"time"
)

// 2-3 Downloads downloading at a time
// Rest are queued
// Appears on the Downloads list
type Resource struct {
	Name        string
	Description string
	Note        string
	Tier        int
	Guide       Guide

	UrlGetter        UrlGetter
	FileName         string
	Location         string
	InitiateDownload func(path string, logFunction func(string), infoStruct *Resource) error
	Info             ResourceInformation
	Status           DownloadStatus
	Error            error

	CustomResource bool
}

// Download info to construct the UI
type ResourceInformation struct {
	Done      float64   // Bytes downloaded so far
	Size      float64   // Total size of the download in bytes
	Bandwidth float64   // Current download speed in bytes per second
	ETA       uint64    // Estimated time remaining (formatted as a string, e.g., "10.5 seconds")
	StartTime time.Time // Time when the download started
	EndTime   time.Time // Time when the download completed
}

func (d ResourceInformation) ProgressPercent() float64 {
	if d.Size == 0 {
		return 0
	}
	return (d.Done / d.Size) * 100
}

type DownloadStatus string

const (
	StatusQueued      DownloadStatus = "Queued"
	StatusDownloading DownloadStatus = "Downloading"
	StatusPaused      DownloadStatus = "Paused"
	StatusCompleted   DownloadStatus = "Completed"
	StatusFailed      DownloadStatus = "Failed"
)

var EmptyResource = Resource{
	Name:        "example",
	Description: "example",
	Note:        "example",
	Tier:        1,
	UrlGetter: UrlGetter{
		Key:           "example",
		UpdatedURLURL: []string{},
		DefaultURLs: []string{
			"https://www.example.com",
		},
	},
	FileName:         "example.txt",
	InitiateDownload: func(path string, logFunction func(string), infoStruct *Resource) error { return nil },
	Info:             ResourceInformation{},
	Status:           StatusQueued,
	Error:            nil,
}
