package main

import (
	"fmt"
)

func main() {
a := []int{1,9,-5,3,0,-10}
maxtillnow := 0
maxsofar := 0
for i := 0 ; i<len(a); i++{
 maxtillnow = maxtillnow + a[i]
if a[i] > maxtillnow{
	maxtillnow = a[i]
}
if maxtillnow > maxsofar{
	maxsofar = maxtillnow
} 
 


}
fmt.Println("The max sum possible is", maxsofar)
}
