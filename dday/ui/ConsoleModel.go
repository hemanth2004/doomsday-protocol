package ui

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/debug"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"
	"github.com/hemanth2004/doomsday-protocol/dday/util"
)

type ConsoleModel struct {
	Width  int
	Height int

	ConsoleOpened bool
	Focused       bool
	LogsContent   [][2]string
	Viewport      viewport.Model
}

func (m ConsoleModel) Init() tea.Cmd {
	return nil
}

func (m ConsoleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

		m.Viewport = viewport.New(m.Width, m.Height-3)
		m.Viewport.SetContent(debug.SimpleSpread(m.LogsContent, true, styles.TertiaryInvertedStyle, styles.SecondaryStyle))
		if len(m.LogsContent) > 0 {
			m.Viewport.GotoBottom()
		}

	case tea.KeyMsg:
		if m.Focused {
			if msg.String() == "enter" {
				m.ConsoleOpened = !m.ConsoleOpened
				debug.Log("Console opened: " + strconv.FormatBool(m.ConsoleOpened))
				if m.ConsoleOpened {
					m.Viewport = viewport.New(m.Width, m.Height-3)
					m.Viewport.SetContent(debug.SimpleSpread(m.LogsContent, true, styles.TertiaryInvertedStyle, styles.SecondaryStyle))
					if len(m.LogsContent) > 0 {
						m.Viewport.GotoBottom()
					}
				}
			} else if msg.String() == "up" || msg.String() == "down" {
				m.Viewport, cmd = m.Viewport.Update(msg)
				cmds = append(cmds, cmd)
			}
		}

	case core.LoggedMsg:
		m.Viewport.SetContent(debug.SimpleSpread(msg, true, styles.TertiaryInvertedStyle, styles.SecondaryStyle))
		m.Viewport.GotoBottom()
		m.LogsContent = [][2]string(msg)
		if m.ConsoleOpened {
			if len(m.LogsContent) > 0 {
				m.Viewport.GotoBottom()
			}
		}
	}

	return m, tea.Batch(cmds...)
}

func (m ConsoleModel) View() string {
	var s string

	consoleState := ""
	if m.ConsoleOpened {
		consoleState = styles.DebugStyle.Render("[Opened]")
	} else {
		consoleState = styles.DebugStyle.Render("[Closed]")
	}
	latestUpdated := ""
	if len(m.LogsContent) > 0 && !m.ConsoleOpened {
		latestUpdated = styles.DebugStyle.Render("[last=" + m.LogsContent[len(m.LogsContent)-1][1] + "]")
	}

	content := fmt.Sprintf("CONSOLE "+consoleState+" "+latestUpdated+"\n%s", styles.DebugStyle.Render(util.DrawLine(m.Width))+"\n")

	if m.ConsoleOpened {
		content += m.Viewport.View()
	} else {
		logs := m.LogsContent
		if len(logs) > 0 {
			shortLogs := debug.TruncateContent("> "+logs[len(logs)-1][0], m.Width-5, 2)
			content += styles.DebugStyle.Render(shortLogs)
		} else {
			content += styles.DebugStyle.Render("> No logs yet.")
		}
	}

	selectedHeight := m.Height
	if !m.ConsoleOpened {
		selectedHeight = 4
	}

	if m.Focused {
		// Highlight Window if active
		s += styles.PanelHighlightStyle.Width(m.Width).Height(selectedHeight).Render(content)
	} else {
		s += styles.PanelStyle.Width(m.Width).Height(selectedHeight).Render(content)
	}

	return s
}
