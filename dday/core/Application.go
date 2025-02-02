package core

import (
	"errors"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hemanth2004/doomsday-protocol/dday/core/guides"
)

// Everything that the entire application should keep track of
type Application struct {
	Config     *Config
	TeaProgram *tea.Program

	GuidesFolderPath string

	ProtocolInitiated bool
	ProtocolPaused    bool

	ResourceList ResourceList

	LogFunction func(string)
	LogsContent Logs
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
	a.TeaProgram.Send(guides.ChangeViewingGuideMsg(path))
}

// Initiate every resource download
// TODO: Limit to N at a time and queue the rest
func (a *Application) InitiateProtocol() {
	a.ProtocolInitiated = true
	a.ProtocolPaused = false
	a.LogFunction("Initiating doomsday-protocol.")
	for i, r := range a.ResourceList.DefaultResources {
		if r.Tier == 0 {
			go a.ResourceList.DefaultResources[i].InitiateDownload("downloads/tier0/", a.LogFunction, &a.ResourceList.DefaultResources[i])
		} else if r.Tier == 1 {
			go a.ResourceList.DefaultResources[i].InitiateDownload("downloads/tier1/", a.LogFunction, &a.ResourceList.DefaultResources[i])
		}

	}
}

func (a *Application) OrderToInitiateProtocol() {
	go a.InitiateProtocol()
}

func (a *Application) PauseProtocol() {
	a.ProtocolPaused = true
	a.ResourceList.PauseAllResources()
}

func (a *Application) ResumeProtocol() {
	a.ProtocolPaused = false
	a.ResourceList.ResumeAllResources()
}

var CurrentApplicationInstance *Application

func GetWorkingDirectory() string {
	mydir, err := os.Getwd()
	if err == nil {
		return mydir
	} else {
		panic(errors.New("failed to get working directory"))
	}

}
