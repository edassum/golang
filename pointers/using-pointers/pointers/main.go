package main

import (
	"fmt"
)

func foo(z *int) {
	*z = 5
}

func main() {
	x := 10
	foo(&x)
	fmt.Println(x)
}
