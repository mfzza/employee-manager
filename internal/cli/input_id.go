package cli

import (
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) update_inputId(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
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
		m.inputState = false

	case tea.KeyEsc.String():
		m.inputState = false
		m.state = optMenu
	case "ctrl+c":
		m.quitting = true
		return m, tea.Quit
	}

	return m, cmd
}

func (m Model) view_inputId() string {
	render := renderHeader("Input Employee ID")
	render += m.idInput.View()
	render += renderFooter("[Esc] Main menu. [ctrl+c] Quit")
	render += "\n\n"
	render += m.message
	return render
}

func (m Model) init_textInputId() textinput.Model {
	ti := textinput.New()
	ti.Prompt = "enter ID: "
	ti.Focus()
	ti.CharLimit = 3
	// NOTE: bubbletea textinput doesnt have ability to restrict digit-only
	// so it should be handle in other way

	return ti
}
