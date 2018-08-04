package main

import (
	"fmt"
)

func foo(z int) {
	fmt.Println("The address of x in foo is:", &z)
	z = 0
}
func main() {
	x := 20
	fmt.Println("The address of x is:", &x)
	foo(x)
	fmt.Println(x)
}
