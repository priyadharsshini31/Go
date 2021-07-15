package main

import (
	"fmt"
	"time"
)

func sum(a int, b int) {
	fmt.Println(a+b)
}

func multiply(a int, b int) {
	fmt.Println(a*b)
}

func main() {
	fmt.Println("This is my first example on go concurrency")
	go func(){
		fmt.Println("Immediately execute")
	}()
	
	go sum(3,4)
	
	go multiply(2,3)
	
	time.Sleep(5)
	
	fmt.Println("Program over")
}