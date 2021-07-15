package main

import (
	"fmt"
)

func lcs(text1 string, text2 string) {
	m := len(text1)
	n := len(text2)
	//to create 2 dimensional array
	a := make([][]int, m+1)
    	for i := range a {
        	a[i] = make([]int, n+1)
		}
		
	for i := 0; i <= m ; i++ {
		for j := 0; j <= n;j++ {
			if i == 0 || j == 0 {
				a[i][j] = 0
			} else if text1[i-1] == text2[j-1] {
				a[i][j] = 1+a[i-1][j-1]
			} else {
				if a[i-1][j] > a[i][j-1]{
					a[i][j] = a[i-1][j]
				} else {
					a[i][j] = a[i][j-1]
				}
			}
		}
	}
	fmt.Println(a[m][n])
}

func main() {
	text1 := "longest"
	text2 := "stone"
	lcs(text1, text2)
}
