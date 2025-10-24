package cli

import tea "github.com/charmbracelet/bubbletea"

func (m Model) updateStateEdit(msg tea.KeyMsg) (tea.Model, tea.Cmd) {

	switch m.inputState {
	case inputIdState:
		return m.updateIdInput(msg)
	case inputEmployeeState:
		return m.updateEmployeeInputs(msg)
	}

	if err := m.service.ModifyEmployee(
		m.selectedEmployee.Id,
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
	m.message = "Employee Information Updated!"

	return m, nil
}

func (m Model) viewStateEdit() string {
	switch m.inputState {
	case inputIdState:
		return m.viewIdInput()
	}
	render := renderHeader("Update Employee Information")
	// render += "ID for new employee is \n"
	render += m.viewEmployeeInputs()
	return render
}
