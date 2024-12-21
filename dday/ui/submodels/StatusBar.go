package submodels

import (
	"unicode/utf8"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"
	"github.com/hemanth2004/doomsday-protocol/dday/util"
)

var (
	statusBarStyle = lipgloss.NewStyle().Foreground(styles.White)
)

type StatusbarModel struct {
	Width int
}

func (m StatusbarModel) Init() tea.Cmd {
	return nil
}

func (m StatusbarModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil

}
func (m StatusbarModel) View() string {
	navText := "BATTERY: [▓▓▓▓▓░░░░░] 50%  |  ↓ bps"
	navText += util.Repl(" ", m.Width-utf8.RuneCountInString(navText)-4)

	return statusBarStyle.Render(lipgloss.PlaceHorizontal(m.Width, lipgloss.Center, navText))
}
