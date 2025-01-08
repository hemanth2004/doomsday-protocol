package guides

import glamour "github.com/charmbracelet/glamour"

type SpecificFormatter struct {
	Extension string
	Formatter func(path, content string, width, height int) string
}

var FurtherFormattedFormats = []SpecificFormatter{
	{
		Extension: ".md",
		Formatter: func(_, content string, width, height int) string {
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
		},
	},
	{
		Extension: ".png",
		Formatter: RenderImage,
	},
	{
		Extension: ".jpg",
		Formatter: RenderImage,
	},
	{
		Extension: ".jpeg",
		Formatter: RenderImage,
	},
}
