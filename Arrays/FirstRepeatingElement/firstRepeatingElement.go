package main

import (
	"fmt"
)

func main() {
	input := []int{1,2,3,4,5,2,1,2,3,4,5,6,7,8,9}
	firstRepeating := findFirstRepeating(input)
	fmt.Println(firstRepeating)
}

func findFirstRepeating(input []int)int{
	repeating := make(map[int]bool)
	firstRepeating := input[0]
	for _ , element := range input {
		if _, ok := repeating[element]; !ok {
			repeating[element] = true
		} else {
			firstRepeating = element
			break
		}
	}
	return firstRepeating
}