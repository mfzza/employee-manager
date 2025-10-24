package main

import (
	"fmt"
	"os"

	"employee-management/internal/cli"
	"employee-management/internal/employee"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	repo := employee.NewRepository("./test.json")
	srv := employee.NewService(*repo)

	var err error
	srv.Employees, err = repo.Load()
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
