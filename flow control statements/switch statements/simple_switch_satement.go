package main

import (
	"fmt"
)

func main() {
	//a simple single case switch statement showing the days of the week.
	dayOfTheWeek := 5
	switch dayOfTheWeek {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
		//the default keyword -> specifies the code to run if ther is no case match.
	default:
		fmt.Println("Not a day of the week")
	}
	fmt.Println("have a nice day")
}
