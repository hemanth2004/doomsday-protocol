package ui

import (
	"strconv"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"
	"github.com/hemanth2004/doomsday-protocol/dday/util"
)

// script to handle detailed view of a download thats happening currently

var (
	headStyle = lipgloss.NewStyle().Bold(true).Foreground(styles.Black).Background(styles.Yellow)

	statusSuccessStyle     = lipgloss.NewStyle().Foreground(styles.Black).Background(styles.Green3)
	statusDownloadingStyle = lipgloss.NewStyle().Foreground(styles.Black).Background(styles.Yellow)
	statusFailStyle        = lipgloss.NewStyle().Foreground(styles.Black).Background(styles.Red)
	statusWaitStyle        = lipgloss.NewStyle().Foreground(styles.Black).Background(styles.Blue)

	detailsStyle = lipgloss.NewStyle().Foreground(styles.Black).Background(styles.Grey78)

	underlineStyle = lipgloss.NewStyle().Underline(true).Foreground(styles.PaleTurquoise4)

	button1Style = lipgloss.NewStyle().Foreground(styles.Aquamarine1).Background(styles.Blue)
	buttonBorder = lipgloss.NewStyle().Border(lipgloss.RoundedBorder())
)

type InspectModel struct {
	Width  int
	Height int

	InspectingDownload *core.Resource
	viewport           viewport.Model
}

func (m InspectModel) Init() tea.Cmd {
	return nil
}

func (m InspectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewport = viewport.New(m.Width, m.Height-3)

	case tea.KeyMsg:
		m.viewport, cmd = m.viewport.Update(msg)
	}

	m.viewport.SetContent(m.UpdateContent())

	return m, cmd
}

func (m InspectModel) View() string {

	guideButton := "\n" + lipgloss.Place(m.Width, 3, lipgloss.Center, lipgloss.Center, button1Style.Render(buttonBorder.Render(" Open Guide (ctrl+g) ")))

	horizontalLine := styles.DebugStyle.Render(util.DrawLine(m.Width))
	return m.viewport.View() + horizontalLine + guideButton
}

func (m InspectModel) UpdateContent() string {

	var s string

	if m.InspectingDownload.Name == "" || m.InspectingDownload.Name == "example" {
		s = "No download selected on the table."
		return s
	}

	horizontalLine := styles.DebugStyle.Render(util.DrawLine(m.Width))

	var header string
	header += headStyle.Render(util.MarginHor(m.InspectingDownload.Name, 1))
	header += "\nTYPE: " + util.IfElse(m.InspectingDownload.CustomResource, "Custom", "Default")
	header += util.IfElse(!m.InspectingDownload.CustomResource, "\nTIER: "+strconv.Itoa(m.InspectingDownload.Tier), "") + "\n"
	var statusStyle lipgloss.Style
	switch m.InspectingDownload.Status {
	case core.StatusCompleted:
		statusStyle = statusSuccessStyle
	case core.StatusDownloading:
		statusStyle = statusDownloadingStyle
	case core.StatusFailed:
		statusStyle = statusFailStyle
	default:
		statusStyle = statusWaitStyle
	}
	header += "STATUS: " + statusStyle.Render(util.MarginHor(string(m.InspectingDownload.Status), 1)) + "\n\n"

	var fileDetails string
	fileDetails += "Source URL:\n" + underlineStyle.Render(m.InspectingDownload.UrlGetter.RecentURLUsed)
	fileDetails += "\nLocation:\n" + underlineStyle.Render(m.InspectingDownload.Location) + "\n\n"

	var resourceDetails string
	resourceDetails += "\n\n" + detailsStyle.Render("# Description ") + "\n" +
		m.InspectingDownload.Description
	resourceDetails += "\n\n" + detailsStyle.Render("# Note ") + "\n" +
		util.IfElse(m.InspectingDownload.Note == "", "-", m.InspectingDownload.Note) + "\n\n"

	var allSources string
	allSources += "\n\n" + "All Sources:\n"

	s += header + fileDetails + horizontalLine + resourceDetails + horizontalLine + allSources

	return s
}
