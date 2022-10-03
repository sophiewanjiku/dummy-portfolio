package main

import (
	"fmt"
)
/* go data types
1. boolean- give a true or false
*/
func main(){

	var m bool = true
	fmt .Printf("%v, %T\n",m ,m)

	/*2.integers*/
	//singned integers -> take both negative and positive numbers(keyword=int)

	var x int =-29
	fmt .Printf("%v,%T\n", x, x )

	//unsigned integers-> take only positive numbers  (keyword=uint)

	var y uint= 50
	fmt .Printf("%v, %T\n", y, y)

	//arithmetic operations on integers
	a := 10  //1010
	b := 3   //0011
	fmt .Println(a + b)
	fmt .Println(a - b)
	fmt .Println(a / b)
	fmt .Println(a * b)
	fmt .Println(a % b)

	//bitwise operations on integers.
	//and operator
	fmt .Println(a & b) //0010
	//or operator
	fmt .Println(a | b) //1011
	//exclusive or operator
	fmt .Println(a ^ b) //1001
	//and not operator
	fmt .Println(a &^ b) //0100
	//shift operators on integers.
	j := 8
	fmt .Println(j << 3)  //= 2^3 * 2^3
	fmt .Println(j >> 3)  //= 2^3 / 2^3
}