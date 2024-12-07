package debug

import (
	"github.com/charmbracelet/lipgloss"
)

func SimpleSpread(logs [][2]string, showTime bool, timeStyle lipgloss.Style, messageStyle lipgloss.Style) string {
	var s string
	for _, log := range logs {
		if showTime {
			s += "> " + timeStyle.Render(log[1]) + ": " + messageStyle.Render(log[0]) + "\n"
		} else {
			s += "> " + messageStyle.Render(log[0]) + "\n"
		}
	}
	return s
}

func TruncateContent(content string, width int, maxLines int) string {
	if len(content) == 0 {
		return content
	}

	if len(content) <= width*maxLines {
		return content
	}

	truncated := content[:(width * maxLines)]
	return truncated + "..."
}
