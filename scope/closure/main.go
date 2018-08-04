package main

import (
	"fmt"
)

var x = 0

func increment() {
	x++
	//	return x
}

func main() {
	//fmt.Println(increment())
	//fmt.Println(increment())
	increment()
	x = x + 2
	fmt.Println(x)
}
