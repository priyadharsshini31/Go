package main

import (
	"fmt"
)

func main() {
	a := []int{4, 3, 7, 8, 6, 2, 1}
	largest, smallest := largestAndSmallest(a)
	fmt.Println(largest)
	fmt.Println(smallest)
}

func largestAndSmallest(a []int) (int, int) {
	largest := a[0]
	smallest := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > largest {
			largest = a[i]
		}
		if a[i] < smallest {
			smallest = a[i]
		}

	}
	return largest, smallest

}
