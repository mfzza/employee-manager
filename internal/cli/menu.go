package cli

import (
	"fmt"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) updateStateMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	// NOTE: remove unnecessary message
	case strconv.Itoa(int(optAdd)):
		m.message = "You chose: Add"
		m.state = optAdd
		m.inputState = true
		m.employeeInputs = m.initEmployeeInputs()

	case strconv.Itoa(int(optList)):
		m.message = "You chose: List"
		m.state = optList
		// init table when user select: List
		m.table = m.initTable("id")

	case strconv.Itoa(int(optView)):
		m.message = "You chose: View"
		m.inputState = true
		m.state = optView
		m.idInput = m.initInputId()

	case strconv.Itoa(int(optEdit)):
		m.message = "You chose: Edit"
		m.inputState = true
		m.state = optEdit
		m.idInput = m.initInputId()

	case strconv.Itoa(int(optDelete)):
		m.message = "You chose: Delete"
		m.inputState = true
		m.state = optDelete
		m.idInput = m.initInputId()

	case "q", "Q", "ctrl+c":
		m.message = "Exiting program..."
		m.quitting = true
		return m, tea.Quit
	default:
		m.message = "Invalid option. Press 0-6."
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

func renderMainMenu() string {
	return fmt.Sprintf(
		`[%d] Add
[%d] List
[%d] View
[%d] Edit
[%d] Delete
`, optAdd, optList, optView, optEdit, optDelete)
}
