package ui

import (
	"math"
	"strconv"
	"unicode/utf8"

	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"
	"github.com/hemanth2004/doomsday-protocol/dday/util"
	"github.com/hemanth2004/doomsday-protocol/dday/util/tree"

	"fmt"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

type DownloadsModel struct {
	AllottedHeight   int
	AllottedWidth    int
	WindowDimensions [3]BoxResolution

	// Core
	LogFunction  *func(string)
	LogsContent  string
	ResourceList *core.ResourceList

	// UI
	// 0-Downloads(topRight)
	// 1-Resources(left)
	// 2-Console(bottomRight)
	CurrentWindow   *util.StateHandler[int]
	ConsoleOpened   bool
	ResourceTree    tree.Model
	DownloadsTable  table.Model
	ConsoleViewport viewport.Model
}

func (m DownloadsModel) GetWindowDimensions() (int, int, int, int) {
	leftWidth := m.AllottedWidth / 4
	rightWidth := m.AllottedWidth - leftWidth
	rightHeightPrimary := int(0.7*float64(m.AllottedHeight)) - 2
	rightHeightSecondary := m.AllottedHeight - rightHeightPrimary - 2

	if !m.ConsoleOpened {
		rightHeightPrimary = m.AllottedHeight - 2 - 3
		rightHeightSecondary = 3
	}

	return leftWidth, rightWidth, rightHeightPrimary, rightHeightSecondary
}

type RowsMsg []table.Row

func InitRows(m DownloadsModel, tableWidth int) ([]table.Row, [][]string) {
	var rows []table.Row
	var rowsString [][]string
	for i, resource := range m.ResourceList.DefaultResources {

		width := calculateColumnWidth(downloadTableColumns, tableWidth, getColumnFromKey(downloadTableColumns, progressBarPair.Key))

		bar := progress.New(
			progress.WithWidth(width-2),
			progress.WithDefaultGradient(),
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
		row := table.NewRow(table.RowData{
			idPair.Key:          rowString[0],
			namePair.Key:        rowString[1],
			progressBarPair.Key: rowString[2],
			statusPair.Key:      rowString[3],
			sizePair.Key:        rowString[4],
			speedPair.Key:       rowString[5],
			etaPair.Key:         rowString[6],
		})

		rows = append(rows, row)
		rowsString = append(rowsString, rowString)
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
	rows, rowsString := InitRows(m, rightWidth)

	m.DownloadsTable, cmd = m.DownloadsTable.Update(msg)
	cmds = append(cmds, cmd)

	tableHeight := rightHeightPrimary - 9
	m.DownloadsTable = m.DownloadsTable.
		WithRows(rows).
		WithTargetWidth(rightWidth).
		Focused(m.CurrentWindow.Index() == 2)

	updateTableHeightAndFooter(&m, rowsString, rightWidth, tableHeight)

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:

		leftHeight := int(float64(m.AllottedHeight)/2) - 1
		m.ResourceTree.Resize(leftWidth, leftHeight)
		m.WindowDimensions[0] = [2]int{leftWidth, leftHeight}

		m.WindowDimensions[1] = [2]int{rightWidth, tableHeight}

		m.ConsoleViewport = viewport.New(rightWidth, rightHeightSecondary-2)
		m.ConsoleViewport.SetContent(m.LogsContent)
		m.ConsoleViewport.GotoBottom()
		m.WindowDimensions[2] = [2]int{rightWidth, rightHeightSecondary - 2}

	case core.LoggedMsg:
		m.ConsoleViewport.SetContent(string(msg))
		m.ConsoleViewport.GotoBottom()
		m.LogsContent = string(msg)

	case tea.KeyMsg:
		if msg.String() == "enter" {
			if m.CurrentWindow.Index() == 0 {
				m.ConsoleOpened = !m.ConsoleOpened
			}
		}
		if msg.String() == "tab" {
			m.CurrentWindow.NextState()
		}
		if msg.String() == "shift+tab" {
			m.CurrentWindow.PrevState()
		}
	}

	// Update Resource Tree
	if m.CurrentWindow.Index() == 1 {
		m.ResourceTree, cmd = m.ResourceTree.Update(msg)
	}
	m.ResourceTree.SetNodes(m.GenerateResourceTree())
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m DownloadsModel) GenerateResourceTree() []tree.Node {
	// Initialize the tree structure
	a := []tree.Node{
		{
			Value: "Default Resources",
			Desc:  "Resources that the protocol recommends.",
			Children: []tree.Node{
				{
					Value:    "Tier 0",
					Desc:     "Absolutely necessary.",
					Children: []tree.Node{},
				},
				{
					Value:    "Tier 1",
					Desc:     "Slightly less absolutely necessary.",
					Children: []tree.Node{},
				},
			},
		},
		{
			Value:    "Custom Resources",
			Desc:     "Resources that the user has added.",
			Children: []tree.Node{},
		},
	}

	// Populate default resources with their respective tiers based on their Tier field
	for _, resource := range m.ResourceList.DefaultResources {
		newNode := tree.Node{
			Value:    resource.Name,
			Desc:     resource.Description,
			Children: []tree.Node{},
		}

		// Check the Tier and add to the corresponding tier node
		if resource.Tier == 0 {
			a[0].Children[0].Children = append(a[0].Children[0].Children, newNode)
		} else if resource.Tier == 1 {
			a[0].Children[1].Children = append(a[0].Children[1].Children, newNode)
		}
	}

	// If no resources are added under a specific tier, add an empty node
	if len(a[0].Children[0].Children) == 0 {
		a[0].Children[0].Children = append(a[0].Children[0].Children, tree.Node{
			Value: "-",
			Desc:  "No resources under here.",
		})
	}
	if len(a[0].Children[1].Children) == 0 {
		a[0].Children[1].Children = append(a[0].Children[1].Children, tree.Node{
			Value: "-",
			Desc:  "No resources under here.",
		})
	}

	// Populate custom resources if there are any, otherwise add an empty node
	if len(m.ResourceList.CustomResources) > 0 {
		for _, resource := range m.ResourceList.CustomResources {
			newNode := tree.Node{
				Value:    resource.Name,
				Children: []tree.Node{},
			}
			a[1].Children = append(a[1].Children, newNode)
		}
	} else {
		a[1].Children = append(a[1].Children, tree.Node{
			Value: "-",
			Desc:  "No custom resources.",
		})
	}

	return a
}

func (m DownloadsModel) View() string {
	var s string

	// Specifying dimensions of windows
	leftWidth, rightWidth, rightHeightPrimary, rightHeightSecondary := m.GetWindowDimensions()

	//----------------
	// Rendering content in each window

	// Console Viewport content
	bottomRightContent := fmt.Sprintf("CONSOLE\n%s\n", styles.DebugStyle.Render(util.DrawLine(rightWidth)))
	if m.ConsoleOpened {
		bottomRightContent += m.ConsoleViewport.View()
	}

	// Downloads Table content
	topRightContent := fmt.Sprintf("DOWNLOADS\n%s", "")
	topRightContent += m.DownloadsTable.View()

	// Resource Tree content
	leftContent := fmt.Sprintf("RESOURCES\n%s\n%s", styles.DebugStyle.Render(util.DrawLine(leftWidth)), m.ResourceTree.View())
	leftContent += "\n" + styles.DebugStyle.Render(util.DrawLine(leftWidth))
	curNode := m.ResourceTree.NodeAtCursor()
	if curNode != nil {
		leftContent += "\n" + curNode.Desc
	}

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

	if m.CurrentWindow.CurrentState() == 2 {
		// Highlight Console panel if active
		bottomRightWindow = styles.PanelHighlightStyle.Width(rightWidth).Height(rightHeightSecondary).Render(bottomRightContent)
	} else {
		bottomRightWindow = styles.PanelStyle.Width(rightWidth).Height(rightHeightSecondary).Render(bottomRightContent)
	}

	rightSection := lipgloss.JoinVertical(lipgloss.Top, topRightWindow, bottomRightWindow)
	s += lipgloss.JoinHorizontal(lipgloss.Top, leftWindow, rightSection)

	return s + "\n"
}

// table utils

func updateTableHeightAndFooter(m *DownloadsModel, rowsString [][]string, width, viewportHeight int) {
	linesUsedByEachTableRow := calculateExtraMultilineRows(downloadTableColumns, rowsString, width-2)
	pageSize := calculatePaginationSize(linesUsedByEachTableRow, len(rowsString), viewportHeight)

	startIndex, endIndex := m.DownloadsTable.VisibleIndices()
	visibleRowsHeight := 0
	for i := startIndex; i < endIndex; i++ {
		visibleRowsHeight += linesUsedByEachTableRow[i]
	}
	unusedHeight := viewportHeight - visibleRowsHeight
	customFooter := fmt.Sprintf("%d / %d ", m.DownloadsTable.CurrentPage(), m.DownloadsTable.MaxPages())
	for i := 0; i < unusedHeight-2; i++ {
		customFooter += "\n"
	}

	m.DownloadsTable = m.DownloadsTable.
		WithTargetWidth(width).
		WithPageSize(pageSize).
		WithStaticFooter(customFooter)
}

func getColumnFromKey(columns []table.Column, key string) table.Column {
	for _, col := range columns {
		if col.Key() == key {
			return col
		}
	}
	return table.NewColumn("nil,", "nil", 1)
}

func calculateColumnWidth(columns []table.Column, totalWidth int, targetColumn table.Column) int {

	fixedTotal := 0
	for _, pair := range columns {
		if !pair.IsFlex() {
			fixedTotal += pair.Width()
		}
	}

	remainingWidth := totalWidth - fixedTotal

	totalFlexFactors := 0
	for _, factor := range columns {
		if factor.IsFlex() {
			totalFlexFactors += factor.FlexFactor()
		}
	}

	totalFlexFactorsWithoutInputFactor := totalFlexFactors - targetColumn.FlexFactor()

	return int(float64(remainingWidth) - ((float64(totalFlexFactorsWithoutInputFactor) / float64(totalFlexFactors)) * float64(remainingWidth)))
	// the -1 is accounting for the table border
}

func calculateExtraMultilineRows(columns []table.Column, rows [][]string, totalWidth int) []int {
	var allowedWidths []int
	for _, col := range columns {
		if col.IsFlex() {
			allowedWidths = append(allowedWidths, calculateColumnWidth(columns, totalWidth, col))
		} else {
			allowedWidths = append(allowedWidths, col.Width())
		}
	}

	util.SetColumn(rows, 2, " ")
	extraInEachCell := make([][]int, len(rows))
	for i := range extraInEachCell {
		extraInEachCell[i] = make([]int, len(allowedWidths))
	}

	for i, row := range rows {
		for j, colWidth := range allowedWidths {
			cellContent := row[j]
			lines := int(math.Ceil(float64(utf8.RuneCountInString(cellContent)) / float64(colWidth)))
			extraInEachCell[i][j] = lines
		}
	}

	//debug.Log("Extra In Each Cell:\n" + printMatrix(extraInEachCell))

	maxInEachRow := make([]int, len(rows))
	for i := range rows {
		var maximumExtra int = 0
		for _, extra := range extraInEachCell[i] {
			if extra > maximumExtra {
				maximumExtra = extra
			}
		}
		maxInEachRow[i] = maximumExtra
	}

	return maxInEachRow
}

// Problem of optimization
// Starting info: given in arguements
// Conditions: maximising the page size
// Solution:
//  1. Iterate from pagination size 1 till number of rows
//  2. Example case
//     - total rows = 5 rows
//     - lines used by each row = {1, 2, 3, 2, 1}
//     - allowed viewport = 6 lines
//  Then the page size would be 3 because rows 1, 2 and 3 will be 6, and row 4 and 5 will be 3

func calculatePaginationSize(linesUsedByEachTableRow []int, totalRows, viewportHeight int) (size int) {
	for i := 1; i <= totalRows; i++ { // i is the potential page size
		currentHeight := 0
		valid := true

		// Check if the current page size (i) fits within the viewport height
		for j := 0; j < i; j++ {
			currentHeight += linesUsedByEachTableRow[j]
			if currentHeight > viewportHeight {
				valid = false
				break
			}
		}

		// If valid, update the size; otherwise, break the loop as no larger page size can work
		if valid {
			size = i
		} else {
			break
		}
	}

	//fmt.Print("Lines In Each Row:", linesUsedByEachTableRow, "\n\r", "Size: ", size, "\n\r")
	return
}

func printMatrix(matrix [][]int) (s string) {
	for _, row := range matrix { // Loop through each row
		for _, col := range row { // Loop through each column in the row
			s += strconv.Itoa(col) + "  " // Print each element with a space
		}
		s += "\n"
	}
	return
}

func sum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}
