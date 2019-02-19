package main

import (
	"fmt"
	"sync"
)



func main() {

    	
	var counter = struct{
    sync.RWMutex
    m map[string]int
}{m: make(map[string]int)}

	counter.RLock()
	var priyaAge = counter.m["priya"]
	counter.RUnlock()
	fmt.Println(priyaAge)
	counter.Lock()
	counter.m["priya"] = 27
	counter.Unlock()
	fmt.Println("The value of priya is ", counter.m["priya"])

	}
