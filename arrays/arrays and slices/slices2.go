package main

import (
	"fmt"
)

func main() {
	//
	a := []int{}
	fmt.Println(a)
	fmt.Printf("length: %v\n", len(a))
	fmt.Printf("capacity: %v\n\n", cap(a))

	//adding values to a slice using the append() function.
	a = append(a, 1)
	fmt.Println(a)
	fmt.Printf("length: %v\n", len(a))
	fmt.Printf("capacity: %v\n\n", cap(a))

	//concatenation of slices
	a = append(a, []int{2, 3, 4, 5}...)
	fmt.Println(a)
	fmt.Printf("length: %v\n", len(a))
	fmt.Printf("capacity: %v\n\n", cap(a))

	//stack operations on slices.
	//removing elements from the front of the slice/stack.
	x := []int{1, 2, 3, 4, 5}
	y := x[1:]
	fmt.Println(y)

	//deletion of elements from behind the stack.
	k := x[:len(a)-1]
	fmt.Println(k)

	//deletion of elements from the middle.
	z := append(a[:2], a[3:]...)
	fmt.Println(z)

}
