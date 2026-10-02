package main

import (
	"fmt"
)

func celtofah(fah float64) float64 {
	fahrenheit := (fah * 1.8) + 32

	return fahrenheit
}

func main() {

	var celsius float64
	fmt.Scan(&celsius)

	fahrenheit := celtofah(celsius)

	fmt.Printf(" Weather in Celsius : %.2f", celsius)
	fmt.Printf(" Weather in Fahrenheit : %.2f", fahrenheit)
}
