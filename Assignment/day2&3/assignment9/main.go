package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius int
}

func (c Circle) Area() float64 {
	return math.Pi * float64(c.Radius) * float64(c.Radius)
}

type Rectangle struct {
	Length int
	Breath int
}

func (r Rectangle) Area() float64 {
	return float64(r.Length) * float64(r.Breath)
}

func main() {
	var c Shape = Circle{6}
	var r Shape = Rectangle{10, 20}
	fmt.Println("The area of Rectangle ", r.Area())
	fmt.Println("The area of Circle", c.Area())
}
