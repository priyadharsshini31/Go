package main

import (
	"fmt"
	//"time"
)

func add(chan1 chan int){
	chan1 <- 3
}

func sub(chan2 chan int){
	chan2 <- 2
}

func main() {
	fmt.Println("Program to demonstrate that DEFAULT in select statement is non-blocking and also it makes the full select statment also non blocking")
	chan1 := make(chan int)
	chan2 := make(chan int)
	
	go add(chan1)
	go sub(chan2)
	
	//We can make the case to execute , by adding a time.sleep
	
	//time.Sleep(2 * time.Second)
	
	select{
		case res := <- chan1:
			fmt.Println("The value from chan1 is", res)
		case res := <- chan2:
			fmt.Println("The value from chan2 is", res)
		default:
			fmt.Println("Default response")
	}

}
