package main

import (
	"fmt"
)

func main() {
	if !true {
		fmt.Println("It wont run")
	}
	if !false {
		fmt.Println("It will run")
	}
}
