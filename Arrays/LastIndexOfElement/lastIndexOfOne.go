package main

import (
	"fmt"
)

func main() {
	input := []int{1,2,3,4,5,3,4,2,3,1,1,2,3}
	lastIndex := findLastIndex(input, 1)
	fmt.Println(lastIndex)
}

func findLastIndex(input []int, element int) int{
		lastIndex := 0
		for i := len(input)-1 ; i >= 0 ; i-- {
			if input[i] == element {
				lastIndex = i
				break
			} else {
				lastIndex = -1
			}
		}
		return lastIndex
		
}