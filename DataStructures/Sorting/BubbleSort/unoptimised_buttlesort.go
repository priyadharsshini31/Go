package main

import "fmt"

func main(){
	input := []int{10,7,9,4,3,5}
	result := bubbleSort(input)
	fmt.Println(result)
}

// unoptimized way -- swapping and checking (length of input) times

func bubbleSort(input []int) []int {
	for i := 0 ; i< len(input) ; i ++ {
		for j := 0; j< len(input)-1 ;j++ {
			if input[j] > input[j+1] {
				input[j],input[j+1] = input[j+1],input[j]
			}
	}
	}
	return input
}