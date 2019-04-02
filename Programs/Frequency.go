package main

import(
   "fmt"
   "strings"
)


func findFrequency(name string) int{
	times := strings.Count(name,"s")
	return times
	
}

func main(){
	frequency := findFrequency("priyadharsshini")
	fmt.Println(frequency)
}
