package cli

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) update_list(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// NOTE: this function is called only on keypress trigger,
	// which is correct for handle shortcut,
	// but not sure if update table should be done here
	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)

	// key shortcut
	switch msg.String() {
	case "m", "M":
		m.state = optMenu
		return m, nil
	case "i", "I":
		m.table = m.initTable("id")
		return m, nil
	case "n", "N":
		m.table = m.initTable("name")
		return m, nil
	case "p", "P":
		m.table = m.initTable("phone")
		return m, nil
	case "q", "Q", "ctrl+c":
		m.quitting = true
		return m, tea.Quit
	}
	return m, cmd
}

func (m Model) view_list() string {
	render := renderHeader("LIST")
	render += m.table.View()
	render += renderFooter("sort by: [I]d, [N]ame or [P]hone. [Q]uit. [M]ain menu")
	return render
}

func (m Model) initTable(sorting string) table.Model {
	columns := []table.Column{
		{Title: "ID", Width: 3},
		{Title: "Name", Width: 20},
		{Title: "Phone", Width: 20},
	}

	employees := m.service.List(sorting)
	// Convert to []table.Row
	var rows []table.Row
	for _, e := range employees {
		rows = append(rows, table.Row{
			fmt.Sprintf("%03d", e.Id),
			e.Name,
			e.Phone,
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
