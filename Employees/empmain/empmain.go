package main

import (
	"Assignment1/Employees"
)

func main() {
	company := employees.NewCompany()

	company.AddEmployee(employees.FullTimeEmployee{ID: 1, Name: "Yerassyl", Salary: 500000000}, "1")
	company.AddEmployee(employees.PartTimeEmployee{ID: 2, Name: "Mbappe", HourlyRate: 2000, HoursWorked: 20.5}, "2")

	company.ListEmployees()
}
