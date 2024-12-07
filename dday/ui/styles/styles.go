package styles

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

// Colors
const (
	Black         = lipgloss.Color("0")
	Red           = lipgloss.Color("1")
	Green         = lipgloss.Color("2")
	Yellow        = lipgloss.Color("3")
	Blue          = lipgloss.Color("4")
	Magenta       = lipgloss.Color("5")
	Cyan          = lipgloss.Color("6")
	White         = lipgloss.Color("7")
	BrightBlack   = lipgloss.Color("8")
	BrightRed     = lipgloss.Color("9")
	BrightGreen   = lipgloss.Color("10")
	BrightYellow  = lipgloss.Color("11")
	BrightBlue    = lipgloss.Color("12")
	BrightMagenta = lipgloss.Color("13")
	BrightCyan    = lipgloss.Color("14")
	BrightWhite   = lipgloss.Color("15")
)

const (
	ColorModeMono = iota
	ColorModeANSI16
)

// TODO:
// 1.Use only ANSI-16 since linux virtual consoles only support 16 colors
// 2.Provide option to revert to Mono-color
// lipgloss Styles
var (
	Border     = lipgloss.NormalBorder()
	DebugStyle = lipgloss.NewStyle().Foreground(BrightBlack)
	MainStyle  = lipgloss.NewStyle().Foreground(White)

	HelpStyle1 = lipgloss.NewStyle().Foreground(White)
	HelpStyle2 = lipgloss.NewStyle().Foreground(BrightBlack)

	PanelStyle          = lipgloss.NewStyle().Border(Border).BorderForeground(BrightBlack)
	PanelHighlightStyle = lipgloss.NewStyle().Border(Border).BorderForeground(White)

	TreeDescriptionTitle = lipgloss.NewStyle().Foreground(Black).Background(Yellow).Bold(true)

	TableStyle                    = lipgloss.NewStyle().Foreground(White).BorderForeground(White)
	DefaultResourceHeaderRowStyle = lipgloss.NewStyle().Foreground(Yellow).Bold(true).Background(Green)
	CustomResourceHeaderRowStyle  = lipgloss.NewStyle().Foreground(Yellow).Bold(true).Background(Blue)

	UnderlineStyle = lipgloss.NewStyle().Underline(true)

	CurrentColorMode = ColorModeANSI16

	// Colors
	PrimaryColor   = BrightWhite
	SecondaryColor = White
	TertiaryColor  = BrightBlack
	Accent1Color   = BrightYellow
	Accent2Color   = BrightMagenta

	// Styles
	PrimaryStyle   = lipgloss.NewStyle().Foreground(BrightWhite)
	SecondaryStyle = lipgloss.NewStyle().Foreground(White)
	TertiaryStyle  = lipgloss.NewStyle().Foreground(BrightBlack)
	Accent1Style   = lipgloss.NewStyle().Foreground(BrightYellow)
	Accent2Style   = lipgloss.NewStyle().Foreground(BrightMagenta)

	PrimaryInvertedStyle   = lipgloss.NewStyle().Foreground(Black).Background(BrightWhite)
	SecondaryInvertedStyle = lipgloss.NewStyle().Foreground(Black).Background(White)
	TertiaryInvertedStyle  = lipgloss.NewStyle().Foreground(Black).Background(BrightBlack)
	Accent1InvertedStyle   = lipgloss.NewStyle().Foreground(Black).Background(BrightYellow)
	Accent2InvertedStyle   = lipgloss.NewStyle().Foreground(Black).Background(BrightMagenta)
)

var (
	CustomBorder = table.Border{
		Top:    "─",
		Left:   "",
		Right:  " ",
		Bottom: "─",

		TopRight:    "─",
		TopLeft:     "─",
		BottomRight: "─",
		BottomLeft:  "─",

		TopJunction:    "┬",
		LeftJunction:   "─",
		RightJunction:  "─",
		BottomJunction: "┴",
		InnerJunction:  "┼",

		InnerDivider: "│",
	}
)
