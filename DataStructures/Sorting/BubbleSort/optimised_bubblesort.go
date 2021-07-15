package main

import (
	"fmt"
)

func bubblesort(a []int) {
	for i := 0; i < len(a); i++ {
		swap := false
		fmt.Println(i)
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
				swap = true
			}
		}

		if swap == false {
			break
		}
	}
	fmt.Println("The sorted array is")
	fmt.Println(a)
}
func main() {
	arr := []int{3, 1, 5, 6, 7, 8, 10}
	bubblesort(arr)
}
