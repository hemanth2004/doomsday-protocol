package core

import (
	"time"

	"github.com/hemanth2004/doomsday-protocol/dday/core/guides"
)

// N number of Downloads downloading at a time
// Rest are queued
// Appears on the Downloads list
type Resource struct {
	Name        string
	Description string
	Note        string
	Tier        int
	Guide       guides.Guide

	UrlGetter        UrlGetter
	FileName         string
	Location         string
	InitiateDownload func(path string, logFunction func(string), downloadStruct *Resource) error
	Info             ResourceInformation
	Status           DownloadStatus
	Error            error
	ControlChannel   chan DownloadControl // Channel for controlling the download

	CustomResource bool
}

func (r *Resource) SpacePressed() {
	if r.Status == StatusDownloading {
		r.PauseResource()
	} else if r.Status == StatusPaused {
		r.ResumeResource()
	}
}

func (r *Resource) EnterPressed() {

}

func (r *Resource) PauseResource() {
	r.ControlChannel <- Pause
}

func (r *Resource) ResumeResource() {
	r.ControlChannel <- Start
}

type DownloadStatus string

const (
	StatusQueued      DownloadStatus = "Queued"
	StatusDownloading DownloadStatus = "Downloading"
	StatusPaused      DownloadStatus = "Paused"
	StatusCompleted   DownloadStatus = "Completed"
	StatusFailed      DownloadStatus = "Failed"
	StatusIdle        DownloadStatus = "Idle"
)

type DownloadControl int

const (
	Start DownloadControl = iota // Synonymous with Continue
	Pause
	Cancel
)

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

var FillerResource = Resource{
	Name:             "tableFiller",
	Description:      "tableFiller",
	Note:             "tableFiller",
	Tier:             1,
	UrlGetter:        UrlGetter{},
	FileName:         "example.txt",
	InitiateDownload: func(path string, logFunction func(string), downloadStruct *Resource) error { return nil },
	Info:             ResourceInformation{},
	Status:           StatusQueued,
	Error:            nil,
}
