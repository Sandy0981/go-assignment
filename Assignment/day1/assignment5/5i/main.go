package main

import "fmt"

func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6}
	doubleValue(&numbers)
	fmt.Println("Double Slice :", numbers)
}

func doubleValue(numbers *[]int) {
	s := *numbers
	for i := range s {
		s[i] *= 2
	}
}
