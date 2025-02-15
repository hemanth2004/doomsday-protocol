package main

import (
	"fmt"
	"log"

	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/core/guides"
	"github.com/hemanth2004/doomsday-protocol/dday/debug"
	"github.com/hemanth2004/doomsday-protocol/dday/resources"
	"github.com/hemanth2004/doomsday-protocol/dday/ui"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	Application core.Application = core.Application{
		Config:            core.LoadConfig("."),
		ProtocolInitiated: false,
		GuidesFolderPath:  guides.GuidesFolderFromWorkingDirectory(),
		ResourceList: core.ResourceList{
			DefaultResources: resources.CoreResources,
		},
		LogsContent: make(core.Logs, 0),
	}

	p *tea.Program
)

// Options
const useAlternateBuffer = false
const tickInterval = 1000

func main() {

	// Set the terminal to use the alternate screen buffer
	if useAlternateBuffer {
		fmt.Print("\033[?1049h")
		// Reset the terminal to use the main screen buffer
		defer fmt.Print("\033[?1049l")
	}

	defer debug.Close()

	p = tea.NewProgram(ui.InitialTeaModel(&Application))
	Application.TeaProgram = p
	Application.LogFunction = DebugPrintGoroutine

	go Application.StartPeriodicTicks(tickInterval) // = 1 FPS
	core.CurrentApplicationInstance = &Application

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func DebugPrintGoroutine(message string) {
	go DebugPrint(&Application, message)
}
func DebugPrint(a *core.Application, message string) {
	a.Log(message)
}
