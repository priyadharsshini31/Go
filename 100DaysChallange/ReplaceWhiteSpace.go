package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Have a Good Day"
	sentence = strings.Replace(sentence," ", "@", -1)
	fmt.Println(sentence)
}
