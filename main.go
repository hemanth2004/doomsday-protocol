//go dday.Resources[0].InitiateDownload("downloads/", DebugPrint, &dday.Resources[0])
//go dday.Resources[1].InitiateDownload("downloads/", DebugPrint, &dday.Resources[1])

package main

import (
	"log"

	"github.com/hemanth2004/doomsday-protocol/dday"
	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/debug"
	"github.com/hemanth2004/doomsday-protocol/dday/ui"

	tea "github.com/charmbracelet/bubbletea"
)

var Application core.Application = core.Application{
	ResourceList: core.ResourceList{
		DefaultResources: dday.DefaultResources,
	},
	Logs: "",
}

var p *tea.Program

func main() {
	defer debug.Close()

	p = tea.NewProgram(ui.InitialTeaModel(&Application))
	Application.TeaProgram = p
	Application.LogFunction = DebugPrintGoroutine

	go Application.StartPeriodicTicks(500)
	go Application.InitiateProtocol()

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func DebugPrintGoroutine(message string) {
	go DebugPrint(&Application, message)
}
func DebugPrint(a *core.Application, message string) {
	a.Log(message)
	debug.Log(a.Logs)
	p.Send(core.LoggedMsg(a.Logs))
}
