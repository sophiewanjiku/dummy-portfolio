package main

import (
	"fmt"
	"strconv"
)
/*
use all variables declared otherwise it would result to an error
type conversion can be done explicitly only. implicitly won't work
for conversion of integer to string, one must import the "strconv" package
*/

// variable declaration
func main(){
	var i int = 27
	fmt.Printf("%v, %T\n", i, i)

	var j string
	j="blue"
	fmt.Printf("%v, %T\n", j, j)

	k := 50
	fmt.Printf("%v, %T\n", k, k)
//type conversion
	var a int = 42
	fmt.Printf("%v, %T\n", a, a)

	var b string
	b = strconv.Itoa(a)
	fmt.Printf("%v, %T\n", b, b)
}

