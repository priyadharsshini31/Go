package main

import "fmt"

func longestPalindrome(s string) string {
    if s == ""{
        return ""
    }
    start , end := 0,0
    length := 0
    for i := 0; i < len(s); i++ {
        len1 := expand(s,i,i)
        len2 := expand(s,i,i+1)
        if len1 > len2 {
            length = len1
        } else {
            length = len2
        }
        if length > end -start {
            start = i-(length-1)/2
            end = i + length/2
        }
    }
    return s[start:end+1]
}

func expand(s string, left int, right int) int{
    for left >=0 && right < len(s) && s[left] == s[right]{
            left --
        right ++
    }
    return right - left - 1
}

func main(){
	str := "racecar"
	fmt.Println(longestPalindrome(str))
}