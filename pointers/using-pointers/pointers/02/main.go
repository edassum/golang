package main

import "fmt"

func foo(z *int) {
	fmt.Println("The address of x is:", z)
	fmt.Println("The value in z is:", *z)
}
func main() {
	x := 8
	fmt.Println("The value in x is:", x)
	fmt.Println("The address of x is:", &x)
	//Passing the address of x
	foo(&x)
}
