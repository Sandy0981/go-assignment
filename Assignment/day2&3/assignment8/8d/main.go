package main

import (
	"8d/model"
	"8d/shape"
	"fmt"
)

func main() {

	a := model.Circle{
		Radius: 5,
	}
	fmt.Println(shape.Area(a))
}
