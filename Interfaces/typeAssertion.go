package main

import (
	"fmt"
)

type Employee interface {
	getDetails()
}

type Salary interface {
	getSalary()
}

type Manager struct {
	name   string
	salary int
}

type Lead struct {
	name   string
	salary int
}

func (m Manager) getDetails() {
	fmt.Println(m.name + "MANAGER")
}

func (m Manager) getSalary() {
	fmt.Println(m.salary * 2)
}

func (l Lead) getDetails() {
	fmt.Println(l.name + "LEAD")
}

func (l Lead) getSalary() {
	fmt.Println(l.salary)
}

func main() {
	fmt.Println("MULTIPLE INTERFACES")

	var e Employee
	e = Manager{"Priya", 123}
	e.getDetails()
	//TYPE ASSERTION
	//SYNTAX : i.(Type)
	salary, ok := e.(Salary)
	if !ok {
		fmt.Println("does not implement salary interface or doesnot hold any value")
	}
	salary.getSalary()
	e = Lead{"Priya", 635}
	e.getDetails()

}
