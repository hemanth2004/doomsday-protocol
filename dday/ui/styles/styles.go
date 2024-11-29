package styles

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

// Colors
const (
	Black          = lipgloss.Color("0")
	Maroon         = lipgloss.Color("1")
	Green          = lipgloss.Color("2")
	Olive          = lipgloss.Color("3")
	Navy           = lipgloss.Color("4")
	Purple         = lipgloss.Color("5")
	Teal           = lipgloss.Color("6")
	Silver         = lipgloss.Color("7")
	Grey37         = lipgloss.Color("59")
	PaleTurquoise4 = lipgloss.Color("66")
	DarkOrange     = lipgloss.Color("214")
	Cyan           = lipgloss.Color("14")
	Blue           = lipgloss.Color("12")
	Red            = lipgloss.Color("9")
	Pink           = lipgloss.Color("13")
	Yellow         = lipgloss.Color("11")
	Grey39         = lipgloss.Color("241")
	Grey54         = lipgloss.Color("245")
	Aquamarine1    = lipgloss.Color("122")
)

// lipgloss Styles
var (
	Border     = lipgloss.RoundedBorder()
	DebugStyle = lipgloss.NewStyle().Foreground(Grey39)
	MainStyle  = lipgloss.NewStyle().Foreground(Silver)

	PanelStyle          = lipgloss.NewStyle().Border(Border).BorderForeground(Grey37)
	PanelHighlightStyle = lipgloss.NewStyle().Border(Border).BorderForeground(Aquamarine1)

	TableStyle = lipgloss.NewStyle().Foreground(Silver).BorderForeground(Grey39)
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

		TopJunction:    "─",
		LeftJunction:   "─",
		RightJunction:  "─",
		BottomJunction: "+",
		InnerJunction:  "+",

		InnerDivider: "│",
	}
)
