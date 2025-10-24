package cli

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) updateStateView(msg tea.KeyMsg) (tea.Model, tea.Cmd) {

	if m.inputState == inputIdState {
		return m.updateIdInput(msg)
	}

	switch msg.String() {
	case "m", "M", tea.KeyEnter.String():
		m.state = optMenu
	case "b", "B":
		m.idInput.SetValue("")
		m.inputState = inputIdState
	case "ctrl+c", "q", "Q":
		m.quitting = true
		return m, tea.Quit
	}

	return m, nil
}

func (m Model) viewStateView() string {
	if m.inputState == inputIdState {
		return m.viewInputId()
	}
	render := renderHeader("Employee Details")
	render += m.selectedEmployee.DetailString()
	render += renderFooter("[Q]uit. [M]ain menu. [B]ack")
	return render
}
