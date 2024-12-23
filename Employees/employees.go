package employees

import "fmt"

type Employee interface {
	GetDetails() string
}

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}

func (f FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Salary: %d Tenge", f.ID, f.Name, f.Salary)
}

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

func (p PartTimeEmployee) GetDetails() string {
	totalSalary := float64(p.HourlyRate) * float64(p.HoursWorked)
	return fmt.Sprintf("ID: %d, Name: %s, Total Salary: %.2f Tenge", p.ID, p.Name, totalSalary)
}

type Company struct {
	Employees map[string]Employee
}

func NewCompany() *Company {
	return &Company{Employees: make(map[string]Employee)}
}

func (c *Company) AddEmployee(emp Employee, id string) {
	c.Employees[id] = emp
	fmt.Println("Employee added successfully!")
}

func (c *Company) ListEmployees() {
	fmt.Println("Employees List:")
	for _, emp := range c.Employees {
		fmt.Println(emp.GetDetails())
	}
}
