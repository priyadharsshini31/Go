package main

import (
	"fmt"
)

type Employee struct {
	firstName, lastName string // Multiple declaration of same type
	age int 
	salary Salary  //Nested struct
	Address // Filed promotion eg
	bool //Ananonymous filed
}

type Salary struct {
	HRA int 
	basic int 
}

type Address struct {
	pincode int 
}

func main() {
	ross := Employee {
		firstName: "Ross",
		lastName: "Geller",
		age: 40,
		salary : Salary {
			HRA : 100,
			basic : 200,
		},
		Address: Address{
			pincode : 6000000,
		},
		}
fmt.Println("The firstName is ", ross.firstName)
fmt.Println("The Address is", ross.pincode)
fmt.Println("The salary is",ross.salary.HRA)

monica := &Employee {
		firstName : "Monica",
	}
fmt.Println(monica.firstName)


anonymous := struct{
	id  string
	phone int
}{
	id : "12",
	phone: 1234,
}

fmt.Println(anonymous.phone)
}
