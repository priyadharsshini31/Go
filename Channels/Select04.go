package main

import (
	"fmt"
)

func add(chan3 chan int) {
	chan3 <- 10+10
}

func main() {
	fmt.Println("Program to explain the use of SELECT statement on channels -- BUFFERED CHANNEL")
	chan1 := make(chan int, 3)
	chan2 := make(chan int, 3)
	chan3 := make(chan int)
	
	chan1 <- 1
	chan1 <- 2
	chan2 <- 1
	chan2 <- 2
	
	go add(chan3)
		
	// If some or all of the channel operations are non-blocking, then one of the non-blocking cases will be chosen randomly and executed immediately.
	
	select {
		case res := <- chan1:
			fmt.Println("The value from chan1", res)
		case res := <- chan2:
			fmt.Println("The value from chan2", res)
		case res := <- chan3:
			fmt.Println("The value from the function is", res)
	}

	fmt.Println("main stopped")
}

