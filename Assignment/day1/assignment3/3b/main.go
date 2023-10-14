package main

import (
	"assignment/assignment3/3b/temperature"
	"fmt"
)

var f int

func main() {
	fmt.Println("Enter the value")
	fmt.Scanln(&f)
	fmt.Println("The temperature in celcius is", temperature.FahrenheitToCelcius(f))
}
