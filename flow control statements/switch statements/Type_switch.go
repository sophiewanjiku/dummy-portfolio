package main

import (
	"fmt"
)

func main(){ 
	var i interface{} = 1
	switch i.(type){
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is a float 64 ")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("another type")
	}
	fmt.Println()
}