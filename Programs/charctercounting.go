package main

import (
	"fmt"
)

func main() {
var frequency int
input := "demodemo"
for i := 0; i < len(input); i++{
	if input[i] =='e'{
		frequency ++
}
}
fmt.Println("The character is repeated", frequency, "times")
}
