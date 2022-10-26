package main

import (
	"fmt"
)

/*
1. naming of constants. -> named the same way variables are named.
typed contsants. include the data type of the constants
the value of a consant cannot be modified. constants must be equal to something executed at compietime and not run time.
*/
/*
2.  enumerated constants. declared at package level.
constanst can be declared in block code
example:
*/
const (
	i = iota
	j
	k
)

func main() {
	//typed contsants. include the data type of the constants
	const myConst int = 42
	fmt.Printf("myConst:\n %v, %T\n", myConst, myConst)

	// untyped constants
	const a = 27
	fmt.Printf("a:\n %v, %T\n", a, a)

	// arithmetic operations on constants
	fmt.Printf("sum= %v, %T\n", myConst+a, myConst+a)

	//printing enumerated constants
	fmt.Printf("i: %v\n", i)
	fmt.Printf("j: %v\n", j)
	fmt.Printf("k: %v\n", k)
}
