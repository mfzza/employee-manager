package employee

import (
	"fmt"
	"sort"
)

type Service struct {
	Employees []Employee // TODO: should not store it in-memory(?)
	Repo      *Repository
	NextId    int
}

type employeeSimple struct {
	Id    int
	Name  string
	Phone string
}

func NewService(repo Repository) *Service {
	// empty slice
	// NOTE: id start from 1
	return &Service{Employees: []Employee{}, Repo: &repo, NextId: 1}
}

func (s *Service) Add(name string, phone string, position string, email string) {
	s.updateLastId()
	addedEmployee := NewEmployee(s.NextId, name, phone, position, email)

	s.Employees = append(s.Employees, *addedEmployee)
}

func (s *Service) List(sorting string) []employeeSimple {
	var list []employeeSimple
	for _, e := range s.Employees {
		list = append(list, employeeSimple{e.Id, e.Name, e.Phone})
	}

	sortById := func(i, j int) bool { return list[i].Id < list[j].Id }
	sortByPhone := func(i, j int) bool { return list[i].Phone < list[j].Phone }
	sortByName := func(i, j int) bool { return list[i].Name < list[j].Name }

	switch sorting {
	case "name": sort.Slice(list, sortByName)
	case "phone": sort.Slice(list, sortByPhone)
	default: sort.Slice(list, sortById) // NOTE: can i just remove this? since in json should be already sorted
	}

	return list
}

func (s *Service) View(id int) error{
	idx := s.indexFromId(id)
	if idx < 0 {
		return fmt.Errorf("Cant view data, id: %d not exists", id)
	}

	fmt.Println(s.Employees[idx].DetailString())
	return nil
}

// TODO: implement this, should be able to let user select what they want to edit, maybe wait for cli implementation?
func (s *Service) Edit() {}

func (s *Service) Del(id int) error {
	idx := s.indexFromId(id)
	if idx < 0 {
		return fmt.Errorf("Cant delete data, id: %d not exists", id)
	}

	s.Employees = append(s.Employees[:idx], s.Employees[idx+1:]...)
	return nil
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
