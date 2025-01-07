package core

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Everything that the entire application should keep track of
type Application struct {
	TeaProgram *tea.Program

	GuidesFolderPath string

	ProtocolInitiated bool
	ResourceList      ResourceList

	LogFunction func(string)
	Logs        Log
}

type TickMsg time.Time

func (a *Application) StartPeriodicTicks(deltaTime int) {
	ticker := time.NewTicker(time.Duration(deltaTime) * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		a.TeaProgram.Send(TickMsg(time.Now()))
	}
}

func (a *Application) GuideViewerCallback(path string) {
	a.TeaProgram.Send(ChangeViewingGuideMsg(path))
}

// Initiate every resource download
// TODO: Limit to N at a time and queue the rest
func (a *Application) InitiateProtocol() {
	a.ProtocolInitiated = true
	a.LogFunction("Initiating doomsday-protocol.")
	for i, r := range a.ResourceList.DefaultResources {
		if r.Tier == 0 {
			go a.ResourceList.DefaultResources[i].InitiateDownload("downloads/tier0/", a.LogFunction, &a.ResourceList.DefaultResources[i])
		} else if r.Tier == 1 {
			go a.ResourceList.DefaultResources[i].InitiateDownload("downloads/tier1/", a.LogFunction, &a.ResourceList.DefaultResources[i])
		}

	}
}

func (a *Application) GetProgress() float64 {
	return 0.0
}

// Logs Logic
type Log = [][2]string

// Log a message on the console
func (a *Application) Log(s string) {
	a.Logs = append(a.Logs, [2]string{s, time.Now().Format(time.DateTime)})
	a.TeaProgram.Send(LoggedMsg(a.Logs))
}

type LoggedMsg [][2]string
