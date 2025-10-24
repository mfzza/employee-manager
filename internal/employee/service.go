package employee

import (
	"errors"
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

func NewService(repo *Repository) (*Service, error) {
	employees, err := repo.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load employees: %w", err)
	}
	return &Service{Employees: employees, Repo: repo, NextId: 1}, nil
}

func (s *Service) AddEmployee(name string, phone string, position string, email string) error {
	e := &Employee{
		Name:     name,
		Phone:    phone,
		Position: position,
		Email:    email,
	}

	if err := validateEmployeeInfo(e); err != nil {
		return err
	}

	s.updateLastId()
	e.Id = s.NextId
	s.Employees = append(s.Employees, *e)

	return s.Repo.Save(s.Employees)
}

func (s *Service) GetAllEmployee(sorting string) []employeeSimple {
	var list []employeeSimple
	for _, e := range s.Employees {
		list = append(list, employeeSimple{e.Id, e.Name, e.Phone})
	}

	sortById := func(i, j int) bool { return list[i].Id < list[j].Id }
	sortByPhone := func(i, j int) bool { return list[i].Phone < list[j].Phone }
	sortByName := func(i, j int) bool { return list[i].Name < list[j].Name }

	switch sorting {
	case "name":
		sort.Slice(list, sortByName)
	case "phone":
		sort.Slice(list, sortByPhone)
	default:
		sort.Slice(list, sortById) // NOTE: can i just remove this? since in json should be already sorted
	}

	return list
}

func (s *Service) GetEmployeeById(id int) (*Employee, error) {
	idx := s.indexFromId(id)
	if idx < 0 {
		return nil, fmt.Errorf("Cant view employee, id: %d not exists", id)
	}
	return &s.Employees[idx], nil
}

func (s *Service) ModifyEmployee(id int, name string, phone string, position string, email string) error {
	idx := s.indexFromId(id)
	if idx < 0 {
		return fmt.Errorf("Cant modify employee, id: %d not exists", id)
	}

	e := &Employee{
		Name:     name,
		Phone:    phone,
		Position: position,
		Email:    email,
	}

	if err := validateEmployeeInfo(e); err != nil {
		return err
	}

	s.Employees[idx].Name = name
	s.Employees[idx].Phone = phone
	s.Employees[idx].Position = position
	s.Employees[idx].Email = email

	return s.Repo.Save(s.Employees)
}

func (s *Service) DeleteEmployee(id int) error {
	idx := s.indexFromId(id)
	if idx < 0 {
		return fmt.Errorf("Cant delete data, id: %d not exists", id)
	}

	s.Employees = append(s.Employees[:idx], s.Employees[idx+1:]...)

	return s.Repo.Save(s.Employees)
}

// helper
func validateEmployeeInfo(e *Employee) error {
	var errs []error
	var err error

	// return err, will also trim whitespace for name and position
	e.Name, err = validateName(e.Name)
	if err != nil {
		errs = append(errs, err)
	}
	if err := validatePhone(e.Phone); err != nil {
		errs = append(errs, err)
	}
	e.Position, err = validatePosition(e.Position)
	if err != nil {
		errs = append(errs, err)
	}
	if err := validateEmail(e.Email); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

// NOTE: maybe try to make validate ID as helper
// func (s *Service) validateId() error {
// }

func (s *Service) indexFromId(id int) int {
	for i, e := range s.Employees {
		if e.Id == id {
			return i
		}
	}
	return -1
}

func (s *Service) updateLastId() {
	// update NextId to id + 1 from last employee
	lastIndex := len(s.Employees) - 1
	// return 0 since it empty
	if lastIndex < 0 {
		s.NextId = 1
		return
	}
	s.NextId = s.Employees[lastIndex].Id + 1

}
