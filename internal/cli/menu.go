package cli

import (
	"fmt"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) updateMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case strconv.Itoa(int(optAdd)):
		m.message = "You chose: Add"
	case strconv.Itoa(int(optList)):
		m.message = "You chose: List"
		m.state = optList
		// init table when user select: List
		m.table = m.createTable("id")
	case strconv.Itoa(int(optView)):
		m.message = "You chose: View"
	case strconv.Itoa(int(optEdit)):
		m.message = "You chose: Edit"
	case strconv.Itoa(int(optDelete)):
		m.message = "You chose: Delete"
	case "q", "Q", "ctrl+c":
		m.message = "Exiting program..."
		m.quitting = true
		return m, tea.Quit
	default:
		m.message = "Invalid option. Press 0-6."
	}
	return m, nil
}

func (m Model) viewMenu() string {
		if m.quitting {
			return m.message + "\n"
		}
		render := renderHeader("MAIN MENU")
		render += renderMainMenu() + "\n" + m.message
	render += renderFooter("[Q]uit")

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
