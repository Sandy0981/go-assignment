package main

import "fmt"

type Animal interface {
	MakeSound() string
}

type Dog struct{}

func (d Dog) MakeSound() string {
	return "Bark"
}

type Cat struct{}

func (c Cat) MakeSound() string {
	return "Meow"
}

func main() {
	var dog Animal = Dog{}
	var cat Animal = Cat{}
	fmt.Println(dog.MakeSound())
	fmt.Println(cat.MakeSound())
}
