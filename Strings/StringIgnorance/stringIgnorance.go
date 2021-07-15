package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "It is a long day Dear"
	result := ignoreEvenoccurance(input)
	fmt.Println(result)
}

func ignoreEvenoccurance(input string) string {
	temp := make(map[string]int)
	input1 := strings.ToLower(input)
	result := ""
	for i, value := range input1 {
		temp[string(value)]++
		if temp[string(value)]%2 != 0 {
			result = result + string(input[i])
		}
	}
	return result
}
