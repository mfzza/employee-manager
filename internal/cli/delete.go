package cli

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) update_delete(msg tea.KeyMsg) (tea.Model, tea.Cmd) {

	if m.inputState {
		return m.update_inputId(msg)
	}

	switch msg.String() {
	case "y", "Y":
		m.service.Del(m.selectedEmployee.Id)
		m.state = optMenu
	case "n", "N":
		m.state = optMenu

	case "ctrl+c", "q", "Q":
		m.quitting = true
		return m, tea.Quit
	}

	return m, nil
}

func (m Model) view_delete() string {
	if m.inputState {
		return m.view_inputId()
	}
	render := renderHeader("Delete Employee")
	render += m.selectedEmployee.DetailString()
	render += "\n\nAre you sure you want delete this employee?\n"
	render += "[Y] to delete,\n[N] to back into main menu"
	render += renderFooter("[Q]uit")
	return render
}
