package main

import "fmt"

var sum int
var avg int

func main() {
	var numbers = []int{10, 20, 30, 40, 50}
	fmt.Println("Sum is", Sum(numbers))
	fmt.Println("Average is", Avg(numbers))
}

func Sum(a []int) int {
	for _, v := range a {
		sum += v
	}
	return sum
}

func Avg(a []int) int {
	return sum / len(a)
}
