package main

import "fmt"

type Address struct {
	Street  int
	City    string
	Zipcode int
}

type Student struct {
	Name string
	Address
}

func main() {
	s := Student{
		Name: "Sandeep",
		Address: Address{
			Street:  12,
			City:    "Jodhpur",
			Zipcode: 342001,
		},
	}
	fmt.Print(s)
}
