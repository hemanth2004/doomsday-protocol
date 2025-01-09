package submodels

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hemanth2004/doomsday-protocol/dday/core/guides"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"
	"github.com/hemanth2004/doomsday-protocol/dday/util"
	"github.com/hemanth2004/doomsday-protocol/dday/util/tree"
)

type GuideTreeModel struct {
	Width  int
	Height int

	GuidesPath        string
	ReadGuideCallback func(string)

	Focused  bool
	Viewport viewport.Model
	Tree     tree.Model

	viewportOffset int // Tracks the current viewport scroll position
	visibleLines   int // Number of lines visible in viewport
	Scrollbar      ScrollbarModel
}

func (m GuideTreeModel) Init() tea.Cmd {
	return nil
}

type NavigateToGuideMsg guides.Guide

func (m GuideTreeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	prevCursor := m.Tree.Cursor()

	switch msg := msg.(type) {
	case ResizeMsgL2:
		m.Width = msg.Width
		m.Height = msg.Height
		m.Viewport = viewport.New(m.Width-1, m.Height-2)
		m.visibleLines = m.Height - 2 // Set visible lines
		m.Tree.SetCursor(0)
		m.viewportOffset = 0
		m.Scrollbar.Height = m.Height - 2
		m.Scrollbar.ViewHeight = m.Height - 2

	case NavigateToGuideMsg:
		m.Tree.SetCursor(0)
		m.viewportOffset = 0

	case tea.KeyMsg:
		if m.Focused {
			// Handle tree navigation first
			var cmd tea.Cmd
			m.Tree, cmd = m.Tree.Update(msg)
			if cmd != nil {
				cmds = append(cmds, cmd)
			}

			// Adjust viewport based on cursor position
			cursor := m.Tree.Cursor()
			// Scroll down if cursor is near bottom
			if cursor >= m.viewportOffset+m.visibleLines-2 {
				m.viewportOffset = cursor - m.visibleLines + 2
				if m.viewportOffset < 0 {
					m.viewportOffset = 0
				}
			}
			// Scroll up if cursor is near top
			if cursor <= m.viewportOffset+1 {
				m.viewportOffset = cursor - 1
				if m.viewportOffset < 0 {
					m.viewportOffset = 0
				}
			}
			// Update viewport position
			m.Viewport.SetYOffset(m.viewportOffset)

			// Handle cursor change callback
			if prevCursor != cursor {
				node, err := m.Tree.GetSelectedNode()
				if err != nil {
					panic(err)
				}
				go m.ReadGuideCallback(node.Desc)
			}
		}
	}

	m.Tree.SetNodes(tree.GenerateGuideTree(m.GuidesPath))
	m.Viewport.SetContent(m.Tree.View())

	// Update scrollbar state
	m.Scrollbar.ContentHeight = strings.Count(m.Tree.View(), "\n") + 1
	m.Scrollbar.ScrollOffset = m.viewportOffset

	return m, tea.Batch(cmds...)
}

func (m GuideTreeModel) View() string {
	var s string
	content := fmt.Sprintf("GUIDES "+"\n%s", styles.DebugStyle.Render(util.DrawLine(m.Width))+"\n")
	//content += m.Viewport.View()

	// Render main content with scrollbar
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
