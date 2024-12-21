package submodels

import (
	"fmt"
	"os"
	"strings"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"
	"github.com/hemanth2004/doomsday-protocol/dday/util"
)

type TextViewerModel struct {
	Width  int
	Height int

	Focused  bool
	Viewport viewport.Model
	Content  string
}

func (m TextViewerModel) Init() tea.Cmd {
	return nil
}

func (m TextViewerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case ResizeMsgL2:
		m.Width = msg.Width
		m.Height = msg.Height

		m.Viewport = viewport.New(m.Width-2, m.Height-5)

	case tea.KeyMsg:
		if m.Focused {
			m.Viewport, cmd = m.Viewport.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	path := "README.md"
	source, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	m.Content = string(markdown.Render(
		string(source),
		m.Width-2,
		0,
	))
	// Replace escape sequences with their text representation
	//m.Content = strings.ReplaceAll(m.Content, "\n", "\\n")
	m.Content = strings.ReplaceAll(m.Content, "\r", "")
	//m.Content = strings.ReplaceAll(m.Content, "\t", "\\t")
	//m.Content = strings.ReplaceAll(m.Content, "\b", "\\b")
	//m.Content = strings.ReplaceAll(m.Content, "\f", "\\f")
	//m.Content = strings.ReplaceAll(m.Content, "\v", "\\v")
	//m.Content = strings.ReplaceAll(m.Content, "\a", "\\a")
	m.Viewport.SetContent(m.Content)

	return m, tea.Batch(cmds...)
}

func (m TextViewerModel) View() string {
	var s string
	content := fmt.Sprintf("CONTENT "+"\n%s", styles.DebugStyle.Render(util.DrawLine(m.Width))+"\n")
	content += m.Viewport.View()

	if m.Focused {
		// Highlight Window if active
		s += styles.PanelHighlightStyle.Width(m.Width).Height(m.Height).Render(content)
	} else {
		s += styles.PanelStyle.Width(m.Width).Height(m.Height).Render(content)
	}

	return s
}
