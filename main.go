package main

import (
	"fmt"
	"os"

	"employee-management/internal/cli"
	"employee-management/internal/employee"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	repo := employee.NewRepository("./MOCK_DATA.json")
	srv, err := employee.NewService(repo)
	if err != nil {
		fmt.Println("Error loading employees:", err)
		os.Exit(1)
	}

	p := tea.NewProgram(cli.InitialModel(srv))
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
