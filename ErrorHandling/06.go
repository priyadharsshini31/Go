
package main

import (
	"fmt"
)

type methodError struct{
	status int
	message string 
}

type argumentError struct {
	status int
	message string
}

func (m methodError) Error() string{
	return "This ran into an error and the status is"
}

func (a argumentError) Error() string{
	return "This ran into an arg error"
}

func calculate(val int) error{
	return &methodError{400,"error"}
}


func main() {
	err := calculate(3)
	//we can also type switch when we donot know the underlying error type
	switch err.(type){
	case *methodError:
		fmt.Println("method error")
	case *argumentError:
		fmt.Println("arg error")
	}
	
}