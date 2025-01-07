package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/submodels"
	"github.com/hemanth2004/doomsday-protocol/dday/util"
)

func InitHomeHelpSet() []HelpSet {
	return []HelpSet{
		{
			{"↑/↓", "navigate"},
			{"tab", "switch tabs"},
		},
		{
			{"↑/↓", "navigate"},
			{"tab", "switch tabs"},
		},
		{
			{"ctrl+d+p", "initiate/pause/resume protocol"},
			{"tab", "switch tabs"},
		},
	}
}

type HomeModel struct {
	Height int
	Width  int

	CurrentWindow *util.StateHandler[int]
	TextViewer    submodels.TextViewerModel
	GuideTree     submodels.GuideTreeModel
	StatusModel   submodels.StatusModel
	HelpSet       []HelpSet // Will be used by parent model to render help
}

func (m HomeModel) Init() tea.Cmd {
	return nil
}

func (m HomeModel) GetPanelDimensions() (int, int, int, int, int) {
	leftWidth := int(0.25 * float64(m.Width))
	rightWidth := m.Width - leftWidth

	leftPrimaryHeight := m.Height - 4
	leftSecondaryHeight := 3

	rightHeight := m.Height

	return leftWidth, rightWidth, leftPrimaryHeight - 1, leftSecondaryHeight, rightHeight
}

func (m HomeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//var cmd Tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case ResizeMsgL1:
		m.Width = msg.Width
		m.Height = msg.Height

		leftWidth, rightWidth, leftPrimaryHeight, leftSecondaryHeight, rightHeight := m.GetPanelDimensions()

		if updatedGuides, _ := m.GuideTree.Update(submodels.ResizeMsgL2{Width: leftWidth, Height: leftPrimaryHeight}); updatedGuides != nil {
			m.GuideTree = updatedGuides.(submodels.GuideTreeModel)
		}

		if updatedViewer, _ := m.TextViewer.Update(submodels.ResizeMsgL2{Width: rightWidth, Height: rightHeight}); updatedViewer != nil {
			m.TextViewer = updatedViewer.(submodels.TextViewerModel)
		}

		if updatedStatus, _ := m.StatusModel.Update(submodels.ResizeMsgL2{Width: leftWidth, Height: leftSecondaryHeight}); updatedStatus != nil {
			m.StatusModel = updatedStatus.(submodels.StatusModel)
		}

	case tea.KeyMsg:
		if msg.String() == "tab" {
			m.CurrentWindow.NextState()
		}
		if msg.String() == "shift+tab" {
			m.CurrentWindow.PrevState()
		}

		m.GuideTree.Focused = m.CurrentWindow.Index() == 0
		m.TextViewer.Focused = m.CurrentWindow.Index() == 1
		m.StatusModel.Focused = m.CurrentWindow.Index() == 2

		if m.GuideTree.Focused {
			if updatedGuideTree, cmd := m.GuideTree.Update(msg); updatedGuideTree != nil {
				m.GuideTree = updatedGuideTree.(submodels.GuideTreeModel)
				cmds = append(cmds, cmd)
			}
		}
		if m.TextViewer.Focused {
			if updatedViewer, cmd := m.TextViewer.Update(msg); updatedViewer != nil {
				m.TextViewer = updatedViewer.(submodels.TextViewerModel)
				cmds = append(cmds, cmd)
			}
		}
		if m.StatusModel.Focused {
			if updatedStatus, cmd := m.StatusModel.Update(msg); updatedStatus != nil {
				m.StatusModel = updatedStatus.(submodels.StatusModel)
				cmds = append(cmds, cmd)
			}
		}

	case core.ChangeViewingGuideMsg:
		if updatedViewer, cmd := m.TextViewer.Update(msg); updatedViewer != nil {
			m.TextViewer = updatedViewer.(submodels.TextViewerModel)
			cmds = append(cmds, cmd)
		}
	}

	m.GuideTree.Focused = m.CurrentWindow.Index() == 0
	m.TextViewer.Focused = m.CurrentWindow.Index() == 1
	m.StatusModel.Focused = m.CurrentWindow.Index() == 2

	return m, tea.Batch(cmds...)
}

func (m HomeModel) View() string {
	var s string

	// Joining the windows
	var topLeftWindow, bottomLeftWindow, rightWindow string
	topLeftWindow = m.GuideTree.View()
	bottomLeftWindow = m.StatusModel.View()
	rightWindow = m.TextViewer.View()
	leftSection := lipgloss.JoinVertical(lipgloss.Top, topLeftWindow, bottomLeftWindow)
	s += lipgloss.JoinHorizontal(lipgloss.Top, leftSection, rightWindow)

	return s + "\n"
}
