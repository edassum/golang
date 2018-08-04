package main

import (
	"fmt"
)

func main() {

	a := 20
	fmt.Println("The actual value of a is::", a)

	fmt.Println("The actual address of a is:", &a)

	var b *int = &a
	fmt.Println("The actual address of a is:", b)

	//dereferencing the value
	fmt.Println("The actual value of a is::", *b)
}
