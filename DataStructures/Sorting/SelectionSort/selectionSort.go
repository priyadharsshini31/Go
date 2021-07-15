package main

import (
	"fmt"
)

func main() {
	input := []int{10, 9, 7, 90, 0}
	selectionSort(input)
	fmt.Println(input)
}

func selectionSort(input []int) {
	j := 0
	temp := 0
	tempindex := 0
	for i := 0; i <= len(input)-1; i++ {
		temp = input[j]
		for ; j <= len(input)-1; j++ {
			if input[j] < temp {
				temp = input[j]
				tempindex = j
			}
		}
		input[i], input[tempindex] = temp, input[i]
		j = i + 1
	}
}