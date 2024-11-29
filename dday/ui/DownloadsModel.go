package ui

import (
	"strconv"

	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/debug"
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
	LogsContent  string
	ResourceList *core.ResourceList

	// UI
	// 0-Downloads(topRight)
	// 1-Resources(left)
	// 2-Console(bottomRight)
	CurrentWindow   *util.StateHandler[int]
	ResourceTree    tree.Model
	DownloadsTable  table.Model
	ConsoleViewport viewport.Model
}

func (m DownloadsModel) GetWindowDimensions() (int, int, int, int) {
	leftWidth := m.AllottedWidth / 4
	rightWidth := m.AllottedWidth - leftWidth
	rightHeightPrimary := int(0.7*float64(m.AllottedHeight)) - 2
	rightHeightSecondary := m.AllottedHeight - rightHeightPrimary - 2

	return leftWidth, rightWidth, rightHeightPrimary, rightHeightSecondary
}

type RowsMsg []table.Row

func InitRows(m DownloadsModel) []table.Row {
	var rows []table.Row
	for i, resource := range m.ResourceList.DefaultResources {

		width := 10
		if m.WindowDimensions[1][0] > 10 {
			width = util.CalculateFlexWidth(
				m.WindowDimensions[1][0],
				[]int{3},
				[]int{2, 1, 1, 1},
				2,
			)
		}

		bar := progress.New(
			progress.WithWidth(width-3),
			progress.WithDefaultGradient(),
			progress.WithSolidFill("#00FF00"),
		)
		bar.SetPercent(1)
		row := table.NewRow(table.RowData{
			columnKeyID:          strconv.Itoa(i + 1),
			columnKeyName:        resource.Name,
			columnKeyProgressBar: bar.View(),
			columnKeyStatus:      resource.Status,
			columnKeySpeed:       util.FormatSpeed(int(resource.Info.Bandwidth)),
			columnKeyETA:         util.FormatTime(resource.Info.ETA),
		})

		rows = append(rows, row)
	}

	// Return rows as a RowsMsg
	return rows
}

func (m DownloadsModel) Init() tea.Cmd {

	return func() tea.Msg {
		return InitRows(m)
	}
}

func (m DownloadsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd = nil
	var cmds []tea.Cmd

	m.DownloadsTable = m.DownloadsTable.WithRows(InitRows(m))

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		leftWidth, rightWidth, rightHeightPrimary, rightHeightSecondary := m.GetWindowDimensions()
		leftHeight := int(float64(m.AllottedHeight)/2) - 1
		m.ResourceTree.Resize(leftWidth, leftHeight)
		m.WindowDimensions[0] = [2]int{leftWidth, leftHeight}

		m.DownloadsTable = m.DownloadsTable.WithTargetWidth(rightWidth).WithMinimumHeight(rightHeightPrimary - 1)
		m.WindowDimensions[1] = [2]int{rightWidth, rightHeightPrimary - 1}
		// m.DownloadsTable.SetWidth(rightWidth)
		// m.DownloadsTable.SetHeight(rightHeightPrimary - 2)

		m.ConsoleViewport = viewport.New(rightWidth, rightHeightSecondary-2)
		m.ConsoleViewport.SetContent(m.LogsContent)

	case core.LoggedMsg:
		debug.Log(string(msg))
		m.ConsoleViewport.SetContent(string(msg))
		m.LogsContent = string(msg)

	case tea.KeyMsg:
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

	// Update Resource table
	m.DownloadsTable, cmd = m.DownloadsTable.Update(msg)
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
	bottomRightContent += m.ConsoleViewport.View()

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
