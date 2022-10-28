package main

import (
	"fmt"
)

func main() {
	// if_else statement to determine the room temperature of a certain day.
	//the else statement has to be on the same line as the closing calibracket for the if statement. if not, you'll receive an error.
	roomTemp := 27
	if roomTemp > 25 {
		fmt.Println("It is so hot today!")
	} else {
		fmt.Println("it is cool today")
	}
	fmt.Println()
}
