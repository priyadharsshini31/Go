package main

import (
	"fmt"
)

func findEquilibrium(a []int){
	sum := 0
	leftsum := 0
	for _,val := range a{
		sum += val
	}
	for i, val := range a{
		sum = sum - val
		if leftsum == sum {
			fmt.Println("The equilibrium point is", i)
			return
		}
		leftsum = leftsum + val
		}
fmt.Println("no equilib point")
}
		
func main() {
	arr := []int{-7, 1, 5, 2, -4, 3, 0} 
	findEquilibrium(arr)
}
