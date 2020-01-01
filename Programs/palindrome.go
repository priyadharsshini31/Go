package main

import "fmt"

func check_palindrome(input string)bool{
for i :=0 ; i<len(input)/2; i++{
 if input[i] != input [len(input)-i-1]{
	return false
}
}
return true
	
}

func main() {

value := check_palindrome("ABCDEDCBA")
if value == false{
fmt.Println("The String is not a palindrome")
}else{
fmt.Println("The String is a palindrome")
}


}
