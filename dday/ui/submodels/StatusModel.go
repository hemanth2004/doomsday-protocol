package submodels

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"
)

type StatusModel struct {
	Width  int
	Height int

	Focused bool
}

func (m StatusModel) Init() tea.Cmd {
	return nil
}

func (m StatusModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case ResizeMsgL2:
		m.Width = msg.Width
		m.Height = msg.Height

	case tea.KeyMsg:
		if m.Focused {
			switch msg.String() {
			}
		}
	}

	return m, tea.Batch(cmds...)
}

func (m StatusModel) View() string {
	var s string
	content := fmt.Sprintf("STATUS " + "\n")
	content += "Shortened status of all current downloads."

	if m.Focused {
		// Highlight Window if active
		s += styles.PanelHighlightStyle.Width(m.Width).Height(m.Height).Render(content)
	} else {
		s += styles.PanelStyle.Width(m.Width).Height(m.Height).Render(content)
	}
	return s
}
