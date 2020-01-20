package main

import (
	"fmt"
)

func main() {
    var A = []int {0,0,9,9,9}

i := 0

if A[i] == 0{
 A = removeleadingzeros(A)
}


	
fmt.Println("The elements in A are", A)
}

func removeleadingzeros(A []int) (B []int){
zeroremove := 0
i := 0
for i = 0 ; i < len(A); i++ {
if A[i] == 0{
zeroremove = zeroremove + 1
}
}
B = A[zeroremove:]
return B
}



