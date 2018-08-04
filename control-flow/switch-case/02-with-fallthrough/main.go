package main

import (
	"fmt"
)

func main() {
	switch "Suman" {
	case "Sumit":
		fmt.Println("Hello Sumit")
		fallthrough
	case "Subho":
		fmt.Println("Hello Subho")
	case "Dinku":
		fmt.Println("Hello Dinku")
	case "Suman":
		fmt.Println("Hello Suman")
		fallthrough
	case "Arnab":
		fmt.Println("Hello Arnab")
		fallthrough
	case "Sayan":
		fmt.Println("Hello Sayan")
	case "Abhi":
		fmt.Println("Hello Abhi")
	}
}
