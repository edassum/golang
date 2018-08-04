package main

import (
	"fmt"
)

var x int = 34

func main() {
	fmt.Println("Inside main function")
	fmt.Println(x)
	y := 345
	fmt.Println("Printing the y")
	fmt.Println(y)
	foo()
}
func foo() {
	fmt.Println("Inside foo function")
	fmt.Println(x)
}
