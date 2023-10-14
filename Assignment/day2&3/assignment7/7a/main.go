package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	Age  int
	City string
}

func main() {
	e1 := Employee{
		Id:   100,
		Name: "Sandeep",
		Age:  25,
		City: "Jodhpur",
	}
	e2 := Employee{
		Id:   101,
		Name: "Satyam",
		Age:  28,
		City: "Varanasi",
	}
	fmt.Printf("The employee 1 details are %+v\n", e1)
	fmt.Printf("The employee 2 details are %+v", e2)
}
