package main
import "fmt"
func main(){
const x int = 5
fmt.Println(x)
const name string = "Priya"
//name = "dharsshini" will throw error cos you cannot reassign
fmt.Println(name)
const age = 25 //const can be given without datatype also
fmt.Println(age)
}
