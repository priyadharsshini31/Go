package main

import (
	"fmt"
	//"strings"
)

func main() {
	var output []byte
	input := "priyadharsshini"
	input1 := []byte(input)
	for i := 0; i< len(input); i++{
	if input1[i] =='a'{
	output = append(input1[:i],input1[i+1:]...)
	}
	}
	fmt.Println("the final output is", string(output))
	
}
