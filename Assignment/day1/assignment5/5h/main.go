package main

import "fmt"

func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	o, e := countEvenOdd(numbers)
	fmt.Println("Count of Odd num is ", o, "\nCount of Even num is ", e)
}

func countEvenOdd(a []int) (int, int) {
	codd := 0
	ceven := 0
	for _, v := range a {
		if v%2 == 0 {
			ceven++
		} else {
			codd++
		}
	}
	return codd, ceven
}
