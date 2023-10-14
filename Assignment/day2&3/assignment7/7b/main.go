package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func main() {
	a := Rectangle{
		Width:  10,
		Height: 20,
	}
	fmt.Println("Area of Rectangle is :", a.Area())
	fmt.Println("Perimeter of Rectangle is :", a.Perimeter())
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.Width + r.Height)
}
