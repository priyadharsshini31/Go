package main

import (
	"fmt"
)

func main() {
	numbers := [5]int {2,3,4,5,6}
	sum := 12
	
	for i := 0 ; i < len(numbers); i++ {
		add := 0
		for j := i ; j < len(numbers); j++ {
			add = add + numbers[j]
			if add == sum{
				fmt.Println(i,j)
				
			}
		
}
}
}
