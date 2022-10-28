package main

import (
	"fmt"
)

func main() {
	fmt.Println("have fun playing")

	// simple guessing game
	number := 45
	guess := 10
	//or logical operator

	if guess < 1 || guess > 100 {
		fmt.Println(" the guess must be between 1 and 100")
	}
	//and operator
	// you can use the else key word instead of the many if's
	if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("too low! Try again")
		}
		if guess > number {
			fmt.Println("too high! Try again")
		}
		if guess == number {
			fmt.Println("Congratulations! You got it")
		}
		fmt.Println(number <= guess, number >= guess, number == guess)
	}

}
