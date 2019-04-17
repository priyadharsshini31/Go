package main

import (
	"fmt"
)

func main() {
	numbers := [7]int {2,3,4,5,6,7,8}
	sum := 12
	
	for i := 0 ; i < len(numbers); i++ {
		add := 0
		for j := i ; j < len(numbers); j++ {
			add = add + numbers[j]
			if add == sum{
				fmt.Println(i,j)
				
			}else if add > sum{
				fmt.Println(add)
				break
			}
		}
}
}
