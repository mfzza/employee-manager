package main

import (
	"fmt"
	"os"

	"employee-management/internal/employee"
	"employee-management/internal/cli"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	repo := employee.NewRepository("./test.json")
	srv := employee.NewService(*repo)

	var err error
	srv.Employees, err = repo.Load()
	if err != nil {
		fmt.Println("Error loading file:", err)
	}

	p := tea.NewProgram(cli.InitialModel(srv))
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
