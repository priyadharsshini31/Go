package main

import (
	"fmt"
)

func paranthesisChecker(s string) bool {
	a := make([]string, len(s))
	characters := []rune(s)
	for _, val1 := range characters {
		val := string(val1)
		if val == "{" || val == "(" || val == "[" {
			a = append(a, val)
		} else {

			if len(a) == 0 {
				return false
			}

			lastelement := len(a) - 1
			element := string(a[lastelement])
			a = a[0:lastelement]

			switch val {

			case ")":
				if element != "(" {
					return false
				}

			case "}":
				if element != "{" {
					return false
				}

			case "]":
				if element != "[" {
					return false
				}

			}
		}
	}

	//checking the first element for zero value of string because, initially we created string with make function.. so the length will always be that length
	
	if a[0] == "" {

		return true
	}
	return false
}

func main() {
	s := "{()[]}"
	fmt.Println(s)
	isBalanced := paranthesisChecker(s)
	if isBalanced {
		fmt.Println("The expression is balanced")
	} else {
		fmt.Println("The expression is not balanced")
	}
}
