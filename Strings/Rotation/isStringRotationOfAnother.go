package main

import (
	"fmt"
	"strings"
	
)

func main() {
	input := "ABCDEFG"
	input2 := "CDEFGAB"
	fmt.Println(rotation(input, input2))
}

func rotation(input string, input2 string) bool {
	if len(input) != len(input2){
		return false
	}
	temp := input + input 
	return strings.Contains(temp, input2)
}