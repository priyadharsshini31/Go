package main

import (
	"fmt"
)

func split(s string) {
	words := make([]string, 100)
	j := 0
	letters := []byte(s)
	var finalword string
	for i, _ := range letters {
		if string(letters[i]) != " " {
			finalword += string(letters[i])
		} else {
			fmt.Println(finalword)
			words[j] = finalword
			j++
			finalword = ""
		}
	}
	fmt.Println(words)
}

func main() {
	s := "Twinkle twinkle little star "
	split(s)
}
