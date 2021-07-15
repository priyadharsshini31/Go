package main

import (
	"fmt"
)

type methodError struct{}

func (m methodError) Error() string{
	return "Method is wrong"
}


func main() {
	err := methodError{}
	fmt.Println(err)
}


/*
EXPLANATION
--------------
type error interface {
  Error() string
}

so methodError implemets error , so err (under main) is of type error 

println is designed in such a way that it calls the Error() method if the value is error type
*/