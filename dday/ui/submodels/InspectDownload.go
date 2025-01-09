package submodels

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"
	"github.com/hemanth2004/doomsday-protocol/dday/util"
)

// Component to handle detailed view of a download thats happening currently

var (
	headStyle = lipgloss.NewStyle().Foreground(styles.Black).Background(styles.BrightYellow)

	detailsStyle = lipgloss.NewStyle().Foreground(styles.Black).Background(styles.White)

	progressBackStyle = lipgloss.NewStyle().Foreground(styles.BrightBlack)
	progressFillStyle = lipgloss.NewStyle().Foreground(styles.BrightBlue)
)

type InspectModel struct {
	Width  int
	Height int

	Focused            bool
	ProgressBar        MultilineProgress
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
		m.viewport = viewport.New(m.Width, m.Height-3)
	case tea.KeyMsg:
		if msg.String() == "up" {
			m.viewport.HalfViewDown()
		} else if msg.String() == "down" {
			m.viewport.HalfViewUp()
		}

	}

	m.ProgressBar = NewMultilineProgress(m.Width, 1, progressBackStyle, progressFillStyle)
	if m.InspectingDownload != nil {
		m.ProgressBar.SetPercent(m.InspectingDownload.Info.ProgressPercent())
	} else {
		m.ProgressBar.SetPercent(0)
	}
	m.viewport.SetContent(m.UpdateContent())

	return m, cmd
}

func (m InspectModel) View() string {
	return m.viewport.View()
}

func (m InspectModel) UpdateContent() string {
	var s string
	if m.InspectingDownload != nil {
		if m.InspectingDownload.Name == "" || m.InspectingDownload.Name == "tableFiller" {
			s = "No download selected on the table.\nUse ↑/↓ to navigate table."
			return s
		}

		horizontalLine := styles.DebugStyle.Render(util.DrawLine(m.Width))

		//
		//
		//
		var header string
		header += headStyle.Render(util.MarginHor(m.InspectingDownload.Name, 1))

		header += "\n" + m.ProgressBar.View(fmt.Sprintf("%.2f%%", m.ProgressBar.GetPercent()*100))

		var rsrctype string
		if m.InspectingDownload.CustomResource {
			rsrctype = "Custom"
		} else {
			switch m.InspectingDownload.Tier {
			case 0:
				rsrctype = "Default-Core"
			case 1:
				rsrctype = "Default-Additional-Lvl1"
			case 2:
				rsrctype = "Default-Additional-Lvl2"
			}
		}
		header += "\nTYPE: " + rsrctype

		//header += util.IfElse(!m.InspectingDownload.CustomResource, "\nTIER: "+strconv.Itoa(m.InspectingDownload.Tier), "") + "\n"
		header += "\nSIZE: " + util.FormatSize(int(m.InspectingDownload.Info.Done)) + " / " + util.FormatSize(int(m.InspectingDownload.Info.Size))

		header += "\nSTATUS: " + m.PrintResourceStatus()

		var fileDetails string
		fileDetails += "\nSOURCE: " + m.InspectingDownload.UrlGetter.RecentURLUsed
		fileDetails += "\nLOCATION: " + m.InspectingDownload.Location + "\n"

		//
		//
		//
		var resourceDetails string
		resourceDetails += "\n" + detailsStyle.Render(" Description ") + "\n" +
			m.InspectingDownload.Description + "\n"

		//
		//
		//
		var actionSection string
		actionSection += "\n"
		if m.InspectingDownload.AssociatedGuidePath.CheckIfExists() {
			actionSection += styles.HelpStyle1.Render("ctrl+g") + styles.HelpStyle2.Render(" open associated guide") + "\n"
		} else {
			actionSection += styles.HelpStyle2.Render("ctrl+g") + styles.HelpStyle2.Render(" no associated guide") + "\n"
		}

		if m.InspectingDownload.Status == core.StatusDownloading {
			actionSection += styles.HelpStyle1.Render("space") + styles.HelpStyle2.Render(" pause this resource") + "\n"
		} else if m.InspectingDownload.Status == core.StatusPaused {
			actionSection += styles.HelpStyle1.Render("space") + styles.HelpStyle2.Render(" resume this resource") + "\n"
		} else if m.InspectingDownload.Status == core.StatusFailed {
			actionSection += styles.HelpStyle1.Render("enter") + styles.HelpStyle2.Render(" retry this resource") + "\n"
		}

		s += header + fileDetails + horizontalLine + resourceDetails + horizontalLine + actionSection
	}

	return s
}

func (m InspectModel) PrintResourceStatus() string {
	var statusStyle lipgloss.Style
	switch m.InspectingDownload.Status {
	case core.StatusCompleted:
		statusStyle = styles.StatusSuccessStyle
	case core.StatusDownloading:
		statusStyle = styles.StatusDownloadingStyle
	case core.StatusFailed:
		statusStyle = styles.StatusFailStyle
	default:
		statusStyle = styles.StatusWaitStyle
	}
	return statusStyle.Render(util.MarginHor(string(m.InspectingDownload.Status), 1))
}

type ResizeMsgL2 struct {
	Width  int
	Height int
}

type ResizeMsgL3 struct {
	Width  int
	Height int
}
