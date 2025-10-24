package cli

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) updateEmployeeInputs(msg tea.KeyMsg) (tea.Model, tea.Cmd) {

	totalInput := len(m.employeeInputs)
	cmds := make([]tea.Cmd, totalInput)
	for i := range m.employeeInputs {
		m.employeeInputs[i], cmds[i] = m.employeeInputs[i].Update(msg)
	}

	switch msg.String() {
	case tea.KeyEsc.String():
		m.inputState = inputDisabled
		m.state = optMenu
	case "ctrl+c":
		m.quitting = true
		return m, tea.Quit
	case tea.KeyEnter.String():
		m.inputState = inputDisabled

		// NOTE: it working but look weird
		switch m.state {
		case optAdd:
			return m.updateStateAdd(msg)
		case optEdit:
			return m.updateStateEdit(msg)
		}

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
		for i := range m.employeeInputs {
			if i == m.focusedInfo {
				cmds[i] = m.employeeInputs[i].Focus()
			} else {
				m.employeeInputs[i].Blur()
			}
		}

		return m, tea.Batch(cmds...)
	}

	return m, tea.Batch(cmds...)

}

func (m Model) viewEmployeeInputs() string {
	// render := renderHeader("Input Employee Information")
	render := ""
	for i := range m.employeeInputs {
		render += m.employeeInputs[i].View() + "\n"
	}
	render += renderFooter("[Enter] Proceed. [Esc] Main menu. [ctrl+c] Quit")
	render += "\n\n"
	render += m.message
	return render
}

func (m Model) initEmployeeInputs() []textinput.Model {
	// Initialize inputs
	tis := make([]textinput.Model, 4)
	var ti textinput.Model

	for i := range tis {
		ti = textinput.New()
		ti.CharLimit = 32

		switch i {
		case 0:
			ti.Prompt = "Enter name: "
			ti.Focus() // focus on first input
		case 1:
			ti.Prompt = "Enter phone: "
		case 2:
			ti.Prompt = "Enter position: "
		case 3:
			ti.Prompt = "Enter email: "
		}
		tis[i] = ti
	}
	return tis
}
