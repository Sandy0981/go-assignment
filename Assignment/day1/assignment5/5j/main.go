package main

import "fmt"

func main() {
	var slice1 = []int{1, 2, 3}
	var slice2 = []int{4, 5, 6}
	fmt.Println("Concatenated slice is", sliceConcat(slice1, slice2))
}

func sliceConcat(a []int, b []int) []int {
	a = append(a[0:], b[0:]...)
	return a
}
