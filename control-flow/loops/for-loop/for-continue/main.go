package main

import (
	"fmt"
)

//Program to print odd numbers
func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i > 50 {
			break
		}
	}
}
