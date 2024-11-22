package main

import (
	"log"
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/jroimartin/gocui"

	"github.com/hemanth2004/doomsday-protocol/dday"
)

var (
	app  = kingpin.New("dday-prtcl", "A command-line emergency application.").Terminate(nil)
	post = app.Command("post", "Post a message to a channel.")
)

func main() {
	g, _ := SetupTui()
	defer g.Close()

	var SimpleLog func(s string) = func(s string) {
		LogToConsole(g, s)
	}

	go dday.Resources[0].InitiateDownload("downloads/", SimpleLog, &dday.Resources[0])
	go dday.Resources[1].InitiateDownload("downloads/", SimpleLog, &dday.Resources[1])

	// Update loop
	go func(g *gocui.Gui) {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				g.Update(func(g *gocui.Gui) error {
					return dday.RenderDownloads(g)
				})
			}
		}
	}(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {

		log.Panicln(err)
	}

}
