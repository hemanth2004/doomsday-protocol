package submodels

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type MultilineProgress struct {
	width     int
	height    int
	percent   float64 // 0-1
	style     lipgloss.Style
	fillStyle lipgloss.Style
}

func NewMultilineProgress(width, height int, style, fillStyle lipgloss.Style) MultilineProgress {
	return MultilineProgress{
		width:     width,
		height:    height,
		percent:   0,
		style:     style,
		fillStyle: fillStyle,
	}
}

func (m *MultilineProgress) SetStyles(background, fill lipgloss.Style) {
	m.style = background
	m.fillStyle = fill
}

func (m *MultilineProgress) GetPercent() float64 {
	return m.percent
}
func (m *MultilineProgress) SetPercent(percent float64) {
	m.percent = percent
}

func (m *MultilineProgress) View(overlay string) string {
	if m.width == 0 || m.height == 0 {
		return ""
	}

	fillWidth := int(float64(m.width) * m.percent)
	filled := strings.Repeat("█", fillWidth)
	empty := strings.Repeat("░", m.width-fillWidth)

	line := m.fillStyle.Render(filled) + m.style.Render(empty)

	if m.height == 1 {
		if overlay != "" {
			// Create inverted styles for overlay text
			invertedFillStyle := m.fillStyle.
				Background(m.fillStyle.GetForeground()).
				Foreground(m.fillStyle.GetBackground())
			invertedBackStyle := m.style.
				Background(m.style.GetForeground()).
				Foreground(m.style.GetBackground())

			// Center the overlay text
			padding := (m.width - lipgloss.Width(overlay)) / 2
			if padding > 0 {
				overlay = strings.Repeat(" ", padding) + overlay + strings.Repeat(" ", padding)
			}

			// Render the overlay blended with the progress bar
			var renderedOverlay string
			for j, ch := range overlay {
				if j >= m.width {
					break
				}
				if j < fillWidth {
					renderedOverlay += invertedFillStyle.Render(string(ch))
				} else {
					renderedOverlay += invertedBackStyle.Render(string(ch))
				}
			}

			// Merge overlay with progress bar line
			if lipgloss.Width(renderedOverlay) < m.width {
				renderedOverlay += m.style.Render(strings.Repeat(" ", m.width-lipgloss.Width(renderedOverlay)))
			}

			return renderedOverlay
		}
		// If no overlay, return the progress bar line
		return line
	}

	lines := make([]string, m.height)
	for i := 0; i < m.height; i++ {
		lines[i] = line
	}

	// Create inverted styles for overlay text
	invertedFillStyle := m.fillStyle.
		Background(m.fillStyle.GetForeground()).
		Foreground(m.fillStyle.GetBackground())
	invertedBackStyle := m.style.
		Background(m.style.GetForeground()).
		Foreground(m.style.GetBackground())

	// Split overlay into lines
	overlayLines := strings.Split(overlay, "\n")

	// Process each line of overlay text
	for i, line := range overlayLines {
		if i >= m.height {
			break
		}

		// Render each character with appropriate style based on position
		var renderedLine string
		for j, ch := range line {
			if j >= m.width {
				break
			}
			if j < fillWidth {
				renderedLine += invertedFillStyle.Render(string(ch))
			} else {
				renderedLine += invertedBackStyle.Render(string(ch))
			}
		}

		padding := (m.width - lipgloss.Width(line)) / 2
		if padding > 0 {
			renderedLine = strings.Repeat(" ", padding) + renderedLine
		}

		pos := (m.height-len(overlayLines))/2 + i
		if pos >= 0 && pos < m.height {
			lines[pos] = renderedLine
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}
