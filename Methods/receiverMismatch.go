package main

import(
   "fmt"
)

type name struct{
	firstname, lastname string
}

func printFullName(n name){
	fmt.Println(n.firstname + n.lastname);
}

func(n name) printFullName(){
        fmt.Println(n.firstname + n.lastname);
}

func(n *name) printFullNameWithSpace(){
	fmt.Println(n.firstname +" "+ n.lastname);
}

func printFullNameWithSpace(n *name){
	fmt.Println(n.firstname +" "+ n.lastname);
}

func main(){
	n := name{"alice","wonder"}
	m := &n
	printFullName(n)
	//printFullName(m) func with value arg can accept only value arg
	//methods receivers with value receivers can accept both value and pointer receivers
	n.printFullName()
	m.printFullName() //Internally converted by go as (*m).printFullName() as it has value receiver
	
	printFullNameWithSpace(m)
	//printFullNameWithSpace(n) func with pointer arg can accept only pointer arg
	//methods receivers with pointer receivers can accept both value and pointer receivers
	n.printFullNameWithSpace() //Internally converted by go as (&n).printFullName() as it has pointer receiver
	m.printFullNameWithSpace()

}
