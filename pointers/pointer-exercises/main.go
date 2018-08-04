package main

import (
	"fmt"
)

func main() {
	a := 30

	fmt.Println("The actual value of a is:", a)

	//Printing the address where a is stored
	fmt.Println("The address of a is ::", &a)

	//Assigning a pointer to int for a
	var b *int = &a
	//Printing the address where a is stored via b
	fmt.Println("The address of a is ::", b)

	//Printing the actual value stored in the address of b is
	fmt.Println("The actual value stored at address of b is::", *b)

	//Changing the value of b
	*b = 12
	//Printing the value of a
	fmt.Println("The value of a is::", a)

}
