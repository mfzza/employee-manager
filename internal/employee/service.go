package employee

import (
	"fmt"
)

type Service struct {
	Employees []Employee // TODO: should not store it in-memory(?)
	Repo      Repository
	NextId    int
}

func NewService(repo Repository) *Service {
	// empty slice
	// NOTE: id start from 1
	return &Service{Employees: []Employee{}, Repo: repo, NextId: 1}
}

func (s *Service) Add(name string, phone string, position string, email string) {
	s.updateLastId()
	addedEmployee := NewEmployee(s.NextId, name, phone, position, email)

	s.Employees = append(s.Employees, *addedEmployee)
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
		if e.Id == id {
			return i
		}
	}
	return -1
}

func (s *Service) updateLastId() int {
	lastIndex := len(s.Employees) - 1
	// return 0 since it empty
	if lastIndex < 0 {
		s.NextId = 1
		return 1
	}
	s.NextId = s.Employees[lastIndex].Id + 1
	return s.NextId

}
