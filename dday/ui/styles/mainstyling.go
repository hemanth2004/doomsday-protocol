package styles

import "github.com/charmbracelet/lipgloss"

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

// TODO:
// 1.Use only ANSI-16 since linux virtual consoles only support 16 colors
// 2.Provide option to revert to Mono-color

var (
	Border        = lipgloss.NormalBorder()
	RoundedBorder = lipgloss.RoundedBorder()
	CtrlBorder    = lipgloss.InnerHalfBlockBorder()

	DebugStyle = lipgloss.NewStyle().Foreground(TertiaryColor)
	MainStyle  = lipgloss.NewStyle().Foreground(SecondaryColor)

	HelpStyle1 = lipgloss.NewStyle().Foreground(SecondaryColor)
	HelpStyle2 = lipgloss.NewStyle().Foreground(TertiaryColor)

	PanelStyle              = lipgloss.NewStyle().Border(RoundedBorder).BorderForeground(BrightBlack)
	PanelHighlightStyle     = lipgloss.NewStyle().Border(RoundedBorder).BorderForeground(White)
	CtrlPanelStyle          = lipgloss.NewStyle().Border(CtrlBorder).BorderForeground(BrightBlack)
	CtrlPanelHighlightStyle = lipgloss.NewStyle().Border(CtrlBorder).BorderForeground(BrightBlue)

	TreeDescriptionTitle = lipgloss.NewStyle().Foreground(Black).Background(Yellow).Bold(true)

	TableStyle               = lipgloss.NewStyle().Foreground(SecondaryColor).BorderForeground(SecondaryColor)
	TableRowSeperationStyle  = lipgloss.NewStyle().Foreground(Black).Background(Yellow).Align(lipgloss.Center)
	DefaultResourceHeadStyle = lipgloss.NewStyle().Foreground(Black).Background(White).Align(lipgloss.Center)
	CustomResourceHeadStyle  = lipgloss.NewStyle().Foreground(Yellow).Bold(true).Background(Blue)

	UnderlineStyle = lipgloss.NewStyle().Underline(true)

	// Colors
	PrimaryColor   = BrightWhite
	SecondaryColor = White
	TertiaryColor  = BrightBlack
	Accent1Color   = BrightYellow  // Bright Highlight
	Accent2Color   = BrightMagenta // Slightly less important highlight
	Accent3Color   = Yellow        // Least important highlight
	Accent4Color   = BrightBlue    // Least important highlight

	// Styles
	PrimaryStyle   = lipgloss.NewStyle().Foreground(PrimaryColor)
	SecondaryStyle = lipgloss.NewStyle().Foreground(SecondaryColor)
	TertiaryStyle  = lipgloss.NewStyle().Foreground(TertiaryColor)
	Accent1Style   = lipgloss.NewStyle().Foreground(Accent1Color)
	Accent2Style   = lipgloss.NewStyle().Foreground(Accent2Color)
	Accent3Style   = lipgloss.NewStyle().Foreground(Accent3Color)

	PrimaryInvertedStyle   = lipgloss.NewStyle().Foreground(Black).Background(PrimaryColor)
	SecondaryInvertedStyle = lipgloss.NewStyle().Foreground(Black).Background(SecondaryColor)
	TertiaryInvertedStyle  = lipgloss.NewStyle().Foreground(Black).Background(TertiaryColor)
	Accent1InvertedStyle   = lipgloss.NewStyle().Foreground(Black).Background(Accent1Color)
	Accent2InvertedStyle   = lipgloss.NewStyle().Foreground(Black).Background(Accent2Color)
	Accent3InvertedStyle   = lipgloss.NewStyle().Foreground(Black).Background(Accent3Color)
)
