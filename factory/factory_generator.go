package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// functional approach
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// structure approach
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func main() {
	developerFactory := NewEmployeeFactory("Developer", 60000)
	managerFactory := NewEmployeeFactory("Manager", 80000)

	developer := developerFactory("Adam")
	manager := managerFactory("Jane")

	fmt.Println(developer)
	fmt.Println(manager)

	bossFactory := NewEmployeeFactory2("CEO", 100000)
	bossFactory.AnnualIncome = 110000
	boss := bossFactory.Create("Sam")
	fmt.Println(boss)
}
