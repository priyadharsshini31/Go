package main

import (
	"fmt"
)

//SLICE holds values of same data type
// slice is created on the top of an underlying array

func main() {
	x := []int{1, 2, 3, 4, 5} //creating slice using COMPOSITE LITERAL
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[3])
	for _, v := range x { // if you need only the second item in the range , use blank identifier to discard
		fmt.Println(v)
	}
	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := range x { // if you need only the first item in the range, leave out the second
		fmt.Println(i)
	}

	//slicing a slice

	fmt.Println(x[:])   // : operator is used to slice, here prints everything
	fmt.Println(x[1:])  // from first position till end
	fmt.Println(x[1:3]) // first, second only. last position not included

	for i := 0; i < len(x); i++ { //accessing elements in the slice using the for loop
		fmt.Println(x[i])
	}

	//APPEND func append(slice []T, elements ...T) []T where ...(variadic parameter) is any number of elements
	// ... T means any number of values of type T , T ... means take everything from the data structure T

	x = append(x, 100, 200, 300)
	fmt.Println(x)
	y := []int{31, 26, 07, 15, 28}
	x = append(x, y...)
	fmt.Println(x)

	//DELETE from slice

	x = append(x[:2], x[5:]...)
	fmt.Println(x)

	// ALLOCATION WITH MAKE
	/* With composite literal , when we append and increase the size of the slice, the underlying array is recreated and old values are copied everytime when we do that, thus taking up
	cosiderable run time in between, but when we know the exact size of the array , then we can create a slice using make as the underlying array will be of
	that capacity and no need to increase later */

	xa := make([]int, 5, 15)  
	// xa := make([]int, 5) // 5 becomes both length and capacity
	fmt.Println(xa)
	fmt.Println(len(xa))
	fmt.Println(cap(xa))

	xa[2] = 3
	xa[4] = 7
	fmt.Println(xa)
	fmt.Println(len(xa))
	fmt.Println(cap(xa))

	xa = append(xa, 10, 11, 12) // we can append also when we create a slice using make, new array double the capacity is created when the cap fills up
	fmt.Println(xa)
	fmt.Println(len(xa)) // length increases
	fmt.Println(cap(xa))
	
	//MULTI DIMENSIONAL SLICE
	
	dd := []string{"Divya", "darsini", "Good", "UNO"}
	fmt.Println(dd)
	pd := []string{"Priya", "Dharsshini", "Good", "BLINK"}
	fmt.Println(pd)

	dp := [][]string{dd, pd}
	fmt.Println(dp)

}
