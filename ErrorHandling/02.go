package main

import (
	"fmt"
	"errors"
)

func calculate(val int) error {
	 	return errors.New("0 is not accepted")
}

func main() {
	err := calculate(0)
	if err != nil {
		fmt.Println(err)
	}
}

/*
Instead of creating a type that implements the error type, go provides a package called 'errors' that has New method that will
take a string and return an error
*/

/*NEW BEHIND SCENES
type errorString stuct {
		s string
}

func (e *errorString) Error() string{
		return &e.s
}

func New(text String)error {
		return &errorString{text}
}

*/