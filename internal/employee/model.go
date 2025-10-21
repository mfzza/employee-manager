package employee

import "fmt"

type Employee struct {
	id       int
	name     string
	phone    string
	position string
	email    string
}

func NewEmployee(id int, name string, phone string, position string, email string) *Employee {
	return &Employee{id, name, phone, position, email}
}

func (e *Employee) idString() string {
	return fmt.Sprintf("%03d", e.id)
}

func (e *Employee) simpleString() string {
	return fmt.Sprintf("%s\t|\t%s\t|\t%s", e.idString(), e.name, e.phone)
}

func (e Employee) DetailString() string {
	return fmt.Sprintf(
		"=============================\n"+
			"ID: %s\n"+
			"Name: %s\n"+
			"Phone: %s\n"+
			"Position: %s\n"+
			"Email: %s\n"+
			"=============================",
		e.idString(), e.name, e.phone, e.position, e.email)
}
