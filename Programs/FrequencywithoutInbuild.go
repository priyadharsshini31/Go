package main

import(
   "fmt"
)



func findFrequency(name string) int{
	times := 0
	var alphabet byte
	alphabet = 's'
	bs := ([]byte)(name)
	for i :=0 ; i < len(bs); i++{
		if bs[i] == alphabet{
			times++
			}
		}
		return times
	
	
}

func main(){
	frequency := findFrequency("priyadharsshini")
	fmt.Println(frequency)
}
