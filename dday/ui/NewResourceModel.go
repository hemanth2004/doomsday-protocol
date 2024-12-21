package ui

import tea "github.com/charmbracelet/bubbletea"

type NewResourceModel struct {
	Height int
	Width  int

	HelpSet HelpSet // Will be used by parent model to render help
}

func (m NewResourceModel) Init() tea.Cmd {
	return nil
}

func (m NewResourceModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//var cmd tea.Cmd
	var cmds []tea.Cmd

	return m, tea.Batch(cmds...)
}

func (m NewResourceModel) View() string {
	return ""
}
