package main

import (
	"fmt"

	"../stringutil"
)

func main() {
	fmt.Println("insde main")
	fmt.Println(stringutil.Name)
	stringutil.Reverse()
	//stringutil.Reversename()
}
