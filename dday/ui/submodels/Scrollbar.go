package submodels

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"
)

const (
	ScrollbarTrack = "░"
	ScrollbarThumb = "█"
)

type ScrollbarModel struct {
	Height        int  // Total height of the scrollbar
	ContentHeight int  // Total height of the content being scrolled
	ViewHeight    int  // Height of the visible viewport
	ScrollOffset  int  // Current scroll position
	ShowScrollbar bool // Whether to show the scrollbar
}

func NewScrollbar() ScrollbarModel {
	return ScrollbarModel{
		ShowScrollbar: true,
	}
}

func (m ScrollbarModel) Init() tea.Cmd {
	return nil
}

func (m ScrollbarModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m ScrollbarModel) View() string {
	if !m.ShowScrollbar || m.ContentHeight <= m.ViewHeight {
		return ""
	}

	// Calculate thumb size and position
	thumbSize := max(1, (m.ViewHeight*m.Height)/m.ContentHeight)
	thumbPos := (m.ScrollOffset * (m.Height - thumbSize)) / (m.ContentHeight - m.ViewHeight)

	// Ensure thumb position stays within bounds
	thumbPos = max(0, min(m.Height-thumbSize, thumbPos))

	bar := make([]string, m.Height)

	for i := 0; i < m.Height; i++ {
		bar[i] = ScrollbarTrack
	}
	for i := thumbPos; i < thumbPos+thumbSize; i++ {
		bar[i] = ScrollbarThumb
	}

	return styles.ScrollbarStyle.Render(strings.Join(bar, "\n"))
}
