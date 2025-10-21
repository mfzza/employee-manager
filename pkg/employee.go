package pkg

import (
	"encoding/json"
	"fmt"
	"os"
)

type Employee struct {
	Id       int
	Name     string
	Phone    string
	Position string
	Email    string
}

func idToString(id int) string {
	return fmt.Sprintf("%03d", id)
}

func (emp *Employee) PrintDetail() {
	fmt.Println("=============================")
	fmt.Println("ID:", idToString(emp.Id))
	fmt.Println("name:", emp.Name)
	fmt.Println("Phone:", emp.Phone)
	fmt.Println("Position:", emp.Position)
	fmt.Println("Email:", emp.Email)
	fmt.Println("=============================")
}

type EmployeeService struct {
	Employees []Employee
}

var nextId int

func (emps *EmployeeService) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Errorf("Failed to read file: %w", err)
	}

	if len(data) == 0 {
		return fmt.Errorf("file is empty")
	}

	err = json.Unmarshal(data, &emps.Employees)
	if err != nil {
		return fmt.Errorf("failed to unmarshal json: %w", err)
	}

	// TODO: is it safe?
	nextId = emps.Employees[len(emps.Employees)-1].Id

	return nil
}

func (emps *EmployeeService) Save(filename string) error {
	data, err := json.Marshal(emps.Employees)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (emps *EmployeeService) Add(name string, phone string, position string, email string) {
	emp := Employee{nextId, name, phone, position, email}
	// update idNext
	nextId++

	emps.Employees = append(emps.Employees, emp)
}

func (emps *EmployeeService) List(){
	fmt.Println(emps.Employees)
}

// TODO: view detail, list, edit, delete
// TODO: interactive logic
