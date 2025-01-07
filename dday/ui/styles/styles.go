package styles

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
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

	StatusSuccessStyle     = lipgloss.NewStyle().Foreground(Black).Background(Green)
	StatusDownloadingStyle = lipgloss.NewStyle().Foreground(Black).Background(BrightYellow)
	StatusFailStyle        = lipgloss.NewStyle().Foreground(Black).Background(Red)
	StatusWaitStyle        = lipgloss.NewStyle().Foreground(Black).Background(Blue)
)

var ScrollbarStyle = lipgloss.NewStyle().
	Foreground(SecondaryColor)

var StatusBackgroundStyle = lipgloss.NewStyle().
	Foreground(Accent2Color).
	Background(White)

var StatusFillStyle = lipgloss.NewStyle().
	Foreground(Accent1Color).
	Background(BrightBlack)
