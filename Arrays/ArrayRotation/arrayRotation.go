package main

import (
	"fmt"
)

func main() {
	input := []int{1,2,3,4,5,6,7,8}
	rotation := 2

	tmp := rotate(input, rotation)

	
	/*
	append function expects the first argument to be a slice of type Type, while there can be a variable number of arguments after that. 
	If we have a slice s2 that we want to append to a slice s1, how that will work?
	As from append function syntax, we can’t pass another slice as an argument, it has to be one or many arguments of type Type. 
	Hence, instead, we will use the unpack operator ... to unpack slice into the series of arguments (which is acceptable by append function).
	
	... signifies both pack and unpack operator but if three dots are in the tail position, it will unpack a slice.
	Here s1 and s2 are two slices of the same type. Usually, we know function parameters and how many arguments a function can accept. 
	Then how append function knows how many parameters has passed to it?
	If you look at the signature append function,
	
	func append(slice []Type, elems ...Type) []Type
	You will see elems ...Type which means pack all incoming arguments into elems slice after the first argument.
💡 	One important thing to notice is that only the last argument of a function is allowed to be variadic.

	*/

	fmt.Println(tmp)
	
}

func rotate(input []int, rotation int) []int{
		tmp := append(input[rotation:] , input[:rotation]...)
		return tmp
}
