package main

import "fmt"

func main(){

	var m map[string]int
	m = make(map[string]int)
	name := "goopher"
	alph := []byte(name)
	for _,ch := range alph{
	value, ok := m[string(ch)]
	if ok{
		 m[string(ch)] = value + 1
	}else{
		 m[string(ch)] = 1
	}
	}
	for key,val := range m{
	if m[key] > 1{
		fmt.Println(key,val)
	}
	}
}
