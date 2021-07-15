package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "I am :IronnorI Ma, i"
	fmt.Println(saveIronMan(input))
}

func saveIronMan(input string) bool {
	input2 := strings.ToLower(input)
	temp := ""
	for _, value := range input2 {
		if value >= 97 && value <= 122 {
			temp = temp + string(value)
		}
	}

	if len(temp)%2 != 0 {
		return false
	} else {

		temp1 := temp[:len(temp)/2]
		temp2 := temp[len(temp)/2:]
		temp3 := ""
		for i := len(temp2) - 1; i >= 0; i-- {
			temp3 = temp3 + string(temp2[i])
		}
		if temp1 == temp3 {
			return true
		}

	}
	return false
}
