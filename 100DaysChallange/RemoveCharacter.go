package main

import (
	"fmt"
)

func RemoveCharacter(word string){
var final string
char := []byte(word)
for _,ch := range char{
if ch != 'd'{
	final += string(ch)
}
}
fmt.Println("The final string is ",final)
}

func main(){
RemoveCharacter("donald duck")
}
