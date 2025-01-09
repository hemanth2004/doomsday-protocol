package guides

import glamour "github.com/charmbracelet/glamour"

func MarkdownFormatting(_, content string, width, _ int) string {
	// Format markdown content
	renderer, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width-2),
	)

	result, err := renderer.Render(content)
	if err != nil {
		return content
	}
	return result
}
