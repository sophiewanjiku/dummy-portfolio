package main

import (
	"fmt"
)

func main() {
	a := 35
	b := 35
	if a > b{
		fmt.Println("a is greater than b")
	} else if a < b{
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a is equal to b")
	}
	
}
