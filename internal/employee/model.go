package employee

import "fmt"

type Employee struct {
	Id       int
	Name     string
	Phone    string
	Position string
	Email    string
}

func NewEmployee(id int, name string, phone string, position string, email string) *Employee {
	return &Employee{id, name, phone, position, email}
}

func (e *Employee) idString() string {
	return fmt.Sprintf("%03d", e.Id)
}

func (e *Employee) simpleString() string {
	return fmt.Sprintf("%s\t|\t%s\t|\t%s", e.idString(), e.Name, e.Phone)
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
		e.idString(), e.Name, e.Phone, e.Position, e.Email)
}
