package main

import "fmt"

func main() {
	var fruit []string
	fruit = append(fruit, "Apple")
	fruit = append(fruit, "Banana")
	fruit = append(fruit, "Grapes")
	fruit = append(fruit, "Mango", "Guava")
	for _, v := range fruit {
		fmt.Println(v)
	}
}
