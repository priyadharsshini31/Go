package main

import (
	"fmt"
	"strings"
)

func main() {
	
	input := "priyadharsshini"
	output := strings.Replace(input,"a","",-1)
	fmt.Println("the final output is", output)
	
}
