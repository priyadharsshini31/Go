package main

import (
	"fmt"
)

func main() {
	input := []int{10, 20, 9, 7, 1, 90}
	insertionSort(input)
	fmt.Println(input)
}


func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i-1
		for ; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}
}
