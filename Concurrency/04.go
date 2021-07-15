package main

import (
	"fmt"
)

//DEADLOCK--- Program will throw runtime error
//fatal error: all goroutines are asleep - deadlock!

func main() {
	fmt.Println("Main started")
	c := make(chan int)
	
	// DEADLOCK ---> write to a channel block the go routine until someother go routine reads from the channel
	c <- 10
	fmt.Println("Main stopped")
	
}
