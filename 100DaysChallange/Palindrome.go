package main

import (
	"fmt"
)

func palindrome(word string){
var temp_word string
letters := []byte(word)
for i := len(letters)-1; i >= 0; i--{
temp_word += string(letters[i])
}
fmt.Println(temp_word)
if temp_word == word{
	fmt.Println("It is a palindrome")
}else{
	fmt.Println("not a palindrome")
}

}

func main() {

palindrome("mom")
palindrome("good")
}
