package main

import (
	"fmt"
)

type Doctor struct {
	Number     int
	ActorName  string
	Companions []string
}

func main() {
	aDoctor := Doctor{
		Number:    3,
		ActorName: "jayjay",
		Companions: []string{
			"liz shaw",
			"jo grant",
		},
	}

	//accessing a specific value in a struct
	fmt.Println(aDoctor.ActorName)

	//anonymous structures.
	famer := struct{ name string }{name: "falal"}
	fmt.Println(famer)

	//passing of structures
	anotherFamer := famer
	anotherFamer.name = "leah"
	fmt.Println(anotherFamer)
	//to prevent copying of the uderlying data, use pointers.

}
