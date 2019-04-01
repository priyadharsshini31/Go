package main

import(
"fmt"
)

type square struct{
    side float64
}

func(s square) Area() float64 {
	return s.side * s.side
}

func main() {
	sqr := square{4}
	fmt.Println(sqr.Area()) //Pay attention to the calling
}
