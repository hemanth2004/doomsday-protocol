package tree

import (
	"fmt"
	"strings"

	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	bottomLeft string = " └─"

	white  = lipgloss.Color("#ffffff")
	black  = lipgloss.Color("#000000")
	purple = lipgloss.Color("#bd93f9")
)

type Styles struct {
	Shapes     lipgloss.Style
	Selected   lipgloss.Style
	Unselected lipgloss.Style
	Help       lipgloss.Style
}

func defaultStyles() Styles {
	return Styles{
		Shapes:     lipgloss.NewStyle().Margin(0, 0, 0, 0).Foreground(lipgloss.Color("200")),
		Selected:   lipgloss.NewStyle().Margin(0, 0, 0, 0).Background(styles.Grey37),
		Unselected: lipgloss.NewStyle().Margin(0, 0, 0, 0).Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"}),
		Help:       lipgloss.NewStyle().Margin(0, 0, 0, 0).Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"}),
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

func New(nodes []Node) Model {
	// Initialize with default dimensions
	return Model{
		KeyMap: DefaultKeyMap(),
		Styles: defaultStyles(),

		width:  80, // default width
		height: 25, // default height
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
		switch {
		case key.Matches(msg, m.KeyMap.Up):
			m.NavUp()
		case key.Matches(msg, m.KeyMap.Down):
			m.NavDown()
		}
	}

	return m, nil
}

func (m Model) View() string {
	// Here, we allow dynamic updates of availableHeight based on the current height
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

	for _, node := range remainingNodes {

		var str string

		// If we aren't at the root, we add the arrow shape to the string
		if indent > 0 {
			shape := strings.Repeat(" ", (indent-1)*2) + m.Styles.Shapes.Render(bottomLeft) + " "
			str += shape
		}

		// Generate the correct index for the node
		idx := *count
		*count++

		// Format the string with fixed width for the value and description fields
		valueWidth := 10
		valueStr := fmt.Sprintf("%-*s", valueWidth, node.Value)

		// If we are at the cursor, we add the selected style to the string
		if m.cursor == idx {
			str += fmt.Sprintf("%s\n", m.Styles.Selected.Render(valueStr))
		} else {
			str += fmt.Sprintf("%s\n", m.Styles.Unselected.Render(valueStr))
		}

		b.WriteString(str)

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

	return kb
}

func (m Model) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			m.KeyMap.Up,
			m.KeyMap.Down,
		},
	}
}

func (m *Model) Resize(width, height int) {
	m.SetWidth(width)
	m.SetHeight(height)
}

// Returns the node at where the cursor is present
func (m Model) NodeAtCursor() *Node {
	// Initialize a counter for the node index
	index := 0

	// Recursive function to traverse the tree and find the node at the cursor
	var findNode func([]Node, int) *Node
	findNode = func(nodes []Node, cursor int) *Node {
		for i := 0; i < len(nodes); i++ {
			// Check if the current index matches the cursor
			if index == cursor {
				return &nodes[i]
			}

			index++ // Increment index for each node traversed

			// Recursively check the children of the current node
			if len(nodes[i].Children) > 0 {
				result := findNode(nodes[i].Children, cursor)
				if result != nil {
					return result
				}
			}
		}
		return nil // Node not found
	}

	// Start the traversal with the root nodes
	return findNode(m.nodes, m.cursor)
}
