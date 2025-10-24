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

type inputState int
const (
	inputDisabled inputState = iota
	inputIdState
	inputEmployeeState
)

type Model struct {
	state    optState
	message  string
	quitting bool
	service  *employee.Service

	// handle list state
	table table.Model

	// handle id input (view, edit, delete state)
	inputState inputState
	idInput    textinput.Model

	// handle employee input (name, email and so on) (add and edit state)
	employeeInputs []textinput.Model
	focusedInfo    int

	selectedEmployee *employee.Employee
}

func InitialModel(s *employee.Service) Model {
	// return Model{service: s, state: optMenu, employeeInfoInput: make([]textinput.Model, 4)}
	return Model{service: s, state: optMenu}
}

func (m Model) Init() tea.Cmd {
	// m.table = m.initTable("id")
	// m.idInput = m.initInputId()
	// m.employeeInputs = m.initEmployeeInputs()
	return nil
}
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// for checking what kinda of input
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// check what state we are in
		switch m.state {
		case optMenu:
			return m.updateStateMenu(msg)
		case optAdd:
			return m.updateStateAdd(msg)
		case optList:
			return m.updateStateList(msg)
		case optView:
			return m.updateStateView(msg)
		case optEdit:
			return m.updateStateEdit(msg)
		case optDelete:
			return m.updateStateDelete(msg)
		}
	}

	return m, nil
}

func (m Model) View() string {
	switch m.state {
	case optMenu:
		return m.viewStateMenu()
	case optAdd:
		return m.viewStateAdd()
	case optList:
		return m.viewStateList()
	case optView:
		return m.viewStateView()
	case optEdit:
		return m.viewStateEdit()
	case optDelete:
		return m.viewStateDelete()
	}

	return ""
}
