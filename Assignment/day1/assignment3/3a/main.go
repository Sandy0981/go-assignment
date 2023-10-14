package main

import (
	"assignment/assignment3/3a/calculator"
	"fmt"
)

var i int = 0
var x int

func main() {
	fmt.Println("1.Addition 2.Subtraction 3.Multiply 4.Division 5.Square 5.Exit")
	var x int
	fmt.Scanln(&x)
	switch x {
	case 1:
		fmt.Println(calculator.Sum(10, 20))
	case 2:
		fmt.Println(calculator.Sub(20, 10))
	case 3:
		fmt.Println(calculator.Mul(20, 10))
	case 4:
		x, y := calculator.Div(20, 10)
		fmt.Println("Quotient is ", x)
		fmt.Println("Remainder is ", y)
	case 5:
		fmt.Println(calculator.Square(20))
	default:
		fmt.Println("Invalid operation.")
	}
}
