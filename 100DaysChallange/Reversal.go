package main

import (
	"fmt"
)	
	
func main() {
	var reversed string
	alphabet := []rune("donald duck") //Using rune to not miss on encoding 
	for i := len(alphabet)-1; i >= 0; i--{
		reversed += (string)(alphabet[i])
	}
	fmt.Println(reversed)
}
