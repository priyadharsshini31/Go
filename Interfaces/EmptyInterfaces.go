package main

import (
	"fmt"
)

type myString string

func explain(i interface{}){
	fmt.Println(i)
}



func main() {
	fmt.Println("EMPTY INTERFACES")
	
	/*
	Have you wondered how does the Println function from the fmt built-in package accepts the different types of value as arguments? 
	This is possible because of an empty interface. Let’s see the signature of Println function.
	func Println(a ...interface{}) (n int, err error)
	*/
	
	a := "hello"
	explain(a)
	b := 5
	explain(b)
	myString := "user defined type"
	explain(myString)
	
	
}
