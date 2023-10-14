package main

import "fmt"

func main() {
	var numbers []int
	numbers = append(numbers, 5)
	fmt.Println(numbers)
	numbers = append(numbers, 10)
	fmt.Println(numbers)
	numbers = append(numbers, 15)
	fmt.Println(numbers)
	numbers = append(numbers, 20)
	fmt.Println(numbers)
	numbers = append(numbers, 25)
	fmt.Println(numbers)
	fmt.Println("Len and Cap of numbers slice is", len(numbers), cap(numbers))
	numbers = append(numbers[:2], numbers[3:]...)
	fmt.Println("After removal", numbers)
}
