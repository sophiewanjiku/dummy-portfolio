package main

import (
	"fmt"
)

func main() {
	//initialising an array.
	grades := [3]int{43, 35, 92}
	fmt.Printf("grades: %v\n", grades)

	//2.
	var students [3]string
	students[0] = "sophie"
	fmt.Printf("students: %v\n", students)

	//getting the size of an array
	fmt.Printf("number of students: %v\n", len(students))

	//array of arrays.
	/*
		this will also work:
		var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
		but for simplicity and readability;*/

	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Printf("identityMatrix:\n %v\n %v\n %v\n", identityMatrix[0], identityMatrix[1], identityMatrix[2])

	/* passing arrays into a function.
	arrays in go are considered as values. therefore in printing the application below,
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	the outcome is a copy of the first array.
	to prevent this we use pointers when passing arrays to show that b is pointing to the same values of a. hence a change in b is therefore reflected in a. */

	a := [...]int{1, 2, 3}
	b := &a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
}
