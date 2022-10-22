// embbeding supports composition in go
package main

import (
	"fmt"
)

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal   // embbed animal struct to the bird struct
	SpeedKPH float32
	CanFly   bool
}

func main() {
	b := Bird{}
	b.Name = "EMU"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)
}
