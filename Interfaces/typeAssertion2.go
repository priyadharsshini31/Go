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

type Fun interface {
	getfun()
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

	e = Lead{"Priya", 635}
	e.getDetails()
	
	//How would we know if the underlying value of an interface implements any other interfaces? This is also possible using type assertion
	_ , ok := e.(Fun)
	if !ok {
		fmt.Println("does not implement")
	} 

}
