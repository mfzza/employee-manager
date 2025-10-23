package cli

import tea "github.com/charmbracelet/bubbletea"

func (m Model) update_add(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.inputState {
		return m.update_inputEmployee(msg)
	}

	m.service.Add(
		m.employeeInfoInput[0].Value(),
		m.employeeInfoInput[1].Value(),
		m.employeeInfoInput[2].Value(),
		m.employeeInfoInput[3].Value(),
		)

	return m, nil
}

func (m Model) view_add() string {
	render := renderHeader("Add Employee")
	render += m.view_inputEmployee()
	// render += "\n\nAre you sure you want add this employee?\n"
	// render += "[Y] to add,\n[N] to back into main menu"
	// render += renderFooter("[Q]uit")
	return render
}
