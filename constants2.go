package main

import (
	"fmt"
)

// a simple application to show the use of enumerated constants.
const (
	catSpecialist = iota
	dogSpecialist
	cowSpecialist
)

func main() {

	var specialistType int = catSpecialist
	fmt.Printf("%v", specialistType == catSpecialist)
}

/*  operations that can be determined at compile time are allowed.
they include : arithmetic
            : bitwise operations
			: bit shifting
*/
