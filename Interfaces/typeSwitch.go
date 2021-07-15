package main

import (
	"fmt"
	"strings"
)


func explain(i interface{}) {
	switch i.(type) {
		case string: 
			fmt.Println(strings.ToUpper(i.(string)))
		case int:
			fmt.Println(i)
		default:
			fmt.Println("bye")
	}
}

func main() {
	fmt.Println("TYPE SWITCH")

	a := "priya"
	explain(a)
	b := 4
	explain(b)	
}
