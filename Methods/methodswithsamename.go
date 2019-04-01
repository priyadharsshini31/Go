package main

import(
"fmt"
)

type square struct{
    side float64
}

type rectangle struct {
     length, breadth float64
}

func(s square) Area() float64 {
	return s.side * s.side
}

func(r rectangle) Area() float64 {
        return r.length * r.breadth
}

func main() {
	sqr := square{4}
	rect := rectangle{3,4}
	fmt.Println(sqr.Area())
	fmt.Println(rect.Area())
}
