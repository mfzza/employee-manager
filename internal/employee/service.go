package employee

import (
	"fmt"
)

type Service struct {
	Employees []Employee
	NextId    int
}

func NewService() *Service {
	// empty slice
	// NOTE: id start from 1
	return &Service{Employees: []Employee{}, NextId: 1}
}

func (s *Service) Add(name string, phone string, position string, email string) {
	emp := NewEmployee(s.NextId, name, phone, position, email)

	// update idNext
	s.NextId++

	s.Employees = append(s.Employees, *emp)
}

func (s *Service) List() {
	// TODO: better format, just return whole employees and let other interface to handle it?
	fmt.Printf("%s\t|\t%s\t|\t%s", "ID", "NAME", "PHONE")
	fmt.Println()
	for _, e := range s.Employees {
		fmt.Println(e.simpleString())
	}
}

func (s *Service) View(id int) {
	// TODO: bound check, maybe just check for -1 value?
	idx := s.indexFromId(id)
	fmt.Println(s.Employees[idx].DetailString())
}

// TODO: implement this
func (s *Service) Edit() {}

func (s *Service) Del(id int) {
	// TODO: bound check
	idx := s.indexFromId(id)
	s.Employees = append(s.Employees[:idx], s.Employees[idx+1:]...)
}

// helper
func (s *Service) indexFromId(id int) int {
	for i, e := range s.Employees {
		if e.id == id {
			return i
		}
	}
	return -1
}
