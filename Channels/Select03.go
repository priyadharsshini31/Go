package main

import (
	"fmt"
)

func main() {
	fmt.Println("Program to explain the use of SELECT statement on channels -- BUFFERED CHANNEL")
	chan1 := make(chan int, 3)
	chan2 := make(chan int, 3)
	
	chan1 <- 1
	chan1 <- 2
	chan2 <- 1
	chan2 <- 2
		
	//when all the operations on a select statement is non blocking, the select statment will schedule one randomly
	
	select {
		case res := <- chan1:
			fmt.Println("The value from chan1", res)
		case res := <- chan2:
			fmt.Println("The value from chan2", res)
	}

	fmt.Println("main stopped")
}

