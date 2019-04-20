package main

import (
	"fmt"
	"strings"
)

func countWords(sentence string){
	var m map[string]int
	m = make(map[string]int)
	sent := strings.Fields(sentence) //Fields splits the string s around each instance of one or more consecutive white space characters, returning an array of substrings of s or an empty list if s contains only white space.
	for _, word := range sent {
		value , ok := m[word]
			if ok{
				m[word] = value + 1
			}else{
				m[word] = 1
			}
		}
	
	for key, value := range m{
		fmt.Println(key, value)
	}
}

func main() {
	countWords("hello how are you hello")
	
}
