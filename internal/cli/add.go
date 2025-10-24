package cli

import tea "github.com/charmbracelet/bubbletea"

func (m Model) updateStateAdd(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.inputState == inputEmployeeState {
		return m.updateEmployeeInputs(msg)
	}

	if err := m.service.AddEmployee(
		m.employeeInputs[0].Value(),
		m.employeeInputs[1].Value(),
		m.employeeInputs[2].Value(),
		m.employeeInputs[3].Value(),
		); err != nil {
		m.message = err.Error()
		m.inputState = inputEmployeeState
		return m, nil
	}

	m.state = optMenu
	m.message = "Employee successfully added!"

	return m, nil
}

func (m Model) viewStateAdd() string {
	render := renderHeader("Add Employee")
	// render += "ID for new employee is \n"
	render += m.viewEmployeeInputs()
	return render
}
