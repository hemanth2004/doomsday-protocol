package ui

import (
	"github.com/hemanth2004/doomsday-protocol/dday/core"
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
	Guides       GuidesModel
	HelpSet      HelpSet
}

// bubbletea methods
func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	// Just setting window properties
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width - 4
		m.height = msg.Height

		m.Guides.AllottedHeight = m.height - topSectionHeight - bottomSectionHeight
		m.Guides.AllottedWidth = m.width

		m.Downloads.AllottedHeight = m.height - topSectionHeight - bottomSectionHeight
		m.Downloads.AllottedWidth = m.width

		m.NewResource.AllottedHeight = m.height - topSectionHeight - bottomSectionHeight
		m.NewResource.AllottedWidth = m.width

		cmds = append(cmds, NewMainResizedCmd())
	}

	if m.CurrentState.Index() == 0 {
		updatedGuides, cmd := m.Guides.Update(msg)
		if updatedGuides, ok := updatedGuides.(GuidesModel); ok {
			m.Guides = updatedGuides
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

	// Handle messages for child models
	switch msg := msg.(type) {

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
	stateNames := []string{"guides", "downloads", "new resource(+)"}
	var stateLine string
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
	s += lipgloss.Place(m.width, topSectionHeight, lipgloss.Center, lipgloss.Center, stateLine) + "\n"

	defaultHelpText := m.HelpSet.View("    ")
	guidesHelpText := m.Guides.HelpSet.View("    ")
	downloadsHelpText := m.Downloads.HelpSet[m.Downloads.CurrentWindow.Index()].View("    ")
	newresourceHelpText := m.NewResource.HelpSet.View("    ")

	sep := " | "
	helpBoxContent := ""

	if m.CurrentState.Index() == 0 {
		s += m.Guides.View()
		helpBoxContent += guidesHelpText + sep + defaultHelpText
	} else if m.CurrentState.Index() == 1 {
		s += m.Downloads.View()
		helpBoxContent += downloadsHelpText + sep + defaultHelpText
	} else if m.CurrentState.Index() == 2 {
		s += m.NewResource.View()
		helpBoxContent += newresourceHelpText + sep + defaultHelpText
	}

	return s + lipgloss.Place(m.width, bottomSectionHeight-2, lipgloss.Center, lipgloss.Bottom, helpBoxContent)
}

// tea.Msg defined to propogate window resize changes downwards
type MainResizedMsg struct{}

// NewMainResizedCmd creates a command that returns a MainResizedMsg
func NewMainResizedCmd() tea.Cmd {
	return func() tea.Msg {
		return MainResizedMsg{}
	}
}
