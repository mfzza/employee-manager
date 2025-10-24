package cli

import (
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) updateIdInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.idInput, cmd = m.idInput.Update(msg)

	switch msg.String() {
	case tea.KeyEnter.String():
		inputId, err := strconv.Atoi(m.idInput.Value())

		// if input is not a number, then keep user in input
		if err != nil {
			m.message = "Please input number only!"
			return m, cmd
		}
		m.selectedEmployee, err = m.service.GetEmployeeById(inputId)
		if err != nil {
			m.message = err.Error()
			return m, cmd
		}

		switch m.state {
		case optView, optDelete:
			m.inputState = inputDisabled
		case optEdit:
			m.inputState = inputEmployeeState
			m.employeeInputs = m.initEmployeeInputs()
			m.employeeInputs[0].SetValue(m.selectedEmployee.Name)
			m.employeeInputs[1].SetValue(m.selectedEmployee.Phone)
			m.employeeInputs[2].SetValue(m.selectedEmployee.Position)
			m.employeeInputs[3].SetValue(m.selectedEmployee.Email)

		}
	case tea.KeyEsc.String():
		m.inputState = inputDisabled
		m.state = optMenu
	case "ctrl+c":
		m.quitting = true
		return m, tea.Quit
	}

	return m, cmd
}

func (m Model) viewInputId() string {
	render := renderHeader("Input Employee ID")
	render += m.idInput.View()
	render += renderFooter("[Esc] Main menu. [ctrl+c] Quit")
	render += "\n\n"
	render += m.message
	return render
}

func (m Model) initInputId() textinput.Model {
	ti := textinput.New()
	ti.Prompt = "enter ID: "
	ti.Focus()
	ti.CharLimit = 3
	// NOTE: bubbletea textinput doesnt have ability to restrict digit-only
	// so it should be handle in other way

	return ti
}
