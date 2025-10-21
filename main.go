package main

import (
	"employee-management/pkg"
)

func main() {
	// NOTE: Test, delete later
	emps := pkg.EmployeeService{}
	emps.Load("test.json")
	emps.List()
	emps.Add("fikri", "boss", "081-1234-1234", "fikri@gmail.com")
	emps.Add("fikro", "boss", "081-1234-1234", "fikri@gmail.com")
	emps.Save("test.json")

}
