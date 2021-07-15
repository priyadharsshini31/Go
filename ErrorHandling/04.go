package main

import (
	"fmt"
)

type methodError struct{
	status int
	message string 
}

func (m methodError) Error() string{
	return "This ran into an error and the status is "
}

func calculate(val int) error{
	return methodError{400,"error"}
}


func main() {
	err := calculate(3)
	fmt.Println(err)
	errval := err.(methodError) //Type assertion to access the dynamic values of an interface
	if errval.status == 400 {	
		fmt.Println("status errror")
	}
}