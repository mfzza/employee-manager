package cli

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	message  string
	quitting bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "1":
			m.message = "You chose: Add"
		case "2":
			m.message = "You chose: List"
		case "3":
			m.message = "You chose: View"
		case "4":
			m.message = "You chose: Edit"
		case "5":
			m.message = "You chose: Delete"
		case "0", "q", "ctrl+c":
			m.message = "Exiting program..."
			m.quitting = true
			return m, tea.Quit
		default:
			m.message = "Invalid option. Press 0-6."
		}
	}
	return m, nil
}

func (m Model) View() string {
	if m.quitting {
		return m.message + "\n"
	}
	return printMainMenu() + "\n" + m.message
}

// helper
func printMainMenu() string {
	return `===== Employee Management Program =====

+-------------------------------------+
|------------- MAIN MENU -------------|
+-------------------------------------+

1. Add
2. List
3. View
4. Edit
5. Delete
0. Quit
`
}
