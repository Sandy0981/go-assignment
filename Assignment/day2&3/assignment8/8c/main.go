package main

import (
	"8c/model"
	"8c/person"
)

func main() {
	a := model.Person{
		Name: "Sandeep",
		Age:  27,
	}
	person.PrintPerson(a)
}
