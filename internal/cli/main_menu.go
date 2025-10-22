package cli

import (
	"employee-management/internal/employee"
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type optState int

const (
	optMenu optState = iota
	optAdd
	optList
	optView
	optEdit
	optDelete
)

type Model struct {
	state    optState
	message  string
	quitting bool
	service  *employee.Service
	table    table.Model
}

func InitialModel(s *employee.Service) Model {

	return Model{service: s, state: optMenu}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case strconv.Itoa(int(optAdd)):
			m.message = "You chose: Add"
		case strconv.Itoa(int(optList)):
			m.message = "You chose: List"
			m.state = optList
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

		case "m", "M":
			m.state = optMenu

		default:
			m.message = "Invalid option. Press 0-6."
		}
	}
	return m, nil
}

func (m Model) View() string {
	render := ""
	switch m.state {
	case optMenu:
		render += renderHeader("MAIN MENU")
		if m.quitting {
			return m.message + "\n"
		}
		render += renderMainMenu() + "\n" + m.message

	case optList:
		render += renderHeader("LIST")
		render += m.createTable().View() + "\n"
	}
	return render
}

// helper

func renderHeader(title string) string {
	boxWidth := 40
	title = " " + title + " "

	// NOTE: should be 2?
	innerWidth := boxWidth - 3

	// Calculate padding for centering
	paddingLeft := (innerWidth - len(title)) / 2
	paddingRight := innerWidth - len(title) - paddingLeft

	line := "|" + strings.Repeat("-", paddingLeft) + title + strings.Repeat("-", paddingRight) + "|"

	return fmt.Sprintf(`===== Employee Management Program =====

+-------------------------------------+
%s
+-------------------------------------+

`, line)
}

func renderMainMenu() string {
	return fmt.Sprintf(`%d. Add
%d. List
%d. View
%d. Edit
%d. Delete
Q. Quit
`, optAdd, optList, optView, optEdit, optDelete)
}

func (m Model) createTable() table.Model {
	columns := []table.Column{
		{Title: "ID", Width: 3},
		{Title: "Name", Width: 20},
		{Title: "Phone", Width: 20},
	}

	employees := m.service.List("id")
	// Convert to []table.Row
	var rows []table.Row
	for _, emp := range employees {
		rows = append(rows, table.Row{
			strconv.Itoa(emp.Id),
			emp.Name,
			emp.Phone,
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(11),
	)
	return t
}
