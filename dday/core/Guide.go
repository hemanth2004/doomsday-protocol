package core

import (
	glamour "github.com/charmbracelet/glamour"
)

type Guide struct {
	Title   string // To display on the guide viewer
	Content string // To display on the guide viewer
	Format  string // ".md" for markdown, ".txt" for plain text etc.
}

type ChangeViewingGuideMsg string

type SpecificFormatter struct {
	Extension string
	Formatter func(content string, width int, height int) string
}

var FurtherFormattedFormats = []SpecificFormatter{
	{
		Extension: ".md",
		Formatter: func(content string, width int, height int) string {
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
}

var SupportedFormats = []string{
	// Plain Text and Documentation
	".txt",  // Plain text
	".md",   // Markdown
	".rst",  // reStructuredText
	".adoc", // AsciiDoc
	".tex",  // LaTeX documents

	// Markup and Hypertext
	".html",  // HTML
	".htm",   // Alternative for HTML
	".xhtml", // XHTML
	".xml",   // XML
	".svg",   // Scalable Vector Graphics
	".mml",   // MathML
	".kml",   // Keyhole Markup Language (GIS)

	// Stylesheets and Presentation
	".css",  // CSS stylesheets
	".scss", // SCSS stylesheets
	".sass", // SASS stylesheets

	// Configuration Files
	".ini",  // INI configuration files
	".conf", // General configuration files
	".cfg",  // Another general configuration format
	".json", // JSON
	".yaml", // YAML
	".yml",  // Alternative YAML extension
	".toml", // TOML configuration
	".env",  // Environment variable files

	// Data and Tabular Formats
	".csv",    // Comma-separated values
	".tsv",    // Tab-separated values
	".psv",    // Pipe-separated values
	".ndjson", // Newline-delimited JSON

	// Query and Scripts
	".sql",     // SQL scripts
	".sparql",  // SPARQL queries
	".graphql", // GraphQL queries

	// Scripting and Programming Files
	".py",    // Python scripts
	".js",    // JavaScript files
	".ts",    // TypeScript files
	".go",    // Go programming language
	".java",  // Java programming language
	".c",     // C programming language
	".cpp",   // C++ programming language
	".cs",    // C# programming language
	".rs",    // Rust programming language
	".php",   // PHP scripts
	".rb",    // Ruby scripts
	".pl",    // Perl scripts
	".lua",   // Lua scripts
	".kt",    // Kotlin programming language
	".sh",    // Shell scripts
	".bat",   // Batch scripts
	".zsh",   // Zsh shell scripts
	".fish",  // Fish shell scripts
	".swift", // Swift programming language
	".scala", // Scala programming language
	".erl",   // Erlang scripts
	".ex",    // Elixir scripts
	".exs",   // Alternative for Elixir scripts
	".clj",   // Clojure scripts
	".dart",  // Dart programming language
	".vb",    // Visual Basic
	".asm",   // Assembly language

	// Markup and Specialty Formats
	".bib", // BibTeX (bibliography)
	".rtf", // Rich Text Format
	".man", // Unix manual pages
	".me",  // Troff markup
	".ms",  // Another Troff format

	// Miscellaneous
	".r",     // R scripts
	".tex",   // TeX/LaTeX
	".dtd",   // Document Type Definitions
	".wxml",  // WeChat Mini Program Markup
	".plist", // macOS/iOS Property List XML
}
