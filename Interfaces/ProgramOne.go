package main

import (
	"fmt"
)

type Shape interface{
	calculateArea()
	calculatePerimeter()
}

//Since Rect implements Shape interface, this is perfectly valid. From the above result, we can see that, dynamic type of s is now Rect and dynamic value of s 
//is the value of the struct Rect which is {5 4}.

//We call it dynamic because we can assign s with a new struct of a different struct type which also implements the interface Shape.

type Rectangle struct {
	width int
	height int
}

type Square struct {
	side int
}

func (r Rectangle) calculateArea() {
	fmt.Println( r.width * r.height)
}

func (r Rectangle) calculatePerimeter() {
	fmt.Println( r.width + r.width + r.height + r.height)
}

func (s Square) calculateArea() {
	fmt.Println(  s.side*s.side)
}

func (s Square) calculatePerimeter() {
	fmt.Println(  4*s.side)
}

func main() {
	fmt.Println("Introduction to interfaces - Program ONE")
	var shape Shape
	r := Rectangle{10,20}
	shape = r
	shape.calculateArea()
	shape.calculatePerimeter()
	shape = Square{10}
	shape.calculateArea()
	shape.calculatePerimeter()
}
