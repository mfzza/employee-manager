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
	optInputId
)

type Model struct {
	state    optState
	message  string
	quitting bool
	service  *employee.Service
	table    table.Model
	textInput    textinput.Model
	selectedEmployee   employee.Employee
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
			return m.updateMenu(msg)
		case optList:
			return m.updateList(msg)
		case optView:
			return m.update_view(msg)
		case optInputId:
			return m.update_inputId(msg)
		}
	}

	return m, nil
}

func (m Model) View() string {
	switch m.state {
	case optMenu:
		return m.viewMenu()

	case optList:
		return m.viewList()

	case optView:
		return m.view_view()

	case optInputId:
		return m.view_inputId()
	}

	return ""
}
