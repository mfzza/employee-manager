package main

import (
	"fmt"
	"os"

	"employee-management/internal/employee"
	"employee-management/internal/cli"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// NOTE: Test, delete later
	srv := employee.NewService()
	srv.Add("gopher", "123-2134-1234", "cool dude", "gopher@gomail.com")
	srv.Add("gopher", "123-2134-1234", "cool dude", "gopher@gomail.com")
	srv.Add("gopher", "123-2134-1234", "cool dude", "gopher@gomail.com")
	srv.Add("gopher", "123-2134-1234", "cool dude", "gopher@gomail.com")
	srv.Add("gopher", "123-2134-1234", "cool dude", "gopher@gomail.com")

	srv.List()
	srv.Del(2)
	srv.List()
	l := srv.List("phone")
	for _, e := range l {
		fmt.Printf("ID: %d\t Name: %s  Phone: %s\n", e.Id, e.Name, e.Phone)
	}
	p := tea.NewProgram(cli.Model{})
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
