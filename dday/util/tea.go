package util

import tea "github.com/charmbracelet/bubbletea"

// UpdateTeaModel safely updates a Bubble Tea model and returns the concrete type
func UpdateTeaModel[T interface{ tea.Model }](model T, msg tea.Msg) (T, tea.Cmd) {
	var m tea.Model = model
	updatedModel, cmd := m.Update(msg)
	if updatedModel != nil {
		return updatedModel.(T), cmd
	}
	return model, cmd
}
