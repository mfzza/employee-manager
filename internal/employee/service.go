package employee

import "fmt"

type Service struct {
	Employees []Employee
}

var nextId int

func (s *Service) Add(name string, phone string, position string, email string) {
	emp := Employee{nextId, name, phone, position, email}
	// update idNext
	nextId++

	s.Employees = append(s.Employees, emp)
}

func (s *Service) List(){
	for _, v := range s.Employees {
		fmt.Printf("%+v\n", v)
	}
}
