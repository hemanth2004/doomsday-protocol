package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jroimartin/gocui"
)

var (
	viewArr = []string{"console", "downloads", "packages"}
	active  = 0
)

func LogToConsole(g *gocui.Gui, s string) {
	g.Update(func(g *gocui.Gui) error {
		out, err := g.View("console")
		if err != nil {
			return err
		}
		fmt.Fprintln(out, s)
		return nil
	})
}

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func scroll(g *gocui.Gui, v *gocui.View, dy int) error {

	// Get the size and position of the view.
	_, y := v.Size()
	ox, oy := v.Origin()

	// If we're at the bottom...
	if oy+dy > strings.Count(v.ViewBuffer(), "\n")-y-1 {
		// Set autoscroll to normal again.
		v.Autoscroll = true
	} else {
		// Set autoscroll to false and scroll.
		v.Autoscroll = false
		v.SetOrigin(ox, oy+dy)
	}

	return nil
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (active + 1) % len(viewArr)
	name := viewArr[nextIndex]

	out, err := g.View("downloads")
	if err != nil {
		return err
	}
	fmt.Fprintln(out, "Going from view "+v.Name()+" to "+name)

	if _, err := setCurrentViewOnTop(g, name); err != nil {
		return err
	}

	if nextIndex == 0 || nextIndex == 3 {
		g.Cursor = true
	} else {
		g.Cursor = false
	}

	active = nextIndex
	return nil
}

const lrRatio = 3.5

func layout(g *gocui.Gui) error {
	_maxX, _maxY := g.Size()
	maxX, maxY := float64(_maxX), float64(_maxY)
	if v, err := g.SetView("console", int(maxX/lrRatio), int(maxY/2), int(maxX-1), int(maxY-1)); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Console"
		v.Editable = true
		v.Wrap = true
		v.Autoscroll = true

		if _, err = setCurrentViewOnTop(g, "console"); err != nil {
			return err
		}
	}

	if v, err := g.SetView("downloads", int(maxX/lrRatio), 0, int(maxX-1), int(maxY/2-1)); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Downloads"
		v.Wrap = true
		v.Autoscroll = true
	}
	if v, err := g.SetView("packages", 0, 0, int(maxX/lrRatio-1), int(maxY-1)); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Doomsday Packages"
		v.Wrap = true
		v.Autoscroll = true
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func SetupTui() (*gocui.Gui, error) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyPgup, gocui.ModNone, msgUp); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyPgdn, gocui.ModNone, msgDown); err != nil {
		log.Panicln(err)
	}

	return g, err
}

func msgUp(g *gocui.Gui, v *gocui.View) error {
	return msgScroll(g, v, -5)
}
func msgDown(g *gocui.Gui, v *gocui.View) error {
	return msgScroll(g, v, 5)
}
func msgScroll(g *gocui.Gui, v *gocui.View, delta int) error {

	// Current position
	_, viewHeight := v.Size()
	ox, oy := v.Origin()
	if viewHeight+oy+delta > len(v.ViewBufferLines())-1 {
		// We are at the bottom, enable Autoscroll
		v.Autoscroll = true
	} else {
		// Set autoscroll to false and scroll.
		v.Autoscroll = false
		v.SetOrigin(ox, oy+delta)
	}
	return nil
}
