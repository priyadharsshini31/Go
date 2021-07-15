package main

import (
	"fmt"
	"strings"
)

func findOccurances(sent string) {
	var words map[string]int
	words = make(map[string]int)
	wordArray := strings.Split(sent, " ")
	for i, _ := range wordArray {
		words[wordArray[i]]++
	}
	
	for key , value := range words {
		fmt.Println("the word", key, "is repeated", value, "times")
	}
	}
			
	

func main() {
	sent := "Twinkle Twinkle little star"
	findOccurances(sent)
	
}
