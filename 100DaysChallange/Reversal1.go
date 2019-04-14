package main

import (
	"fmt"
)	
	
func main() {
	var reversed string
	alphabet := ([]byte)("donald duck")
	for i := len(alphabet)-1; i >= 0; i--{
		reversed += (string)(alphabet[i])
	}
	fmt.Println(reversed)
}
