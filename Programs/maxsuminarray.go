package main

import (
	"fmt"
)

func main() {
a := []int{1,-9,-1,3,4,-10}
maxtillnow := 0
maxsofar := 0
for i := 0 ; i<len(a); i++{
maxtillnow = 0
for j := i ; j < len(a); j++{
maxtillnow = maxtillnow + a[j]

if maxtillnow > maxsofar{

	maxsofar = maxtillnow
} 
}
 



}
fmt.Println("The max sum possible is", maxsofar)
}
