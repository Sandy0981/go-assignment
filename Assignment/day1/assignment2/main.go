package main

import (
	"fmt"
	"math"
)

var (
	x, y int
)

func main() {
	var i int = 42
	var j float64 = 3.14
	const pi = math.Pi
	x, y = 3, 4
	const a, b, c = 10, 11, 23

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(pi)
	fmt.Println(x, y)
	fmt.Println(a, b, c)

}
