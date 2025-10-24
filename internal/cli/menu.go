package cli

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) updateStateMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	// NOTE: remove unnecessary message
	case fmt.Sprint(optAdd):
		m.message = "You chose: Add"
		m.inputState = inputEmployeeState
		m.state = optAdd
		m.employeeInputs = m.initEmployeeInputs()

	case fmt.Sprint(optList):
		m.message = "You chose: List"
		m.state = optList
		// init table when user select: List
		m.table = m.initTable("id")

	case fmt.Sprint(optView):
		m.message = "You chose: View"
		m.inputState = inputIdState
		m.state = optView
		m.idInput = m.initIdInput()

	case fmt.Sprint(optEdit):
		m.message = "You chose: Edit"
		m.inputState = inputIdState
		m.state = optEdit
		m.idInput = m.initIdInput()

	case fmt.Sprint(optDelete):
		m.message = "You chose: Delete"
		m.inputState = inputIdState
		m.state = optDelete
		m.idInput = m.initIdInput()

	case "q", "Q", "ctrl+c":
		m.message = "Exiting program..."
		m.quitting = true
		return m, tea.Quit

	default:
		m.message = "Invalid option. Please select one of options above."
	}
	return m, nil
}

func (m Model) viewStateMenu() string {
	if m.quitting {
		return m.message + "\n"
	}
	render := renderHeader("MAIN MENU")
	render += renderMainMenu()
	render += renderFooter("[Q]uit")
	render += "\n\n" + m.message

	return render
}

