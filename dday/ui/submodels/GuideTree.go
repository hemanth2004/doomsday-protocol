package submodels

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/hemanth2004/doomsday-protocol/dday/debug"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"
	"github.com/hemanth2004/doomsday-protocol/dday/util"
	"github.com/hemanth2004/doomsday-protocol/dday/util/tree"
)

type GuideTreeModel struct {
	Width  int
	Height int

	Focused  bool
	Viewport viewport.Model
	Tree     tree.Model
}

func (m GuideTreeModel) Init() tea.Cmd {
	return nil
}

func (m GuideTreeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case ResizeMsgL2:
		m.Width = msg.Width
		m.Height = msg.Height
		debug.Log("GuideTreeModel" + "WindowSizeMsg" + fmt.Sprintf("%+v", msg))

	case tea.KeyMsg:
		if m.Focused {
			switch msg.String() {
			}
		}
	}

	return m, tea.Batch(cmds...)
}

func (m GuideTreeModel) View() string {
	var s string
	content := fmt.Sprintf("GUIDES "+"\n%s", styles.DebugStyle.Render(util.DrawLine(m.Width))+"\n")
	content += "Guide Tree Content"

	if m.Focused {
		// Highlight Window if active
		s += styles.PanelHighlightStyle.Width(m.Width).Height(m.Height).Render(content)
	} else {
		s += styles.PanelStyle.Width(m.Width).Height(m.Height).Render(content)
	}

	return s
}
