package main

import "fmt"

func Add(nums ...int) int {
	var sum int = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {

	fmt.Println("The sum of 1st Set", Add(1, 2, 3, 4, 5))
	fmt.Println("The sum of 2nd Set", Add(1, 2, 3, 4, 5, 6, 7, 8, 9))
}
