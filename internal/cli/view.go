package cli

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) update_view(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "m", "M", tea.KeyEnter.String():
		m.state = optMenu
	case "v", "V", tea.KeyEnter.String():
		m.textInput.SetValue("")
		m.state = optInputId
		return m, nil
	case "ctrl+c", "q", "Q":
		m.quitting = true
		return m, tea.Quit
	}

	return m, nil
}

func (m Model) view_view() string {
	render := renderHeader("Employee Details")
	render += m.selectedEmployee.DetailString()
	render += renderFooter("[Q]uit. [M]ain menu. [V]iew employee details again")
	return render
}

