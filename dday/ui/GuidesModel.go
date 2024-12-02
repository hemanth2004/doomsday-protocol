package ui

import tea "github.com/charmbracelet/bubbletea"

type GuidesModel struct {
	AllottedHeight int
	AllottedWidth  int

	HelpSet HelpSet // Will be used by parent model to render help
}

func (m GuidesModel) Init() tea.Cmd {
	return nil
}

func (m GuidesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//var cmd tea.Cmd
	var cmds []tea.Cmd

	return m, tea.Batch(cmds...)
}

func (m GuidesModel) View() string {
	return ""
}
