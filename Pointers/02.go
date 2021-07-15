package main

import (
	"fmt"
)

func pass(a *int) {
	*a = 6
}

func change(a []int){
	fmt.Println(a)
	a[0] = 3
}

	

func main() {
	a := 5 
	fmt.Println(a)
	pass(&a)
	fmt.Println(a)
	b := &a
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(*b)
	var stringpointer *string 
	st := "priya"
	fmt.Println("Initail value", st)
	stringpointer = &st
	*stringpointer = "mayank"
	fmt.Println("changed the value of string", *stringpointer)
	var intpointer *int
	fmt.Println(intpointer)
	
	newpointer := new(int)
	fmt.Println(newpointer)
	fmt.Println(&newpointer)
	
	arr := []int{1,2,3}
	change(arr)
	fmt.Println(arr)	
	
	
}
