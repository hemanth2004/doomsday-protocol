package guides

type SpecificFormatter struct {
	Extension string
	Formatter func(path, content string, width, height int) string
}

var FurtherFormattedFormats = []SpecificFormatter{
	{
		Extension: ".md",
		Formatter: MarkdownFormatting,
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
