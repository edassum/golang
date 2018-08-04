package main

import (
	"fmt"
)

func main() {
	for i := 50; i < 51; i++ {
		fmt.Println(i, "-", string(i), "-", []byte(string(i)))
	}

	foo := 'a'
	fmt.Println(foo)
}
