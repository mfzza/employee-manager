package cli

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) update_inputEmployee(msg tea.KeyMsg) (tea.Model, tea.Cmd) {

	totalInput := len(m.employeeInfoInput)
	cmds := make([]tea.Cmd, totalInput)
	for i := range m.employeeInfoInput {
		m.employeeInfoInput[i], cmds[i] = m.employeeInfoInput[i].Update(msg)
	}

	switch msg.String() {
	case tea.KeyEsc.String():
		m.inputState = false
		m.state = optMenu
	case "ctrl+c":
		m.quitting = true
		return m, tea.Quit
	// TODO: handle adding employee here
	case tea.KeyEnter.String():
		m.inputState = false
		// return m, tea.Quit

	case "tab", "shift+tab", "down", "up":
		s := msg.String()

		// Move focus
		if s == "up" || s == "shift+tab" {
			m.focusedInfo--
		} else {
			m.focusedInfo++
		}

		if m.focusedInfo > totalInput-1 {
			m.focusedInfo = 0
		} else if m.focusedInfo < 0 {
			m.focusedInfo = totalInput - 1
		}

		cmds := make([]tea.Cmd, totalInput)
		for i := range m.employeeInfoInput {
			if i == m.focusedInfo {
				cmds[i] = m.employeeInfoInput[i].Focus()
			} else {
				m.employeeInfoInput[i].Blur()
			}
		}

		return m, tea.Batch(cmds...)
	}

	return m, tea.Batch(cmds...)

}

func (m Model) view_inputEmployee() string {
	// render := renderHeader("Input Employee Information")
	render := ""
	for i := range m.employeeInfoInput {
		render += m.employeeInfoInput[i].View() + "\n"
	}
	render += renderFooter("[Enter] Proceed. [Esc] Main menu. [ctrl+c] Quit")
	render += "\n\n"
	render += m.message
	return render
}

func (m Model) init_textInputEmployee() []textinput.Model {
	// Initialize inputs
	tis := make([]textinput.Model, 4)
	var t textinput.Model

	for i := range tis {
		t = textinput.New()
		t.CharLimit = 32

		switch i {
		case 0:
			t.Prompt = "Enter name: "
			t.Focus() // focus on first input
		case 1:
			t.Prompt = "Enter phone: "
		case 2:
			t.Prompt = "Enter position: "
		case 3:
			t.Prompt = "Enter email: "
		}
		tis[i] = t
	}
	return tis
}
