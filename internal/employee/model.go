package employee

import "fmt"

type Employee struct {
	id       int
	name     string
	phone    string
	position string
	email    string
}

func (e *Employee) idString() string {
	return fmt.Sprintf("%03d", e.id)
}

func (e *Employee) printDetail() {
	fmt.Println("=============================")
	fmt.Println("ID:", e.idString())
	fmt.Println("name:", e.name)
	fmt.Println("Phone:", e.phone)
	fmt.Println("Position:", e.position)
	fmt.Println("Email:", e.email)
	fmt.Println("=============================")
}
