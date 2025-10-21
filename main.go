package main

import (
	"employee-management/internal/employee"
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
}
