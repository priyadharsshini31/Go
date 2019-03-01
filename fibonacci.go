package main

import (
	"fmt"
)

func main() {
	
 var max_no = 100
 var first = 0
 var second = 1
 var sum = 0
 fmt.Println(first)
 fmt.Println(second)

for ;sum <= max_no; {   // can drop in the semicolons, for is while loop in go 
    sum = first + second
    first = second
    second = sum
    fmt.Println(sum)
}
}

