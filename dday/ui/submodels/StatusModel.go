package submodels

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/debug"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"
)

type StatusModel struct {
	Width  int
	Height int

	Focused           bool
	ApplicationObject *core.Application
	Progress          MultilineProgress
	FillStyle         lipgloss.Style
	BackStyle         lipgloss.Style
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
		m.Progress = NewMultilineProgress(m.Width, m.Height, styles.StatusFillStyle, styles.StatusBackgroundStyle)

	case tea.KeyMsg:
		if m.Focused {
			switch msg.String() {
			case "alt+d":
				if m.ApplicationObject.ProtocolInitiated {
					if m.ApplicationObject.ProtocolPaused {
						m.ApplicationObject.ResumeProtocol()
					} else {
						m.ApplicationObject.PauseProtocol()
					}
				} else {
					m.ApplicationObject.OrderToInitiateProtocol()
				}
			}
		}
	}

	m.Progress.SetPercent(m.ApplicationObject.ResourceList.GetOverallProgress())

	return m, tea.Batch(cmds...)
}

func (m StatusModel) View() string {
	var s string

	overlay := ""

	if !m.ApplicationObject.ProtocolInitiated {
		overlay += "PROTOCOL UNINITIATED\n" +
			strconv.Itoa(len(m.ApplicationObject.ResourceList.DefaultResources)) + " Default" + ", " +
			strconv.Itoa(len(m.ApplicationObject.ResourceList.CustomResources)) + " Custom" + "\n" +
			"INITIATE [ALT + D]"
	} else {
		overlay += strconv.FormatFloat(m.ApplicationObject.ResourceList.GetOverallProgress(), 'f', 2, 64) + "%" + "\n" +
			//"ETA: 13 hours 6 mins" + "\n" +
			"PAUSE [ALT + D]"
	}

	debug.Log(overlay)
	overlayExpanded := lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, overlay)
	content := m.Progress.View(overlayExpanded)

	if m.Focused {
		// Highlight Window if active
		s += styles.CtrlPanelHighlightStyle.Width(m.Width).Height(m.Height).Render(content)
	} else {
		s += styles.CtrlPanelStyle.Width(m.Width).Height(m.Height).Render(content)
	}
	return s
}
