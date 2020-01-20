package main

import (
	"fmt"
)

func main() {
    A := [5]int{1,1,9,9,9}
index := len(A)-1
repeat := true

for repeat{
if A[index] != 9 {
	A[index] = A[index] + 1
	repeat = false
}else{
	A[index] = 0
	index = index -1 
	repeat = true
	
}
}
	
fmt.Println("The elements in A are", A)
}
