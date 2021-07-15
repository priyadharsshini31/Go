package main

import (
	"fmt"
)

//GO IS PASS BY VALUE 

type Person struct {
	first_name string 
	last_name string
}

func changeName(p Person) {
	p.first_name = "mayank"
}

func changeNamePointer( p *Person) {
	fmt.Println(&p)
	p.first_name = "mayank"
	
}

func main() {
	person := Person {
		first_name : "Priya",
		last_name : "dharsshini",
	}
	changeName(person)
	fmt.Println(person)
	
	changeNamePointer(&person)
	fmt.Println(person)
	
	fmt.Printf("Address of person %p", &person)
	
}

/*
{Priya dharsshini}
0xc000082020
{mayank dharsshini}
Address of person 0xc00008c000
*/