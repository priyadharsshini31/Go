package main

import(
   "fmt"
)
type myInt int

func(a myInt) add(b myInt){
	fmt.Println(a+b);
}

func main(){
a := myInt(5)
a.add(10)
}
