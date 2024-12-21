package ui

import (
	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"
)

// Alternative to bubbles/key and bubbles/help
// The default bubble-tea help tries to list every keymap as help
// With 'HelpSet', I'm trying to compress two or more similar instructions to one "help" entry
// Example:  "ctrl+q/e: switch tabs" will take a lot longer with bubbles/key

type HelpSet [][2]string

func (hSet HelpSet) View(seperator string) (s string) {
	for i, h := range hSet {
		s += styles.HelpStyle1.Render(h[0]) + styles.HelpStyle2.Render(" "+h[1])
		if i < len(hSet)-1 {
			s += seperator
		}
	}
	return
}
