package ui

import (
	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/core/guides"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"
	"github.com/hemanth2004/doomsday-protocol/dday/util"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainModel struct {
	width  int
	height int

	Application  *core.Application
	ResourceList *core.ResourceList

	CurrentState *util.StateHandler[string]
	Downloads    DownloadsModel
	NewResource  NewResourceModel
	Home         HomeModel
	HelpSet      HelpSet
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	if m.Downloads.NavigateToCtrlPanel == nil {
		m.Downloads.NavigateToCtrlPanel = m.NavigateToCtrlPanel
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width - 4
		m.height = msg.Height

		// pass down a new viewport message to child models
		viewportMsg := ResizeMsgL1{
			Width:  m.width,
			Height: m.height - topSectionHeight - bottomSectionHeight,
		}

		updatedGuides, cmd := m.Home.Update(viewportMsg)
		if updatedGuides, ok := updatedGuides.(HomeModel); ok {
			m.Home = updatedGuides
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}

		updatedDownloads, cmd := m.Downloads.Update(viewportMsg)
		if updatedDownloads, ok := updatedDownloads.(DownloadsModel); ok {
			m.Downloads = updatedDownloads
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}

		updatedNewForm, cmd := m.NewResource.Update(viewportMsg)
		if updatedNewForm, ok := updatedNewForm.(NewResourceModel); ok {
			m.NewResource = updatedNewForm
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}

	case tea.KeyMsg:
		if msg.String() == "ctrl+e" {
			m.CurrentState.NextState()
		}
		if msg.String() == "ctrl+q" {
			m.CurrentState.PrevState()
		}
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	// if msg is not a window size msg,
	// then pass down msg to child models
	//if _, ok := msg.(tea.WindowSizeMsg); !ok {
	if m.CurrentState.Index() == 0 {
		updatedGuides, cmd := m.Home.Update(msg)
		if updatedGuides, ok := updatedGuides.(HomeModel); ok {
			m.Home = updatedGuides
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}
	} else if m.CurrentState.Index() == 1 {
		updatedDownloads, cmd := m.Downloads.Update(msg)
		if updatedDownloads, ok := updatedDownloads.(DownloadsModel); ok {
			m.Downloads = updatedDownloads
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}
	} else if m.CurrentState.Index() == 2 {

		updatedNewResource, cmd := m.NewResource.Update(msg)
		if updatedNewResource, ok := updatedNewResource.(NewResourceModel); ok {
			m.NewResource = updatedNewResource
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}
	}
	//}

	return m, tea.Batch(cmds...)
}

// Some constant variables that'll help me render UI better
const (
	topSectionHeight    int = 1
	bottomSectionHeight int = 4 // -2 for borders
)

func (m MainModel) View() string {

	var s string

	// Rendering first state line
	stateNames := []string{"home", "downloads", "new resource(+)"}
	var stateLine string
	//stateLine += styles.Accent3InvertedStyle.Render("[^Q] ← ") + " "
	for i, name := range stateNames {
		if i == m.CurrentState.Index() {
			stateLine += styles.MainStyle.Render(name)
		} else {
			stateLine += styles.DebugStyle.Render(name)
		}
		if i < len(stateNames)-1 {
			stateLine += "  •  "
		}
	}
	//stateLine += " " + styles.Accent3InvertedStyle.Render(" → [^E]")
	s += lipgloss.Place(m.width, topSectionHeight, lipgloss.Center, lipgloss.Center, stateLine) + "\n"

	defaultHelpText := m.HelpSet.View("  ")

	sep := " | "
	helpBoxContent := ""

	if m.CurrentState.Index() == 0 {

		s += m.Home.View()
		homeHelpText := m.Home.HelpSet[m.Home.CurrentWindow.Index()].View("  ")
		helpBoxContent += homeHelpText + sep + defaultHelpText
	} else if m.CurrentState.Index() == 1 {
		s += m.Downloads.View()
		downloadsHelpText := m.Downloads.HelpSet[m.Downloads.CurrentWindow.Index()].View("  ")
		helpBoxContent += downloadsHelpText + sep + defaultHelpText
	} else if m.CurrentState.Index() == 2 {
		s += m.NewResource.View()
		newresourceHelpText := m.NewResource.HelpSet.View("  ")
		helpBoxContent += newresourceHelpText + sep + defaultHelpText
	}

	return s + lipgloss.Place(m.width, bottomSectionHeight-3, lipgloss.Center, lipgloss.Bottom, helpBoxContent) + "\n"
}

func (m *MainModel) NavigateToGuide(guide guides.Guide) {
	m.CurrentState.SetState("home")
	m.Home.CurrentWindow.SetState(2)

}

func (m *MainModel) NavigateToCtrlPanel() {
	m.CurrentState.SetState("home")
	m.Home.CurrentWindow.SetState(0)
}

type ResizeMsgL1 struct {
	Width  int
	Height int
}
