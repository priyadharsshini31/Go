package main

import (
	"fmt"
)

func main() {
	input := []int{10, 20, 30, 40, 50}
	element := 40
	fmt.Println(binarySearch(input, element))

}

func binarySearch(arr []int, element int) int {
	mid := len(arr) / 2
	if element == arr[mid] {
		return mid
	} else if element < arr[mid] {
		for i := 0; i < mid; i++ {
			if arr[i] == element {
				return i
			}
		}
	} else {
		for i := mid; i < len(arr); i++ {
			if arr[i] == element {
				return i
			}
		}
	}
	return 0

}