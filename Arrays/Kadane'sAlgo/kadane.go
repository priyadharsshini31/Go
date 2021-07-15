package main

import (
	"fmt"
)

func kadane(a []int){
	maxsum := 0
	cursum := 0
	for i := 0 ; i < len(a); i++ {
		cursum = cursum + a[i]
		if maxsum < cursum {
			maxsum = cursum
		} 
		
		if cursum < 0 {
			cursum = 0
		}
}

fmt.Println(maxsum)
}

func main() {
	a := []int{-2,-3,4,-1,-2,1,5,-3}
	kadane(a)
	
}
