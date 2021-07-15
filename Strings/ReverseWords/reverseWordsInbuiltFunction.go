package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "Hello How are you"
	fmt.Println(reverseWords(input))
}

func reverseWords(input string) string {
	words := strings.Split(input, " ")
	result := ""
	for i := len(words)-1; i >= 0; i-- {
		result = result + " " + words[i]
	}
	return result
}