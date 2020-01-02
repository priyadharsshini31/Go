package main

import (
	"fmt"
	"strings"
)

func main() {
	
	input := "priyadharsshini"
	output := strings.Replace(input,"a","",-1) //If n < 0, there is no limit on the number of replacements or n number if replacements

	fmt.Println("the final output is", output)
	
}
