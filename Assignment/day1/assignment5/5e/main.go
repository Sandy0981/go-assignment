package main

import (
	"fmt"
	"math"
)

var max = math.MinInt

func main() {
	var numbers = []int{2, 4, 5, 7, 55, 77, 33, 9}
	fmt.Println("The max value is", maxElement(numbers))
}

func maxElement(n []int) int {
	for _, v := range n {
		if max < v {
			max = v
		}
	}
	return max
}
