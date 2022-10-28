package main

import (
	"fmt"
)

// nested if example.
func main() {
	x := 25
	if x > 20 {
		fmt.Println("x is greater than 20")
		if x < 100 {
			fmt.Println("x is less than 100")
		}
	} else {
		fmt.Println("both are false")
	}
	fmt.Println("hello")
}
