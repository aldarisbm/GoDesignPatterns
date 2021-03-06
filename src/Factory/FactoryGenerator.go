package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome int
}

// functional approach
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee{
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

type EmployeeFactory struct {
	Position string
	AnnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func main() {
	developerFactory := NewEmployeeFactory("developer", 100000)
	managerFactory := NewEmployeeFactory("manager", 80000)

	developer := developerFactory("Adamn")
	manager := managerFactory("Jay")

	fmt.Println(developer)
	fmt.Println(manager)

	bossFactory := NewEmployeeFactory2("CEO", 300000)
	//allows for some modification
	bossFactory.AnnualIncome = 350000
	pepito := bossFactory.Create("Pepito")
	fmt.Println(pepito)
}
