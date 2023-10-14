package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1, 5, 5, 6, 6, 3, 5, 7, 8, 9, 2, 3}
	previousNumber := numbers[0]
	sort.Ints(numbers)
	fmt.Println("Deduplicated Slice : ", removeDuplicates(numbers, previousNumber))
}

func removeDuplicates(numbers []int, previousNumber int) []int {
	deduplicatedNums := make([]int, 0)
	for i := range numbers {
		if numbers[i] != previousNumber {
			deduplicatedNums = append(deduplicatedNums, numbers[i])
			previousNumber = numbers[i]
		}
	}
	return deduplicatedNums
}
