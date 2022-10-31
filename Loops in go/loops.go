package main

import (
	"fmt"
)

// in go there exists only one loop which is the for loop.
func main() {

	for i := 0; i < 10; i++ {
		//the continue statement. skips a number.
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
	// break statement stops the execution of the loop.

}
