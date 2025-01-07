package submodels

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"
	"github.com/hemanth2004/doomsday-protocol/dday/util"
)

type TextViewerModel struct {
	Width  int
	Height int

	Focused   bool
	Path      string
	Viewport  viewport.Model
	Heading   string
	Content   string
	Scrollbar ScrollbarModel
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

		m.Viewport = viewport.New(m.Width-1, m.Height-2)

		m.Scrollbar.Height = m.Height - 2
		m.Scrollbar.ViewHeight = m.Height - 2

	case core.ChangeViewingGuideMsg:
		m.Path = string(msg)
		m.GetSetViewerContent()

	case tea.KeyMsg:
		if m.Focused {
			m.Viewport, cmd = m.Viewport.Update(msg)
			cmds = append(cmds, cmd)
		}

	}

	m.GetSetViewerContent()

	content := m.Content
	m.Scrollbar.ContentHeight = strings.Count(content, "\n") + 1
	m.Scrollbar.ScrollOffset = m.Viewport.YOffset

	return m, tea.Batch(cmds...)
}

func (m TextViewerModel) View() string {
	var s string
	Heading := util.TrimString(fmt.Sprintf("CONTENT | "+"%s", m.Heading), m.Width-2)
	Path := "" //lipgloss.NewStyle().Foreground(styles.TertiaryColor).Width(m.Width).Height(1).Render(fmt.Sprintf("[%s]", m.Path))
	content := fmt.Sprintf("%s\n%s%s", Heading, Path, styles.DebugStyle.Render(util.DrawLine(m.Width))+"\n")

	// Add scrollbar if content exceeds viewport
	viewportContent := m.Viewport.View()
	mainContent := viewportContent
	if m.Scrollbar.ContentHeight > m.Scrollbar.ViewHeight {
		mainContent = lipgloss.JoinHorizontal(
			lipgloss.Top,
			viewportContent,
			m.Scrollbar.View(),
		)
	}

	mainContent = content + mainContent
	if m.Focused {
		s += styles.PanelHighlightStyle.Width(m.Width).Height(m.Height).Render(mainContent)
	} else {
		s += styles.PanelStyle.Width(m.Width).Height(m.Height).Render(mainContent)
	}

	return s
}

func (m *TextViewerModel) GetSetViewerContent() {
	path := m.Path

	// Check if path is a directory
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	if fileInfo.IsDir() {
		// If it's a directory, just show the folder name
		m.Heading = filepath.Base(path)
		m.Content = m.Heading
	} else {
		// If it's a file, read and format content
		source, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}

		m.Heading = filepath.Base(path)
		fileExt := filepath.Ext(path)

		// Check if format is supported
		isSupported := false
		for _, format := range core.SupportedFormats {
			if format == fileExt {
				isSupported = true
				break
			}
		}

		if !isSupported {
			m.Content = "Unsupported format"
		} else {
			// Check if format needs special formatting
			for _, formatter := range core.FurtherFormattedFormats {
				if formatter.Extension == fileExt {
					m.Content = formatter.Formatter(string(source), m.Width-3, m.Height-2)
					m.Viewport.SetContent(m.Content)
					return
				}
			}
			// If no special formatting needed, display raw content
			m.Content = string(source)
		}
	}

	m.Viewport.SetContent(m.Content)
}
