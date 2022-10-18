package main

import (
	"fmt"
)

func main() {
	//initialising a slice.
	a := []int{1, 2, 3}
	fmt.Println(a)

	//getting the length of a slice.
	fmt.Printf("length: %v\n", len(a))

	//getting the capacity of a slice
	fmt.Printf("capacity: %v\n", cap(a))

	//passing a value to a slice. pointers not required.
	i := []int{0, 5, 6}
	b := i
	b[1] = 7
	fmt.Println(i)
	fmt.Println(b)

	//other ways of declaring slices.
	j := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k := j[:]   //slice of all the elements
	x := j[3:]  //slice from fourth element to end
	y := j[:6]  //slice of the first 6 elements
	z := j[3:6] //slice of the 4th, 5th and 6th elements.

	fmt.Printf("j is : %v\n", j)
	fmt.Printf("k is : %v\n", k)
	fmt.Printf("x is : %v\n", x)
	fmt.Printf("y is : %v\n", y)
	fmt.Printf("z is : %v\n", z)

	//creating a slice using the make() function.
	mySlice := make([]int, 3, 100)
	fmt.Println(mySlice)
	fmt.Printf("length: %v\n", len(mySlice))
	fmt.Printf("capacity: %v\n", cap(mySlice))
}
