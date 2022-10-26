package main

import (
	"fmt"
)

func main() {
	// creating a map
	statePopulations := map[string]int{
		"california":    39520017,
		"texas":         27862596,
		"florida":       20612439,
		"new york":      19745289,
		"pennsylvannia": 12802503,
		"illinois":      12801539,
		"ohio":          11614373,
	}
	fmt.Println(statePopulations)

	//to get the value of a specific key say, texas,
	fmt.Println(statePopulations["texas"])

	//to add an element to an existing map,
	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations)

	//to delete an element from a map, use the delete() function
	cars := make(map[string]string)
	cars = map[string]string{
		"brand": "ford",
		"model": "mustang",
		"year":  "1964",
	}
	fmt.Println(cars)
	delete(cars, "year")
	fmt.Println(cars)
}
