package main

import "fmt"

type employee struct {
	lastName  string
	firstName string
	role      string
	salary    float64
}

type Employee interface {
	// GetLastName() string
	// GetFirstName() string
	SetLastname(ln string)
	SetFirstName(fn string)
	SetRole(r string)
	GetRole() string
	SetSalary(s float64)
	GetSalary() float64
	String() string
}

type partTimeEmployee struct {
	employee
	hourlyWage float64
}

type PartTimeEmployee interface {
	Employee
	SethourlyWage(w float64)
	GetHourlyWage() float64
	// String() string
}

func (p *employee) SetLastname(ln string) {
	p.lastName = ln
}

func (p *employee) SetFirstName(fn string) {
	p.firstName = fn
}

func (p employee) GetFirstName() string {
	return p.firstName
}

func (p employee) GetLastName() string {

	return p.lastName
}
func (p *employee) SetRole(r string) {
	p.role = r
}

func (p employee) GetRole() string {
	return p.role
}

func (p *employee) SetSalary(s float64) {
	p.salary = s
}

func (p employee) GetSalary() float64 {
	return p.salary
}

func (p employee) String() string {
	result := "Name: " + p.firstName + " " + p.lastName + "\n"
	result += "Role: " + p.role + "\n"
	result += "Annual salary: $" + fmt.Sprintf("%0.2f", p.salary) + "\n"
	return result
}

func (p partTimeEmployee) String() string {
	result := "Name: " + p.firstName + " " + p.lastName + "\n"
	result += "Role: " + p.role + "\n"
	result += "HourlyWage: $" + fmt.Sprintf("%0.2f", p.hourlyWage) + "\n"
	return result
}

func (p *partTimeEmployee) SethourlyWage(wage float64) {
	p.hourlyWage = wage
}

func (p partTimeEmployee) GetHourlyWage() float64 {
	return p.hourlyWage
}

func main() {
	person := new(employee)
	person.SetFirstName("Helen")
	person.SetLastname("Rose")
	person.SetRole("Technical Lead")
	person.SetSalary(12_644.0)
	fmt.Println(person.String())

	hourlyWorker := new(partTimeEmployee)
	hourlyWorker.SetFirstName("Mark")
	hourlyWorker.SetLastname("Smith")
	hourlyWorker.SetRole("Software Developer")
	hourlyWorker.SethourlyWage(85.00)
	fmt.Println(hourlyWorker.String())
}
