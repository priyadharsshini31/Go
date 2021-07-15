
package main

import (
	"fmt"
)

type methodError struct{
	status int
	message string 
}

func (m methodError) Error() string{
	return "This ran into an error and the status is"
}

func calculate(val int) error{
	return &methodError{400,"error"}
}


func main() {
	err := calculate(3)
	erVal := err.(*methodError) // type assertion to access the dynamic value of an interface
	if erVal.status == 400 {
		fmt.Println("it is a 400 error status")
	}
	
}