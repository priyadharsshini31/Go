package main

import (
	"fmt"
)

func main() {
	input := []int{1, 1,1,2,5,1,4,5,3}
	firstNonRepeating := findFirstNonRepeating(input)
	fmt.Println(firstNonRepeating)
}

func findFirstNonRepeating(input []int) int {
	flag := false
	firstNonRepeating := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if (i != j && input[i] == input[j]) {
				
				break
			}

			if j == len(input)-1 {
				firstNonRepeating = input[i]
				flag = true
			
			}
		}
		if flag == true {
			break
		}

	}
	return firstNonRepeating
}