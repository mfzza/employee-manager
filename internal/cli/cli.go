package cli

import (
	"employee-management/internal/employee"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
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
	state            optState
	inputState       bool
	message          string
	quitting         bool
	service          *employee.Service
	table            table.Model
	textInput        textinput.Model
	selectedEmployee employee.Employee
}

func InitialModel(s *employee.Service) Model {

	return Model{service: s, state: optMenu}
}

func (m Model) Init() tea.Cmd {
	return nil
}
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// for checking what kinda of input
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// check what state we are in
		switch m.state {
		case optMenu:
			return m.update_menu(msg)
		case optList:
			return m.update_list(msg)
		case optView:
			return m.update_view(msg)
		case optDelete:
			return m.update_delete(msg)
		}
	}

	return m, nil
}

func (m Model) View() string {
	switch m.state {
	case optMenu:
		return m.view_menu()

	case optList:
		return m.view_list()

	case optView:
		return m.view_view()

	case optDelete:
		return m.view_delete()
	}

	return ""
}
