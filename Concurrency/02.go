package main

import (
	"fmt"
)

func squares(c chan int) {

	//write to a channel is blocking until some other go routine reads from the channel
	for i := 0; i <= 10; i++ {
		c <- i*i
	}
	close(c)
}

func main() {
	fmt.Println("Main started")
	c := make(chan int)
	go squares(c)
	
	// when we use range with for loop for channels, it will automatically check for channel close also 
	for val := range c { // read call from an empty channel is blocking
		fmt.Println(val)
	}
	fmt.Println("Main over")
	
}
