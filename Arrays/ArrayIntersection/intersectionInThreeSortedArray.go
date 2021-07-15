package main

import (
	"fmt"
)

func main() {
	a := []int{2,4,6,8,10}
	b := []int{2,3,4,5,6}
	c := []int{2,3,4}
	intersection(a,b,c)
}

func intersection(a []int, b []int, c []int) {
		i := 0
		j := 0
		k := 0
		for (i < len(a) && j < len(b) && k < len(c)){
			if (a[i] == b[j] && a[i] == c[k]) {
				fmt.Println(a[i])
				i++
				j++
				k++
			} else if a[i] <= b[j] && a[i] <= c[k] {
				i++
			} else if b[j] <= a[i] && b[j] <= c[k] {
				j++
			} else {
				k++
			}
		}
				
}

