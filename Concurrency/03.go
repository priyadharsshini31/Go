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
	
	for {
	
		//ok when used with channels says if the channel is still open for read/write or not 
		val, ok := <-c 
		if ok == true {	
			fmt.Println(val)
		} else {
			fmt.Println("The channel is closed ")
			break
		}
	}
			
	fmt.Println("Main over")
	
}
