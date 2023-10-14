package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := os.Args[1:]
	numToSearch, _ := strconv.Atoi(a[0])
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	ok := searchNum(numbers, numToSearch)
	if ok {
		fmt.Println("The number", numToSearch, "was found in the slice.")
		return
	}
	fmt.Println("The number", numToSearch, "was not found in the slice.")
}

func searchNum(numbers []int, n int) bool {
	for i := range numbers {
		if numbers[i] == n {
			return true
		}
	}
	return false
}
