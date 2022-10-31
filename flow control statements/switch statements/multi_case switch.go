package main

import (
	"fmt"
)

func main() {
	month := 8
	switch month {
	case 1, 3, 5:
		fmt.Println("31 day month in the first half of the year")
	case 4, 6:
		fmt.Println("30 day month in the first half of the year")
	case 7, 8, 10, 12:
		fmt.Println("31 day month in the second half of the year")
	case 9, 11:
		fmt.Println("30 day month in the second half of the year")
	default:
		fmt.Println("february. Has 28 days")
	}
}
