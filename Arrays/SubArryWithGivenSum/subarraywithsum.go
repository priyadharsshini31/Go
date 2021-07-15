package main

import (
	"fmt"
)

func main() {
	input := []int{1,-1,0}
	i := subArrayWithSum(input)
	fmt.Println(i)
}

func subArrayWithSum(input []int) int{
sum := 0
   tempsum := 0
	count := 0
	for i := 0; i < len(input); i++ {
		tempsum = 0
		for j := i ; j< len(input); j++ {
			
			tempsum = tempsum + input[j]
			if tempsum == sum{
				count++
				
			} 
		}
	}
return count
}
		