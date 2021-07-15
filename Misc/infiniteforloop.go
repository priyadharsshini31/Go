package main

import "fmt"

func main() {
	ans := "priya"
	i := "priya"
	
	//instead of hardcoding i, we can use fmt.Scan and get it from user, for program to make more sense
	for {
		if i == ans {
			fmt.Println("well done")
			break
		}
		if i != ans {
			fmt.Println("Nope")
		}
	}
}

//Till I get the ans as priya , I will continue looping 

//we use break or return to exit from an infinite for loop