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
	Path     string
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

		m.Viewport = viewport.New(m.Width-2, m.Height-2)

	case tea.KeyMsg:
		if m.Focused {
			m.Viewport, cmd = m.Viewport.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	m.GetSetViewerContent()

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

func (m *TextViewerModel) GetSetViewerContent() {
	path := m.Path

	source, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if strings.HasSuffix(path, ".md") {
		m.Content = string(markdown.Render(
			string(source),
			m.Width-2,
			0,
		))

		// Remove carriage returns
		m.Content = strings.ReplaceAll(m.Content, "\r", "")

	} else {
		m.Content = string(source)
	}

	m.Viewport.SetContent(m.Content)
}
