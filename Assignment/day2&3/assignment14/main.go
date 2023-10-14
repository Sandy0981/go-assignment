package main

import "fmt"

func divide(num, den int) int {
	if den == 0 {
		panic("Cannot do the operation")
	}
	return num / den
	fmt.Println("")
}

func safeDivide(num, den int) int {
	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println("Error", msg)
			return
		}
	}()
	return divide(num, den)
}

func main() {
	result := safeDivide(10, 0)
	fmt.Println(result)
	result = safeDivide(10, 5)
	fmt.Println(result)
	fmt.Println("End of program")
}
