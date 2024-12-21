package submodels

import (
	"strconv"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/debug"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"
	"github.com/hemanth2004/doomsday-protocol/dday/util"
)

// script to handle detailed view of a download thats happening currently

var (
	headStyle = lipgloss.NewStyle().Foreground(styles.Black).Background(styles.BrightYellow)

	statusSuccessStyle     = lipgloss.NewStyle().Foreground(styles.Black).Background(styles.BrightGreen)
	statusDownloadingStyle = lipgloss.NewStyle().Foreground(styles.Black).Background(styles.Yellow)
	statusFailStyle        = lipgloss.NewStyle().Foreground(styles.Black).Background(styles.Red)
	statusWaitStyle        = lipgloss.NewStyle().Foreground(styles.Black).Background(styles.Blue)

	detailsStyle = lipgloss.NewStyle().Foreground(styles.Black).Background(styles.White)

	underlineStyle = lipgloss.NewStyle().Underline(true).Foreground(styles.BrightBlue)

	button1Style = lipgloss.NewStyle().Foreground(styles.BrightWhite).Background(styles.Blue)
)

type InspectModel struct {
	Width  int
	Height int

	Focused            bool
	InspectingDownload *core.Resource
	viewport           viewport.Model
}

func (m InspectModel) Init() tea.Cmd {
	return nil
}

func (m InspectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case ResizeMsgL2:
		m.Width = msg.Width
		m.Height = msg.Height
		debug.Log("InspectModel: Resize: " + strconv.Itoa(msg.Width) + " " + strconv.Itoa(msg.Height))
		m.viewport = viewport.New(m.Width, m.Height-3)
	case tea.KeyMsg:
		debug.Log("Movingviewport " + msg.String())
		if msg.String() == "up" {
			m.viewport.HalfViewDown()
		} else if msg.String() == "down" {
			m.viewport.HalfViewUp()
		}

	}

	m.viewport.SetContent(m.UpdateContent())

	return m, cmd
}

func (m InspectModel) View() string {

	var bottomSection string

	if m.InspectingDownload != nil {
		buttonRender := ""
		if m.InspectingDownload.Name != "example" {
			buttonRender = button1Style.Render(util.MarginHor(" Open Guide (ctrl+g) ", 3))
		} else {
			buttonRender = styles.TertiaryInvertedStyle.Render(util.MarginHor(" Open Guide (ctrl+g) ", 3))
		}
		guideButton := "\n" + lipgloss.Place(m.Width, 3, lipgloss.Center, lipgloss.Center, buttonRender)
		horizontalLine := styles.PrimaryStyle.Render(util.DrawLine(m.Width))
		bottomSection = horizontalLine + guideButton
	}

	return m.viewport.View() + bottomSection
}

func (m InspectModel) UpdateContent() string {
	var s string
	if m.InspectingDownload != nil {

		if m.InspectingDownload.Name == "" || m.InspectingDownload.Name == "example" {
			s = "No download selected on the table."
			return s
		}

		horizontalLine := styles.DebugStyle.Render(util.DrawLine(m.Width))

		var header string
		header += headStyle.Render(util.MarginHor(m.InspectingDownload.Name, 1))
		header += "\nTYPE: " + util.IfElse(m.InspectingDownload.CustomResource, "Custom", "Default")
		header += util.IfElse(!m.InspectingDownload.CustomResource, "\nTIER: "+strconv.Itoa(m.InspectingDownload.Tier), "") + "\n"
		header += "\nSIZE: " + util.FormatSize(int(m.InspectingDownload.Info.Size))
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
		header += "\nSTATUS: " + statusStyle.Render(util.MarginHor(string(m.InspectingDownload.Status), 1)) + "\n\n"

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
	}

	return s
}

type ResizeMsgL2 struct {
	Width  int
	Height int
}

type ResizeMsgL3 struct {
	Width  int
	Height int
}
