package tree

import (
	"errors"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"
	"github.com/hemanth2004/doomsday-protocol/dday/util"
)

const (
	bottomLeft string = " └──"
	pipe       string = "│"
	tee        string = "├──"
	bottomTee  string = "└──"
)

type Styles struct {
	Shapes     lipgloss.Style
	Selected   lipgloss.Style
	Unselected lipgloss.Style
	Help       lipgloss.Style
}

func defaultStyles() Styles {
	return Styles{
		Shapes:     styles.Accent1InvertedStyle,
		Selected:   styles.Accent1InvertedStyle,
		Unselected: styles.PrimaryStyle,
		Help:       styles.DebugStyle,
	}
}

type Node struct {
	Value    string
	Desc     string
	Children []Node
}

type Model struct {
	KeyMap KeyMap
	Styles Styles

	width  int
	height int
	nodes  []Node
	cursor int

	Help     help.Model
	showHelp bool

	AdditionalShortHelpKeys func() []key.Binding
}

func New(nodes []Node, width int, height int) Model {
	return Model{
		KeyMap: DefaultKeyMap(),
		Styles: defaultStyles(),

		width:  width,
		height: height,
		nodes:  nodes,

		showHelp: true,
		Help:     help.New(),
	}
}

// KeyMap holds the key bindings for the table.
type KeyMap struct {
	Bottom      key.Binding
	Top         key.Binding
	SectionDown key.Binding
	SectionUp   key.Binding
	Down        key.Binding
	Up          key.Binding
	Quit        key.Binding

	ShowFullHelp  key.Binding
	CloseFullHelp key.Binding
}

// DefaultKeyMap is the default key bindings for the table.
func DefaultKeyMap() KeyMap {
	return KeyMap{
		Bottom: key.NewBinding(
			key.WithKeys("bottom"),
			key.WithHelp("end", "bottom"),
		),
		Top: key.NewBinding(
			key.WithKeys("top"),
			key.WithHelp("home", "top"),
		),
		SectionDown: key.NewBinding(
			key.WithKeys("secdown"),
			key.WithHelp("secdown", "section down"),
		),
		SectionUp: key.NewBinding(
			key.WithKeys("secup"),
			key.WithHelp("secup", "section up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("↓", "down"),
		),
		Up: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("↑", "up"),
		),

		ShowFullHelp: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "more"),
		),
		CloseFullHelp: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "close help"),
		),

		Quit: key.NewBinding(
			key.WithKeys("q", "esc"),
			key.WithHelp("q", "quit"),
		),
	}
}

func (m Model) Nodes() []Node {
	return m.nodes
}

func (m *Model) SetNodes(nodes []Node) {
	m.nodes = nodes
}

func (m *Model) NumberOfNodes() int {
	count := 0

	var countNodes func([]Node)
	countNodes = func(nodes []Node) {
		for _, node := range nodes {
			count++
			if node.Children != nil {
				countNodes(node.Children)
			}
		}
	}

	countNodes(m.nodes)

	return count

}

func (m Model) Width() int {
	return m.width
}

func (m Model) Height() int {
	return m.height
}

func (m *Model) SetSize(width, height int) {
	m.width = width
	m.height = height
}

func (m *Model) SetWidth(newWidth int) {
	m.SetSize(newWidth, m.height)
}

func (m *Model) SetHeight(newHeight int) {
	m.SetSize(m.width, newHeight)
}

func (m Model) Cursor() int {
	return m.cursor
}

func (m *Model) SetCursor(cursor int) {
	m.cursor = cursor
}

func (m *Model) SetShowHelp() bool {
	return m.showHelp
}

func (m *Model) NavUp() {
	m.cursor--

	if m.cursor < 0 {
		m.cursor = 0
		return
	}

}

func (m *Model) NavDown() {
	m.cursor++

	if m.cursor >= m.NumberOfNodes() {
		m.cursor = m.NumberOfNodes() - 1
		return
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			m.NavUp()
		case "down":
			m.NavDown()
		}
	}

	return m, nil
}

func (m Model) View() string {
	availableHeight := m.height
	var sections []string

	nodes := m.Nodes()

	var help string
	if m.showHelp {
		help = m.helpView()
		availableHeight -= lipgloss.Height(help)
	}

	count := 0 // This is used to keep track of the index of the node we are on (important because we are using a recursive function)
	sections = append(sections, lipgloss.NewStyle().Height(availableHeight).Render(m.renderTree(m.nodes, 0, &count)), help)

	if len(nodes) == 0 {
		return "No data"
	}
	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

func (m *Model) renderTree(remainingNodes []Node, indent int, count *int) string {
	var b strings.Builder

	for i, node := range remainingNodes {
		var str string
		isLast := i == len(remainingNodes)-1

		// Build the indent prefix
		if indent > 0 {
			// Just add spaces for indentation instead of pipes
			str += strings.Repeat("    ", indent-1)

			// Add the appropriate connector (tee or bottomTee)
			if isLast {
				str += styles.TertiaryStyle.Render(bottomTee)
			} else {
				str += styles.TertiaryStyle.Render(tee)
			}
		}

		// Generate the correct index for the node
		idx := *count
		*count++

		// Format the node name with cursor if selected
		nodeStr := fmt.Sprintf("%s%s",
			util.IfElse[string](m.cursor == idx, styles.Accent1InvertedStyle.Render(" →"), ""),
			util.IfElse[string](m.cursor == idx, styles.Accent1InvertedStyle.Render(node.Value), styles.PrimaryStyle.Render(node.Value)),
		)

		str += nodeStr

		b.WriteString(str + "\n")

		// If this node has children, render them with the appropriate pipe continuation
		if node.Children != nil {
			childStr := m.renderTree(node.Children, indent+1, count)
			b.WriteString(childStr)
		}
	}

	return b.String()
}

func (m Model) helpView() string {
	return m.Styles.Help.Render(m.Help.View(m))
}

func (m Model) ShortHelp() []key.Binding {
	kb := []key.Binding{
		m.KeyMap.Up,
		m.KeyMap.Down,
	}

	if m.AdditionalShortHelpKeys != nil {
		kb = append(kb, m.AdditionalShortHelpKeys()...)
	}

	return append(kb,
		m.KeyMap.Quit,
	)
}

func (m Model) FullHelp() [][]key.Binding {
	kb := [][]key.Binding{{
		m.KeyMap.Up,
		m.KeyMap.Down,
	}}

	return append(kb,
		[]key.Binding{
			m.KeyMap.Quit,
			m.KeyMap.CloseFullHelp,
		})
}

func (m *Model) GetSelectedNode() (Node, error) {

	count := 0
	return m.findCursorNode(m.nodes, 0, &count)
}

// GetSelectedNode returns the node at the current cursor position
func (m *Model) findCursorNode(remainingNodes []Node, indent int, count *int) (Node, error) {

	for _, node := range remainingNodes {

		// Generate the correct index for the node
		idx := *count
		*count++

		if idx == m.cursor {
			return node, nil
		}

		// If this node has children, render them with the appropriate pipe continuation
		if node.Children != nil {
			childNode, err := m.findCursorNode(node.Children, indent+1, count)
			if err == nil {
				return childNode, nil
			}
		}
	}

	return Node{}, errors.New("node not found")
}
