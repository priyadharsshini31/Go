package main

import (
	"fmt"
)

func main() {
	input := []int{1,1,2,2,3,3,4,4,5,5,6,6,7,7,8,8,9,9}
	Result := removeDuplicatesWithoutOrder(input)
	fmt.Println(Result)
}

func removeDuplicatesWithoutOrder(input []int) []int {
	unique := make(map[int]int,100)
	var result []int
	for _, i := range input {
		unique[i]++
	}
	for key, _ := range unique {
		result = append(result, key)
	}
	return result
}