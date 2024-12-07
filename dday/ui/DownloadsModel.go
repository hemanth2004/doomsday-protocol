package ui

import (
	"strconv"

	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"
	"github.com/hemanth2004/doomsday-protocol/dday/util"
	"github.com/hemanth2004/doomsday-protocol/dday/util/tableutils"
	"github.com/hemanth2004/doomsday-protocol/dday/util/tree"

	"fmt"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

type DownloadsModel struct {
	AllottedHeight int
	AllottedWidth  int

	// Core
	LogFunction  *func(string)
	LogsContent  [][2]string
	ResourceList *core.ResourceList

	// UI
	// 0-Downloads(topRight)
	// 1-Resources(left)
	// 2-Console(bottomRight)
	CurrentWindow  *util.StateHandler[int]
	ResourceTree   tree.Model
	InspectModel   InspectModel
	DownloadsTable table.Model
	ConsoleModel   ConsoleModel

	HelpSet []HelpSet
}

func (m DownloadsModel) GetWindowDimensions() (int, int, int, int) {
	leftWidth := m.AllottedWidth / 4
	rightWidth := m.AllottedWidth - leftWidth
	rightHeightPrimary := int(0.75*float64(m.AllottedHeight)) - 2
	rightHeightSecondary := m.AllottedHeight - rightHeightPrimary - 2

	if !m.ConsoleModel.ConsoleOpened {
		rightHeightSecondary = 4
		rightHeightPrimary = m.AllottedHeight - 2 - rightHeightSecondary
	}

	return leftWidth, rightWidth, rightHeightPrimary, rightHeightSecondary
}

type RowsMsg []table.Row

func InitRows(m *DownloadsModel, tableWidth int) ([]table.Row, [][]string) {
	convertStringToRow := func(s []string, resource *core.Resource) table.Row {
		return table.NewRow(table.RowData{
			idPair.Key:          s[0],
			namePair.Key:        s[1],
			progressBarPair.Key: s[2],
			statusPair.Key:      s[3],
			sizePair.Key:        s[4],
			speedPair.Key:       s[5],
			etaPair.Key:         s[6],
			"resourceObject":    resource,
		})
	}
	var rows []table.Row
	var rowsString [][]string

	// Adding default resources rows
	for tier := 0; tier < 2; tier++ {

		defaultResourcesHeader := []string{
			" ",
			"Default Resources",
			"Tier " + strconv.Itoa(tier),
			" ",
			" ",
			" ",
			" ",
		}
		rows = append(rows, convertStringToRow(defaultResourcesHeader, &core.EmptyResource).WithStyle(styles.DefaultResourceHeaderRowStyle))
		rowsString = append(rowsString, defaultResourcesHeader)
		countInThisTier := 0
		for i, resource := range m.ResourceList.DefaultResources {

			if resource.Tier == tier {
				width := tableutils.CalculateColumnWidth(downloadTableColumns, tableWidth, tableutils.GetColumnFromKey(downloadTableColumns, progressBarPair.Key))

				bar := progress.New(
					progress.WithWidth(width-5),
					progress.WithSolidFill(string(styles.Accent2Color)),
				)
				rowString := []string{
					strconv.Itoa(i + 1),
					resource.Name,
					bar.ViewAs(float64(resource.Info.Done) / float64(resource.Info.Size)),
					string(resource.Status),
					util.FormatSize(int(resource.Info.Done)) + " / " + util.FormatSize(int(resource.Info.Size)),
					util.FormatSpeed(int(resource.Info.Bandwidth)),
					util.FormatTime(resource.Info.ETA),
				}
				row := convertStringToRow(rowString, &m.ResourceList.DefaultResources[i])

				rows = append(rows, row)
				rowsString = append(rowsString, rowString)
				countInThisTier++

			}
		}
		if countInThisTier == 0 {
			rows = util.DeleteElement[table.Row](rows, len(rows)-1)
			rowsString = util.DeleteElement[[]string](rowsString, len(rowsString)-1)
			break
		}
	}

	return rows, rowsString
}

func (m DownloadsModel) Init() tea.Cmd {
	return nil
}

func (m DownloadsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd = nil
	var cmds []tea.Cmd

	leftWidth, rightWidth, rightHeightPrimary, rightHeightSecondary := m.GetWindowDimensions()
	rows, rowsString := InitRows(&m, rightWidth)

	m.DownloadsTable, cmd = m.DownloadsTable.Update(msg)
	cmds = append(cmds, cmd)

	// Add the check here
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		if selectedResource := m.DownloadsTable.HighlightedRow().Data["resourceObject"].(*core.Resource); selectedResource.Name == "example" {
			if keyMsg.String() == "down" || keyMsg.String() == "up" {
				m.DownloadsTable, cmd = m.DownloadsTable.Update(msg)
				cmds = append(cmds, cmd)
			}
		}
	}

	tableHeight := rightHeightPrimary - 10
	m.DownloadsTable = m.DownloadsTable.
		WithRows(rows).
		WithTargetWidth(rightWidth).
		Focused(m.CurrentWindow.Index() == 2)

	m.DownloadsTable = tableutils.UpdateTableHeightAndFooter(m.DownloadsTable, rowsString, downloadTableColumns, rightWidth, tableHeight)

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:

		leftHeight := m.AllottedHeight - 3
		m.ResourceTree.Resize(leftWidth, leftHeight)

		rows[1] = rows[1].Selected(true)
		m.DownloadsTable = m.DownloadsTable.WithRows(rows)

		// Update Console Viewport
		if updatedConsole, _ := m.ConsoleModel.Update(tea.WindowSizeMsg{Width: rightWidth, Height: rightHeightSecondary}); updatedConsole != nil {
			m.ConsoleModel = updatedConsole.(ConsoleModel)
		}
		m.InspectModel.Width = leftWidth
		m.InspectModel.Height = leftHeight

	case core.LoggedMsg:
		if updatedConsole, _ := m.ConsoleModel.Update(msg); updatedConsole != nil {
			m.ConsoleModel = updatedConsole.(ConsoleModel)
		}

	case tea.KeyMsg:
		if msg.String() == "tab" {
			m.CurrentWindow.NextState()
		}
		if msg.String() == "shift+tab" {
			m.CurrentWindow.PrevState()
		}

		m.ConsoleModel.Focused = m.CurrentWindow.Index() == 0
		if updatedConsole, _ := m.ConsoleModel.Update(msg); updatedConsole != nil {
			m.ConsoleModel = updatedConsole.(ConsoleModel)
		}

	}

	// Update Inpsect Model
	selectedResource := m.DownloadsTable.HighlightedRow().Data["resourceObject"].(*core.Resource)

	//(*m.LogFunction)(selectedResource.Name)
	m.InspectModel.InspectingDownload = selectedResource
	updatedInspect, cmd := m.InspectModel.Update(msg)
	if updatedInspect, ok := updatedInspect.(InspectModel); ok {
		m.InspectModel = updatedInspect
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m DownloadsModel) View() string {
	var s string

	// Specifying dimensions of windows
	leftWidth, rightWidth, rightHeightPrimary, _ := m.GetWindowDimensions()

	//----------------
	// Rendering content in each window

	// Console Viewport content

	// Downloads Table content
	topRightContent := fmt.Sprintf("DOWNLOADS\n%s", "")
	topRightContent += m.DownloadsTable.View()

	// Resource Tree content
	leftContent := fmt.Sprintf("RESOURCES\n%s\n", styles.DebugStyle.Render(util.DrawLine(leftWidth)))
	leftContent += m.InspectModel.View()

	//-------------

	// Joining the windows
	var leftWindow, topRightWindow, bottomRightWindow string
	if m.CurrentWindow.CurrentState() == 0 {
		// Highlight Downloads panel if active
		topRightWindow = styles.PanelHighlightStyle.Width(rightWidth).Height(rightHeightPrimary).Render(topRightContent)
	} else {
		topRightWindow = styles.PanelStyle.Width(rightWidth).Height(rightHeightPrimary).Render(topRightContent)
	}

	if m.CurrentWindow.CurrentState() == 1 {
		// Highlight Resources panel if active
		leftWindow = styles.PanelHighlightStyle.Width(leftWidth).Height(m.AllottedHeight).Render(leftContent)
	} else {
		leftWindow = styles.PanelStyle.Width(leftWidth).Height(m.AllottedHeight).Render(leftContent)
	}

	bottomRightWindow = m.ConsoleModel.View()

	rightSection := lipgloss.JoinVertical(lipgloss.Top, topRightWindow, bottomRightWindow)
	s += lipgloss.JoinHorizontal(lipgloss.Top, leftWindow, rightSection)

	return s + "\n"
}
