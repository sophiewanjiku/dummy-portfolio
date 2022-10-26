package main

import (
	"fmt"
)

func main() {
	// if_else statement to determine the room temperature of a certain day
	roomTemp := 27
	if (roomTemp > 25) {
		fmt.Println("It is so hot today!")
	} else {
		fmt.Println ("it is cool today")
	}
	fmt.Println()
}
