package main

import (
	"fmt"
	
)

func calculate(val int) error {
	 	return fmt.Errorf("wrong value: %d", val)
}

func main() {
	err := calculate(0)
	if err != nil {
		fmt.Println(err)
	}
}

/*
We can also have interpolated error messages using the fmt package 
*/