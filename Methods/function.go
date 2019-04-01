/* Why need of method when we have function?
1)Go is not a complete Object Oriented language.Class is not supported. Define type to achieve same functionality
2) Methods can have same name with different types but functions cannot have same name 
*/
package main

import(
"fmt"
)

type square struct{
    side float64
}

func Area(s square) float64 {
	return s.side * s.side
}

func main() {
	sqr := square{4}
	fmt.Println(Area(sqr)) //Pay attention to calling
}
