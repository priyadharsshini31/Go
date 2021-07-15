package main

import (
	"fmt"
	"time"
)
func add(chan1 chan int) {
	//since this sleeps for only 3 seconds this will get executed first
	time.Sleep(3 * time.Second)
	chan1 <- 3+4
}

func sub(chan2 chan int) {
	time.Sleep(5 * time.Second)
	chan2 <- 4-3
}

func main() {
	fmt.Println("Program to explain the use of SELECT statement on channels -- BLOCKING")
	chan1 := make(chan int)
	chan2 := make(chan int)
	
	go add(chan1)
	go sub(chan2)
		
	//when all the operations on a select statement is blocking, the select will block the routine from where it is called 
	//and schedule the other go routines and whichever finished first will be considered and the calling go routine is unblocked
	
	select {
		case res := <- chan1:
			fmt.Println("The value from add function is", res)
		case res := <- chan2:
			fmt.Println("The value from sun function is", res)
	}

	fmt.Println("main stopped")
}

