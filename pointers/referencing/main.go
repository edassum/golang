package main

import "fmt"

func main() {
	a := 10

	fmt.Println("The variable value is::", a)
	fmt.Println("The address of a is::", &a)

	var b *int = &a
	fmt.Println(b)
}
