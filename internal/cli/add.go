package cli

import tea "github.com/charmbracelet/bubbletea"

func (m Model) updateStateAdd(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.inputState {
		return m.updateEmployeeInputs(msg)
	}

	m.service.AddEmployee(
		m.employeeInputs[0].Value(),
		m.employeeInputs[1].Value(),
		m.employeeInputs[2].Value(),
		m.employeeInputs[3].Value(),
		)

	return m, nil
}

func (m Model) viewStateAdd() string {
	render := renderHeader("Add Employee")
	render += m.viewEmployeeInputs()
	// render += "\n\nAre you sure you want add this employee?\n"
	// render += "[Y] to add,\n[N] to back into main menu"
	// render += renderFooter("[Q]uit")
	return render
}
