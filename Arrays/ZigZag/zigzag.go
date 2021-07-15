package main

import (
	"fmt"
)

func main() {
	a := []int{4, 3, 7, 8, 6, 2, 1}
	result := zigzag(a)
	fmt.Println(result)
}

func zigzag(a []int) []int {
	flag := true
	for i := 0; i < len(a) -1; i++ {
		if flag == true{
			flag = false
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		} else {
		
			flag = true
			if a[i] < a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
	return a

}
