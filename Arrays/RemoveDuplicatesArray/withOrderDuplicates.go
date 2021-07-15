package main

import (
	"fmt"
)

func main() {
	input := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	Result := removeDuplicates(input)
	fmt.Println(Result)
}

func removeDuplicates(input []int) []int {
	unique := make(map[int]bool) //for map you need not specify the size in make
	var result []int
	for _, i := range input {
		if _, ok := unique[i]; !ok {

			result = append(result, i)
			unique[i] = true
		}

	}

	return result

	/* to check if an element is present in map already
	   if value , ok : = map[key]; ok {
		fmt.Println("the key is already present")
	   }
	*/
}