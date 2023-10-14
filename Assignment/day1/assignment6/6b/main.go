package main

import "fmt"

func main() {
	var fruit = make(map[string]int)
	fruit["apple"] = 10
	fruit["banana"] = 20
	fruit["grapes"] = 12
	fruit["orange"] = 8
	fruit["mango"] = 16
	for i, v := range fruit {
		fmt.Println("The", i, "quantity is", v)
	}
}
