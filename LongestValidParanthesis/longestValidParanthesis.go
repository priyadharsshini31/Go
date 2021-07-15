package main

import (
	"fmt"
)

/*Instead of finding every possible string and checking its validity, we can make use of stack 
while scanning the given string to check if the string scanned so far is valid, and also the 
length of the longest valid string. In order to do so, we start by pushing -1−1 onto the stack.

For every \text{‘(’}‘(’ encountered, we push its index onto the stack.

For every \text{‘)’}‘)’ encountered, we pop the topmost element and subtract the current 
element's index from the top element of the stack, which gives the length of the currently 
encountered valid string of parentheses. If while popping the element, the stack becomes empty, 
we push the current element's index onto the stack. In this way, we keep on calculating the 
lengths of the valid substrings, and return the length of the longest valid string at the end.*/

func findLongestValidParanthesis(s string) int {
	finalLength := 0
	var inputArray []int
	inputByte := []rune(s)
	inputArray = append(inputArray, -1)
	for i, val := range inputByte {
		if string(val) == "(" {
			inputArray = append(inputArray, i)

		} else {
			inputArray = inputArray[:len(inputArray)-1]
			if len(inputArray) == 0 {
				inputArray = append(inputArray, i)
			}
			if finalLength < i-inputArray[len(inputArray)-1] {
				
				finalLength = i-inputArray[len(inputArray)-1]
			}
		}

	}
	return finalLength
}

func main() {
	input := "()(()"
	length := findLongestValidParanthesis(input)
	fmt.Println(length)
}
